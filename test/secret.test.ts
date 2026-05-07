import { App, RemovalPolicy, SecretValue, Stack } from 'aws-cdk-lib';
import { Match, Template } from 'aws-cdk-lib/assertions';
import {
  AnyPrincipal,
  Effect,
  PolicyStatement,
  Role,
  ServicePrincipal,
} from 'aws-cdk-lib/aws-iam';
import { Key } from 'aws-cdk-lib/aws-kms';
import { Function, InlineCode, Runtime } from 'aws-cdk-lib/aws-lambda';
import { LogGroup, RetentionDays } from 'aws-cdk-lib/aws-logs';
import { Topic } from 'aws-cdk-lib/aws-sns';
import { EmailSubscription } from 'aws-cdk-lib/aws-sns-subscriptions';
import { StringParameter } from 'aws-cdk-lib/aws-ssm';
import {
  SopsSecret,
  SopsSyncProvider,
  UploadType,
  MultiStringParameter,
  RawOutput,
} from '../src';

const keyStatements = [
  {
    Action: 'kms:Decrypt',
    Effect: 'Allow',
    Resource:
      'arn:aws:kms:aws-region-1:123456789011:key/00000000-1234-4321-abcd-1234abcd12ab',
  },
];

test('Upload type ASSET', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');

  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/yaml/sopsfile.enc-kms.yaml',
    uploadType: UploadType.ASSET,
    rawOutput: RawOutput.STRING,
  });
  Template.fromStack(stack).hasResource('Custom::SopsSync', {
    Properties: Match.objectLike({
      SopsS3File: {
        Bucket: {
          'Fn::Sub': 'cdk-hnb659fds-assets-${AWS::AccountId}-${AWS::Region}',
        },
        Key: '6d7a109a504c02a3455d09067a3695a4355cf3c2c914b6f3e949e20d6a741128.yaml',
      },
    }),
  });
});

