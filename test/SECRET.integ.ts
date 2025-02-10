import * as cdk from 'aws-cdk-lib';
import { SopsSecret } from '../src/index';

const app = new cdk.App();
const stack = new cdk.Stack(app, 'SECRET');

interface TestCase {
  name: string;
  sopsFilePath: string;
  additionalProperties: Record<string, any>;
}

const tc = [
  {
    name: 'Json2Json',
    sopsFilePath: 'test-secrets/testsecret.sops.json',
    additionalProperties: {},
  },
  {
    name: 'Yaml2Json',
    sopsFilePath: 'test-secrets/testsecret.sops.yaml',
    additionalProperties: {},
  },
  {
    name: 'Json2Raw',
    sopsFilePath: 'test-secrets/testsecret.sops.json',
    additionalProperties: {
      rawOutput: true,
    },
  },
  {
    name: 'Yaml2Raw',
    sopsFilePath: 'test-secrets/testsecret.sops.yaml',
    additionalProperties: {
      rawOutput: true,
    },
  },
  {
    name: 'DotEnv2Json',
    sopsFilePath: 'test-secrets/testsecret.sops.env',
    additionalProperties: {},
  },
  {
    name: 'DotEnv2Raw',
    sopsFilePath: 'test-secrets/README.sops.binary',
    additionalProperties: {
      rawOutput: true,
    },
  },
  {
    name: 'Binary2Raw',
    sopsFilePath: 'test-secrets/README.sops.binary',
    additionalProperties: {
      rawOutput: true,
    },
  },
] satisfies TestCase[];

tc.forEach((t) => {
  new SopsSecret(stack, t.name, {
    secretName: t.name,
    sopsFilePath: t.sopsFilePath,
    sopsAgeKey: cdk.SecretValue.unsafePlainText(
      'AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3',
    ),
    ...t.additionalProperties,
  });
});

app.synth();
