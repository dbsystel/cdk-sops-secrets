import * as cdk from 'aws-cdk-lib';
import { ConstantAssetHashAspect } from './ConstantAssetHashAspect';
import { RawOutput, SopsSecret } from '../src/index';

const app = new cdk.App();
const stack = new cdk.Stack(app, 'SECRET');

// Apply aspect for constant asset IDs
cdk.Aspects.of(stack).add(new ConstantAssetHashAspect());

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
    name: 'Json2RawString',
    sopsFilePath: 'test-secrets/testsecret.sops.json',
    additionalProperties: {
      rawOutput: RawOutput.STRING,
    },
  },
  {
    name: 'Yaml2Json',
    sopsFilePath: 'test-secrets/testsecret.sops.yaml',
    additionalProperties: {},
  },
  {
    name: 'Yaml2RawString',
    sopsFilePath: 'test-secrets/testsecret.sops.yaml',
    additionalProperties: {
      rawOutput: RawOutput.STRING,
    },
  },
  {
    name: 'DotEnv2Json',
    sopsFilePath: 'test-secrets/testsecret.sops.env',
    additionalProperties: {},
  },
  {
    name: 'DotEnv2RawString',
    sopsFilePath: 'test-secrets/README.sops.binary',
    additionalProperties: {
      rawOutput: RawOutput.STRING,
    },
  },
  {
    name: 'Binary2RawString',
    sopsFilePath: 'test-secrets/README.sops.binary',
    additionalProperties: {
      rawOutput: RawOutput.STRING,
    },
  },
  {
    name: 'Binary2RawBinary',
    sopsFilePath: 'test-secrets/README.gz.sops.binary',
    additionalProperties: {
      rawOutput: RawOutput.BINARY,
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
