/**
 * Caution - this test can only be used for snapshots - you can't deploy this
 */
import { App, SecretValue, Stack } from 'aws-cdk-lib';
import { Key } from 'aws-cdk-lib/aws-kms';
import { SopsSecret, UploadType } from '../src/index';

const app = new App();

const stack = new Stack(app, 'SecretMultiKms');

new SopsSecret(stack, 'SopsSecretOwnKmsMey', {
  sopsFilePath: 'test-secrets/json/sopsfile.enc-age.json',
  uploadType: UploadType.ASSET,
  // see test-secrets/README.md for further information regarding the test file
  sopsAgeKey: SecretValue.plainText(
    'AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3',
  ),
  encryptionKey: new Key(stack, 'CustomKey'),
});

new SopsSecret(stack, 'SopsSecretForeignKmsMey', {
  sopsFilePath: 'test-secrets/json/sopsfile.enc-age.json',
  uploadType: UploadType.ASSET,
  // see test-secrets/README.md for further information regarding the test file
  sopsAgeKey: SecretValue.plainText(
    'AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3',
  ),
  encryptionKey: Key.fromKeyArn(
    stack,
    'ForeignKey',
    'arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012',
  ),
});
