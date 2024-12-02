import * as fs from 'fs';
import * as path from 'path';
import {
  SecretValue,
  Duration,
  Lazy,
  Stack,
  Annotations,
  CustomResource,
  FileSystem,
} from 'aws-cdk-lib';
import { ISecurityGroup, IVpc, SubnetSelection } from 'aws-cdk-lib/aws-ec2';
import {
  IGrantable,
  IRole,
  ManagedPolicy,
  PolicyStatement,
} from 'aws-cdk-lib/aws-iam';
import { IKey, Key } from 'aws-cdk-lib/aws-kms';
import { SingletonFunction, Code, Runtime } from 'aws-cdk-lib/aws-lambda';
import { Asset } from 'aws-cdk-lib/aws-s3-assets';
import { ISecret } from 'aws-cdk-lib/aws-secretsmanager';
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

export enum CreationType {
  /**
   * Create or update a single secret/parameter
   */
  SINGLE = 'SINGLE',
  /**
   * Create or update a multiple secrets/parameters by flattening the SOPS file
   */
  MULTI = 'MULTI',
}

export enum ResourceType {
  SECRET = 'SECRET',
  PARAMETER = 'PARAMETER',
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
  readonly sopsFileFormat?: undefined | 'json' | 'yaml' | 'dotenv' | 'binary';

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
   * If the structure should be flattened use the provided separator between keys.
   *
   * @default '.'
   */
  readonly flattenSeparator?: string;

  /**
   * Add this prefix to parameter names.
   */
  readonly parameterKeyPrefix?: string;

  /**
   * Shall all values be flattened? This is usefull for dynamic references, as there
   * are lookup errors for certain float types
   */
  readonly stringifyValues?: boolean;

  readonly creationType?: CreationType;
  readonly resourceType?: ResourceType;
}

/**
 * The configuration options extended by the target Secret
 */
export interface SopsSyncProps extends SopsSyncOptions {
  /**
   * The secret that will be populated with the encrypted sops file content.
   */
  readonly secret?: ISecret;

  /**
   * The parameter name. If set this creates an encrypted SSM Parameter instead of a secret.
   */
  readonly parameterName?: string;

  /**
   * The parameter name. If set this creates an encrypted SSM Parameter instead of a secret.
   */
  readonly parameterNames?: string[];

  /**
   * The encryption key used for encrypting the ssm parameter if `parameterName` is set.
   */
  readonly encryptionKey?: IKey;
}

export interface SopsSyncProviderProps {
  readonly vpc?: IVpc;
  readonly vpcSubnets?: SubnetSelection;
  readonly securityGroups?: ISecurityGroup[];
}

export class SopsSyncProvider extends SingletonFunction implements IGrantable {
  private sopsAgeKeys: SecretValue[];

