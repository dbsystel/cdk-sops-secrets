import * as fs from 'fs';
import { IKey } from '@aws-cdk/aws-kms';
import { Code, IFunction, Runtime, SingletonFunction } from '@aws-cdk/aws-lambda';
import { Asset } from '@aws-cdk/aws-s3-assets';
import { ISecret } from '@aws-cdk/aws-secretsmanager';
import { Annotations, Construct, CustomResource } from '@aws-cdk/core';

export interface SopsSecretProps {
  readonly provider?: IFunction;
  readonly secret: ISecret;
  readonly sopsFilePath: string;
  readonly sopsFileFormat?: string;
  readonly sopsKmsKey: IKey;
  readonly sopsAgeKey?: string;
}
export class SopsSecrets extends Construct {
  public constructor(scope: Construct, id: string, props: SopsSecretProps) {
    super(scope, id);

    const sopsFileFormat = props.sopsFileFormat ?? props.sopsFilePath.split('.').pop;

    if (!fs.existsSync(props.sopsFilePath)) {
      throw new Error(`File ${props.sopsFilePath} does not exist!`);
    }

    const provider = props.provider ?? new SingletonFunction(this, 'Function', {
      code: Code.fromAsset('./assets/cdk-sops-lambda.zip'),
      runtime: Runtime.GO_1_X,
      handler: 'cdk-sops-lambda',
      uuid: 'cdk-sops-lambda',
    });

    if ( provider.role !== undefined ) {
      props.secret.grantWrite(provider);
      props.sopsKmsKey?.grantEncrypt(provider);
    } else {
      Annotations.of(this).addWarning('Please ensure propper permissions for the passed lambda function:\n  - write Access to the secret\n  - encrypt with the sopsKmsKey');
    }

    const sopsAsset = new Asset(this, 'Asset', {
      path: props.sopsFilePath,
    });

    new CustomResource(this, 'Resource', {
      serviceToken: provider.functionArn,
      resourceType: 'Custom::SOPSSecretsManager',
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