test('Upload type INLINE', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');

  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/yaml/sopsfile.enc-kms.yaml',
    uploadType: UploadType.INLINE,
  });
  Template.fromStack(stack).hasResource('Custom::SopsSync', {
    Properties: Match.objectLike({
      SopsInline: {
        Content:
          'aGVsbG86IEVOQ1tBRVMyNTZfR0NNLGRhdGE6c3N3YlJIY2NJdWlzYkd5TyttRy9FV2IrMUpjdHFOVW9yT3pRVWFNUm4zSGtBeVd0cmVOdVp5aFlZbGt2Y2c9PSxpdjpYcEhpbGh5Q1JYYThjbEdpRWNrcHF1Q3FTWkdXQTgwSDBlRFVtUjgweGhFPSx0YWc6V3N1eEk0bFVId0pFYXM3Nk9EUXhVQT09LHR5cGU6c3RyXQpleGFtcGxlX2tleTogRU5DW0FFUzI1Nl9HQ00sZGF0YTovNE90aUgzbm9vVFZadXpDK2c9PSxpdjpSNnFyTlk3R0F4Y0l3UjQyYnZFMjc3NHVmK0lmUTlwdjlVQTc4TG1pdFUwPSx0YWc6akZvZ1ZzeWN1ZlNYUnZCS1hMVkE5Zz09LHR5cGU6c3RyXQojRU5DW0FFUzI1Nl9HQ00sZGF0YTphRmpIbCtLSzM3T0lzRmtoQnFiV0FhMD0saXY6WlNRUDY1UWg2aXNnUUw2YVFxazlRNGJQQ3hxZC9rMHcrOXF6bmF1bTFQQT0sdGFnOjdDVXVxbmRHNGIzSjN2L2c1T0tPMlE9PSx0eXBlOmNvbW1lbnRdCmV4YW1wbGVfYXJyYXk6CiAgICAtIEVOQ1tBRVMyNTZfR0NNLGRhdGE6VXhXZnNCYUdjZ3JlZ21EZktNST0saXY6WEJzMk9Db0hXT21vYmoxT1BkL2NsYlVKRmJQVGRBZG90RkhWR1FtMVRQMD0sdGFnOkNUVDE1WllkMW5HOGtKYXJERExaTEE9PSx0eXBlOnN0cl0KICAgIC0gRU5DW0FFUzI1Nl9HQ00sZGF0YTpRUDhaS2E4OG43U25KYUsvWkxrPSxpdjpRN2FCc1dySXYwRzhEMzNJdzY1S1JxT1J3cXlVV2phQXZJQlNiV05IQVd3PSx0YWc6d0JCN3RSY2M2eVJ2NWNBNTNrNWpvZz09LHR5cGU6c3RyXQpleGFtcGxlX251bWJlcjogRU5DW0FFUzI1Nl9HQ00sZGF0YTo0MVF2Sm9DSjVNdUVMUT09LGl2OlFleEpNK3J3aHVaY3hETDNUREVzd3BrOHF3VWVVNmJHazVpNVRDZHVpRWs9LHRhZzpyenlrTVlybUh5djNVSGpFR1JyTkhnPT0sdHlwZTpmbG9hdF0KZXhhbXBsZV9ib29sZWFuczoKICAgIC0gRU5DW0FFUzI1Nl9HQ00sZGF0YToyc0ZoRGc9PSxpdjpoQk1aMHgvMVJKQy8yTmVRL0JFdWdQRFZDQ29udlZIQ1prMVRhazl2bG44PSx0YWc6Q0pxeXg3eUF0Tm5GQUR2RGFrZlhiQT09LHR5cGU6Ym9vbF0KICAgIC0gRU5DW0FFUzI1Nl9HQ00sZGF0YTo1MVQzaDlBPSxpdjowVnV3TUp3VVJUOEF5WkVBcHRRQzVrZDVrK2RmWkRIUjU4TXI0WjRUVklJPSx0YWc6bVRLMXZRRTllWGxoQWtMUW9namRmdz09LHR5cGU6Ym9vbF0Kc29wczoKICAgIGttczoKICAgICAgICAtIGFybjogYXJuOmF3czprbXM6YXdzLXJlZ2lvbi0xOjEyMzQ1Njc4OTAxMTprZXkvMDAwMDAwMDAtMTIzNC00MzIxLWFiY2QtMTIzNGFiY2QxMmFiCiAgICAgICAgICBjcmVhdGVkX2F0OiAiMjAyMi0wNC0wM1QxNzozNDo0NVoiCiAgICAgICAgICBlbmM6IEFRSUNBSGlTZ1pvTFA2ZkRyVUJZWVBVMm9KT0IvM3FGQVI1bUVZdVpZMkRRcXpZckJnR0FQUytTaU81eWIvYlhiVWRvVVBlWkFBQUFmakI4QmdrcWhraUc5dzBCQndhZ2J6QnRBZ0VBTUdnR0NTcUdTSWIzRFFFSEFUQWVCZ2xnaGtnQlpRTUVBUzR3RVFRTVhaUHBQVTNHaWJJT05LNlZBZ0VRZ0R1MTU0WXBuWW9lMmY4WUZ1V2VCcEdYZkRkYXVkNW9NRGZxdXdxWTJVV0c4Y2xuWlY5MzU1eitWN2NxQ2krNFBFQm9hdmVNTExjTFlzTE9BQT09CiAgICAgICAgICBhd3NfcHJvZmlsZTogIiIKICAgIGdjcF9rbXM6IFtdCiAgICBhenVyZV9rdjogW10KICAgIGhjX3ZhdWx0OiBbXQogICAgYWdlOiBbXQogICAgbGFzdG1vZGlmaWVkOiAiMjAyMi0wNC0wM1QxNzozNDo1NloiCiAgICBtYWM6IEVOQ1tBRVMyNTZfR0NNLGRhdGE6cURpL0NhQUdmWGJjQ3hQWHhIUlZuRGlQM2F5TGhMNStPcmR2K3JDeUYvTTBpQmxPNlBFQmV6NG1XKzBKRTRGOEU4YTNsTUpPUmhhRUdUbVJ0R3l6UFhtNjJEb05McFZlcHFqOGZkQk5nenhEUWRiVFgyeWZQb1NsdFNteUNwN2xwbmJVakd3U0hWSHNHSzgrMUVFb29jVG44MWVmUDBoeXRSU29jWUVLT05ZPSxpdjo4dlo0MTdOaHA2Qk8ya0U1SWM0SFJrektqUzc5ZW5yZlZKbHNZZXE1OUZvPSx0YWc6bXBFK090bytIUExkMnRzV3V0ZXVTQT09LHR5cGU6c3RyXQogICAgcGdwOiBbXQogICAgdW5lbmNyeXB0ZWRfc3VmZml4OiBfdW5lbmNyeXB0ZWQKICAgIHZlcnNpb246IDMuNy4yCg==',
        Hash: '6d7a109a504c02a3455d09067a3695a4355cf3c2c914b6f3e949e20d6a741128',
      },
    }),
  });
});

test('Throw exception on non existent sops secret', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');
  expect(
    () =>
      new SopsSecret(stack, 'SopsSecret', {
        sopsFilePath: 'test-secrets/does-not-exist.json',
      }),
  ).toThrow('File test-secrets/does-not-exist.json does not exist!');
});

