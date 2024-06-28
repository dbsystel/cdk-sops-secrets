import * as fs from 'fs';
import { IKey } from 'aws-cdk-lib/aws-kms';
import { ParameterTier, StringParameter } from 'aws-cdk-lib/aws-ssm';
import { ResourceEnvironment, Stack } from 'aws-cdk-lib/core';
import { Construct } from 'constructs';
import * as YAML from 'yaml';
import { SopsStringParameterProps } from './SopsStringParameter';
import {
  CreationType,
  ResourceType,
  SopsSync,
  SopsSyncOptions,
} from './SopsSync';

interface JSONObject {
  [key: string]: any;
}

export interface MultiStringParameterProps extends SopsStringParameterProps {
  readonly keySeperator?: string;
  readonly keyPrefix?: string;
}

function flattenJSON(
  data: JSONObject,
  parentKey: string = '',
  result: JSONObject = {},
  keySeparator = '',
): JSONObject {
  for (let key in data) {
    if (data.hasOwnProperty(key)) {
      let newKey = parentKey ? `${parentKey}${keySeparator}${key}` : key;
      if (Array.isArray(data[key])) {
        data[key].forEach((item: JSONObject | null, index: any) => {
          let arrayKey = `${newKey}[${index}]`;
          if (typeof item === 'object' && item !== null) {
            flattenJSON(item, arrayKey, result, keySeparator);
          } else {
            result[arrayKey] = item;
          }
        });
      } else if (typeof data[key] === 'object' && data[key] !== null) {
        flattenJSON(data[key], newKey, result, keySeparator);
      } else {
        result[newKey] = data[key];
      }
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
    this.keySeparator = props.keySeperator ?? '/';

    const keys = this.parseFile(props.sopsFilePath!, this.keySeparator, this.keyPrefix).filter(
      (key) => !key.startsWith('sops'),
    );

    keys.forEach((key) => {
      new StringParameter(this, 'Resource' + key, {
        parameterName: key,
        tier: ParameterTier.STANDARD,
        stringValue: ' ',
      });
    });

    this.sync = new SopsSync(this, 'SopsSync', {
      encryptionKey: this.encryptionKey,
      resourceType: ResourceType.PARAMETER,
      creationType: CreationType.MULTI,
      flatten: true,
      flattenSeparator: this.keySeparator,
      parameterKeyPrefix: this.keyPrefix,
      parameterNames: keys,
      ...(props as SopsSyncOptions),
    });
  }

  private parseFile(sopsFilePath: string, keySeparator: string, keyPrefix: string): string[] {
    const _sopsFileFormat = sopsFilePath.split('.').pop();
    var result: string[] = [];
    switch (_sopsFileFormat) {
      case 'json': {
        result = Object.keys(
          flattenJSON(JSON.parse(fs.readFileSync(sopsFilePath, 'utf-8')), '', undefined, keySeparator),
        );
        break;
      }
      case 'yaml': {
        const content = fs.readFileSync(sopsFilePath, 'utf-8');
        const data = YAML.parse(content) as JSONObject;
        result = Object.keys(flattenJSON(data, '', undefined, keySeparator));
        break;
      }
      default: {
        throw new Error(
          `Unsupported sopsFileFormat for multiple parameters: ${_sopsFileFormat}`,
        );
      }
    }
    return result.map(value => `${keyPrefix}${value}`);
  }
}
