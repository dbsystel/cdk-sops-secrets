import { App, Stack } from 'aws-cdk-lib';
import { SopsSecret } from '../src/index';

const app = new App();

const stack = new Stack(app, 'SecretIntegrationInline');

new SopsSecret(stack, 'SopsSecret', {
  sopsFilePath: 'test-secrets/yaml/sopsfile.enc-multikms.yaml',
  // see test-secrets/README.md for further information regarding the test file
});
