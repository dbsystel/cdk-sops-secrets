import { App, SecretValue, Stack } from 'aws-cdk-lib';
import { Code, Function, Runtime } from 'aws-cdk-lib/aws-lambda';
import { SopsSecret, UploadType } from '../src/index';
import { Key } from 'aws-cdk-lib/aws-kms';

const app = new App();

const stack = new Stack(app, 'SecretIntegrationAsset');

new SopsSecret(stack, 'SopsSecretOwnKmsMey', {
  sopsFilePath: 'test-secrets/json/sopsfile.enc-age.json',
  uploadType: UploadType.ASSET,
  // see test-secrets/README.md for further information regarding the test file
  sopsAgeKey: SecretValue.plainText(
    'AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3',
  ),
  encryptionKey: new Key(stack, 'CustomKey'),
});
