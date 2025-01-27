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

export enum ResourceType {
  SECRET = 'SECRET',
  PARAMETER = 'PARAMETER',
  PARAMETER_MULTI = 'PARAMETER_MULTI',
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

  /**
   * Remove any JSON/Yaml object structure and store this secret as a plaintext secret.
   * SOPS files must contain a JSON/Yaml object on the toplevel, they cannot store only a single string.
   * As a workaround, you can enable this flag and use the following format in the SOPS file:
   * `data: "secret value"`. Only the secret value will be synced to AWS Secrets Manager.
   */
  readonly plaintext?: boolean;

  /**
   * Should this construct automatically create IAM permissions?
   *
   * @default true
   */
  readonly autoGenerateIamPermissions?: boolean;
}

/**
 * The configuration options extended by the target Secret / Parameter
 */
export interface SopsSyncProps extends SopsSyncOptions {
  /**
   * The secret that will be populated with the encrypted sops file content.
   */
  readonly secret?: ISecret;

  /**
   * The parameter names. If set this creates encrypted SSM Parameters instead of a secret.
   */
  readonly parameterNames?: string[];

  /**
   * The encryption key used for encrypting the ssm parameter if `parameterName` is set.
   */
  readonly encryptionKey?: IKey;

  /**
   * Will this Sync deploy a Secret or Parameter(s)
   */
  readonly resourceType?: ResourceType;
}

/**
 * Configuration options for a custom SopsSyncProvider.
 */
export interface SopsSyncProviderProps {
  /**
   * VPC network to place Lambda network interfaces.
   *
   * @default - Lambda function is not placed within a VPC.
   */
  readonly vpc?: IVpc;
  /**
   * Where to place the network interfaces within the VPC.
   *
   * @default - Subnets will be chosen automatically.
   */
  readonly vpcSubnets?: SubnetSelection;
  /**
   * Only if `vpc` is supplied: The list of security groups to associate with the Lambda's network interfaces.
   *
   * @default - A dedicated security group will be created for the lambda function.
   */
  readonly securityGroups?: ISecurityGroup[];

  /**
   * The role that should be used for the custom resource provider.
   * If you don't specify any, a new role will be created with all required permissions
   *
   * @default - a new role will be created
   */
  readonly role?: IRole;
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
      role: props?.role,
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
      (props.sopsFilePath == undefined &&
        (props.sopsS3Bucket == undefined || props.sopsS3Key == undefined)) ||
      (props.sopsFilePath !== undefined &&
        props.sopsS3Bucket !== undefined &&
        props.sopsS3Key !== undefined)
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
      const sopsFileContent = fs.readFileSync(props.sopsFilePath);

      switch (uploadType) {
        case UploadType.INLINE: {
          sopsInline = {
            Content: sopsFileContent.toString('base64'),
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

      if (
        // Is allways true, but to satisfy TS we check explicitly
        provider.role !== undefined &&
        // Check if user has disabled automatic generation
        props.autoGenerateIamPermissions !== false
      ) {
        Permissions.sopsKeys(this, {
          userDefinedKeys: props.sopsKmsKey,
          role: provider.role,
          sopsFileContent: sopsFileContent.toString(),
        });
        Permissions.assetBucket(sopsAsset, provider.role);
        Permissions.encryptionKey(props.encryptionKey, provider.role);
        Permissions.secret(props.secret, provider.role);
        Permissions.parameters(this, props.parameterNames, provider.role);
      } else {
        Annotations.of(this).addWarning(
          [
            'Please ensure proper permissions for the passed lambda function:',
            '  - write Access to the secret/parameters',
            '  - encrypt with the sopsKmsKey',
            '  - download from asset bucket',
          ].join('\n'),
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
        Plaintext: props.plaintext ?? false,
        ParameterKeyPrefix: props.parameterKeyPrefix,
        Format: sopsFileFormat,
        StringifiedValues: this.stringifiedValues,
        ParameterName:
          // Dirty Workaround, refactor ...
          props.parameterNames && props.parameterNames.length == 1
            ? props.parameterNames[0]
            : undefined,
        EncryptionKey:
          props.secret !== undefined ? undefined : props.encryptionKey?.keyId,
        ResourceType: props.resourceType
          ? props.resourceType.toString()
          : ResourceType.SECRET.toString(),
      },
    });
    this.versionId = cr.getAttString('VersionId');
  }
}

export namespace Permissions {
  /**
   * Grants the necessary permissions for encrypt/decrypt on the customer managed encryption key
   * for the secrets / parameters.
   */
  export function encryptionKey(key: IKey | undefined, target: IGrantable) {
    if (key === undefined) {
      return;
    }
    key.grantEncryptDecrypt(target);
  }

