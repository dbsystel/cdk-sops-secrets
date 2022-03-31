import { App, Stack } from '@aws-cdk/core';
import { SopsSecrets } from '../src/index';

const app = new App();

const stack = new Stack(app, 'SecretIntegration');

new SopsSecrets(stack, 'SopsSecret', {
  sopsFilePath: 'test-secrets/sopsfile.enc-age.json',
  sopsFileFormat: 'json',
  // see test-secrets/README.md for further information regarding the test file
  sopsAgeKey: 'AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3',
});