test('Age Key passed', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');

  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/yaml/sopsfile.enc-kms.yaml',
    sopsAgeKey: SecretValue.unsafePlainText('SOME-KEY'),
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

test('Age Key add', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');

  const provider = new SopsSyncProvider(stack, 'Provider');
  provider.addAgeKey(SecretValue.unsafePlainText('SOME-KEY'));
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

test('Exception when derive format: notsupported', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');

  expect(
    () =>
      new SopsSecret(stack, 'SopsSecret', {
        sopsFilePath: 'test-secrets/testsecret.notsupported',
      }),
  ).toThrow('You have to specify sopsFileFormat!');
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

test('Correct Lambda Policy if KMS key for secrets encryption is an IKey ', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');
  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/json/sopsfile.enc-age.json',
    sopsFileFormat: 'yaml',
    encryptionKey: Key.fromKeyArn(
      stack,
      'EncryptionKey',
      'arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab',
    ),
  });
  Template.fromStack(stack).hasResource('AWS::IAM::Policy', {
    Properties: Match.objectLike({
      PolicyDocument: Match.objectLike({
        Statement: Match.arrayWith([
          {
            Action: [
              'kms:Decrypt',
              'kms:Encrypt',
              'kms:ReEncrypt*',
              'kms:GenerateDataKey*',
            ],
            Effect: 'Allow',
            Resource:
              'arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab',
          },
        ]),
      }),
    }),
  });
});

