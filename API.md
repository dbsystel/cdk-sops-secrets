# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### SopsSecrets <a name="SopsSecrets" id="cdk-sops-secrets.SopsSecrets"></a>

#### Initializers <a name="Initializers" id="cdk-sops-secrets.SopsSecrets.Initializer"></a>

```typescript
import { SopsSecrets } from 'cdk-sops-secrets'

new SopsSecrets(scope: Construct, id: string, props: SopsSecretProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSecrets.Initializer.parameter.scope">scope</a></code> | <code>@aws-cdk/core.Construct</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSecrets.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSecrets.Initializer.parameter.props">props</a></code> | <code><a href="#cdk-sops-secrets.SopsSecretProps">SopsSecretProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="cdk-sops-secrets.SopsSecrets.Initializer.parameter.scope"></a>

- *Type:* @aws-cdk/core.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="cdk-sops-secrets.SopsSecrets.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="cdk-sops-secrets.SopsSecrets.Initializer.parameter.props"></a>

- *Type:* <a href="#cdk-sops-secrets.SopsSecretProps">SopsSecretProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSecrets.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#cdk-sops-secrets.SopsSecrets.applyRemovalPolicy">applyRemovalPolicy</a></code> | Apply the given removal policy to this resource. |
| <code><a href="#cdk-sops-secrets.SopsSecrets.addReplicaRegion">addReplicaRegion</a></code> | Adds a replica region for the secret. |
| <code><a href="#cdk-sops-secrets.SopsSecrets.addRotationSchedule">addRotationSchedule</a></code> | Adds a rotation schedule to the secret. |
| <code><a href="#cdk-sops-secrets.SopsSecrets.addTargetAttachment">addTargetAttachment</a></code> | Adds a target attachment to the secret. |
| <code><a href="#cdk-sops-secrets.SopsSecrets.addToResourcePolicy">addToResourcePolicy</a></code> | Adds a statement to the IAM resource policy associated with this secret. |
| <code><a href="#cdk-sops-secrets.SopsSecrets.attach">attach</a></code> | Attach a target to this secret. |
| <code><a href="#cdk-sops-secrets.SopsSecrets.denyAccountRootDelete">denyAccountRootDelete</a></code> | Denies the `DeleteSecret` action to all principals within the current account. |
| <code><a href="#cdk-sops-secrets.SopsSecrets.grantRead">grantRead</a></code> | Grants reading the secret value to some role. |
| <code><a href="#cdk-sops-secrets.SopsSecrets.grantWrite">grantWrite</a></code> | Grants writing and updating the secret value to some role. |
| <code><a href="#cdk-sops-secrets.SopsSecrets.secretValueFromJson">secretValueFromJson</a></code> | Interpret the secret as a JSON object and return a field's value from it as a `SecretValue`. |

---

##### `toString` <a name="toString" id="cdk-sops-secrets.SopsSecrets.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="cdk-sops-secrets.SopsSecrets.applyRemovalPolicy"></a>

```typescript
public applyRemovalPolicy(policy: RemovalPolicy): void
```

Apply the given removal policy to this resource.

The Removal Policy controls what happens to this resource when it stops
being managed by CloudFormation, either because you've removed it from the
CDK application or because you've made a change that requires the resource
to be replaced.

The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).

###### `policy`<sup>Required</sup> <a name="policy" id="cdk-sops-secrets.SopsSecrets.applyRemovalPolicy.parameter.policy"></a>

- *Type:* @aws-cdk/core.RemovalPolicy

---

##### `addReplicaRegion` <a name="addReplicaRegion" id="cdk-sops-secrets.SopsSecrets.addReplicaRegion"></a>

```typescript
public addReplicaRegion(region: string, encryptionKey?: IKey): void
```

Adds a replica region for the secret.

###### `region`<sup>Required</sup> <a name="region" id="cdk-sops-secrets.SopsSecrets.addReplicaRegion.parameter.region"></a>

- *Type:* string

The name of the region.

---

###### `encryptionKey`<sup>Optional</sup> <a name="encryptionKey" id="cdk-sops-secrets.SopsSecrets.addReplicaRegion.parameter.encryptionKey"></a>

- *Type:* @aws-cdk/aws-kms.IKey

The customer-managed encryption key to use for encrypting the secret value.

---

##### `addRotationSchedule` <a name="addRotationSchedule" id="cdk-sops-secrets.SopsSecrets.addRotationSchedule"></a>

```typescript
public addRotationSchedule(id: string, options: RotationScheduleOptions): RotationSchedule
```

Adds a rotation schedule to the secret.

###### `id`<sup>Required</sup> <a name="id" id="cdk-sops-secrets.SopsSecrets.addRotationSchedule.parameter.id"></a>

- *Type:* string

---

###### `options`<sup>Required</sup> <a name="options" id="cdk-sops-secrets.SopsSecrets.addRotationSchedule.parameter.options"></a>

- *Type:* @aws-cdk/aws-secretsmanager.RotationScheduleOptions

---

##### ~~`addTargetAttachment`~~ <a name="addTargetAttachment" id="cdk-sops-secrets.SopsSecrets.addTargetAttachment"></a>

```typescript
public addTargetAttachment(id: string, options: AttachedSecretOptions): SecretTargetAttachment
```

Adds a target attachment to the secret.

###### `id`<sup>Required</sup> <a name="id" id="cdk-sops-secrets.SopsSecrets.addTargetAttachment.parameter.id"></a>

- *Type:* string

---

###### `options`<sup>Required</sup> <a name="options" id="cdk-sops-secrets.SopsSecrets.addTargetAttachment.parameter.options"></a>

- *Type:* @aws-cdk/aws-secretsmanager.AttachedSecretOptions

---

##### `addToResourcePolicy` <a name="addToResourcePolicy" id="cdk-sops-secrets.SopsSecrets.addToResourcePolicy"></a>

```typescript
public addToResourcePolicy(statement: PolicyStatement): AddToResourcePolicyResult
```

Adds a statement to the IAM resource policy associated with this secret.

If this secret was created in this stack, a resource policy will be
automatically created upon the first call to `addToResourcePolicy`. If
the secret is imported, then this is a no-op.

###### `statement`<sup>Required</sup> <a name="statement" id="cdk-sops-secrets.SopsSecrets.addToResourcePolicy.parameter.statement"></a>

- *Type:* @aws-cdk/aws-iam.PolicyStatement

---

##### `attach` <a name="attach" id="cdk-sops-secrets.SopsSecrets.attach"></a>

```typescript
public attach(target: ISecretAttachmentTarget): ISecret
```

Attach a target to this secret.

###### `target`<sup>Required</sup> <a name="target" id="cdk-sops-secrets.SopsSecrets.attach.parameter.target"></a>

- *Type:* @aws-cdk/aws-secretsmanager.ISecretAttachmentTarget

The target to attach.

---

##### `denyAccountRootDelete` <a name="denyAccountRootDelete" id="cdk-sops-secrets.SopsSecrets.denyAccountRootDelete"></a>

```typescript
public denyAccountRootDelete(): void
```

Denies the `DeleteSecret` action to all principals within the current account.

##### `grantRead` <a name="grantRead" id="cdk-sops-secrets.SopsSecrets.grantRead"></a>

```typescript
public grantRead(grantee: IGrantable, versionStages?: string[]): Grant
```

Grants reading the secret value to some role.

###### `grantee`<sup>Required</sup> <a name="grantee" id="cdk-sops-secrets.SopsSecrets.grantRead.parameter.grantee"></a>

- *Type:* @aws-cdk/aws-iam.IGrantable

---

###### `versionStages`<sup>Optional</sup> <a name="versionStages" id="cdk-sops-secrets.SopsSecrets.grantRead.parameter.versionStages"></a>

- *Type:* string[]

---

##### `grantWrite` <a name="grantWrite" id="cdk-sops-secrets.SopsSecrets.grantWrite"></a>

```typescript
public grantWrite(grantee: IGrantable): Grant
```

Grants writing and updating the secret value to some role.

###### `grantee`<sup>Required</sup> <a name="grantee" id="cdk-sops-secrets.SopsSecrets.grantWrite.parameter.grantee"></a>

- *Type:* @aws-cdk/aws-iam.IGrantable

---

##### `secretValueFromJson` <a name="secretValueFromJson" id="cdk-sops-secrets.SopsSecrets.secretValueFromJson"></a>

```typescript
public secretValueFromJson(jsonField: string): SecretValue
```

Interpret the secret as a JSON object and return a field's value from it as a `SecretValue`.

###### `jsonField`<sup>Required</sup> <a name="jsonField" id="cdk-sops-secrets.SopsSecrets.secretValueFromJson.parameter.jsonField"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSecrets.isConstruct">isConstruct</a></code> | Return whether the given object is a Construct. |
| <code><a href="#cdk-sops-secrets.SopsSecrets.isResource">isResource</a></code> | Check whether the given construct is a Resource. |
| <code><a href="#cdk-sops-secrets.SopsSecrets.fromSecretArn">fromSecretArn</a></code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSecrets.fromSecretAttributes">fromSecretAttributes</a></code> | Import an existing secret into the Stack. |
| <code><a href="#cdk-sops-secrets.SopsSecrets.fromSecretCompleteArn">fromSecretCompleteArn</a></code> | Imports a secret by complete ARN. |
| <code><a href="#cdk-sops-secrets.SopsSecrets.fromSecretName">fromSecretName</a></code> | Imports a secret by secret name; |
| <code><a href="#cdk-sops-secrets.SopsSecrets.fromSecretNameV2">fromSecretNameV2</a></code> | Imports a secret by secret name. |
| <code><a href="#cdk-sops-secrets.SopsSecrets.fromSecretPartialArn">fromSecretPartialArn</a></code> | Imports a secret by partial ARN. |

---

##### `isConstruct` <a name="isConstruct" id="cdk-sops-secrets.SopsSecrets.isConstruct"></a>

```typescript
import { SopsSecrets } from 'cdk-sops-secrets'

