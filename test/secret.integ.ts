import { Key } from '@aws-cdk/aws-kms';
import { Secret } from '@aws-cdk/aws-secretsmanager';
import { App, Stack } from '@aws-cdk/core';
import path from 'path';
import { SopsSecrets } from '../src/index';

const app = new App();

const stack = new Stack(app, 'SecretIntegration');

const key = new Key(stack, 'Key', {
  alias: 'sops-test-key',
});

const secret = new Secret(stack, 'Secret', {
  secretName: 'sops-test-secret',
  encryptionKey: key,
});

new SopsSecrets(stack, 'SopsSecret', {
  secret,
  sopsFilePath: path.join(__dirname, 'sopsfile.json'),
  sopsKmsKey: key,
  sopsFileFormat: 'json',
  /**
   * For testing purpose, we provide an age encrypted file,
   * encrypted for the following public key
   * 
   * age1djllw2pzuprrqc0en5m8vc8k5ge3tm0f6g7cj0c0glfzp44vdc4ql8ngvu
   * 
   * and can be decrypted with the following private key
   */
  // age1djllw2pzuprrqc0en5m8vc8k5ge3tm0f6g7cj0c0glfzp44vdc4ql8ngvu
  sopsAgeKey: 'AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3',
});