test('Methods of SopsSync implemented', () => {
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
      secretValueFromJson: secret.secretValueFromJson('test').toString(),
      currentVersionId: secret.currentVersionId().toString(),
      secretValue: secret.secretValue.toString(),
    },
  });

  secret.addToResourcePolicy(
    new PolicyStatement({
      actions: ['*'],
      effect: Effect.ALLOW,
      principals: [new AnyPrincipal()],
      resources: ['*'],
    }),
  );

  secret.denyAccountRootDelete();

  const testRole = new Role(stack, 'TestRole', {
    roleName: 'GrantReadRole',
    assumedBy: new ServicePrincipal('testservice'),
  });

  secret.grantRead(testRole);

  secret.applyRemovalPolicy(RemovalPolicy.DESTROY);

  Template.fromStack(stack).hasResource('AWS::SecretsManager::Secret', {
    UpdateReplacePolicy: 'Delete',
    DeletionPolicy: 'Delete',
  });

  Template.fromStack(stack).hasResource('AWS::Lambda::Function', {
    Properties: Match.objectLike({
      Handler: 'test',
      Runtime: 'nodejs14.x',
      Environment: {
        Variables: {
          secretValueFromJson: {
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
          currentVersionId: {
            'Fn::GetAtt': ['SopsSecretSopsSync7D825417', 'VersionId'],
          },
          secretValue: {
            'Fn::Join': [
              '',
              [
                '{{resolve:secretsmanager:',
                {
                  Ref: 'SopsSecretF929FB43',
                },
                ':SecretString:::',
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

  Template.fromStack(stack).hasResource('AWS::SecretsManager::ResourcePolicy', {
    Properties: {
      ResourcePolicy: {
        Statement: [
          // addToResourcePolicy
          {
            Action: '*',
            Effect: 'Allow',
            Principal: {
              AWS: '*',
            },
            Resource: '*',
          },
          // denyAccountRootDelete
          {
            Action: 'secretsmanager:DeleteSecret',
            Effect: 'Deny',
            Principal: {
              AWS: {
                'Fn::Join': [
                  '',
                  [
                    'arn:',
                    {
                      Ref: 'AWS::Partition',
                    },
                    ':iam::',
                    {
                      Ref: 'AWS::AccountId',
                    },
                    ':root',
                  ],
                ],
              },
            },
            Resource: '*',
          },
        ],
        Version: '2012-10-17',
      },
      SecretId: {
        Ref: 'SopsSecretF929FB43',
      },
    },
  });

  Template.fromStack(stack).hasResource('AWS::IAM::Role', {
    Properties: {
      RoleName: 'GrantReadRole',
      AssumeRolePolicyDocument: {
        Statement: [
          {
            Action: 'sts:AssumeRole',
            Effect: 'Allow',
            Principal: {
              Service: 'testservice.amazonaws.com',
            },
          },
        ],
      },
    },
  });

  Template.fromStack(stack).hasResource('AWS::IAM::Policy', {
    Type: 'AWS::IAM::Policy',
    Properties: {
      PolicyDocument: {
        Statement: [
          {
            Action: [
              'secretsmanager:GetSecretValue',
              'secretsmanager:DescribeSecret',
            ],
            Effect: 'Allow',
            Resource: {
              Ref: 'SopsSecretF929FB43',
            },
          },
        ],
        Version: '2012-10-17',
      },
      Roles: [
        {
          Ref: 'TestRole6C9272DF',
        },
      ],
    },
  });
});

test('Methods of SopsSync not implemented', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');
  const secret = new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/json/sopsfile.enc-age.json',
  });

  expect(() => secret.addRotationSchedule('something', {})).toThrow(
    `Method addRotationSchedule('something', {}) not allowed as this secret is managed by SopsSync`,
  );
  expect(() =>
    secret.grantWrite(
      Role.fromRoleArn(
        stack,
        'Role',
        'arn:aws:iam::123456789012:role/SecretAccess',
      ),
    ),
  ).toThrow(
    `Method grantWrite(...) not allowed as this secret is managed by SopsSync`,
  );
});

test('Allowed options for SopsSync', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');
  expect(
    () =>
      new SopsSecret(stack, 'SopsSecret1', {
        sopsFilePath: 'test-secrets/json/sopsfile.enc-age.json',
        sopsS3Key: 'test',
        sopsS3Bucket: 'test',
      }),
  ).toThrow(
    'You can either specify sopsFilePath or sopsS3Bucket and sopsS3Key!',
  );
  expect(
    () =>
      new SopsSecret(stack, 'SopsSecret2', {
        sopsS3Key: 'test',
      }),
  ).toThrow(
    'You can either specify sopsFilePath or sopsS3Bucket and sopsS3Key!',
  );
  expect(
    () =>
      new SopsSecret(stack, 'SopsSecret3', {
        sopsS3Bucket: 'test',
      }),
  ).toThrow(
    'You can either specify sopsFilePath or sopsS3Bucket and sopsS3Key!',
  );
  expect(
    () =>
      new SopsSecret(stack, 'SopsSecret4', {
        sopsS3Key: 'test',
        sopsS3Bucket: 'test',
      }),
  ).toThrow('You have to specify sopsFileFormat!');
});

test('Multiple parameters from yaml file', () => {
  const app = new App();
  const stack = new Stack(app, 'ParameterIntegration');
  new MultiStringParameter(stack, 'SopsSecret1', {
    sopsFilePath: 'test-secrets/yaml/sopsfile-complex-parameters.enc-age.yaml',
    encryptionKey: Key.fromKeyArn(
      stack,
      'Key',
      'arn:aws:kms:eu-central-1:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab',
    ),
  });
  const template = Template.fromStack(stack);

  template.hasResourceProperties('AWS::SSM::Parameter', {
    Name: '/foo/bar/key1',
  });

  template.hasResourceProperties('AWS::SSM::Parameter', {
    Name: '/foo/bar/key2',
  });

  template.hasResourceProperties('AWS::IAM::Policy', {
    PolicyDocument: {
      Statement: [
        {
          Action: [
            'kms:Decrypt',
            'kms:Encrypt',
            'kms:ReEncrypt*',
            'kms:GenerateDataKey*',
          ],
          Effect: 'Allow',
          Resource:
            'arn:aws:kms:eu-central-1:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab',
        },
      ],
    },
  });

  template.hasResourceProperties('AWS::IAM::ManagedPolicy', {
    PolicyDocument: {
      Statement: [
        {
          Action: 'ssm:PutParameter',
          Effect: 'Allow',
          Resource: [
            {
              'Fn::Join': [
                '',
                [
                  'arn:aws:ssm:',
                  { Ref: 'AWS::Region' },
                  ':',
                  { Ref: 'AWS::AccountId' },
                  ':parameter/foo/bar/key1',
                ],
              ],
            },
            {
              'Fn::Join': [
                '',
                [
                  'arn:aws:ssm:',
                  { Ref: 'AWS::Region' },
                  ':',
                  { Ref: 'AWS::AccountId' },
                  ':parameter/foo/bar/key2',
                ],
              ],
            },
          ],
        },
      ],
    },
  });
});

test('Multiple parameters from yaml file with custom key structure', () => {
  const app = new App();
  const stack = new Stack(app, 'ParameterIntegration');
  new MultiStringParameter(stack, 'SopsSecret1', {
    sopsFilePath: 'test-secrets/yaml/sopsfile-complex-parameters.enc-age.yaml',
    keyPrefix: '_',
    keySeparator: '.',
    encryptionKey: Key.fromKeyArn(
      stack,
      'Key',
      'arn:aws:kms:eu-central-1:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab',
    ),
  });
  const template = Template.fromStack(stack);

  template.hasResourceProperties('AWS::SSM::Parameter', {
    Name: '_foo.bar.key1',
  });

  template.hasResourceProperties('AWS::SSM::Parameter', {
    Name: '_foo.bar.key2',
  });

  template.hasResourceProperties('AWS::IAM::Policy', {
    PolicyDocument: {
      Statement: [
        {
          Action: [
            'kms:Decrypt',
            'kms:Encrypt',
            'kms:ReEncrypt*',
            'kms:GenerateDataKey*',
          ],
          Effect: 'Allow',
          Resource:
            'arn:aws:kms:eu-central-1:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab',
        },
      ],
    },
  });

  template.hasResourceProperties('AWS::IAM::ManagedPolicy', {
    PolicyDocument: {
      Statement: [
        {
          Action: 'ssm:PutParameter',
          Effect: 'Allow',
          Resource: [
            {
              'Fn::Join': [
                '',
                [
                  'arn:aws:ssm:',
                  { Ref: 'AWS::Region' },
                  ':',
                  { Ref: 'AWS::AccountId' },
                  ':parameter/_foo.bar.key1',
                ],
              ],
            },
            {
              'Fn::Join': [
                '',
                [
                  'arn:aws:ssm:',
                  { Ref: 'AWS::Region' },
                  ':',
                  { Ref: 'AWS::AccountId' },
                  ':parameter/_foo.bar.key2',
                ],
              ],
            },
          ],
        },
      ],
    },
  });
});

test('Large set of parameters to split in multiple policies', () => {
  const app = new App();
  const stack = new Stack(app, 'ParameterIntegration');
  new MultiStringParameter(stack, 'SopsSecret1', {
    sopsFilePath: 'test-secrets/yaml/sopsfile-parameters-large.yaml',
    keyPrefix: '_',
    keySeparator: '.',
    encryptionKey: Key.fromKeyArn(
      stack,
      'Key',
      'arn:aws:kms:eu-central-1:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab',
    ),
  });
  const template = Template.fromStack(stack);

  template.hasResourceProperties('AWS::SSM::Parameter', {
    Name: '_DefineSomeExtraExtraLongNameForParameterWhichHaveAExtraExtraLongNameToTestPolicySplitting1',
  });

  template.resourceCountIs('AWS::IAM::ManagedPolicy', 3);
});

test('Legacy logRetention', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');

  new SopsSyncProvider(stack, 'Provider', {
    logRetention: RetentionDays.THREE_MONTHS,
  });
  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/yaml/sopsfile.enc-kms.yaml',
    uploadType: UploadType.ASSET,
    rawOutput: RawOutput.STRING,
  });
  Template.fromStack(stack).hasResourceProperties('Custom::LogRetention', {
    RetentionInDays: 90,
  });
});

test('Custom LogGroup', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');

  const logGroup = new LogGroup(stack, 'LogGroup', {
    retention: RetentionDays.THREE_MONTHS,
  });

  new SopsSyncProvider(stack, 'Provider', {
    logGroup: logGroup,
  });
  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/yaml/sopsfile.enc-kms.yaml',
    uploadType: UploadType.ASSET,
    rawOutput: RawOutput.STRING,
  });
  Template.fromStack(stack).hasResourceProperties('AWS::Logs::LogGroup', {
    RetentionInDays: 90,
  });
});