SopsSecrets.isConstruct(x: any)
```

Return whether the given object is a Construct.

###### `x`<sup>Required</sup> <a name="x" id="cdk-sops-secrets.SopsSecrets.isConstruct.parameter.x"></a>

- *Type:* any

---

##### `isResource` <a name="isResource" id="cdk-sops-secrets.SopsSecrets.isResource"></a>

```typescript
import { SopsSecrets } from 'cdk-sops-secrets'

SopsSecrets.isResource(construct: IConstruct)
```

Check whether the given construct is a Resource.

###### `construct`<sup>Required</sup> <a name="construct" id="cdk-sops-secrets.SopsSecrets.isResource.parameter.construct"></a>

- *Type:* @aws-cdk/core.IConstruct

---

##### ~~`fromSecretArn`~~ <a name="fromSecretArn" id="cdk-sops-secrets.SopsSecrets.fromSecretArn"></a>

```typescript
import { SopsSecrets } from 'cdk-sops-secrets'

SopsSecrets.fromSecretArn(scope: Construct, id: string, secretArn: string)
```

###### `scope`<sup>Required</sup> <a name="scope" id="cdk-sops-secrets.SopsSecrets.fromSecretArn.parameter.scope"></a>

- *Type:* constructs.Construct

---

###### `id`<sup>Required</sup> <a name="id" id="cdk-sops-secrets.SopsSecrets.fromSecretArn.parameter.id"></a>

- *Type:* string

---

###### `secretArn`<sup>Required</sup> <a name="secretArn" id="cdk-sops-secrets.SopsSecrets.fromSecretArn.parameter.secretArn"></a>

- *Type:* string

---

##### `fromSecretAttributes` <a name="fromSecretAttributes" id="cdk-sops-secrets.SopsSecrets.fromSecretAttributes"></a>

```typescript
import { SopsSecrets } from 'cdk-sops-secrets'

