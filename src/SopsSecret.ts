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
import {
  CreationType,
  ResourceType,
  SopsSync,
  SopsSyncOptions,
} from './SopsSync';

/**
 * The configuration options of the SopsSecret
 */
export interface SopsSecretProps extends SecretProps, SopsSyncOptions {}

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
    this.secret = new Secret(this, 'Resource', props as SecretProps);

    // Fullfill secret Interface
    this.encryptionKey = this.secret.encryptionKey;
    this.secretArn = this.secret.secretArn;
    this.secretName = this.secret.secretName;
    this.stack = Stack.of(scope);
    this.env = {
      account: this.stack.account,
      region: this.stack.region,
    };

    this.sync = new SopsSync(this, 'SopsSync', {
      secret: this.secret,
      resourceType: ResourceType.SECRET,
      creationType: CreationType.SINGLE,
      flattenSeparator: '.',
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
