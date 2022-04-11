import * as fs from 'fs';
import * as path from 'path';
//import * as crypto from 'crypto';
import {
  Annotations,
  CustomResource,
  FileSystem,
  Lazy,
  RemovalPolicy,
  ResourceEnvironment,
  SecretValue,
  Stack,
} from 'aws-cdk-lib';
import {
  IGrantable,
  Grant,
  PolicyStatement,
  AddToResourcePolicyResult,
} from 'aws-cdk-lib/aws-iam';
import { IKey, Key } from 'aws-cdk-lib/aws-kms';
import { Code, Runtime, SingletonFunction } from 'aws-cdk-lib/aws-lambda';
import { Asset } from 'aws-cdk-lib/aws-s3-assets';
import {
  ISecret,
  ISecretAttachmentTarget,
  RotationSchedule,
  RotationScheduleOptions,
  Secret,
  SecretProps,
} from 'aws-cdk-lib/aws-secretsmanager';
import { Construct } from 'constructs';
import { execSync } from 'child_process';

export enum UploadType {
  /**
   * Pass the secret data inline (base64 encoded and compressed)
   */
  INLINE = 'INLINE',
  /**
   * Uplaod the secert data as asset
   */
  ASSET = 'ASSET',
}

/**
 * Configuration options for the SopsSync
 */
export interface SopsSyncOptions {
  /**
   * The custom resource provider to use. If you don't specify any, a new
   * provider will be created - or if already exists within this stack - reused.
   *
   * @default - A new singleton provider will be created
   */
  readonly sopsProvider?: SopsSyncProvider;

  /**
   * The filepath to the sops file
   */
  readonly sopsFilePath: string;

  /**
   * How should the secret be passed to the CustomResource?
   * @default INLINE
   */
  readonly uploadType?: UploadType;

  /**
   * The format of the sops file.
   *
   * @default - The fileformat will be derived from the file ending
   */
  readonly sopsFileFormat?: undefined | 'json' | 'yaml';

  /**
   * The kmsKey used to encrypt the sops file. Encrypt permissions
   * will be granted to the custom resource provider.
   *
   * @default - The key will be derived from the sops file
   */
  readonly sopsKmsKey?: IKey[];

  /**
   * The age key that should be used for encryption.
   */
  readonly sopsAgeKey?: SecretValue;

  /**
   * Should the encrypted sops value should be converted to JSON?
   * Only JSON can be handled by cloud formations dynamic references.
   *
   * @default true
   */
  readonly convertToJSON?: boolean;

  /**
   * Should the structure be flattened? The result will be a flat structure and all
   * object keys will be replaced with the full jsonpath as key.
   * This is usefull for dynamic references, as those don't support nested objects.
   *
   * @default true
   */
  readonly flatten?: boolean;

  /**
   * Shall all values be flattened? This is usefull for dynamic references, as there
   * are lookup errors for certain float types
   */
  readonly stringifyValues?: boolean;
}

/**
 * The configuration options extended by the target Secret
 */
export interface SopsSyncProps extends SopsSyncOptions {
  /**
   * The secret that will be populated with the encrypted sops file content.
   */
  readonly secret: ISecret;
}

/**
 * The custom resource, that is syncing the content from a sops file to a secret.
 */
export class SopsSync extends Construct {
  /**
   * The current versionId of the secret populated via this resource
   */
  readonly versionId: string;

  /**
   * The format of the input file
   */
  readonly sopsFileFormat: 'json' | 'yaml';

  /**
   * Was the format converted to json?
   */
  readonly converToJSON: boolean;

  /**
   * Was the structure flattened?
   */
  readonly flatten: boolean;

  /**
   * Were the values stringified?
   */
  readonly stringifiedValues: boolean;