SopsSecrets.fromSecretAttributes(scope: Construct, id: string, attrs: SecretAttributes)
```

Import an existing secret into the Stack.

###### `scope`<sup>Required</sup> <a name="scope" id="cdk-sops-secrets.SopsSecrets.fromSecretAttributes.parameter.scope"></a>

- *Type:* constructs.Construct

the scope of the import.

---

###### `id`<sup>Required</sup> <a name="id" id="cdk-sops-secrets.SopsSecrets.fromSecretAttributes.parameter.id"></a>

- *Type:* string

the ID of the imported Secret in the construct tree.

---

###### `attrs`<sup>Required</sup> <a name="attrs" id="cdk-sops-secrets.SopsSecrets.fromSecretAttributes.parameter.attrs"></a>

- *Type:* @aws-cdk/aws-secretsmanager.SecretAttributes

the attributes of the imported secret.

---

##### `fromSecretCompleteArn` <a name="fromSecretCompleteArn" id="cdk-sops-secrets.SopsSecrets.fromSecretCompleteArn"></a>

```typescript
import { SopsSecrets } from 'cdk-sops-secrets'

SopsSecrets.fromSecretCompleteArn(scope: Construct, id: string, secretCompleteArn: string)
```

Imports a secret by complete ARN.

The complete ARN is the ARN with the Secrets Manager-supplied suffix.

###### `scope`<sup>Required</sup> <a name="scope" id="cdk-sops-secrets.SopsSecrets.fromSecretCompleteArn.parameter.scope"></a>

- *Type:* constructs.Construct

---

###### `id`<sup>Required</sup> <a name="id" id="cdk-sops-secrets.SopsSecrets.fromSecretCompleteArn.parameter.id"></a>

- *Type:* string

---

###### `secretCompleteArn`<sup>Required</sup> <a name="secretCompleteArn" id="cdk-sops-secrets.SopsSecrets.fromSecretCompleteArn.parameter.secretCompleteArn"></a>

- *Type:* string

---

##### ~~`fromSecretName`~~ <a name="fromSecretName" id="cdk-sops-secrets.SopsSecrets.fromSecretName"></a>

```typescript
import { SopsSecrets } from 'cdk-sops-secrets'

