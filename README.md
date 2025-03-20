<img src="https://github.com/dbsystel/cdk-sops-secrets/blob/main/img/banner-dl-small.png?raw=true">

![stability](https://img.shields.io/badge/Stability-stable-green)
[![release](https://github.com/dbsystel/cdk-sops-secrets/actions/workflows/release.yml/badge.svg)](https://github.com/dbsystel/cdk-sops-secrets/actions/workflows/release.yml)
[![cdk-construct-hub](https://img.shields.io/badge/CDK-ConstructHub-blue)](https://constructs.dev/packages/cdk-sops-secrets)
[![npm](https://img.shields.io/npm/v/cdk-sops-secrets.svg)](https://www.npmjs.com/package/cdk-sops-secrets)
[![npm downloads](https://img.shields.io/npm/dw/cdk-sops-secrets)](https://www.npmjs.com/package/cdk-sops-secrets)
[![pypi](https://img.shields.io/pypi/v/cdk-sops-secrets.svg)](https://pypi.org/project/cdk-sops-secrets)
[![pypi downloads](https://img.shields.io/pypi/dw/cdk-sops-secrets)](https://pypi.org/project/cdk-sops-secrets)
[![codecov](https://codecov.io/gh/dbsystel/cdk-sops-secrets/branch/main/graph/badge.svg?token=OT7P7HQHXB)](https://codecov.io/gh/dbsystel/cdk-sops-secrets)
[![security-vulnerabilities](https://img.shields.io/github/issues-search/dbsystel/cdk-sops-secrets?color=%23ff0000&label=security-vulnerabilities&query=is%3Aissue%20is%3Aopen%20label%3A%22Mend%3A%20dependency%20security%20vulnerability%22)](https://github.com/dbsystel/cdk-sops-secrets/issues?q=is%3Aissue+is%3Aopen+label%3A%22security+vulnerability%22)

# Introduction

_Create secret values in AWS with infrastructure-as-code easily_

This construct library offers CDK Constructs that facilitate syncing [SOPS-encrypted secrets](https://github.com/getsops/sops) to AWS Secrets Manager and SSM Parameter Store.
It enables secure storage of secrets in Git repositories while allowing seamless synchronization and usage within AWS. Even large sets of SSM Parameters can be created quickly from a single file.

- Create AWS Secrets Manager secrets
- Create single SSM Parameter
- Create multiple SSM Parameter in a batch from a file
- Use SOPS json, yaml or dotenv as input files, as well as binary data
- No need for manual permission setups for the Custom Ressource due to automatic least-privilege generation for the SyncProvider

# Table Of Contents

- [Introduction](#introduction)
- [Table Of Contents](#table-of-contents)
- [Available Constructs](#available-constructs)
  - [SopsSecret — Sops to SecretsManager](#sopssecret--sops-to-secretsmanager)
  - [SopsStringParameter — Sops to single SSM ParameterStore Parameter](#sopsstringparameter--sops-to-single-ssm-parameterstore-parameter)
  - [MultiStringParameter — Sops to multiple SSM ParameterStore Parameters](#multistringparameter--sops-to-multiple-ssm-parameterstore-parameters)
  - [SopsSyncProvider](#sopssyncprovider)
  - [Common configuration options for SopsSecret, SopsStringParameter and MultiStringParameter](#common-configuration-options-for-sopssecret-sopsstringparameter-and-multistringparameter)
- [Considerations](#considerations)
  - [UploadType: INLINE / ASSET](#uploadtype-inline--asset)
  - [Stability](#stability)
- [FAQ](#faq)
  - [How can I migrate to V2](#how-can-i-migrate-to-v2)
    - [SecretsManager](#secretsmanager)
    - [Parameter](#parameter)
    - [MultiParameter](#multiparameter)
  - [It does not work, what can I do?](#it-does-not-work-what-can-i-do)
  - [I get errors with `dotenv` formatted files](#i-get-errors-with-dotenv-formatted-files)
  - [Error: Error getting data key: 0 successful groups required, got 0](#error-error-getting-data-key-0-successful-groups-required-got-0)
  - [Error: Asset of sync lambda not found](#error-asset-of-sync-lambda-not-found)
  - [Can I upload the sops file myself and provide the required information as CloudFormation Parameter?](#can-i-upload-the-sops-file-myself-and-provide-the-required-information-as-cloudformation-parameter)
  - [Can I access older versions of the secret stored in the SecretsManager?](#can-i-access-older-versions-of-the-secret-stored-in-the-secretsmanager)
  - [I want the `raw` content of the sops file, but I always get the content nested in json](#i-want-the-raw-content-of-the-sops-file-but-i-always-get-the-content-nested-in-json)
- [License](#license)

# Available Constructs

The construct library cdk-sops-secrets supports three different Constructs that help you to sync your encrypted sops secrets to secure places in AWS.

Let's assume we want to store the following secret information in AWS:

```json
{
  "apiKey": "sk-1234567890abcdef",
  "database": {
    "user": "admin",
    "password": "P@ssw0rd!",
    "host": "db.example.com"
  },
  "tokens": [
    { "service": "github", "token": "ghp_abcd1234" },
    { "service": "aws", "token": "AKIAIOSFODNN7EXAMPLE" }
  ],
  "someOtherKey": "base64:VGhpcyBpcyBhIHNlY3JldCBrZXk="
}
```

It doesn't matter if this data is in `json`, `yaml` or `dotenv` format, `cdk-sops-secret` can handle them all.
Even binary data is supported with some limitations.

## SopsSecret — Sops to SecretsManager

If you want to store your secret data in the AWS SecretsManager, use the `SopsSecret` construct. This is a "drop-in-replacement" for the [Secret Construct](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_secretsmanager.Secret.html) of the AWS CDK.

Minimal Example:

```ts
const secret = new SopsSecret(stack, 'MySopsSecret', {
  secertName: 'mySecret', // name of the secret in AWS SecretsManager
  sopsFilePath: 'secrets/sopsfile-encrypted-secret.json', // filepath to the sops encrypted file
});
```

The content referenced sops secret file will be synced to the AWS SecretsManager Secret with the name `mySecret`.
For convenience, several transformations apply:

- Nested structures and arrays will be resolved and flattened to a JSONPath notation
- All values will be stored as strings

This is done also because of limitations of CDK in conjunction with
[dynamic references](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references-secretsmanager.html) and limitiations
of the `Key/Value` view of the AWS SecretsManager WebConsole. So the result, saved in the AWS SecretsManager will actually be:

```json
{
  "apiKey": "sk-1234567890abcdef",
  "database.user": "admin",
  "database.password": "P@ssw0rd!",
  "database.host": "db.example.com",
  "tokens[0].service": "github",
  "tokens[0].token": "ghp_abcd1234",
  "tokens[1].service": "aws",
  "tokens[1].token": "AKIAIOSFODNN7EXAMPLE",
  "someOtherKey": "base64:VGhpcyBpcyBhIHNlY3JldCBrZXk="
}
```

This allows you to access the values from your secret via CDK:

```ts
secret.secretValueFromJson('"database.password"').toString(),
  secret.secretValueFromJson('"tokens[0].token"').toString();
```

If you don't want these conversions, you can completely disable them by using the `rawOutput` property.

```ts
const secret = new SopsSecret(stack, 'MySopsSecret', {
  rawOutput: RawOutput.STRING,
  ...
});
```

This will turn off the conversions and just place the decrypted content in the target secret. It's also possible to use
`RawOutput.BINARY` than the AWS SecretsManager Secret will be populted with binary, instead of string data.

## SopsStringParameter — Sops to single SSM ParameterStore Parameter

If you want to sync the whole content of a sops encrypted file to an encrypted AWS SSM ParameterStore Parameter, you can use the SopsStringParameter Construct.

```ts
const parameter = new SopsStringParameter(stack, 'MySopsParameter', {
  encryptionKey: Key.fromLookup(stack, 'DefaultKey', {
    aliasName: 'alias/aws/ssm',
  }),
  sopsFilePath: 'secrets/sopsfile-encrypted-secret.json',
});
```

This will create a Parameter with the value of the decrypted sops file content. No transformations are applied.

## MultiStringParameter — Sops to multiple SSM ParameterStore Parameters

If you have a structured sops file (yaml, json, dotenv) and want to populate the AWS SSM ParameterStore with it, you want to use the MultiStringParameter Construct.

```ts
const multi = new MultiStringParameter(stack, 'MyMultiParameter', {
  encryptionKey: Key.fromLookup(stack, 'DefaultKey', {
    aliasName: 'alias/aws/ssm',
  }),
  sopsFilePath: 'secrets/sopsfile-encrypted-secret.json',
});
```

This will create several AWS SSM ParameterStore Parameters:

```bash
ParameterName       => Value

/apiKey             => "sk-1234567890abcdef"
/database/user      => "admin"
/database/password  => "P@ssw0rd!"
/database/host      => "db.example.com"
/tokens/0/service   => "github"
/tokens/0/token     => "ghp_abcd1234"
/tokens/1/service   => "aws"
/tokens/1/token     => "AKIAIOSFODNN7EXAMPLE"
/someOtherKey       => "base64:VGhpcyBpcyBhIHNlY3JldCBrZXk="
```

You can configure the naming schema via the properties `keySeperator` and `keyPrefix`:

```ts
const multi = new MultiStringParameter(stack, 'MyMultiParameter', {
  keyPrefix: 'mykeyprefix.'  // All keys will start with this string, default '/'
  keySeperator: '-'         // This seperator is used when converting to a flat structure, default '/'
})
```

This would lead to Parameters

```bash
ParameterName       => Value

mykeyprefix.apiKey             => "sk-1234567890abcdef"
mykeyprefix.database-user      => "admin"
mykeyprefix.tokens-0-service   => "github"
...
```

## SopsSyncProvider

The SOPS-Provider is the custom resource AWS Lambda Function, that is doing all the work. It downloads, decrypts
and stores the secret content in your desired location. This Lambda Function needs several IAM permissions to do it's work.

For most use cases, you don't need to create it on your own, as the other Constructs try to create this and derive the required IAM permissions from your input.

But there are use cases, that require you to change the defaults of this Provider. If this is the case,
you have to create the provider on your own and add it to the other constructs.

```ts
const provider = new SopsSyncProvider(this, 'MySopsSyncProvider', {
  role: customRole,       // you can pass a custom role

  vpc: customVpc,         // The default SopsSync Provider
  vpcSubnets: {           // won't run in any VPC,
    subnets: [            // as it does not require
      customSubnet1,      // access to any VPC resources.
      customSubnet2,      // But if you want,
    ]                     // you can change this behaviour
  },                      // and set vpc, subnet and
  securityGroups: [       // securitygroups to your
    customSecurityGroup   // needs.
  ],
  logRetention: RetentionDays.INFINITE,  // you can increase the default log retention
});

provider.addToRolePolicy( // You cann pass PolicyStatements
  new PolicyStatement({   // via the addToRolePolicy Method
    actions: ['...'],     //
    resources: ['...'],   //
  })                      //
);                        //

kmsKey.grantDecrypt(      // The provider implements
  provider                // the IGrantable interface,
);                        // so you can use it as grant target

const secret = new SopsSecret(this, 'MySecret', {
  sopsProvider: provider, // this property is available in all Constructs
  ...
});
```

## Common configuration options for SopsSecret, SopsStringParameter and MultiStringParameter

```ts

const construct = new Sops...(this, 'My' {
  /**
   * use your own SopsSyncProvider
   * @see SopsSyncProvider
   */
  sopsProvider: myCustomProvider      // default - a new provider will be created

  /**
   * the constructs try to derive the required iam permissions from the sops file
   * and the target. If you don't want this, you can disable this behaviour.
   * You have to take care of all required permissions on your own.
   */
  autoGenerateIamPermissions: false,  // default: true

  /**
   * the default behaviour of passing the sops file content to the provider is
   * by embedding the base64 encoded content in the cloudformation template.
   * Using CKD Assets is also supported. It might be required to switch to
   * Assets, if your sops files are very large.
   */
  uploadType: UploadType.ASSET,       // default: UploadType.INLINE

  /**
   * if you don't want this constructs to take care of passing the encrypted
   * sops file to the sops provider, you can upload them yourself to a
   * S3 bucket.
   * You can pass bucket and key, and the constructs won't pass the content
   * as ASSET or in the CloudFormation Template.
   * As the construct isn't aware of the sopsfile, we can't derive the required
   * permissions to decrypt the sops file. The same applies to the sopsFileFormat.
   * You have to pass them all manually.
   */
  sopsS3Bucket: 'my-custom-bucket',
  sopsS3Key: 'encoded-sops.json',
  sopsKmsKey: [
    kmsKeyUsedForEncryption,
  ]
  sopsFileFormat: 'json',   // Allowed values are json, yaml, dotenv and binary
})

```

# Considerations

## UploadType: INLINE / ASSET

I decided, that the default behavior should be "INLINE" because of the following consideration:

- Fewer permissions

  _If we use inline content instead of a S3 asset, the SopsSyncProvider does not need permissions to access the asset bucket and its KMS key._

- Faster

  _If we don't have to upload and download things from and to S3, it should be a little faster._

- Interchangeable

  _As we use the same information to generate the version of the secret,
  no new version of the secret should be created, if you change from INLINE to ASSET or vice versa,
  even if the CloudFormation resource updates._

## Stability

You can consider this package as stable. Updates will follow [Semantic Versioning](https://semver.org/).

Nevertheless, I would recommend pinning the exact version of this library in your `package.json`.

# FAQ

## How can I migrate to V2

It was required to change some user facing configuration properties. So minor changes are required to make things work again.

### SecretsManager

- Removed property convertToJSON, flatten, stringifiedValues
- Use property rawOutput instaed:
  - `undefined / not set` => (default) convertToJSON and flatten and stringifiedValues = true
  - `RawOutput.STRING` => convertToJSON and flatten and stringifiedValues = false
  - `RawOutput.BINARY` => convertToJSON and flatten and stringifiedValues = false and Secret is binary

### Parameter

- Removed property convertToJSON, flatten, stringifiedValues => all of them made no sense - now only raw output of decrypted secret

### MultiParameter

- Removed property convertToJSON, flatten, stringifiedValues => most of this combinations made no sense
- Allways convertToJson and flatten (as we have to parse it to create multiple parameters)
- You are allowed to chose the flattenSeperator

## It does not work, what can I do?

Even if this construct has some unit and integration tests performed, there can be bugs and issues. As everything is performed by a cloudformation custom resource provider, a good starting point is the log of the corresponding lambda function. It should be located in your AWS Account under Cloudwatch -> Log groups:

`/aws/lambda/<YOUR-STACK-NAME>-SingletonLambdaSopsSyncProvider<SOMETHINGsomething1234>`

## I get errors with `dotenv` formatted files

Only very basic dotenv syntax is working right now. Only single line values are accepted. The format must match:

```dotenv
key=value
```

comments must be a single line, not after value assignments.

## Error: Error getting data key: 0 successful groups required, got 0

This error message (and failed sync) is related to the getsops/sops issues [#948](https://github.com/getsops/sops/issues/948) and [#634](https://github.com/getsops/sops/issues/634). You must not create your secret with the `--aws-profile` flag. This profile will be written to your sops filed and is required in every runtime environment. You have to define the profile to use via the environment variable `AWS_PROFILE` instead, to avoid this.

## Error: Asset of sync lambda not found

The lambda asset code is generated relative to the path of the index.ts in this package. With tools like nx this can lead to wrong results, so that the asset could not be found.

You can override the asset path via the [cdk.json](https://docs.aws.amazon.com/cdk/v2/guide/get_context_var.html) or via the flag `-c`of the cdk cli.

The context used for this override is `sops_sync_provider_asset_path`.

So for example you can use

```bash
cdk deploy -c "sops_sync_provider_asset_path=some/path/asset.zip"
```

or in your cdk.json

```json
{
  "context": {
    "sops_sync_provider_asset_path": "some/path/asset.zip"
  }
}
```

## Can I upload the sops file myself and provide the required information as CloudFormation Parameter?

This should be possible the following way. Ensure, that you have created a custom sops provider,
with proper IAM permissions.

```typescript
const sopsS3BucketParam = new CfnParameter(this, "s3BucketName", {
  type: "String",
  description: "The name of the Amazon S3 bucket where your sopsFile was uploaded."});

const sopsS3KeyParam = new CfnParameter(this, "s3KeyName", {
  type: "String",
  description: "The name of the key of the sopsFile inside the Amazon S3 bucket."});

const sopsKmsKeyArn = new CfnParameter(this, "sopsKeyArn", {
  type: "String",
  description: "The ARN of the KMS Key used for sops encryption"});

const sopsKmsKey = Key.fromKeyArn(this, 'Key', sopsKmsKeyArn.valueAsString)

new SopsSecret(stack, 'SopsSecret', {
  sopsS3Bucket: sopsS3BucketParam.valueAsString,
  sopsS3Key: sopsS3KeyParam.valueAsString,
  sopsKmsKey: [
    sopsKmsKey
  ],
  sopsFileFormat: 'json',
  ...
});
```

## Can I access older versions of the secret stored in the SecretsManager?

While creating the secret or updating the entries of a secret, the native CDK function `cdk.FileSystem.fingerprint(...)` is used
to generate the version information of the AWS SecretsManager secret.
Therefore, it is possible to reference the entries from a specific AWS SecretsManager version.

```typescript
const versionId = cdk.FileSystem.fingerprint(`./sops/SomeSecrets.json`);
const passphrase = ecs.Secret.fromSecretsManagerVersion(
  secretMgmt,
  { versionId: versionId },
  'MY_PRIVATE_PASSPHRASE',
);

const container = TaskDef.addContainer('Container', {
  secrets: {
    MY_PRIVATE_PASSPHRASE: passphrase,
  },
});
```

## I want the `raw` content of the sops file, but I always get the content nested in json

To get the best raw experience, you should encrypt your sops files in binary format:

```bash
sops encrypt ... my-whatever-file --output my-secret-information.sops.binary --input-type binary
```

You will lose features like only encrypting the values, not the keys.
The whole file content will be stored in the sops file.
You can store everything you like as binary, even binary data[^1].

When using binary encrypted secrets with this constructs, ensure the ending is also binary, or override via
`sopsFormat` property.

This does not work for `MultiStringParameter`

[^1] Even if sops can handle binary data, only the AWS SecretsManager allows to store it.

# License

The Apache-2.0 license. Please have a look at the [LICENSE](LICENSE) and [LICENSE-3RD-PARTY](LICENSE-3RD-PARTY).
