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
import {
  SopsSecret,
  SopsSyncProvider,
  UploadType,
  MultiStringParameter,
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
  ).toThrowError('File test-secrets/does-not-exist.json does not exist!');
});

test('Age Key passed', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');

  new SopsSecret(stack, 'SopsSecret', {
    sopsFilePath: 'test-secrets/yaml/sopsfile.enc-kms.yaml',
    sopsAgeKey: SecretValue.plainText('SOME-KEY'),
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

test('Exception when derive format: notsupported', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');

  expect(
    () =>
      new SopsSecret(stack, 'SopsSecret', {
        sopsFilePath: 'test-secrets/json/sopsfile.enc-age.notsupported',
      }),
  ).toThrowError('Unsupported sopsFileFormat notsupported');
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

  expect(() => secret.addRotationSchedule('something', {})).toThrowError(
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
  ).toThrowError(
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
  ).toThrowError(
    'You can either specify sopsFilePath or sopsS3Bucket and sopsS3Key!',
  );
  expect(
    () =>
      new SopsSecret(stack, 'SopsSecret2', {
        sopsS3Key: 'test',
      }),
  ).toThrowError(
    'You have to specify both sopsS3Bucket and sopsS3Key or neither!',
  );
  expect(
    () =>
      new SopsSecret(stack, 'SopsSecret3', {
        sopsS3Bucket: 'test',
      }),
  ).toThrowError(
    'You have to specify both sopsS3Bucket and sopsS3Key or neither!',
  );
  expect(
    () =>
      new SopsSecret(stack, 'SopsSecret4', {
        sopsS3Key: 'test',
        sopsS3Bucket: 'test',
      }),
  ).toThrowError('You have to specify sopsFileFormat!');
});

test('Multiple parameters from yaml file', () => {
  const app = new App();
  const stack = new Stack(app, 'ParameterIntegration');
  new MultiStringParameter(stack, 'SopsSecret1', {
    simpleName: false,
    sopsFilePath: 'test-secrets/yaml/sopsfile-complex-parameters.enc-age.yaml',
    encryptionKey: Key.fromKeyArn(
      stack,
      'Key',
      'arn:aws:kms:eu-central-1:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab',
    ),
    stringValue: ' ',
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
});

test('Multiple parameters from yaml file with custom key structure', () => {
  const app = new App();
  const stack = new Stack(app, 'ParameterIntegration');
  new MultiStringParameter(stack, 'SopsSecret1', {
    simpleName: false,
    sopsFilePath: 'test-secrets/yaml/sopsfile-complex-parameters.enc-age.yaml',
    keyPrefix: '_',
    keySeperator: '.',
    encryptionKey: Key.fromKeyArn(
      stack,
      'Key',
      'arn:aws:kms:eu-central-1:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab',
    ),
    stringValue: ' ',
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
});
