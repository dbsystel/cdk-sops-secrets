import { Match, Template } from '@aws-cdk/assertions';
import { App, Stack } from '@aws-cdk/core';

import { SopsSecrets } from '../src';

const keyStatements = [{
  Action: [
    'kms:Encrypt',
    'kms:ReEncrypt*',
    'kms:GenerateDataKey*',
  ],
  Effect: 'Allow',
  Resource: 'arn:aws:kms:aws-region-1:123456789011:key/00000000-1234-4321-abcd-1234abcd12ab',
}];

test('KMS Key lookup from sopsfile: json', () => {
  const app = new App();
  const stack = new Stack(app, 'SecretIntegration');
  new SopsSecrets(stack, 'SopsSecret', {
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
  new SopsSecrets(stack, 'SopsSecret', {
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