SopsSecrets.fromSecretName(scope: Construct, id: string, secretName: string)
```

Imports a secret by secret name;

the ARN of the Secret will be set to the secret name.
A secret with this name must exist in the same account & region.

###### `scope`<sup>Required</sup> <a name="scope" id="cdk-sops-secrets.SopsSecrets.fromSecretName.parameter.scope"></a>

- *Type:* constructs.Construct

---

###### `id`<sup>Required</sup> <a name="id" id="cdk-sops-secrets.SopsSecrets.fromSecretName.parameter.id"></a>

- *Type:* string

---

###### `secretName`<sup>Required</sup> <a name="secretName" id="cdk-sops-secrets.SopsSecrets.fromSecretName.parameter.secretName"></a>

- *Type:* string

---

##### `fromSecretNameV2` <a name="fromSecretNameV2" id="cdk-sops-secrets.SopsSecrets.fromSecretNameV2"></a>

```typescript
import { SopsSecrets } from 'cdk-sops-secrets'

SopsSecrets.fromSecretNameV2(scope: Construct, id: string, secretName: string)
```

Imports a secret by secret name.

A secret with this name must exist in the same account & region.
Replaces the deprecated `fromSecretName`.

###### `scope`<sup>Required</sup> <a name="scope" id="cdk-sops-secrets.SopsSecrets.fromSecretNameV2.parameter.scope"></a>

- *Type:* constructs.Construct

---

###### `id`<sup>Required</sup> <a name="id" id="cdk-sops-secrets.SopsSecrets.fromSecretNameV2.parameter.id"></a>

- *Type:* string

---

###### `secretName`<sup>Required</sup> <a name="secretName" id="cdk-sops-secrets.SopsSecrets.fromSecretNameV2.parameter.secretName"></a>

- *Type:* string

---

##### `fromSecretPartialArn` <a name="fromSecretPartialArn" id="cdk-sops-secrets.SopsSecrets.fromSecretPartialArn"></a>

```typescript
import { SopsSecrets } from 'cdk-sops-secrets'

