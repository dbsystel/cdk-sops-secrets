import * as fs from 'fs';
import * as path from 'path';
//import * as crypto from 'crypto';
import {
  Annotations,
  CustomResource,
  Duration,
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
  readonly sopsFilePath?: string;

  /**
   * If you want to pass the sops file via s3, you can specify the bucket
   * you can use cfn parameter here
   * Both, sopsS3Bucket and sopsS3Key have to be specified
   */
  readonly sopsS3Bucket?: string;
  /**
   * If you want to pass the sops file via s3, you can specify the key inside the bucket
   * you can use cfn parameter here
   * Both, sopsS3Bucket and sopsS3Key have to be specified
   */
  readonly sopsS3Key?: string;

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
  readonly sopsFileFormat?: undefined | 'json' | 'yaml' | 'dotenv' | 'raw';

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

    const provider = props.sopsProvider ?? new SopsSyncProvider(scope);

    let uploadType = props.uploadType ?? UploadType.INLINE;
    let sopsFileFormat: 'json' | 'yaml' | 'dotenv' | 'raw' | undefined =
      props.sopsFileFormat;
    let sopsAsset: Asset | undefined = undefined;
    let sopsInline: { Content: string; Hash: string } | undefined = undefined;
    let sopsS3File: { Bucket: string; Key: string } | undefined = undefined;

    if (
      props.sopsFilePath !== undefined &&
      (props.sopsS3Bucket !== undefined || props.sopsS3Key !== undefined)
    ) {
      throw new Error(
        'You can either specify sopsFilePath or sopsS3Bucket and sopsS3Key!',
      );
    }

    if (props.sopsFilePath !== undefined) {
      const _sopsFileFormat =
        props.sopsFileFormat ?? props.sopsFilePath.split('.').pop();
      switch (_sopsFileFormat) {
        case 'json': {
          sopsFileFormat = 'json';
          break;
        }
        case 'yaml': {
          sopsFileFormat = 'yaml';
          break;
        }
        case 'yml': {
          sopsFileFormat = 'yaml';
          break;
        }
        case 'dotenv': {
          sopsFileFormat = 'dotenv';
          break;
        }
        case 'env': {
          sopsFileFormat = 'dotenv';
          break;
        }
        case 'raw': {
          sopsFileFormat = 'raw';
          break;
        }
        default: {
          throw new Error(`Unsupported sopsFileFormat ${_sopsFileFormat}`);
        }
      }

      if (!fs.existsSync(props.sopsFilePath)) {
        throw new Error(`File ${props.sopsFilePath} does not exist!`);
      }

      switch (uploadType) {
        case UploadType.INLINE: {
          sopsInline = {
            Content: fs.readFileSync(props.sopsFilePath).toString('base64'),
            // We calculate the hash the same way as it would be done by new Asset(..) - so we can ensure stable version names even if switching from INLINE to ASSET and viceversa.
            Hash: FileSystem.fingerprint(props.sopsFilePath),
          };
          break;
        }
        case UploadType.ASSET: {
          sopsAsset = new Asset(this, 'Asset', {
            path: props.sopsFilePath,
          });
          sopsS3File = {
            Bucket: sopsAsset.bucket.bucketName,
            Key: sopsAsset.s3ObjectKey,
          };
          break;
        }
      }

      if (provider.role !== undefined) {
        if (props.sopsKmsKey !== undefined) {
          props.sopsKmsKey.forEach((key) => key.grantDecrypt(provider.role!));
        }
        const fileContent = fs.readFileSync(props.sopsFilePath);
        // Handle keys
        const regexKey = /arn:aws:kms:[a-z0-9-]+:[\d]+:key\/[a-z0-9-]+/g;
        const resultsKey = fileContent.toString().match(regexKey);
        if (resultsKey !== undefined) {
          resultsKey?.forEach((result, index) =>
            Key.fromKeyArn(this, `SopsKey${index}`, result).grantDecrypt(
              provider.role!,
            ),
          );
        }
        const regexAlias = /arn:aws:kms:[a-z0-9-]+:[\d]+:alias\/[a-z0-9-]+/g;
        const resultsAlias = fileContent.toString().match(regexAlias);
        if (resultsAlias !== undefined) {
          resultsAlias?.forEach((result, index) =>
            Key.fromLookup(this, `SopsAlias${index}`, {
              aliasName: `alias/${result.split('/').slice(1).join('/')}`,
            }).grantDecrypt(provider.role!),
          );
        }
        props.secret.grantWrite(provider);
        if (sopsAsset !== undefined) {
          sopsAsset.bucket.grantRead(provider);
        }
        /**
         * fixes #234
         * If the kms key for secrets encryption is an IKey
         * there will be no permissions otherwise
         */
        if (
          props.secret.encryptionKey !== undefined &&
          !(props.secret.encryptionKey instanceof Key)
        ) {
          props.secret.encryptionKey.grantEncryptDecrypt(provider);
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
    } else if (
      props.sopsS3Bucket !== undefined &&
      props.sopsS3Key !== undefined
    ) {
      sopsS3File = {
        Bucket: props.sopsS3Bucket,
        Key: props.sopsS3Key,
      };
      uploadType = UploadType.ASSET;
      Annotations.of(this).addWarning(
        'You have to manually add permissions to the sops provider to (permission to download file, to decrypt sops file)!',
      );
    } else {
      throw new Error(
        'You have to specify both sopsS3Bucket and sopsS3Key or neither!',
      );
    }

    if (sopsFileFormat === undefined) {
      throw new Error('You have to specify sopsFileFormat!');
    }

    const cr = new CustomResource(this, 'Resource', {
      serviceToken: provider.functionArn,
      resourceType: 'Custom::SopsSync',
      properties: {
        SecretARN: props.secret.secretArn,
        SopsS3File: sopsS3File,
        SopsInline: sopsInline,
        ConvertToJSON: this.converToJSON,
        Flatten: this.flatten,
        Format: sopsFileFormat,
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
        scope.node.tryGetContext('sops_sync_provider_asset_path') ||
          path.join(__dirname, '../assets/cdk-sops-lambda.zip'),
      ),
      runtime: Runtime.PROVIDED_AL2,
      handler: 'bootstrap',
      uuid: 'SopsSyncProvider',
      timeout: Duration.seconds(60),
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
      `Method addTotationSchedule('${id}', ${JSON.stringify(
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
