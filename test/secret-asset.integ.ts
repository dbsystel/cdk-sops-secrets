import { App, SecretValue, Stack } from 'aws-cdk-lib';
import { Code, Function, Runtime } from 'aws-cdk-lib/aws-lambda';
import { SopsSecret, UploadType } from '../src/index';

const app = new App();

const stack = new Stack(app, 'SecretIntegrationAsset');

new SopsSecret(stack, 'SopsSecretJSON', {
  sopsFilePath: 'test-secrets/json/sopsfile.enc-age.json',
  uploadType: UploadType.ASSET,
  // see test-secrets/README.md for further information regarding the test file
  sopsAgeKey: SecretValue.plainText(
    'AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3',
  ),
});

new SopsSecret(stack, 'SopsSecretYAML', {
  sopsFilePath: 'test-secrets/yaml/sopsfile.enc-age.yaml',
  convertToJSON: false,
  uploadType: UploadType.ASSET,
  // see test-secrets/README.md for further information regarding the test file
  sopsAgeKey: SecretValue.plainText(
    'AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3',
  ),
});

new SopsSecret(stack, 'SopsSecretYAMLasJSON', {
  sopsFilePath: 'test-secrets/yaml/sopsfile.enc-age.yaml',
  convertToJSON: true,
  uploadType: UploadType.ASSET,
  // see test-secrets/README.md for further information regarding the test file
  sopsAgeKey: SecretValue.plainText(
    'AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3',
  ),
});

new SopsSecret(stack, 'SopsComplexSecretJSON', {
  sopsFilePath: 'test-secrets/json/sopsfile-complex.enc-age.json',
  flatten: false,
  uploadType: UploadType.ASSET,
  // see test-secrets/README.md for further information regarding the test file
  sopsAgeKey: SecretValue.plainText(
    'AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3',
  ),
});

const sopsComplexSecretJSONFlat = new SopsSecret(
  stack,
  'SopsComplexSecretJSONFlat',
  {
    sopsFilePath: 'test-secrets/json/sopsfile-complex.enc-age.json',
    uploadType: UploadType.ASSET,
    flatten: true,
    // see test-secrets/README.md for further information regarding the test file
    sopsAgeKey: SecretValue.plainText(
      'AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3',
    ),
  },
);

new SopsSecret(stack, 'SopComplexSecretYAML', {
  sopsFilePath: 'test-secrets/yaml/sopsfile-complex.enc-age.yaml',
  convertToJSON: false,
  uploadType: UploadType.ASSET,
  flatten: false,
  // see test-secrets/README.md for further information regarding the test file
  sopsAgeKey: SecretValue.plainText(
    'AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3',
  ),
});

new SopsSecret(stack, 'SopComplexSecretYAMLFlat', {
  sopsFilePath: 'test-secrets/yaml/sopsfile-complex.enc-age.yaml',
  convertToJSON: false,
  uploadType: UploadType.ASSET,
  flatten: true,
  // see test-secrets/README.md for further information regarding the test file
  sopsAgeKey: SecretValue.plainText(
    'AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3',
  ),
});

new SopsSecret(stack, 'SopsComplexSecretYAMLasJSON', {
  sopsFilePath: 'test-secrets/yaml/sopsfile-complex.enc-age.yaml',
  convertToJSON: true,
  uploadType: UploadType.ASSET,
  flatten: false,
  // see test-secrets/README.md for further information regarding the test file
  sopsAgeKey: SecretValue.plainText(
    'AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3',
  ),
});

const sopsComplexSecretYAMLasJSONFlat = new SopsSecret(
  stack,
  'SopsComplexSecretYAMLasJSONFlat',
  {
    sopsFilePath: 'test-secrets/yaml/sopsfile-complex.enc-age.yaml',
    uploadType: UploadType.ASSET,
    convertToJSON: true,
    flatten: true,
    // see test-secrets/README.md for further information regarding the test file
    sopsAgeKey: SecretValue.plainText(
      'AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3',
    ),
  },
);

new SopsSecret(stack, 'SopsBinaryAsBinary', {
  sopsFilePath: 'test-secrets/binary/sopsfile.enc-age.binary',
  uploadType: UploadType.ASSET,
  // see test-secrets/README.md for further information regarding the test file
  sopsAgeKey: SecretValue.plainText(
    'AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3',
  ),
});

new Function(stack, 'TestFunction', {
  code: Code.fromInline('test'),
  handler: 'test',
  runtime: Runtime.NODEJS_14_X,
  environment: {
    ...createMapComplex('SopsComplexSecretJSONFlat', sopsComplexSecretJSONFlat),
    ...createMapComplex(
      'sopsComplexSecretYAMLasJSONFlat',
      sopsComplexSecretYAMLasJSONFlat,
    ),
  },
});

function createMapComplex(prefix: string, secret: SopsSecret) {
  return {
    [`${prefix}_and_now_some_0_basic`]: secret
      .secretValueFromJson('and now.some[0].basic')
      .toString(),
    [`${prefix}_and_now_some_1_nested`]: secret
      .secretValueFromJson('and now.some[1].nested')
      .toString(),
    [`${prefix}_and_now_some_2_type`]: secret
      .secretValueFromJson('and now.some[2].type')
      .toString(),
    [`${prefix}_and_now_some_3_tests`]: secret
      .secretValueFromJson('and now.some[3].tests')
      .toString(),
    [`${prefix}_some_deep_nested_arrays_0`]: secret
      .secretValueFromJson('some.deep.nested.arrays[0]')
      .toString(),
    [`${prefix}_some_deep_nested_arrays_1`]: secret
      .secretValueFromJson('some.deep.nested.arrays[1]')
      .toString(),
    [`${prefix}_some_deep_nested_arrays_2_values_and`]: secret
      .secretValueFromJson('some.deep.nested.arrays[2].values.and')
      .toString(),
    [`${prefix}_some_deep_nested_object`]: secret
      .secretValueFromJson('some.deep.nested.object')
      .toString(),
    [`${prefix}_some_notsodeep`]: secret
      .secretValueFromJson('some.notsodeep')
      .toString(),
  };
}