SopsSecrets.fromSecretPartialArn(scope: Construct, id: string, secretPartialArn: string)
```

Imports a secret by partial ARN.

The partial ARN is the ARN without the Secrets Manager-supplied suffix.

###### `scope`<sup>Required</sup> <a name="scope" id="cdk-sops-secrets.SopsSecrets.fromSecretPartialArn.parameter.scope"></a>

- *Type:* constructs.Construct

---

###### `id`<sup>Required</sup> <a name="id" id="cdk-sops-secrets.SopsSecrets.fromSecretPartialArn.parameter.id"></a>

- *Type:* string

---

###### `secretPartialArn`<sup>Required</sup> <a name="secretPartialArn" id="cdk-sops-secrets.SopsSecrets.fromSecretPartialArn.parameter.secretPartialArn"></a>

- *Type:* string

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSecrets.property.node">node</a></code> | <code>@aws-cdk/core.ConstructNode</code> | The construct tree node associated with this construct. |
| <code><a href="#cdk-sops-secrets.SopsSecrets.property.env">env</a></code> | <code>@aws-cdk/core.ResourceEnvironment</code> | The environment this resource belongs to. |
| <code><a href="#cdk-sops-secrets.SopsSecrets.property.stack">stack</a></code> | <code>@aws-cdk/core.Stack</code> | The stack in which this resource is defined. |
| <code><a href="#cdk-sops-secrets.SopsSecrets.property.secretArn">secretArn</a></code> | <code>string</code> | The ARN of the secret in AWS Secrets Manager. |
| <code><a href="#cdk-sops-secrets.SopsSecrets.property.secretName">secretName</a></code> | <code>string</code> | The name of the secret. |
| <code><a href="#cdk-sops-secrets.SopsSecrets.property.secretValue">secretValue</a></code> | <code>@aws-cdk/core.SecretValue</code> | Retrieve the value of the stored secret as a `SecretValue`. |
| <code><a href="#cdk-sops-secrets.SopsSecrets.property.encryptionKey">encryptionKey</a></code> | <code>@aws-cdk/aws-kms.IKey</code> | The customer-managed encryption key that is used to encrypt this secret, if any. |
| <code><a href="#cdk-sops-secrets.SopsSecrets.property.secretFullArn">secretFullArn</a></code> | <code>string</code> | The full ARN of the secret in AWS Secrets Manager, which is the ARN including the Secrets Manager-supplied 6-character suffix. |

---

##### `node`<sup>Required</sup> <a name="node" id="cdk-sops-secrets.SopsSecrets.property.node"></a>

```typescript
public readonly node: ConstructNode;
```

- *Type:* @aws-cdk/core.ConstructNode

The construct tree node associated with this construct.

---

##### `env`<sup>Required</sup> <a name="env" id="cdk-sops-secrets.SopsSecrets.property.env"></a>

```typescript
public readonly env: ResourceEnvironment;
```

- *Type:* @aws-cdk/core.ResourceEnvironment

The environment this resource belongs to.

For resources that are created and managed by the CDK
(generally, those created by creating new class instances like Role, Bucket, etc.),
this is always the same as the environment of the stack they belong to;
however, for imported resources
(those obtained from static methods like fromRoleArn, fromBucketName, etc.),
that might be different than the stack they were imported into.

---

##### `stack`<sup>Required</sup> <a name="stack" id="cdk-sops-secrets.SopsSecrets.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* @aws-cdk/core.Stack

The stack in which this resource is defined.

---

##### `secretArn`<sup>Required</sup> <a name="secretArn" id="cdk-sops-secrets.SopsSecrets.property.secretArn"></a>

```typescript
public readonly secretArn: string;
```

- *Type:* string

The ARN of the secret in AWS Secrets Manager.

Will return the full ARN if available, otherwise a partial arn.
For secrets imported by the deprecated `fromSecretName`, it will return the `secretName`.

---

##### `secretName`<sup>Required</sup> <a name="secretName" id="cdk-sops-secrets.SopsSecrets.property.secretName"></a>

```typescript
public readonly secretName: string;
```

- *Type:* string

The name of the secret.

For "owned" secrets, this will be the full resource name (secret name + suffix), unless the
'@aws-cdk/aws-secretsmanager:parseOwnedSecretName' feature flag is set.

---

##### `secretValue`<sup>Required</sup> <a name="secretValue" id="cdk-sops-secrets.SopsSecrets.property.secretValue"></a>

```typescript
public readonly secretValue: SecretValue;
```

- *Type:* @aws-cdk/core.SecretValue

Retrieve the value of the stored secret as a `SecretValue`.

---

##### `encryptionKey`<sup>Optional</sup> <a name="encryptionKey" id="cdk-sops-secrets.SopsSecrets.property.encryptionKey"></a>

```typescript
public readonly encryptionKey: IKey;
```

- *Type:* @aws-cdk/aws-kms.IKey

The customer-managed encryption key that is used to encrypt this secret, if any.

When not specified, the default
KMS key for the account and region is being used.

---

##### `secretFullArn`<sup>Optional</sup> <a name="secretFullArn" id="cdk-sops-secrets.SopsSecrets.property.secretFullArn"></a>

```typescript
public readonly secretFullArn: string;
```

- *Type:* string

The full ARN of the secret in AWS Secrets Manager, which is the ARN including the Secrets Manager-supplied 6-character suffix.

This is equal to `secretArn` in most cases, but is undefined when a full ARN is not available (e.g., secrets imported by name).

---


### SopsSync <a name="SopsSync" id="cdk-sops-secrets.SopsSync"></a>

#### Initializers <a name="Initializers" id="cdk-sops-secrets.SopsSync.Initializer"></a>

```typescript
import { SopsSync } from 'cdk-sops-secrets'

