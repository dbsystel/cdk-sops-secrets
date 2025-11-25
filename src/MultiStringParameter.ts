import * as fs from 'fs';
import { IKey } from 'aws-cdk-lib/aws-kms';
import { ParameterTier, StringParameter } from 'aws-cdk-lib/aws-ssm';
import { ResourceEnvironment, Stack } from 'aws-cdk-lib/core';
import { Construct } from 'constructs';
import * as YAML from 'yaml';
import { SopsCommonParameterProps } from './SopsStringParameter';
import { ResourceType, SopsSync, SopsSyncOptions } from './SopsSync';

interface JSONObject {
  [key: string]: unknown;
}

export interface MultiStringParameterProps extends SopsCommonParameterProps {
  /**
   * The seperator used to seperate keys
   *
   * @default - '/'
   */
  readonly keySeparator?: string;
  /**
   * The prefix used for all parameters
   *
   * @default - '/'
   */
  readonly keyPrefix?: string;
}

function flattenJSON(
  data: JSONObject,
  parentKey: string = '',
  result: JSONObject = {},
  keySeparator = '',
): JSONObject {
  for (const key of Object.keys(data)) {
    const value = data[key];
    const newKey = parentKey ? `${parentKey}${keySeparator}${key}` : key;

    if (Array.isArray(value)) {
      value.forEach((item, index) => {
        const arrayKey = `${newKey}[${index}]`;
        if (item !== null && typeof item === 'object' && !Array.isArray(item)) {
          flattenJSON(item as JSONObject, arrayKey, result, keySeparator);
        } else {
          result[arrayKey] = item as unknown;
        }
      });
    } else if (
      value !== null &&
      typeof value === 'object' &&
      !Array.isArray(value)
    ) {
      flattenJSON(value as JSONObject, newKey, result, keySeparator);
    } else {
      result[newKey] = value;
    }
  }
  return result;
}

export class MultiStringParameter extends Construct {
  readonly sync: SopsSync;
  readonly encryptionKey: IKey;
  readonly stack: Stack;
  readonly env: ResourceEnvironment;
  readonly keyPrefix: string;
  readonly keySeparator: string;

  constructor(scope: Construct, id: string, props: MultiStringParameterProps) {
    super(scope, id);

    this.encryptionKey = props.encryptionKey;
    this.stack = Stack.of(scope);
    this.env = {
      account: this.stack.account,
      region: this.stack.region,
    };
    this.keyPrefix = props.keyPrefix ?? '/';
    this.keySeparator = props.keySeparator ?? '/';

    const keys = this.parseFile(props.sopsFilePath!, this.keySeparator)
      .filter((key) => !key.startsWith('sops'))
      .map((value) => {
        // Ass we flatten array to [number] path notations, we have to fix this for parameter store
        let fixedKey = value.replace('[', this.keySeparator);
        fixedKey = fixedKey.replace(']', this.keySeparator);
        if (fixedKey.endsWith(this.keySeparator)) {
          fixedKey = fixedKey.slice(0, -1);
        }
        fixedKey = fixedKey.replace(
          this.keySeparator + this.keySeparator,
          this.keySeparator,
        );

        // The secret name can contain ASCII letters, numbers, and the following characters: /_+=.@-
        const allowedChars = '/_+=.@-';
        for (let i = 0; i < fixedKey.length; i++) {
          const char = fixedKey[i];
          if (
            !(
              (char >= 'a' && char <= 'z') ||
              (char >= 'A' && char <= 'Z') ||
              (char >= '0' && char <= '9') ||
              allowedChars.includes(char)
            )
          ) {
            fixedKey = fixedKey.slice(0, i) + '_' + fixedKey.slice(i + 1);
          }
        }
        return `${this.keyPrefix}${fixedKey}`;
      });

    keys.forEach((key) => {
      new StringParameter(this, 'Resource' + key, {
        parameterName: key,
        description: props.description,
        tier: ParameterTier.STANDARD,
        stringValue: ' ',
      });
    });

    this.sync = new SopsSync(this, 'SopsSync', {
      encryptionKey: this.encryptionKey,
      resourceType: ResourceType.PARAMETER_MULTI,
      flattenSeparator: this.keySeparator,
      parameterNames: keys,
      target: this.keyPrefix,
      ...(props as SopsSyncOptions),
    });
  }

  private parseFile(sopsFilePath: string, keySeparator: string): string[] {
    const _sopsFileFormat = sopsFilePath.split('.').pop();
    switch (_sopsFileFormat) {
      case 'json': {
        return Object.keys(
          flattenJSON(
            JSON.parse(fs.readFileSync(sopsFilePath, 'utf-8')),
            '',
            undefined,
            keySeparator,
          ),
        );
      }
      case 'yaml': {
        const content = fs.readFileSync(sopsFilePath, 'utf-8');
        const data = YAML.parse(content) as JSONObject;
        return Object.keys(flattenJSON(data, '', undefined, keySeparator));
      }
      default: {
        throw new Error(
          `Unsupported sopsFileFormat for multiple parameters: ${_sopsFileFormat}`,
        );
      }
    }
  }
}
