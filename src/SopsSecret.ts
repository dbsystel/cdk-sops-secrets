import {
  AddToResourcePolicyResult,
  Grant,
  IGrantable,
  PolicyStatement,
} from 'aws-cdk-lib/aws-iam';
import { IKey } from 'aws-cdk-lib/aws-kms';
import {
  ISecret,
  ISecretAttachmentTarget,
  ReplicaRegion,
  RotationSchedule,
  RotationScheduleOptions,
  Secret,
  SecretProps,
} from 'aws-cdk-lib/aws-secretsmanager';
import {
  RemovalPolicy,
  ResourceEnvironment,
  SecretValue,
  Stack,
} from 'aws-cdk-lib/core';
import { Construct } from 'constructs';
import { ResourceType, SopsSync, SopsSyncOptions } from './SopsSync';

export enum RawOutput {
  /**
   * Parse the secret as a string
   */
  STRING = 'STRING',
  /**
   * Parse the secret as a binary
   */
  BINARY = 'BINARY',
}

/**
 * The configuration options of the SopsSecret
 */
export interface SopsSecretProps extends SopsSyncOptions {
  /**
   * Should the secret parsed and transformed to json?
   * @default - undefined - no raw output
   */
  readonly rawOutput?: RawOutput;
  /**
   * An optional, human-friendly description of the secret.
   *
   * @default - No description.
   */
  readonly description?: string;
  /**
   * The customer-managed encryption key to use for encrypting the secret value.
   *
   * @default - A default KMS key for the account and region is used.
   */
  readonly encryptionKey?: IKey;
  /**
   * A name for the secret. Note that deleting secrets from SecretsManager does not happen immediately, but after a 7 to
   * 30 days blackout period. During that period, it is not possible to create another secret that shares the same name.
   *
   * @default - A name is generated by CloudFormation.
   */
  readonly secretName?: string;
  /**
   * Policy to apply when the secret is removed from this stack.
   *
   * @default - Not set.
   */
  readonly removalPolicy?: RemovalPolicy;
  /**
   * A list of regions where to replicate this secret.
   *
   * @default - Secret is not replicated
   */
  readonly replicaRegions?: ReplicaRegion[];
}

/**
 * A drop in replacement for the normal Secret, that is populated with the encrypted
 * content of the given sops file.
 */
export class SopsSecret extends Construct implements ISecret {
  private readonly secret: Secret;
  readonly encryptionKey?: IKey | undefined;
  readonly secretArn: string;
  readonly secretFullArn?: string | undefined;
  readonly secretName: string;
  readonly stack: Stack;
  readonly env: ResourceEnvironment;

  readonly sync: SopsSync;
  public constructor(scope: Construct, id: string, props: SopsSecretProps) {
    super(scope, id);
    this.secret = new Secret(this, 'Resource', props satisfies SecretProps);

    // Fullfill secret Interface
    this.encryptionKey = this.secret.encryptionKey;
    this.secretArn = this.secret.secretArn;
    this.secretName = this.secret.secretName;
    this.stack = Stack.of(scope);
    this.env = {
      account: this.stack.account,
      region: this.stack.region,
    };

    let resourceType = ResourceType.SECRET;
    if (props.rawOutput === RawOutput.BINARY) {
      resourceType = ResourceType.SECRET_BINARY;
    }
    if (props.rawOutput === RawOutput.STRING) {
      resourceType = ResourceType.SECRET_RAW;
    }

    this.sync = new SopsSync(this, 'SopsSync', {
      target: this.secret.secretArn,
      resourceType,
      flattenSeparator: '.',
      secret: this.secret,
      ...(props as SopsSyncOptions),
    });
  }

  public currentVersionId(): string {
    return this.sync.versionId;
  }

  public grantRead(grantee: IGrantable, versionStages?: string[]): Grant {
    return this.secret.grantRead(grantee, versionStages);
  }
  public grantWrite(_grantee: IGrantable): Grant {
    throw new Error(
      `Method grantWrite(...) not allowed as this secret is managed by SopsSync`,
    );
  }
  public addRotationSchedule(
    id: string,
    options: RotationScheduleOptions,
  ): RotationSchedule {
    throw new Error(
      `Method addRotationSchedule('${id}', ${JSON.stringify(
        options,
      )}) not allowed as this secret is managed by SopsSync`,
    );
  }
  public addToResourcePolicy(
    statement: PolicyStatement,
  ): AddToResourcePolicyResult {
    return this.secret.addToResourcePolicy(statement);
  }
  public denyAccountRootDelete(): void {
    return this.secret.denyAccountRootDelete();
  }
  public attach(target: ISecretAttachmentTarget): ISecret {
    return this.secret.attach(target);
  }
  public applyRemovalPolicy(policy: RemovalPolicy): void {
    return this.secret.applyRemovalPolicy(policy);
  }

  public secretValueFromJson(jsonField: string) {
    return SecretValue.secretsManager(this.secretArn, {
      jsonField,
      versionId: this.sync.versionId,
    });
  }

  public get secretValue(): SecretValue {
    return this.secretValueFromJson('');
  }
}