test('Age Key from SSM Parameter string', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');
  const cmk = Key.fromKeyArn(
    stack,
    'SsmCmk',
    'arn:aws:kms:us-east-1:111122223333:key/aaaabbbb-1234-5678-abcd-111122223333',
  );

  const provider = new SopsSyncProvider(stack, 'Provider');
  provider.addAgeKeyFromSsmParameter('/sops/age/private-key', cmk);
  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/yaml/sopsfile.enc-kms.yaml',
    sopsProvider: provider,
  });

  const template = Template.fromStack(stack);

  // SOPS_AGE_KEY_PARAMS env var must be set on the Lambda
  template.hasResourceProperties('AWS::Lambda::Function', {
    Environment: Match.objectLike({
      Variables: Match.objectLike({
        SOPS_AGE_KEY_PARAMS: '/sops/age/private-key',
      }),
    }),
  });

  // Lambda role must have ssm:GetParameter for the parameter
  template.hasResource('AWS::IAM::Policy', {
    Properties: Match.objectLike({
      PolicyDocument: Match.objectLike({
        Statement: Match.arrayWith([
          Match.objectLike({
            Action: 'ssm:GetParameter',
            Effect: 'Allow',
            Resource: {
              'Fn::Join': [
                '',
                [
                  'arn:',
                  { Ref: 'AWS::Partition' },
                  ':ssm:',
                  { Ref: 'AWS::Region' },
                  ':',
                  { Ref: 'AWS::AccountId' },
                  ':parameter/sops/age/private-key',
                ],
              ],
            },
          }),
        ]),
      }),
    }),
  });

  // Lambda role must have kms:Decrypt for the encryption key
  template.hasResource('AWS::IAM::Policy', {
    Properties: Match.objectLike({
      PolicyDocument: Match.objectLike({
        Statement: Match.arrayWith([
          Match.objectLike({
            Action: 'kms:Decrypt',
            Effect: 'Allow',
            Resource:
              'arn:aws:kms:us-east-1:111122223333:key/aaaabbbb-1234-5678-abcd-111122223333',
          }),
        ]),
      }),
    }),
  });
});

