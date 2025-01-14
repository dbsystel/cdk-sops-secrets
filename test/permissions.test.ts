import * as fs from 'fs';
import path from 'path';
import { Stack } from 'aws-cdk-lib';
import { Template } from 'aws-cdk-lib/assertions';
import { IRole, Role } from 'aws-cdk-lib/aws-iam';
import { Key } from 'aws-cdk-lib/aws-kms';
import { Permissions } from '../src/SopsSync';

describe('keysFromSopsContent', () => {
  let stack: Stack;

  beforeEach(() => {
    stack = new Stack();
  });

  test('returns an empty array when no keys are found', () => {
    const content = 'no keys here';
    const result = Permissions.keysFromSopsContent(stack, content);
    expect(result).toEqual([]);
  });

  test('returns an array of IKey objects when keys are found', () => {
    const content = fs
      .readFileSync(
        path.join(__dirname, '../test-secrets/yaml/sopsfile.enc-kms.yaml'),
      )
      .toString();
    const result = Permissions.keysFromSopsContent(stack, content);
    expect(result).toHaveLength(1);
    expect(result[0].keyArn).toBe(
      'arn:aws:kms:aws-region-1:123456789011:key/00000000-1234-4321-abcd-1234abcd12ab',
    );
    expect(result[0].keyId).toBe('00000000-1234-4321-abcd-1234abcd12ab');
  });

  test('returns multiple IKey objects when multiple keys are found', () => {
    const content = fs
      .readFileSync(
        path.join(__dirname, '../test-secrets/yaml/sopsfile.enc-multikms.yaml'),
      )
      .toString();
    const result = Permissions.keysFromSopsContent(stack, content);
    expect(result).toHaveLength(4);
    expect(result[0].keyArn).toBe(
      'arn:aws:kms:aws-region-1:123456789011:key/00000000-1234-4321-abcd-1234abcd12ab',
    );
    expect(result[0].keyId).toBe('00000000-1234-4321-abcd-1234abcd12ab');
    expect(result[1].keyArn).toBe(
      'arn:aws:kms:aws-region-1:123456789011:key/00000001-1234-4321-abcd-1234abcd12ab',
    );
    expect(result[1].keyId).toBe('00000001-1234-4321-abcd-1234abcd12ab');
    expect(result[2].keyArn).toBe(
      'arn:aws:kms:aws-region-1:123456789011:key/00000002-1234-4321-abcd-1234abcd12ab',
    );
    expect(result[2].keyId).toBe('00000002-1234-4321-abcd-1234abcd12ab');
    expect(result[3].keyArn).toBe(
      'arn:aws:kms:aws-region-1:123456789011:key/00000003-1234-4321-abcd-1234abcd12ab',
    );
    expect(result[3].keyId).toBe('00000003-1234-4321-abcd-1234abcd12ab');
  });
});

describe('keysFromSopsContentAlias', () => {
  let stack: Stack;

  beforeEach(() => {
    stack = new Stack();
  });

  test('returns an empty array when no keys are found', () => {
    const content = 'no keys here';
    const result = Permissions.keysFromSopsContentAlias(stack, content);
    expect(result).toEqual([]);
  });

  test('returns an array of IKey objects when keys are found', () => {
    jest
      .spyOn(Key, 'fromLookup')
      .mockReturnValue(
        Key.fromKeyArn(
          stack,
          'Key',
          'arn:aws:kms:aws-region-1:123456789011:key/00000000-1234-4321-abcd-1234abcd12ab',
        ),
      );

    const content = fs
      .readFileSync(
        path.join(
          __dirname,
          '../test-secrets/yaml/sopsfile.enc-kms-alias.yaml',
        ),
      )
      .toString();
    const result = Permissions.keysFromSopsContentAlias(stack, content);
    expect(result[0].keyArn).toBe(
      'arn:aws:kms:aws-region-1:123456789011:key/00000000-1234-4321-abcd-1234abcd12ab',
    );
    expect(result[0].keyId).toBe('00000000-1234-4321-abcd-1234abcd12ab');
  });
});

describe('parameters', () => {
  let stack: Stack;
  let role: IRole;

  beforeEach(() => {
    stack = new Stack();
    role = Role.fromRoleArn(
      stack,
      'TestRole',
      'arn:aws:iam::123456789012:role/test-role',
    );
  });

  function genParameters(prefix: string, count: number) {
    const params: string[] = [];
    for (let i = 0; i < count; i++) {
      params.push(`${prefix}${i}`);
    }
    return params;
  }

  test('1 parameter - full snapshot', () => {
    const params = genParameters('parameter', 1);
    Permissions.parameters(stack, params, role);
    expect(
      Template.fromStack(stack).findResources('AWS::IAM::ManagedPolicy'),
    ).toMatchSnapshot();
  });

  test('100 parameter - snapshot count', () => {
    const params = genParameters('parameter', 100);
    Permissions.parameters(stack, params, role);
    expect(
      Object.keys(
        Template.fromStack(stack).findResources('AWS::IAM::ManagedPolicy'),
      ),
    ).toMatchSnapshot();
  });

  test('1000 parameter - snapshot count', () => {
    const params = genParameters('parameter', 1000);
    Permissions.parameters(stack, params, role);
    expect(
      Object.keys(
        Template.fromStack(stack).findResources('AWS::IAM::ManagedPolicy'),
      ),
    ).toMatchSnapshot();
  });
});
