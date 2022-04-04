import * as fs from 'fs';
import { IKey, Key } from '@aws-cdk/aws-kms';
import {
  Code,
  IFunction,
  Runtime,
  SingletonFunction,
} from '@aws-cdk/aws-lambda';
import { Asset } from '@aws-cdk/aws-s3-assets';
import { ISecret, Secret, SecretProps } from '@aws-cdk/aws-secretsmanager';
import {
  Annotations,
  Construct,
  CustomResource,
  SecretValue,
} from '@aws-cdk/core';

/**
 * Configuration options for the SopsSync
 */
export interface SopsSyncOptions {
  /**
   * The custom resource provider to use. If you don't specify any, a new
   * provider will be created - or if already exists within this stack - reused.
   *
   * @default - A new singleton provider will be created
   */
  readonly sopsProvider?: IFunction;

  /**
   * The filepath to the sops file
   */
  readonly sopsFilePath: string;

  /**
   * The format of the sops file.
   *
   * @default - The fileformat will be derived from the file ending
   */
  readonly sopsFileFormat?: undefined | 'json' | 'yaml';

  /**
   * The kmsKey used to encrypt the sops file. Encrypt permissions
   * will be granted to the custom resource provider.
   *
   * @default - The key will be derived from the sops file
   */
  readonly sopsKmsKey?: IKey[];

  /**
   * The age key that should be used for encryption.
   */
  readonly sopsAgeKey?: SecretValue;

  /**
   * Should the encrypted sops value should be converted to JSON?
   * Only JSON can be handled by cloud formations dynamic references.
   *
   * @default true
   */
  readonly convertToJSON?: boolean;

  /**
   * Should the structure be flattened? The result will be a flat structure and all
   * object keys will be replaced with the full jsonpath as key.
   * This is usefull for dynamic references, as those don't support nested objects.
   *
   * @default true
   */
  readonly flatten?: boolean;
}

/**
 * The configuration options extended by the target Secret
 */
export interface SopsSyncProps extends SopsSyncOptions {
  /**
   * The secret that will be populated with the encrypted sops file content.
   */
  readonly secret: ISecret;
}

/**
 * The custom resource, that is syncing the content from a sops file to a secret.
 */
export class SopsSync extends Construct {
  /**
   * The current versionId of the secret populated via this resource
   */
  readonly versionId: string;

  /**
   * The format of the input file
   */
  readonly sopsFileFormat: 'json' | 'yaml';

  /**
   * Was the format converted to json?
   */
  readonly converToJSON: boolean;

  /**
   * Was the structure flattened?
   */
  readonly flatten: boolean;

  constructor(scope: Construct, id: string, props: SopsSyncProps) {
    super(scope, id);

    this.converToJSON = props.convertToJSON ?? true;
    this.flatten = props.flatten ?? true;

    const sopsFileFormat =
      props.sopsFileFormat ?? props.sopsFilePath.split('.').pop();
    switch (sopsFileFormat) {
      case 'json': {
        this.sopsFileFormat = 'json';
        break;
      }
      case 'yaml': {
        this.sopsFileFormat = 'yaml';
        break;
      }
      case 'yml': {
        this.sopsFileFormat = 'yaml';
        break;
      }
      default: {
        throw new Error(`Unsupported sopsFileFormat ${sopsFileFormat}`);
      }
    }

    const code = Code.fromAsset('assets/cdk-sops-lambda.zip');
    const provider =
      props.sopsProvider ??
      new SingletonFunction(this, 'Function', {
        code,
        runtime: Runtime.GO_1_X,
        handler: 'cdk-sops-secrets',
        uuid: 'cdk-sops-secrets',
      });

    if (!fs.existsSync(props.sopsFilePath)) {
      throw new Error(`File ${props.sopsFilePath} does not exist!`);
    }

    const sopsAsset = new Asset(this, 'Asset', {
      path: props.sopsFilePath,
    });

    if (provider.role !== undefined) {
      if (props.sopsKmsKey !== undefined) {
        props.sopsKmsKey.forEach((key) => key.grantDecrypt(provider.role!));
      }
      const fileContent = fs.readFileSync(props.sopsFilePath);
      const regex = /arn:aws:kms:[a-z0-9-]+:[\d]+:key\/[a-z0-9-]+/gm;
      const results = regex.exec(fileContent.toString());
      if (results !== undefined) {
        results?.forEach((result, index) =>
          Key.fromKeyArn(this, `SopsKey${index}`, result).grantDecrypt(
            provider.role!,
          ),
        );
      }
      props.secret.grantWrite(provider);
      sopsAsset.bucket.grantRead(provider);
    } else {
      Annotations.of(this).addWarning(
        'Please ensure propper permissions for the passed lambda function:\n  - write Access to the secret\n  - encrypt with the sopsKmsKey\n  - download from asset bucket',
      );
    }

    const cr = new CustomResource(this, 'Resource', {
      serviceToken: provider.functionArn,
      resourceType: 'Custom::SopsSync',
      properties: {
        SecretARN: props.secret.secretArn,
        SopsS3File: {
          Bucket: sopsAsset.s3BucketName,
          Key: sopsAsset.s3ObjectKey,
        },
        ConvertToJSON: this.converToJSON,
        Flatten: this.flatten,
        Format: this.sopsFileFormat,
        SopsAgeKey: props.sopsAgeKey,
      },
    });
    this.versionId = cr.getAttString('VersionId');
  }
}

/**
 * The configuration options of the SopsSecret
 */
export interface SopsSecretProps extends SecretProps, SopsSyncOptions {}

/**
 * A drop in replacement for the normal Secret, that is populated with the encrypted
 * content of the given sops file.
 */
export class SopsSecret extends Secret {
  readonly sync: SopsSync;
  public constructor(scope: Construct, id: string, props: SopsSecretProps) {
    super(scope, id, props as SecretProps);
    this.sync = new SopsSync(this, 'SopsSync', {
      secret: this,
      ...(props as SopsSyncOptions),
    });
  }

  public secretValueFromJson(jsonField: string) {
    return SecretValue.secretsManager(this.secretArn, {
      jsonField,
      versionId: this.sync.versionId,
    });
  }
}