test('Age Key from SSM IStringParameter reference', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');
  const cmk = Key.fromKeyArn(
    stack,
    'SsmCmk',
    'arn:aws:kms:us-east-1:111122223333:key/aaaabbbb-1234-5678-abcd-111122223333',
  );

  const provider = new SopsSyncProvider(stack, 'Provider');
  const keyParam = StringParameter.fromStringParameterName(
    stack,
    'AgeKeyParam',
    '/sops/age/private-key',
  );
  provider.addAgeKeyFromSsmParameter(keyParam, cmk);
  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/yaml/sopsfile.enc-kms.yaml',
    sopsProvider: provider,
  });

  Template.fromStack(stack).hasResourceProperties('AWS::Lambda::Function', {
    Environment: Match.objectLike({
      Variables: Match.objectLike({
        SOPS_AGE_KEY_PARAMS: '/sops/age/private-key',
      }),
    }),
  });
});

test('Age Key from SSM combined with static age key', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');
  const cmk = Key.fromKeyArn(
    stack,
    'SsmCmk',
    'arn:aws:kms:us-east-1:111122223333:key/aaaabbbb-1234-5678-abcd-111122223333',
  );

  const provider = new SopsSyncProvider(stack, 'Provider');
  provider.addAgeKey(SecretValue.unsafePlainText('STATIC-KEY'));
  provider.addAgeKeyFromSsmParameter('/sops/age/private-key', cmk);

  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/yaml/sopsfile.enc-kms.yaml',
    sopsProvider: provider,
  });

  // Both env vars must be set
  Template.fromStack(stack).hasResourceProperties('AWS::Lambda::Function', {
    Environment: Match.objectLike({
      Variables: Match.objectLike({
        SOPS_AGE_KEY: 'STATIC-KEY',
        SOPS_AGE_KEY_PARAMS: '/sops/age/private-key',
      }),
    }),
  });
});

test('Age Key from SSM Parameter without leading slash gets correct ARN', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');
  const cmk = Key.fromKeyArn(
    stack,
    'SsmCmk',
    'arn:aws:kms:us-east-1:111122223333:key/aaaabbbb-1234-5678-abcd-111122223333',
  );

  const provider = new SopsSyncProvider(stack, 'Provider');
  provider.addAgeKeyFromSsmParameter('sops/age/private-key', cmk);
  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/yaml/sopsfile.enc-kms.yaml',
    sopsProvider: provider,
  });

  Template.fromStack(stack).hasResource('AWS::IAM::Policy', {
    Properties: Match.objectLike({
      PolicyDocument: Match.objectLike({
        Statement: Match.arrayWith([
          Match.objectLike({
            Action: 'ssm:GetParameter',
            Effect: 'Allow',
            Resource: {
              'Fn::Join': [
                '',
                [
                  'arn:',
                  { Ref: 'AWS::Partition' },
                  ':ssm:',
                  { Ref: 'AWS::Region' },
                  ':',
                  { Ref: 'AWS::AccountId' },
                  ':parameter/sops/age/private-key',
                ],
              ],
            },
          }),
        ]),
      }),
    }),
  });
});

// ---- Expiration tests ----

test('Expiration disabled by default - no scheduler/SNS resources', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');

  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/yaml/sopsfile.enc-kms.yaml',
  });

  const template = Template.fromStack(stack);
  template.resourceCountIs('AWS::SNS::Topic', 0);
  template.resourceCountIs('AWS::Scheduler::ScheduleGroup', 0);
  template.resourceCountIs('AWS::Scheduler::Schedule', 0);
  // The singleton SopsSync provider Lambda still needs its execution role.
  template.resourceCountIs('AWS::IAM::Role', 1);
});

test('Expiration disabled by default - config without enabled does not create scheduler/SNS resources', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');

  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/yaml/sopsfile.expiration.enc-kms.yaml',
    expiration: {
      daysBeforeExpiration: 30,
    },
  });

  const template = Template.fromStack(stack);
  template.resourceCountIs('AWS::SNS::Topic', 0);
  template.resourceCountIs('AWS::Scheduler::ScheduleGroup', 0);
  template.resourceCountIs('AWS::Scheduler::Schedule', 0);
});

test('Expiration enabled - auto-creates SNS topic, scheduler role, and schedule group', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');

  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/yaml/sopsfile.enc-kms.yaml',
    expiration: {
      enabled: true,
    },
  });

  const template = Template.fromStack(stack);
  template.resourceCountIs('AWS::SNS::Topic', 1);
  template.resourceCountIs('AWS::Scheduler::ScheduleGroup', 1);
  template.resourceCountIs('AWS::Scheduler::Schedule', 0);
  // Scheduler execution role + SopsSync provider role
  template.hasResourceProperties('AWS::IAM::Role', {
    AssumeRolePolicyDocument: Match.objectLike({
      Statement: Match.arrayWith([
        Match.objectLike({
          Action: 'sts:AssumeRole',
          Principal: { Service: 'scheduler.amazonaws.com' },
        }),
      ]),
    }),
  });
});

