//import { Code, Function, Runtime } from '@aws-cdk/aws-lambda';
import { App, SecretValue, Stack } from '@aws-cdk/core';
import { SopsSecrets } from '../src/index';

const app = new App();

const stack = new Stack(app, 'SecretIntegration');

new SopsSecrets(stack, 'SopsSecretJSON', {
  sopsFilePath: 'test-secrets/json/sopsfile.enc-age.json',
  // see test-secrets/README.md for further information regarding the test file
  sopsAgeKey: SecretValue.plainText('AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3'),
});

new SopsSecrets(stack, 'SopsSecretYAML', {
  sopsFilePath: 'test-secrets/yaml/sopsfile.enc-age.yaml',
  convertToJSON: false,
  // see test-secrets/README.md for further information regarding the test file
  sopsAgeKey: SecretValue.plainText('AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3'),
});

new SopsSecrets(stack, 'SopsSecretYAMLasJSON', {
  sopsFilePath: 'test-secrets/yaml/sopsfile.enc-age.yaml',
  convertToJSON: true,
  // see test-secrets/README.md for further information regarding the test file
  sopsAgeKey: SecretValue.plainText('AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3'),
});

new SopsSecrets(stack, 'SopsComplexSecretJSON', {
  sopsFilePath: 'test-secrets/json/sopsfile-complex.enc-age.json',
  flatten: false,
  // see test-secrets/README.md for further information regarding the test file
  sopsAgeKey: SecretValue.plainText('AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3'),
});

new SopsSecrets(stack, 'SopsComplexSecretJSONFlat', {
  sopsFilePath: 'test-secrets/json/sopsfile-complex.enc-age.json',
  flatten: true,
  // see test-secrets/README.md for further information regarding the test file
  sopsAgeKey: SecretValue.plainText('AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3'),
});

new SopsSecrets(stack, 'SopComplexSecretYAML', {
  sopsFilePath: 'test-secrets/yaml/sopsfile-complex.enc-age.yaml',
  convertToJSON: false,
  flatten: false,
  // see test-secrets/README.md for further information regarding the test file
  sopsAgeKey: SecretValue.plainText('AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3'),
});

new SopsSecrets(stack, 'SopComplexSecretYAMLFlat', {
  sopsFilePath: 'test-secrets/yaml/sopsfile-complex.enc-age.yaml',
  convertToJSON: false,
  flatten: true,
  // see test-secrets/README.md for further information regarding the test file
  sopsAgeKey: SecretValue.plainText('AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3'),
});

new SopsSecrets(stack, 'SopsComplexSecretYAMLasJSON', {
  sopsFilePath: 'test-secrets/yaml/sopsfile-complex.enc-age.yaml',
  convertToJSON: true,
  flatten: false,
  // see test-secrets/README.md for further information regarding the test file
  sopsAgeKey: SecretValue.plainText('AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3'),
});

new SopsSecrets(stack, 'SopsComplexSecretYAMLasJSONFlat', {
  sopsFilePath: 'test-secrets/yaml/sopsfile-complex.enc-age.yaml',
  convertToJSON: true,
  flatten: true,
  // see test-secrets/README.md for further information regarding the test file
  sopsAgeKey: SecretValue.plainText('AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3'),
});
//new Function(stack, 'TestFunction', {
//  code: Code.fromInline('test'),
//  handler: 'test',
//  runtime: Runtime.NODEJS_14_X,
//  environment: {
//    jsonKey1: secretValueJSON.secretValueFromJson('key1').toString(),
//    yamlKey1: secretValueYAML.secretValueFromJson('key1').toString(),
//  },
//});