  constructor(scope: Construct, id: string, props: SopsSyncProps) {
    super(scope, id);

    this.converToJSON = props.convertToJSON ?? true;
    this.flatten = props.flatten ?? true;
    this.stringifiedValues = props.stringifyValues ?? true;

    const sopsFileFormat =
      props.sopsFileFormat ?? props.sopsFilePath.split('.').pop();
    switch (sopsFileFormat) {
      case 'json': {
        this.sopsFileFormat = 'json';
        break;
      }
      case 'yaml': {
        this.sopsFileFormat = 'yaml';
        break;
      }
      case 'yml': {
        this.sopsFileFormat = 'yaml';
        break;
      }
      default: {
        throw new Error(`Unsupported sopsFileFormat ${sopsFileFormat}`);
      }
    }

    const provider = props.sopsProvider ?? new SopsSyncProvider(scope);

    if (!fs.existsSync(props.sopsFilePath)) {
      throw new Error(`File ${props.sopsFilePath} does not exist!`);
    }
    

    let fileHash:string = FileSystem.fingerprint(props.sopsFilePath);
    //try {
      const commit = execSync(`git rev-list ${ process.env.GITHUB_HEAD_REF ?? '--all'} --no-merges -1 ${props.sopsFilePath}`).toString().replace('\n','');
      // Check if the result is a valid git commit
    //  if (commit.match("^[a-fA-F0-9]{40}$")?.length == 1) {
        fileHash = commit
    //  }
    //} catch (e) {}
  

    /**
     * Handle uploadType INLINE or ASSET
     */
    const uploadType = props.uploadType ?? UploadType.INLINE;
    let sopsAsset: Asset | undefined = undefined;
    let sopsInline: { Content: string; Hash: string } | undefined = undefined;
    let sopsS3File: { Bucket: string; Key: string } | undefined = undefined;
    if (uploadType === UploadType.INLINE) {
      sopsInline = {
        Content: fs.readFileSync(props.sopsFilePath).toString('base64'),
        // We calculate the hash the same way as it would be done by new Asset(..) - so we can ensure stable version names even if switching from INLINE to ASSET and viceversa.
        Hash: FileSystem.fingerprint(props.sopsFilePath),
      };
    } else if (uploadType === UploadType.ASSET) {
      sopsAsset = new Asset(this, 'Asset', {
        path: props.sopsFilePath,
      });
      sopsS3File = {
        Bucket: sopsAsset.bucket.bucketName,
        Key: sopsAsset.s3ObjectKey,
      };
    } else {
      throw new Error(`Unsupported UploadType: ${uploadType}`);
    }

    if (provider.role !== undefined) {
      if (props.sopsKmsKey !== undefined) {
        props.sopsKmsKey.forEach((key) => key.grantDecrypt(provider.role!));
      }
      const fileContent = fs.readFileSync(props.sopsFilePath);
      const regex = /arn:aws:kms:[a-z0-9-]+:[\d]+:key\/[a-z0-9-]+/gm;
      const results = regex.exec(fileContent.toString());
      if (results !== undefined) {
        results?.forEach((result, index) =>
          Key.fromKeyArn(this, `SopsKey${index}`, result).grantDecrypt(
            provider.role!,
          ),
        );
      }
      props.secret.grantWrite(provider);
      if (sopsAsset !== undefined) {
        sopsAsset.bucket.grantRead(provider);
      }
    } else {
      Annotations.of(this).addWarning(
        `Please ensure propper permissions for the passed lambda function:\n  - write Access to the secret\n  - encrypt with the sopsKmsKey${
          uploadType === UploadType.ASSET
            ? '\n  - download from asset bucket'
            : ''
        }`,
      );
    }
    if (props.sopsAgeKey !== undefined) {
      provider.addAgeKey(props.sopsAgeKey);
    }

    const cr = new CustomResource(this, 'Resource', {
      serviceToken: provider.functionArn,
      resourceType: 'Custom::SopsSync',
      properties: {
        SecretARN: props.secret.secretArn,
        SopsS3File: sopsS3File,
        SopsInline: sopsInline,
        Hash: fileHash,
        ConvertToJSON: this.converToJSON,
        Flatten: this.flatten,
        Format: this.sopsFileFormat,
        StringifiedValues: this.stringifiedValues,
      },
    });
    this.versionId = cr.getAttString('VersionId');
  }
}

export class SopsSyncProvider extends SingletonFunction implements IGrantable {
  private sopsAgeKeys: SecretValue[];

  constructor(scope: Construct, id?: string) {
    super(scope, id ?? 'SopsSyncProvider', {
      code: Code.fromAsset(
        path.join(__dirname, '../assets/cdk-sops-lambda.zip'),
      ),
      runtime: Runtime.GO_1_X,
      handler: 'cdk-sops-secrets',
      uuid: 'SopsSyncProvider',
      environment: {
        SOPS_AGE_KEY: Lazy.string({
          produce: () =>
            (this.sopsAgeKeys.map((secret) => secret.toString()) ?? []).join(
              '\n',
            ),
        }),
      },
    });
    this.sopsAgeKeys = [];
  }

  public addAgeKey(key: SecretValue) {
    this.sopsAgeKeys.push(key);
  }
}

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
      ...(props as SopsSyncOptions),
    });
  }

  /**
   * Returns the current versionId that was created via the SopsSync
   */
  public currentVersionId(): string {
    return this.sync.versionId;
  }

  public grantRead(grantee: IGrantable, versionStages?: string[]): Grant {
    return this.secret.grantRead(grantee, versionStages);
  }
  public grantWrite(grantee: IGrantable): Grant {
    return this.secret.grantWrite(grantee);
  }
  public addRotationSchedule(
    id: string,
    options: RotationScheduleOptions,
  ): RotationSchedule {
    throw new Error(
      `Method not allowed as this secret is managed by SopsSync!\nid: ${id}\noptions: ${JSON.stringify(
        options,
        null,
        2,
      )}`,
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
