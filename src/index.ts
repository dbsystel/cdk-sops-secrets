import * as fs from 'fs';
import { IKey, Key } from '@aws-cdk/aws-kms';
import { Code, IFunction, Runtime, SingletonFunction } from '@aws-cdk/aws-lambda';
import { Asset } from '@aws-cdk/aws-s3-assets';
import { ISecret, Secret, SecretProps } from '@aws-cdk/aws-secretsmanager';
import { Annotations, Construct, CustomResource, SecretValue } from '@aws-cdk/core';

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
   *
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
export interface SopsSyncProps extends SopsSyncOptions{
  readonly secret: ISecret;
}

export class SopsSync extends Construct {

  readonly versionId: string;
  readonly sopsFileFormat: 'json' | 'yaml';
  readonly converToJSON: boolean;
  readonly flatten: boolean;

  constructor(scope: Construct, id: string, props: SopsSyncProps) {
    super(scope, id);

    this.converToJSON = props.convertToJSON ?? true;
    this.flatten = props.flatten ?? true;

    const sopsFileFormat = props.sopsFileFormat ?? props.sopsFilePath.split('.').pop();
    switch ( sopsFileFormat ) {
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
    const provider = props.sopsProvider ?? new SingletonFunction(this, 'Function', {
      code,
      runtime: Runtime.GO_1_X,
      handler: 'cdk-sops-secrets',
      uuid: 'cdk-sops-secrets',
    });

    if ( provider.role !== undefined ) {
      if ( props.sopsKmsKey !== undefined ) {
        props.sopsKmsKey.forEach( (key) => key.grantDecrypt(provider.role!) );
      }
      const fileContent = fs.readFileSync(props.sopsFilePath);
      const regex = /arn:aws:kms:[a-z0-9-]+:[\d]+:key\/[a-z0-9-]+/gm;
      const results = regex.exec(fileContent.toString());
      if ( results !== undefined ) {
        results?.forEach( (result, index) => Key.fromKeyArn(this, `SopsKey${index}`, result).grantEncrypt(provider.role!));
      }
      props.secret.grantWrite(provider);
    } else {
      Annotations.of(this).addWarning('Please ensure propper permissions for the passed lambda function:\n  - write Access to the secret\n  - encrypt with the sopsKmsKey');
    }

    if (!fs.existsSync(props.sopsFilePath)) {
      throw new Error(`File ${props.sopsFilePath} does not exist!`);
    }

    const sopsAsset = new Asset(this, 'Asset', {
      path: props.sopsFilePath,
    });

    sopsAsset.bucket.grantRead(provider);

    const cr = new CustomResource(this, 'Resource', {
      serviceToken: provider.functionArn,
      resourceType: 'Custom::SecretSopsSync',
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
  };
}

export interface SopsSecretProps extends SecretProps, SopsSyncOptions {}
export class SopsSecrets extends Secret {
  readonly sync: SopsSync;
  public constructor(scope: Construct, id: string, props: SopsSecretProps) {
    super(scope, id, props as SecretProps);
    this.sync = new SopsSync(this, 'SopsSync', { secret: this, ...(props as SopsSyncOptions) });
  }

  public secretValueFromJson(jsonField: string) {
    return SecretValue.secretsManager(this.secretArn, { jsonField, versionId: this.sync.versionId });
  }
}