  export function keysFromSopsContent(ctx: Construct, c: string): IKey[] {
    const regexKey = /arn:aws:kms:[a-z0-9-]+:[\d]+:key\/[a-z0-9-]+/g;
    const resultsKey = c.match(regexKey);
    if (resultsKey !== null) {
      return resultsKey.map((result, index) =>
        Key.fromKeyArn(ctx, `SopsKey${index}`, result),
      );
    }
    return [];
  }

  export function keysFromSopsContentAlias(ctx: Construct, c: string): IKey[] {
    const regexAlias = /arn:aws:kms:[a-z0-9-]+:[\d]+:alias\/[a-z0-9-]+/g;
    const resultsAlias = c.match(regexAlias);
    if (resultsAlias !== null) {
      return resultsAlias.map((result, index) =>
        Key.fromLookup(ctx, `SopsAlias${index}`, {
          aliasName: `alias/${result.split('/').slice(1).join('/')}`,
        }),
      );
    }
    return [];
  }

  /**
   * Grants the necessary permissions to decrypt the given sops file content.
   * Takes user defined keys, and searches the sops file for keys and aliases.
   */
  export function sopsKeys(
    ctx: Construct,
    props: {
      userDefinedKeys?: IKey[];
      sopsFileContent: string;
      role: IRole;
    },
  ) {
    (props.userDefinedKeys ?? [])
      .concat(
        keysFromSopsContent(ctx, props.sopsFileContent),
        keysFromSopsContentAlias(ctx, props.sopsFileContent),
      )
      .forEach((key) => key.grantDecrypt(props.role));
  }

  /**
   * Grants the necessary permissions to write the given secrets.
   */
  export function secret(
    targetSecret: ISecret | undefined,
    target: IGrantable,
  ) {
    if (targetSecret === undefined) {
      return;
    }
    targetSecret.grantWrite(target);
  }

  function sliceParameters(params: string[]): string[][] {
    const result: string[][] = [];
    /**
     * The maximum size of a managed policy is 6.144 bytes -> 1 character = 1 byte
     * bout 300 characters are reserved for the policy apart from resource arns
     * with some buffer, we end with an upper limit of 5750 bytes
     */
    const limit = 5750;

    /**
     * Content for "arn:aws:ssm:ap-southeast-3:<accountnumer>:parameter/
     */
    const prefix = 55;

    let currentSize = 0;
    let currentChunk: string[] = [];
    for (const param of params) {
      const paramLength = param.length + prefix;
      if (currentSize + paramLength > limit) {
        result.push(currentChunk);
        currentChunk = [];
        currentSize = 0;
      }
      currentChunk.push(param);
      currentSize += paramLength;
    }

    if (currentChunk.length > 0) {
      result.push(currentChunk);
    }
    return result;
  }

  /**
   * Grants the necessary permissions to write the given parameters.
   */
  export function parameters(
    ctx: Construct,
    targetParameters: string[] | undefined,
    role: IRole,
  ) {
    if (targetParameters === undefined) {
      return;
    }

    const paramSlices = sliceParameters(targetParameters);

    for (let i = 0; i < paramSlices.length; i++) {
      const putPolicy = new ManagedPolicy(
        ctx,
        `SopsSecretParameterProviderManagedPolicyParameterAccess${i}`,
        {
          description:
            'Policy to grant parameter provider permissions to put parameter',
        },
      );
      putPolicy.addStatements(
        new PolicyStatement({
          actions: ['ssm:PutParameter'],
          resources: paramSlices[i].map(
            (param) =>
              `arn:aws:ssm:${Stack.of(ctx).region}:${
                Stack.of(ctx).account
              }:parameter${param.startsWith('/') ? param : `/${param}`}`,
          ),
        }),
      );
      role.addManagedPolicy(putPolicy);
    }
  }

  /**
   * Grants the necessary permissions to read the given asset from S3.
   */
  export function assetBucket(asset: Asset | undefined, target: IGrantable) {
    if (asset === undefined) {
      return;
    }
    asset.bucket.grantRead(target);
  }
}
