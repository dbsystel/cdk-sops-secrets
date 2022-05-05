import { App, SecretValue, Stack } from 'aws-cdk-lib';
import { Match, Template } from 'aws-cdk-lib/assertions';
import { Key } from 'aws-cdk-lib/aws-kms';
import { Function, InlineCode, Runtime } from 'aws-cdk-lib/aws-lambda';
import { SopsSecret, SopsSyncProvider } from '../src';

const keyStatements = [
  {
    Action: 'kms:Decrypt',
    Effect: 'Allow',
    Resource:
      'arn:aws:kms:aws-region-1:123456789011:key/00000000-1234-4321-abcd-1234abcd12ab',
  },
];

test('Throw exception on non existent sops secret', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');
  expect(
    () =>
      new SopsSecret(stack, 'SopsSecret', {
        sopsFilePath: 'test-secrets/does-not-exist.json',
      }),
  ).toThrowError('File test-secrets/does-not-exist.json does not exist!');
});

test('Age Key add', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');

  const provider = new SopsSyncProvider(stack, 'Provider');
  provider.addAgeKey(SecretValue.plainText('SOME-KEY'));
  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/yaml/sopsfile.enc-kms.yaml',
  });
  Template.fromStack(stack).hasResource('AWS::Lambda::Function', {
    Properties: Match.objectLike({
      Environment: Match.objectLike({
        Variables: Match.objectLike({
          SOPS_AGE_KEY: 'SOME-KEY',
        }),
      }),
    }),
  });
});

test('KMS Key lookup from sopsfile: json', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');
  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/json/sopsfile.enc-kms.json',
  });
  Template.fromStack(stack).hasResource('AWS::IAM::Policy', {
    Properties: Match.objectLike({
      PolicyDocument: Match.objectLike({
        Statement: Match.arrayWith(keyStatements),
      }),
    }),
  });
});

test('KMS Key lookup from sopsfile: yaml', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');
  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/yaml/sopsfile.enc-kms.yaml',
  });
  Template.fromStack(stack).hasResource('AWS::IAM::Policy', {
    Properties: Match.objectLike({
      PolicyDocument: Match.objectLike({
        Statement: Match.arrayWith(keyStatements),
      }),
    }),
  });
});

test('KMS Key given via props', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');
  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/yaml/sopsfile.enc-age.yaml',
    sopsKmsKey: [
      Key.fromKeyArn(
        stack,
        'Key',
        'arn:aws:kms:aws-region-1:110123456789:key/00000000-1234-4321-abcd-1234abcd12ab',
      ),
    ],
  });
  Template.fromStack(stack).hasResource('AWS::IAM::Policy', {
    Properties: Match.objectLike({
      PolicyDocument: Match.objectLike({
        Statement: Match.arrayWith([
          {
            Action: 'kms:Decrypt',
            Effect: 'Allow',
            Resource:
              'arn:aws:kms:aws-region-1:110123456789:key/00000000-1234-4321-abcd-1234abcd12ab',
          },
        ]),
      }),
    }),
  });
});

test('Derive correct format: yml', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');
  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/yaml/sopsfile.enc-age.yml',
  });
  Template.fromStack(stack).hasResource('Custom::SopsSync', {
    Properties: Match.objectLike({
      Format: 'yaml',
    }),
  });
});

test('Derive correct format: yaml', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');
  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/yaml/sopsfile.enc-age.yaml',
  });
  Template.fromStack(stack).hasResource('Custom::SopsSync', {
    Properties: Match.objectLike({
      Format: 'yaml',
    }),
  });
});

test('Derive correct format: json', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');
  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/json/sopsfile.enc-age.json',
  });
  Template.fromStack(stack).hasResource('Custom::SopsSync', {
    Properties: Match.objectLike({
      Format: 'json',
    }),
  });
});

test('Set format: json', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');
  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/yaml/sopsfile.enc-age.yaml',
    sopsFileFormat: 'json',
  });
  Template.fromStack(stack).hasResource('Custom::SopsSync', {
    Properties: Match.objectLike({
      Format: 'json',
    }),
  });
});

test('Set format: yaml', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');
  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/json/sopsfile.enc-age.json',
    sopsFileFormat: 'yaml',
  });
  Template.fromStack(stack).hasResource('Custom::SopsSync', {
    Properties: Match.objectLike({
      Format: 'yaml',
    }),
  });
});

test('Exception when set format: notsupported', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');

  expect(
    () =>
      new SopsSecret(stack, 'SopsSecret', {
        sopsFilePath: 'test-secrets/json/sopsfile.enc-age.notsupported',
      }),
  ).toThrowError('Unsupported sopsFileFormat notsupported');
});

test('secretValueFromJson(...)', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');
  const secret = new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/json/sopsfile.enc-age.json',
  });

  new Function(stack, 'TestParameter', {
    code: InlineCode.fromInline('TEST'),
    runtime: Runtime.NODEJS_14_X,
    handler: 'test',
    environment: {
      stringValue: secret.secretValueFromJson('test').toString(),
    },
  });

  Template.fromStack(stack).hasResource('AWS::Lambda::Function', {
    Properties: Match.objectLike({
      Handler: 'test',
      Runtime: 'nodejs14.x',
      Environment: {
        Variables: {
          stringValue: {
            'Fn::Join': [
              '',
              [
                '{{resolve:secretsmanager:',
                {
                  Ref: 'SopsSecretF929FB43',
                },
                ':SecretString:test::',
                {
                  'Fn::GetAtt': ['SopsSecretSopsSync7D825417', 'VersionId'],
                },
                '}}',
              ],
            ],
          },
        },
      },
    }),
  });
});