  constructor(scope: Construct, id?: string, props?: SopsSyncProviderProps) {
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
      vpc: props?.vpc,
      vpcSubnets: props?.vpcSubnets,
      securityGroups: props?.securityGroups,
    });
    this.sopsAgeKeys = [];
  }

  public addAgeKey(key: SecretValue) {
    this.sopsAgeKeys.push(key);
  }
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
    let sopsFileFormat: 'json' | 'yaml' | 'dotenv' | 'binary' | undefined =
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
        case 'binary': {
          sopsFileFormat = 'binary';
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
        if (props.secret) {
          props.secret.grantWrite(provider);
          props.secret.encryptionKey?.grantEncryptDecrypt(provider);
          if (props.secret?.encryptionKey !== undefined) {
            props.secret.encryptionKey.grantEncryptDecrypt(provider);
          }
        }
        if (props.parameterName) {
          provider.addToRolePolicy(
            new PolicyStatement({
              actions: ['ssm:PutParameter'],
              resources: [
                `arn:aws:ssm:${Stack.of(this).region}:${
                  Stack.of(this).account
                }:parameter${
                  props.parameterName.startsWith('/')
                    ? props.parameterName
                    : `/${props.parameterName}`
                }`,
              ],
            }),
          );
          props.encryptionKey?.grantEncryptDecrypt(provider);
        }
        if (props.parameterNames) {
          this.createReducedParameterPolicy(
            props.parameterNames,
            provider.role,
          );
          props.encryptionKey?.grantEncryptDecrypt(provider);
        }
        if (sopsAsset !== undefined) {
          sopsAsset.bucket.grantRead(provider);
        }
      } else {
        Annotations.of(this).addWarning(
          `Please ensure proper permissions for the passed lambda function:\n  - write Access to the secret\n  - encrypt with the sopsKmsKey${
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
        SecretARN: props.secret?.secretArn,
        SopsS3File: sopsS3File,
        SopsInline: sopsInline,
        ConvertToJSON: this.converToJSON,
        Flatten: this.flatten,
        FlattenSeparator: props.flattenSeparator ?? '.',
        ParameterKeyPrefix: props.parameterKeyPrefix,
        Format: sopsFileFormat,
        StringifiedValues: this.stringifiedValues,
        ParameterName: props.parameterName,
        EncryptionKey:
          props.secret !== undefined ? undefined : props.encryptionKey?.keyId,
        ResourceType: props.resourceType
          ? props.resourceType.toString()
          : ResourceType.SECRET.toString(),
        CreationType: props.creationType
          ? props.creationType.toString()
          : CreationType.SINGLE.toString(),
      },
    });
    this.versionId = cr.getAttString('VersionId');
  }

  private createReducedParameterPolicy(parameters: string[], role: IRole) {
    // Avoid too large policies
    // The maximum size of a managed policy is 6.144 bytes -> 1 character = 1 byte
    const maxPolicyBytes = 6000; // Keep some bytes as a buffer
    const arnPrefixBytes = 55; // Content for "arn:aws:ssm:ap-southeast-3:<accountnumer>:parameter/
    let startAtParameter = 0;
    let currentPolicyBytes = 300; // Reserve some byte space for basic stuff inside the policy
    for (let i = 0; i < parameters.length; i += 1) {
      if (
        // Check if the current parameter would fit into the policy
        arnPrefixBytes + parameters[i].length + currentPolicyBytes <
        maxPolicyBytes
      ) {
        // If so increase the byte counter
        currentPolicyBytes =
          arnPrefixBytes + parameters[i].length + currentPolicyBytes;
      } else {
        const parameterNamesChunk = parameters.slice(
          startAtParameter,
          i, //end of slice is not included
        );
        startAtParameter = i;
        currentPolicyBytes = 300;
        // Create the policy for the selected chunk
        const putPolicy = new ManagedPolicy(
          this,
          `SopsSecretParameterProviderManagedPolicyParameterAccess${i}`,
          {
            description:
              'Policy to grant parameter provider permissions to put parameter',
          },
        );
        putPolicy.addStatements(
          new PolicyStatement({
            actions: ['ssm:PutParameter'],
            resources: parameterNamesChunk.map(
              (param) =>
                `arn:aws:ssm:${Stack.of(this).region}:${
                  Stack.of(this).account
                }:parameter${param.startsWith('/') ? param : `/${param}`}`,
            ),
          }),
        );
        role.addManagedPolicy(putPolicy);
      }
    }
    const parameterNamesChunk = parameters.slice(
      startAtParameter,
      parameters.length,
    );
    // Create the policy for the remaning elements
    const putPolicy = new ManagedPolicy(
      this,
      `SopsSecretParameterProviderManagedPolicyParameterAccess${parameters.length}`,
      {
        description:
          'Policy to grant parameter provider permissions to put parameter',
      },
    );
    putPolicy.addStatements(
      new PolicyStatement({
        actions: ['ssm:PutParameter'],
        resources: parameterNamesChunk.map(
          (param) =>
            `arn:aws:ssm:${Stack.of(this).region}:${
              Stack.of(this).account
            }:parameter${param.startsWith('/') ? param : `/${param}`}`,
        ),
      }),
    );
    role.addManagedPolicy(putPolicy);
  }
}
