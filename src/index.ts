import * as fs from 'fs';
import * as path from 'path';
import { IKey } from '@aws-cdk/aws-kms';
import { Code, IFunction, Runtime, SingletonFunction } from '@aws-cdk/aws-lambda';
import { Asset } from '@aws-cdk/aws-s3-assets';
import { ISecret, Secret, SecretProps } from '@aws-cdk/aws-secretsmanager';
import { Annotations, Construct, CustomResource } from '@aws-cdk/core';

export interface SopsSyncOptions {
  readonly sopsProvider?: IFunction;
  readonly sopsFilePath: string;
  readonly sopsFileFormat?: string;
  readonly sopsKmsKey?: IKey;
  readonly sopsAgeKey?: string;
}
export interface SopsSyncProps extends SopsSyncOptions{
  readonly secret: ISecret;
}

export class SopsSync extends Construct {
  constructor(scope: Construct, id: string, props: SopsSyncProps) {
    super(scope, id);

    const sopsFileFormat = props.sopsFileFormat ?? props.sopsFilePath.split('.').pop;

    const provider = props.sopsProvider ?? new SingletonFunction(this, 'Function', {
      code: Code.fromAsset( path.join(__dirname, '../assets/cdk-sops-lambda.zip')),
      runtime: Runtime.GO_1_X,
      handler: 'cdk-sops-secrets',
      uuid: 'cdk-sops-secrets',
    });

    if ( provider.role !== undefined ) {
      props.secret.grantWrite(provider);
      props.sopsKmsKey?.grantEncrypt(provider);
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

    new CustomResource(this, 'Resource', {
      serviceToken: provider.functionArn,
      resourceType: 'Custom::SecretSopsSync',
      properties: {
        SecretARN: props.secret.secretArn,
        S3SOPSContentFile: {
          Bucket: sopsAsset.s3BucketName,
          Key: sopsAsset.s3ObjectKey,
          Format: sopsFileFormat,
        },
        SOPSAgeKey: props.sopsAgeKey,
      },
    });
  }
}

export interface SopsSecretProps extends SecretProps, SopsSyncOptions {}
export class SopsSecrets extends Secret {
  public constructor(scope: Construct, id: string, props: SopsSecretProps) {
    super(scope, id, props as SecretProps);
    new SopsSync(this, 'SopsSync', { secret: this, ...(props as SopsSyncOptions) });
  }
}