new SopsSync(scope: Construct, id: string, props: SopsSyncProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSync.Initializer.parameter.scope">scope</a></code> | <code>@aws-cdk/core.Construct</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSync.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSync.Initializer.parameter.props">props</a></code> | <code><a href="#cdk-sops-secrets.SopsSyncProps">SopsSyncProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="cdk-sops-secrets.SopsSync.Initializer.parameter.scope"></a>

- *Type:* @aws-cdk/core.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="cdk-sops-secrets.SopsSync.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="cdk-sops-secrets.SopsSync.Initializer.parameter.props"></a>

- *Type:* <a href="#cdk-sops-secrets.SopsSyncProps">SopsSyncProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSync.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="cdk-sops-secrets.SopsSync.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSync.isConstruct">isConstruct</a></code> | Return whether the given object is a Construct. |

---

##### `isConstruct` <a name="isConstruct" id="cdk-sops-secrets.SopsSync.isConstruct"></a>

```typescript
import { SopsSync } from 'cdk-sops-secrets'

SopsSync.isConstruct(x: any)
```

Return whether the given object is a Construct.

###### `x`<sup>Required</sup> <a name="x" id="cdk-sops-secrets.SopsSync.isConstruct.parameter.x"></a>

- *Type:* any

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSync.property.node">node</a></code> | <code>@aws-cdk/core.ConstructNode</code> | The construct tree node associated with this construct. |

---

##### `node`<sup>Required</sup> <a name="node" id="cdk-sops-secrets.SopsSync.property.node"></a>

```typescript
public readonly node: ConstructNode;
```

- *Type:* @aws-cdk/core.ConstructNode

The construct tree node associated with this construct.

---


## Structs <a name="Structs" id="Structs"></a>

### SopsSecretProps <a name="SopsSecretProps" id="cdk-sops-secrets.SopsSecretProps"></a>

#### Initializer <a name="Initializer" id="cdk-sops-secrets.SopsSecretProps.Initializer"></a>

```typescript
import { SopsSecretProps } from 'cdk-sops-secrets'

const sopsSecretProps: SopsSecretProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.description">description</a></code> | <code>string</code> | An optional, human-friendly description of the secret. |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.encryptionKey">encryptionKey</a></code> | <code>@aws-cdk/aws-kms.IKey</code> | The customer-managed encryption key to use for encrypting the secret value. |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.generateSecretString">generateSecretString</a></code> | <code>@aws-cdk/aws-secretsmanager.SecretStringGenerator</code> | Configuration for how to generate a secret value. |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.removalPolicy">removalPolicy</a></code> | <code>@aws-cdk/core.RemovalPolicy</code> | Policy to apply when the secret is removed from this stack. |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.replicaRegions">replicaRegions</a></code> | <code>@aws-cdk/aws-secretsmanager.ReplicaRegion[]</code> | A list of regions where to replicate this secret. |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.secretName">secretName</a></code> | <code>string</code> | A name for the secret. |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.secretStringBeta1">secretStringBeta1</a></code> | <code>@aws-cdk/aws-secretsmanager.SecretStringValueBeta1</code> | Initial value for the secret. |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.sopsFilePath">sopsFilePath</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.sopsAgeKey">sopsAgeKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.sopsFileFormat">sopsFileFormat</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.sopsKmsKey">sopsKmsKey</a></code> | <code>@aws-cdk/aws-kms.IKey</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.sopsProvider">sopsProvider</a></code> | <code>@aws-cdk/aws-lambda.IFunction</code> | *No description.* |

---

##### `description`<sup>Optional</sup> <a name="description" id="cdk-sops-secrets.SopsSecretProps.property.description"></a>

```typescript
public readonly description: string;
```

- *Type:* string
- *Default:* No description.

An optional, human-friendly description of the secret.

---

##### `encryptionKey`<sup>Optional</sup> <a name="encryptionKey" id="cdk-sops-secrets.SopsSecretProps.property.encryptionKey"></a>

```typescript
public readonly encryptionKey: IKey;
```

- *Type:* @aws-cdk/aws-kms.IKey
- *Default:* A default KMS key for the account and region is used.

The customer-managed encryption key to use for encrypting the secret value.

---

##### `generateSecretString`<sup>Optional</sup> <a name="generateSecretString" id="cdk-sops-secrets.SopsSecretProps.property.generateSecretString"></a>

```typescript
public readonly generateSecretString: SecretStringGenerator;
```

- *Type:* @aws-cdk/aws-secretsmanager.SecretStringGenerator
- *Default:* 32 characters with upper-case letters, lower-case letters, punctuation and numbers (at least one from each category), per the default values of ``SecretStringGenerator``.

Configuration for how to generate a secret value.

Only one of `secretString` and `generateSecretString` can be provided.

---

##### `removalPolicy`<sup>Optional</sup> <a name="removalPolicy" id="cdk-sops-secrets.SopsSecretProps.property.removalPolicy"></a>

```typescript
public readonly removalPolicy: RemovalPolicy;
```

- *Type:* @aws-cdk/core.RemovalPolicy
- *Default:* Not set.

Policy to apply when the secret is removed from this stack.

---

##### `replicaRegions`<sup>Optional</sup> <a name="replicaRegions" id="cdk-sops-secrets.SopsSecretProps.property.replicaRegions"></a>

```typescript
public readonly replicaRegions: ReplicaRegion[];
```

- *Type:* @aws-cdk/aws-secretsmanager.ReplicaRegion[]
- *Default:* Secret is not replicated

A list of regions where to replicate this secret.

---

##### `secretName`<sup>Optional</sup> <a name="secretName" id="cdk-sops-secrets.SopsSecretProps.property.secretName"></a>

```typescript
public readonly secretName: string;
```

- *Type:* string
- *Default:* A name is generated by CloudFormation.

A name for the secret.

Note that deleting secrets from SecretsManager does not happen immediately, but after a 7 to
30 days blackout period. During that period, it is not possible to create another secret that shares the same name.

---

##### `secretStringBeta1`<sup>Optional</sup> <a name="secretStringBeta1" id="cdk-sops-secrets.SopsSecretProps.property.secretStringBeta1"></a>

```typescript
public readonly secretStringBeta1: SecretStringValueBeta1;
```

- *Type:* @aws-cdk/aws-secretsmanager.SecretStringValueBeta1
- *Default:* SecretsManager generates a new secret value.

Initial value for the secret.

**NOTE:** *It is **highly** encouraged to leave this field undefined and allow SecretsManager to create the secret value.
The secret string -- if provided -- will be included in the output of the cdk as part of synthesis,
and will appear in the CloudFormation template in the console. This can be secure(-ish) if that value is merely reference to
another resource (or one of its attributes), but if the value is a plaintext string, it will be visible to anyone with access
to the CloudFormation template (via the AWS Console, SDKs, or CLI).

Specifies text data that you want to encrypt and store in this new version of the secret.
May be a simple string value, or a string representation of a JSON structure.

Only one of `secretString` and `generateSecretString` can be provided.

---

##### `sopsFilePath`<sup>Required</sup> <a name="sopsFilePath" id="cdk-sops-secrets.SopsSecretProps.property.sopsFilePath"></a>

```typescript
public readonly sopsFilePath: string;
```

- *Type:* string

---

##### `sopsAgeKey`<sup>Optional</sup> <a name="sopsAgeKey" id="cdk-sops-secrets.SopsSecretProps.property.sopsAgeKey"></a>

```typescript
public readonly sopsAgeKey: string;
```

- *Type:* string

---

##### `sopsFileFormat`<sup>Optional</sup> <a name="sopsFileFormat" id="cdk-sops-secrets.SopsSecretProps.property.sopsFileFormat"></a>

```typescript
public readonly sopsFileFormat: string;
```

- *Type:* string

---

##### `sopsKmsKey`<sup>Optional</sup> <a name="sopsKmsKey" id="cdk-sops-secrets.SopsSecretProps.property.sopsKmsKey"></a>

```typescript
public readonly sopsKmsKey: IKey;
```

- *Type:* @aws-cdk/aws-kms.IKey

---

##### `sopsProvider`<sup>Optional</sup> <a name="sopsProvider" id="cdk-sops-secrets.SopsSecretProps.property.sopsProvider"></a>

```typescript
public readonly sopsProvider: IFunction;
```

- *Type:* @aws-cdk/aws-lambda.IFunction

---

### SopsSyncOptions <a name="SopsSyncOptions" id="cdk-sops-secrets.SopsSyncOptions"></a>

#### Initializer <a name="Initializer" id="cdk-sops-secrets.SopsSyncOptions.Initializer"></a>

```typescript
import { SopsSyncOptions } from 'cdk-sops-secrets'

const sopsSyncOptions: SopsSyncOptions = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSyncOptions.property.sopsFilePath">sopsFilePath</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSyncOptions.property.sopsAgeKey">sopsAgeKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSyncOptions.property.sopsFileFormat">sopsFileFormat</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSyncOptions.property.sopsKmsKey">sopsKmsKey</a></code> | <code>@aws-cdk/aws-kms.IKey</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSyncOptions.property.sopsProvider">sopsProvider</a></code> | <code>@aws-cdk/aws-lambda.IFunction</code> | *No description.* |

---

##### `sopsFilePath`<sup>Required</sup> <a name="sopsFilePath" id="cdk-sops-secrets.SopsSyncOptions.property.sopsFilePath"></a>

```typescript
public readonly sopsFilePath: string;
```

- *Type:* string

---

##### `sopsAgeKey`<sup>Optional</sup> <a name="sopsAgeKey" id="cdk-sops-secrets.SopsSyncOptions.property.sopsAgeKey"></a>

```typescript
public readonly sopsAgeKey: string;
```

- *Type:* string

---

##### `sopsFileFormat`<sup>Optional</sup> <a name="sopsFileFormat" id="cdk-sops-secrets.SopsSyncOptions.property.sopsFileFormat"></a>

```typescript
public readonly sopsFileFormat: string;
```

- *Type:* string

---

##### `sopsKmsKey`<sup>Optional</sup> <a name="sopsKmsKey" id="cdk-sops-secrets.SopsSyncOptions.property.sopsKmsKey"></a>

```typescript
public readonly sopsKmsKey: IKey;
```

- *Type:* @aws-cdk/aws-kms.IKey

---

##### `sopsProvider`<sup>Optional</sup> <a name="sopsProvider" id="cdk-sops-secrets.SopsSyncOptions.property.sopsProvider"></a>

```typescript
public readonly sopsProvider: IFunction;
```

- *Type:* @aws-cdk/aws-lambda.IFunction

---

### SopsSyncProps <a name="SopsSyncProps" id="cdk-sops-secrets.SopsSyncProps"></a>

#### Initializer <a name="Initializer" id="cdk-sops-secrets.SopsSyncProps.Initializer"></a>

```typescript
import { SopsSyncProps } from 'cdk-sops-secrets'

const sopsSyncProps: SopsSyncProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.sopsFilePath">sopsFilePath</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.sopsAgeKey">sopsAgeKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.sopsFileFormat">sopsFileFormat</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.sopsKmsKey">sopsKmsKey</a></code> | <code>@aws-cdk/aws-kms.IKey</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.sopsProvider">sopsProvider</a></code> | <code>@aws-cdk/aws-lambda.IFunction</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.secret">secret</a></code> | <code>@aws-cdk/aws-secretsmanager.ISecret</code> | *No description.* |

---

##### `sopsFilePath`<sup>Required</sup> <a name="sopsFilePath" id="cdk-sops-secrets.SopsSyncProps.property.sopsFilePath"></a>

```typescript
public readonly sopsFilePath: string;
```

- *Type:* string

---

##### `sopsAgeKey`<sup>Optional</sup> <a name="sopsAgeKey" id="cdk-sops-secrets.SopsSyncProps.property.sopsAgeKey"></a>

```typescript
public readonly sopsAgeKey: string;
```

- *Type:* string

---

##### `sopsFileFormat`<sup>Optional</sup> <a name="sopsFileFormat" id="cdk-sops-secrets.SopsSyncProps.property.sopsFileFormat"></a>

```typescript
public readonly sopsFileFormat: string;
```

- *Type:* string

---

##### `sopsKmsKey`<sup>Optional</sup> <a name="sopsKmsKey" id="cdk-sops-secrets.SopsSyncProps.property.sopsKmsKey"></a>

```typescript
public readonly sopsKmsKey: IKey;
```

- *Type:* @aws-cdk/aws-kms.IKey

---

##### `sopsProvider`<sup>Optional</sup> <a name="sopsProvider" id="cdk-sops-secrets.SopsSyncProps.property.sopsProvider"></a>

```typescript
public readonly sopsProvider: IFunction;
```

- *Type:* @aws-cdk/aws-lambda.IFunction

---

##### `secret`<sup>Required</sup> <a name="secret" id="cdk-sops-secrets.SopsSyncProps.property.secret"></a>

```typescript
public readonly secret: ISecret;
```

- *Type:* @aws-cdk/aws-secretsmanager.ISecret

---



