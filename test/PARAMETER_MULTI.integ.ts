import * as cdk from 'aws-cdk-lib';
import { Key } from 'aws-cdk-lib/aws-kms';
import { ConstantAssetHashAspect } from './ConstantAssetHashAspect';
import { MultiStringParameter } from '../src/index';

const app = new cdk.App();
const stack = new cdk.Stack(app, 'PARAMETERMULTI');

// Apply aspect for constant asset IDs
cdk.Aspects.of(stack).add(new ConstantAssetHashAspect());

interface TestCase {
  name: string;
  sopsFilePath: string;
}

const tc = [
  {
    name: 'Json',
    sopsFilePath: 'test-secrets/testsecret.sops.json',
  },
  {
    name: 'Yaml',
    sopsFilePath: 'test-secrets/testsecret.sops.yaml',
  },
  // { NotSupported right now
  //   name: 'DotEnv',
  //   sopsFilePath: 'test-secrets/testsecret.sops.env',
  // },
] satisfies TestCase[];

const encryptionKey = Key.fromKeyArn(
  stack,
  'Key',
  'arn:aws:kms:eu-central-1:505755377845:key/58b75958-883a-440c-842c-85af5d33a5bb',
);

tc.forEach((t) => {
  new MultiStringParameter(stack, t.name, {
    keyPrefix: `/${t.name}/`,
    keySeparator: '/',
    sopsFilePath: t.sopsFilePath,
    sopsAgeKey: cdk.SecretValue.unsafePlainText(
      'AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3',
    ),
    encryptionKey,
  });
});

app.synth();