test('Expiration enabled - reuses one schedule group per stack', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');

  new SopsSecret(stack, 'FirstSecret', {
    secretName: 'first-secret',
    sopsFilePath: 'test-secrets/yaml/sopsfile.expiration.enc-kms.yaml',
    expiration: {
      enabled: true,
    },
  });

  new SopsSecret(stack, 'SecondSecret', {
    secretName: 'second-secret',
    sopsFilePath: 'test-secrets/yaml/sopsfile.expiration.enc-kms.yaml',
    expiration: {
      enabled: true,
    },
  });

  const template = Template.fromStack(stack);
  template.resourceCountIs('AWS::Scheduler::ScheduleGroup', 1);
  template.resourceCountIs('AWS::Scheduler::Schedule', 4);

  const schedules = Object.values(
    template.findResources('AWS::Scheduler::Schedule'),
  ).map((resource) => resource.Properties as Record<string, unknown>);

  expect(new Set(schedules.map((resource) => resource.GroupName)).size).toBe(1);
  expect(
    schedules.some((resource) => resource.Name === 'first-secret-gitlab_token'),
  ).toBe(true);
  expect(
    schedules.some(
      (resource) => resource.Name === 'second-secret-gitlab_token',
    ),
  ).toBe(true);
});

test('Expiration enabled - adds subscriber to auto-created SNS topic', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');

  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/yaml/sopsfile.enc-kms.yaml',
    expiration: {
      enabled: true,
      subscriber: new EmailSubscription('alerts@example.com'),
    },
  });

  const template = Template.fromStack(stack);
  template.resourceCountIs('AWS::SNS::Topic', 1);
  template.resourceCountIs('AWS::SNS::Subscription', 1);
  template.hasResourceProperties('AWS::SNS::Subscription', {
    Protocol: 'email',
    Endpoint: 'alerts@example.com',
  });
});

test('Expiration enabled - uses provided SNS topic instead of auto-creating', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');

  const existingTopic = new Topic(stack, 'MyTopic');

  const secret = new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/yaml/sopsfile.enc-kms.yaml',
    expiration: {
      enabled: true,
      notificationTopic: existingTopic,
    },
  });

  // The expirationNotificationTopic property should point to the provided topic
  expect(secret.expirationNotificationTopic).toBe(existingTopic);

  const template = Template.fromStack(stack);
  // Only 1 topic: the one we created manually (not auto-created)
  template.resourceCountIs('AWS::SNS::Topic', 1);
  template.resourceCountIs('AWS::Scheduler::Schedule', 0);
});

test('Expiration enabled - adds subscriber to provided SNS topic', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');

  const existingTopic = new Topic(stack, 'MyTopic');

  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/yaml/sopsfile.enc-kms.yaml',
    expiration: {
      enabled: true,
      notificationTopic: existingTopic,
      subscriber: new EmailSubscription('alerts@example.com'),
    },
  });

  const template = Template.fromStack(stack);
  template.resourceCountIs('AWS::SNS::Topic', 1);
  template.resourceCountIs('AWS::SNS::Subscription', 1);
  template.hasResourceProperties('AWS::SNS::Subscription', {
    Protocol: 'email',
    Endpoint: 'alerts@example.com',
  });
});

test('Expiration enabled - synthesizes schedules from unencrypted expiration keys', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');

  new SopsSecret(stack, 'SopsSecret', {
    secretName: 'my-secret',
    sopsFilePath: 'test-secrets/yaml/sopsfile.expiration.enc-kms.yaml',
    expiration: {
      enabled: true,
    },
  });

  const template = Template.fromStack(stack);
  template.resourceCountIs('AWS::Scheduler::Schedule', 2);

  const schedules = Object.values(
    template.findResources('AWS::Scheduler::Schedule'),
  ).map((resource) => resource.Properties as Record<string, unknown>);

  const gitlabSchedule = schedules.find(
    (resource) => resource.Name === 'my-secret-gitlab_token',
  );
  expect(gitlabSchedule).toBeDefined();
  expect(gitlabSchedule?.Description).toContain(
    'Notify about expiration of SOPS secret key "gitlab_token"',
  );
  expect(gitlabSchedule?.ScheduleExpression).toBe('at(2099-12-17T00:00:00)');
  expect(JSON.stringify(gitlabSchedule?.Target)).toContain(
    'SOPS_SECRET_EXPIRATION_NOTIFICATION',
  );
  expect(JSON.stringify(gitlabSchedule?.Target)).toContain(
    'SOPS secret expiration notification',
  );
  expect(JSON.stringify(gitlabSchedule?.Target)).toContain('stackName');
  expect(JSON.stringify(gitlabSchedule?.Target)).toContain('secretName');
  expect(JSON.stringify(gitlabSchedule?.Target)).toContain('gitlab_token');
  expect(JSON.stringify(gitlabSchedule?.Target)).toContain('2099-12-31');
  expect(JSON.stringify(gitlabSchedule?.Target)).toContain('2099-12-17');
  expect(JSON.stringify(gitlabSchedule?.Target)).toContain(
    '2099-12-31T00:00:00.000Z',
  );
  expect(JSON.stringify(gitlabSchedule?.Target)).toContain('scheduleGroupName');

  const nestedSchedule = schedules.find(
    (resource) => resource.Name === 'my-secret-nested-api_key',
  );
  expect(nestedSchedule).toBeDefined();
  expect(nestedSchedule?.Description).toContain(
    'Notify about expiration of SOPS secret key "nested.api_key"',
  );
  expect(nestedSchedule?.ScheduleExpression).toBe('at(2099-05-18T12:00:00)');
  expect(JSON.stringify(nestedSchedule?.Target)).toContain('nested.api_key');
  expect(JSON.stringify(nestedSchedule?.Target)).toContain('2099-06-01');
  expect(JSON.stringify(nestedSchedule?.Target)).toContain('2099-05-18');
  expect(JSON.stringify(nestedSchedule?.Target)).toContain(
    '2099-06-01T12:00:00.000Z',
  );
});

