import * as fs from 'fs';
import path from 'path';
import { Stack } from 'aws-cdk-lib';
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
