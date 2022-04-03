import { App, Stack } from '@aws-cdk/core';
import { SopsSecrets } from '../src/index';
import { Code, Function, Runtime } from '@aws-cdk/aws-lambda';

const app = new App();

const stack = new Stack(app, 'SecretIntegration');

const secretValue = new SopsSecrets(stack, 'SopsSecret', {
  sopsFilePath: 'test-secrets/sopsfile.enc-age.json',
  // see test-secrets/README.md for further information regarding the test file
  sopsAgeKey: 'AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3',
});

new Function(stack, 'TestFunction', {
  code: Code.fromInline('test'),
  handler: 'test',
  runtime: Runtime.NODEJS_14_X,
  environment: {
    SecretValue: secretValue.secretValueFromJson('key1').toString(),
  }
})