test('Expiration enabled - respects custom daysBeforeExpiration', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');

  new SopsSecret(stack, 'SopsSecret', {
    secretName: 'my-secret',
    sopsFilePath: 'test-secrets/yaml/sopsfile.expiration.enc-kms.yaml',
    expiration: {
      enabled: true,
      daysBeforeExpiration: 30,
    },
  });

  const schedules = Object.values(
    Template.fromStack(stack).findResources('AWS::Scheduler::Schedule'),
  ).map((resource) => resource.Properties as Record<string, unknown>);

  const gitlabSchedule = schedules.find(
    (resource) => resource.Name === 'my-secret-gitlab_token',
  );
  expect(gitlabSchedule).toBeDefined();
  expect(gitlabSchedule?.Description).toContain('30 days before expiration');
  expect(gitlabSchedule?.ScheduleExpression).toBe('at(2099-12-01T00:00:00)');
  expect(JSON.stringify(gitlabSchedule?.Target)).toContain('2099-12-01');
  expect(JSON.stringify(gitlabSchedule?.Target)).toContain('30');
  expect(JSON.stringify(gitlabSchedule?.Target)).toContain(
    '2099-12-01T00:00:00.000Z',
  );

  const nestedSchedule = schedules.find(
    (resource) => resource.Name === 'my-secret-nested-api_key',
  );
  expect(nestedSchedule).toBeDefined();
  expect(nestedSchedule?.ScheduleExpression).toBe('at(2099-05-02T12:00:00)');
  expect(JSON.stringify(nestedSchedule?.Target)).toContain('2099-05-02');
  expect(JSON.stringify(nestedSchedule?.Target)).toContain('30');
});

test('Expiration enabled - CustomResource no longer gets Expiration config', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');

  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/yaml/sopsfile.enc-kms.yaml',
    expiration: {
      enabled: true,
    },
  });

  const resources = Template.fromStack(stack).findResources('Custom::SopsSync');
  const properties = Object.values(resources)[0].Properties as Record<
    string,
    unknown
  >;
  expect(properties).not.toHaveProperty('Expiration');
});

test('Expiration enabled - Lambda role no longer gets scheduler permissions', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');

  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/yaml/sopsfile.expiration.enc-kms.yaml',
    expiration: {
      enabled: true,
    },
  });

  const policies = Template.fromStack(stack).findResources('AWS::IAM::Policy');
  expect(JSON.stringify(policies)).not.toContain('scheduler:CreateSchedule');
  expect(JSON.stringify(policies)).not.toContain('iam:PassRole');
});

test('Expiration enabled - invalid expiration date throws during synthesis setup', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');

  expect(
    () =>
      new SopsSecret(stack, 'SopsSecret', {
        sopsFilePath:
          'test-secrets/yaml/sopsfile.invalid-expiration.enc-kms.yaml',
        expiration: {
          enabled: true,
        },
      }),
  ).toThrow('unsupported date format: "not-a-date"');
});

test('Expiration enabled - s3 source is rejected', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');

  expect(
    () =>
      new SopsSecret(stack, 'SopsSecret', {
        sopsS3Bucket: 'bucket',
        sopsS3Key: 'secret.yaml',
        sopsFileFormat: 'yaml',
        expiration: {
          enabled: true,
        },
      }),
  ).toThrow(
    'Expiration scheduling requires a local sopsFilePath and does not support sopsS3Bucket/sopsS3Key.',
  );
});
