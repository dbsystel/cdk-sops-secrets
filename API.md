# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### MultiStringParameter <a name="MultiStringParameter" id="cdk-sops-secrets.MultiStringParameter"></a>

#### Initializers <a name="Initializers" id="cdk-sops-secrets.MultiStringParameter.Initializer"></a>

```typescript
import { MultiStringParameter } from 'cdk-sops-secrets'

new MultiStringParameter(scope: Construct, id: string, props: SopsStringParameterProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.MultiStringParameter.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.MultiStringParameter.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.MultiStringParameter.Initializer.parameter.props">props</a></code> | <code><a href="#cdk-sops-secrets.SopsStringParameterProps">SopsStringParameterProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="cdk-sops-secrets.MultiStringParameter.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="cdk-sops-secrets.MultiStringParameter.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="cdk-sops-secrets.MultiStringParameter.Initializer.parameter.props"></a>

- *Type:* <a href="#cdk-sops-secrets.SopsStringParameterProps">SopsStringParameterProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-sops-secrets.MultiStringParameter.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="cdk-sops-secrets.MultiStringParameter.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-sops-secrets.MultiStringParameter.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="cdk-sops-secrets.MultiStringParameter.isConstruct"></a>

```typescript
import { MultiStringParameter } from 'cdk-sops-secrets'

MultiStringParameter.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="cdk-sops-secrets.MultiStringParameter.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.MultiStringParameter.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#cdk-sops-secrets.MultiStringParameter.property.encryptionKey">encryptionKey</a></code> | <code>aws-cdk-lib.aws_kms.IKey</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.MultiStringParameter.property.env">env</a></code> | <code>aws-cdk-lib.ResourceEnvironment</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.MultiStringParameter.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.MultiStringParameter.property.sync">sync</a></code> | <code><a href="#cdk-sops-secrets.SopsSync">SopsSync</a></code> | *No description.* |

---

##### `node`<sup>Required</sup> <a name="node" id="cdk-sops-secrets.MultiStringParameter.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `encryptionKey`<sup>Required</sup> <a name="encryptionKey" id="cdk-sops-secrets.MultiStringParameter.property.encryptionKey"></a>

```typescript
public readonly encryptionKey: IKey;
```

- *Type:* aws-cdk-lib.aws_kms.IKey

---

##### `env`<sup>Required</sup> <a name="env" id="cdk-sops-secrets.MultiStringParameter.property.env"></a>

```typescript
public readonly env: ResourceEnvironment;
```

- *Type:* aws-cdk-lib.ResourceEnvironment

---

##### `stack`<sup>Required</sup> <a name="stack" id="cdk-sops-secrets.MultiStringParameter.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

---

##### `sync`<sup>Required</sup> <a name="sync" id="cdk-sops-secrets.MultiStringParameter.property.sync"></a>

```typescript
public readonly sync: SopsSync;
```

- *Type:* <a href="#cdk-sops-secrets.SopsSync">SopsSync</a>

---


### SopsSecret <a name="SopsSecret" id="cdk-sops-secrets.SopsSecret"></a>

- *Implements:* aws-cdk-lib.aws_secretsmanager.ISecret

A drop in replacement for the normal Secret, that is populated with the encrypted content of the given sops file.

#### Initializers <a name="Initializers" id="cdk-sops-secrets.SopsSecret.Initializer"></a>

```typescript
import { SopsSecret } from 'cdk-sops-secrets'

new SopsSecret(scope: Construct, id: string, props: SopsSecretProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSecret.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSecret.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSecret.Initializer.parameter.props">props</a></code> | <code><a href="#cdk-sops-secrets.SopsSecretProps">SopsSecretProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="cdk-sops-secrets.SopsSecret.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="cdk-sops-secrets.SopsSecret.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="cdk-sops-secrets.SopsSecret.Initializer.parameter.props"></a>

- *Type:* <a href="#cdk-sops-secrets.SopsSecretProps">SopsSecretProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSecret.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#cdk-sops-secrets.SopsSecret.addRotationSchedule">addRotationSchedule</a></code> | Adds a rotation schedule to the secret. |
| <code><a href="#cdk-sops-secrets.SopsSecret.addToResourcePolicy">addToResourcePolicy</a></code> | Adds a statement to the IAM resource policy associated with this secret. |
| <code><a href="#cdk-sops-secrets.SopsSecret.applyRemovalPolicy">applyRemovalPolicy</a></code> | Apply the given removal policy to this resource. |
| <code><a href="#cdk-sops-secrets.SopsSecret.attach">attach</a></code> | Attach a target to this secret. |
| <code><a href="#cdk-sops-secrets.SopsSecret.currentVersionId">currentVersionId</a></code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSecret.denyAccountRootDelete">denyAccountRootDelete</a></code> | Denies the `DeleteSecret` action to all principals within the current account. |
| <code><a href="#cdk-sops-secrets.SopsSecret.grantRead">grantRead</a></code> | Grants reading the secret value to some role. |
| <code><a href="#cdk-sops-secrets.SopsSecret.grantWrite">grantWrite</a></code> | Grants writing and updating the secret value to some role. |
| <code><a href="#cdk-sops-secrets.SopsSecret.secretValueFromJson">secretValueFromJson</a></code> | Interpret the secret as a JSON object and return a field's value from it as a `SecretValue`. |

---

##### `toString` <a name="toString" id="cdk-sops-secrets.SopsSecret.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `addRotationSchedule` <a name="addRotationSchedule" id="cdk-sops-secrets.SopsSecret.addRotationSchedule"></a>

```typescript
public addRotationSchedule(id: string, options: RotationScheduleOptions): RotationSchedule
```

Adds a rotation schedule to the secret.

###### `id`<sup>Required</sup> <a name="id" id="cdk-sops-secrets.SopsSecret.addRotationSchedule.parameter.id"></a>

- *Type:* string

---

###### `options`<sup>Required</sup> <a name="options" id="cdk-sops-secrets.SopsSecret.addRotationSchedule.parameter.options"></a>

- *Type:* aws-cdk-lib.aws_secretsmanager.RotationScheduleOptions

---

##### `addToResourcePolicy` <a name="addToResourcePolicy" id="cdk-sops-secrets.SopsSecret.addToResourcePolicy"></a>

```typescript
public addToResourcePolicy(statement: PolicyStatement): AddToResourcePolicyResult
```

Adds a statement to the IAM resource policy associated with this secret.

If this secret was created in this stack, a resource policy will be
automatically created upon the first call to `addToResourcePolicy`. If
the secret is imported, then this is a no-op.

###### `statement`<sup>Required</sup> <a name="statement" id="cdk-sops-secrets.SopsSecret.addToResourcePolicy.parameter.statement"></a>

- *Type:* aws-cdk-lib.aws_iam.PolicyStatement

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="cdk-sops-secrets.SopsSecret.applyRemovalPolicy"></a>

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

###### `policy`<sup>Required</sup> <a name="policy" id="cdk-sops-secrets.SopsSecret.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

##### `attach` <a name="attach" id="cdk-sops-secrets.SopsSecret.attach"></a>

```typescript
public attach(target: ISecretAttachmentTarget): ISecret
```

Attach a target to this secret.

###### `target`<sup>Required</sup> <a name="target" id="cdk-sops-secrets.SopsSecret.attach.parameter.target"></a>

- *Type:* aws-cdk-lib.aws_secretsmanager.ISecretAttachmentTarget

---

##### `currentVersionId` <a name="currentVersionId" id="cdk-sops-secrets.SopsSecret.currentVersionId"></a>

```typescript
public currentVersionId(): string
```

##### `denyAccountRootDelete` <a name="denyAccountRootDelete" id="cdk-sops-secrets.SopsSecret.denyAccountRootDelete"></a>

```typescript
public denyAccountRootDelete(): void
```

Denies the `DeleteSecret` action to all principals within the current account.

##### `grantRead` <a name="grantRead" id="cdk-sops-secrets.SopsSecret.grantRead"></a>

```typescript
public grantRead(grantee: IGrantable, versionStages?: string[]): Grant
```

Grants reading the secret value to some role.

###### `grantee`<sup>Required</sup> <a name="grantee" id="cdk-sops-secrets.SopsSecret.grantRead.parameter.grantee"></a>

- *Type:* aws-cdk-lib.aws_iam.IGrantable

---

###### `versionStages`<sup>Optional</sup> <a name="versionStages" id="cdk-sops-secrets.SopsSecret.grantRead.parameter.versionStages"></a>

- *Type:* string[]

---

##### `grantWrite` <a name="grantWrite" id="cdk-sops-secrets.SopsSecret.grantWrite"></a>

```typescript
public grantWrite(_grantee: IGrantable): Grant
```

Grants writing and updating the secret value to some role.

###### `_grantee`<sup>Required</sup> <a name="_grantee" id="cdk-sops-secrets.SopsSecret.grantWrite.parameter._grantee"></a>

- *Type:* aws-cdk-lib.aws_iam.IGrantable

---

##### `secretValueFromJson` <a name="secretValueFromJson" id="cdk-sops-secrets.SopsSecret.secretValueFromJson"></a>

```typescript
public secretValueFromJson(jsonField: string): SecretValue
```

Interpret the secret as a JSON object and return a field's value from it as a `SecretValue`.

###### `jsonField`<sup>Required</sup> <a name="jsonField" id="cdk-sops-secrets.SopsSecret.secretValueFromJson.parameter.jsonField"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSecret.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="cdk-sops-secrets.SopsSecret.isConstruct"></a>

```typescript
import { SopsSecret } from 'cdk-sops-secrets'

SopsSecret.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="cdk-sops-secrets.SopsSecret.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSecret.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#cdk-sops-secrets.SopsSecret.property.env">env</a></code> | <code>aws-cdk-lib.ResourceEnvironment</code> | The environment this resource belongs to. |
| <code><a href="#cdk-sops-secrets.SopsSecret.property.secretArn">secretArn</a></code> | <code>string</code> | The ARN of the secret in AWS Secrets Manager. |
| <code><a href="#cdk-sops-secrets.SopsSecret.property.secretName">secretName</a></code> | <code>string</code> | The name of the secret. |
| <code><a href="#cdk-sops-secrets.SopsSecret.property.secretValue">secretValue</a></code> | <code>aws-cdk-lib.SecretValue</code> | Retrieve the value of the stored secret as a `SecretValue`. |
| <code><a href="#cdk-sops-secrets.SopsSecret.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this resource is defined. |
| <code><a href="#cdk-sops-secrets.SopsSecret.property.sync">sync</a></code> | <code><a href="#cdk-sops-secrets.SopsSync">SopsSync</a></code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSecret.property.encryptionKey">encryptionKey</a></code> | <code>aws-cdk-lib.aws_kms.IKey</code> | The customer-managed encryption key that is used to encrypt this secret, if any. |
| <code><a href="#cdk-sops-secrets.SopsSecret.property.secretFullArn">secretFullArn</a></code> | <code>string</code> | The full ARN of the secret in AWS Secrets Manager, which is the ARN including the Secrets Manager-supplied 6-character suffix. |

---

##### `node`<sup>Required</sup> <a name="node" id="cdk-sops-secrets.SopsSecret.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `env`<sup>Required</sup> <a name="env" id="cdk-sops-secrets.SopsSecret.property.env"></a>

```typescript
public readonly env: ResourceEnvironment;
```

- *Type:* aws-cdk-lib.ResourceEnvironment

The environment this resource belongs to.

For resources that are created and managed by the CDK
(generally, those created by creating new class instances like Role, Bucket, etc.),
this is always the same as the environment of the stack they belong to;
however, for imported resources
(those obtained from static methods like fromRoleArn, fromBucketName, etc.),
that might be different than the stack they were imported into.

---

##### `secretArn`<sup>Required</sup> <a name="secretArn" id="cdk-sops-secrets.SopsSecret.property.secretArn"></a>

```typescript
public readonly secretArn: string;
```

- *Type:* string

The ARN of the secret in AWS Secrets Manager.

Will return the full ARN if available, otherwise a partial arn.
For secrets imported by the deprecated `fromSecretName`, it will return the `secretName`.

---

##### `secretName`<sup>Required</sup> <a name="secretName" id="cdk-sops-secrets.SopsSecret.property.secretName"></a>

```typescript
public readonly secretName: string;
```

- *Type:* string

The name of the secret.

For "owned" secrets, this will be the full resource name (secret name + suffix), unless the
'@aws-cdk/aws-secretsmanager:parseOwnedSecretName' feature flag is set.

---

##### `secretValue`<sup>Required</sup> <a name="secretValue" id="cdk-sops-secrets.SopsSecret.property.secretValue"></a>

```typescript
public readonly secretValue: SecretValue;
```

- *Type:* aws-cdk-lib.SecretValue

Retrieve the value of the stored secret as a `SecretValue`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="cdk-sops-secrets.SopsSecret.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this resource is defined.

---

##### `sync`<sup>Required</sup> <a name="sync" id="cdk-sops-secrets.SopsSecret.property.sync"></a>

```typescript
public readonly sync: SopsSync;
```

- *Type:* <a href="#cdk-sops-secrets.SopsSync">SopsSync</a>

---

##### `encryptionKey`<sup>Optional</sup> <a name="encryptionKey" id="cdk-sops-secrets.SopsSecret.property.encryptionKey"></a>

```typescript
public readonly encryptionKey: IKey;
```

- *Type:* aws-cdk-lib.aws_kms.IKey

The customer-managed encryption key that is used to encrypt this secret, if any.

When not specified, the default
KMS key for the account and region is being used.

---

##### `secretFullArn`<sup>Optional</sup> <a name="secretFullArn" id="cdk-sops-secrets.SopsSecret.property.secretFullArn"></a>

```typescript
public readonly secretFullArn: string;
```

- *Type:* string

The full ARN of the secret in AWS Secrets Manager, which is the ARN including the Secrets Manager-supplied 6-character suffix.

This is equal to `secretArn` in most cases, but is undefined when a full ARN is not available (e.g., secrets imported by name).

---


### SopsStringParameter <a name="SopsStringParameter" id="cdk-sops-secrets.SopsStringParameter"></a>

- *Implements:* aws-cdk-lib.aws_ssm.IStringParameter

A drop in replacement for the normal String Parameter, that is populated with the encrypted content of the given sops file.

#### Initializers <a name="Initializers" id="cdk-sops-secrets.SopsStringParameter.Initializer"></a>

```typescript
import { SopsStringParameter } from 'cdk-sops-secrets'

new SopsStringParameter(scope: Construct, id: string, props: SopsStringParameterProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.SopsStringParameter.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsStringParameter.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsStringParameter.Initializer.parameter.props">props</a></code> | <code><a href="#cdk-sops-secrets.SopsStringParameterProps">SopsStringParameterProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="cdk-sops-secrets.SopsStringParameter.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="cdk-sops-secrets.SopsStringParameter.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="cdk-sops-secrets.SopsStringParameter.Initializer.parameter.props"></a>

- *Type:* <a href="#cdk-sops-secrets.SopsStringParameterProps">SopsStringParameterProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-sops-secrets.SopsStringParameter.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#cdk-sops-secrets.SopsStringParameter.applyRemovalPolicy">applyRemovalPolicy</a></code> | Apply the given removal policy to this resource. |
| <code><a href="#cdk-sops-secrets.SopsStringParameter.grantRead">grantRead</a></code> | Grants read (DescribeParameter, GetParameters, GetParameter, GetParameterHistory) permissions on the SSM Parameter. |
| <code><a href="#cdk-sops-secrets.SopsStringParameter.grantWrite">grantWrite</a></code> | Grants write (PutParameter) permissions on the SSM Parameter. |

---

##### `toString` <a name="toString" id="cdk-sops-secrets.SopsStringParameter.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="cdk-sops-secrets.SopsStringParameter.applyRemovalPolicy"></a>

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

###### `policy`<sup>Required</sup> <a name="policy" id="cdk-sops-secrets.SopsStringParameter.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

##### `grantRead` <a name="grantRead" id="cdk-sops-secrets.SopsStringParameter.grantRead"></a>

```typescript
public grantRead(grantee: IGrantable): Grant
```

Grants read (DescribeParameter, GetParameters, GetParameter, GetParameterHistory) permissions on the SSM Parameter.

###### `grantee`<sup>Required</sup> <a name="grantee" id="cdk-sops-secrets.SopsStringParameter.grantRead.parameter.grantee"></a>

- *Type:* aws-cdk-lib.aws_iam.IGrantable

---

##### `grantWrite` <a name="grantWrite" id="cdk-sops-secrets.SopsStringParameter.grantWrite"></a>

```typescript
public grantWrite(grantee: IGrantable): Grant
```

Grants write (PutParameter) permissions on the SSM Parameter.

###### `grantee`<sup>Required</sup> <a name="grantee" id="cdk-sops-secrets.SopsStringParameter.grantWrite.parameter.grantee"></a>

- *Type:* aws-cdk-lib.aws_iam.IGrantable

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-sops-secrets.SopsStringParameter.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="cdk-sops-secrets.SopsStringParameter.isConstruct"></a>

```typescript
import { SopsStringParameter } from 'cdk-sops-secrets'

SopsStringParameter.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="cdk-sops-secrets.SopsStringParameter.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.SopsStringParameter.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#cdk-sops-secrets.SopsStringParameter.property.encryptionKey">encryptionKey</a></code> | <code>aws-cdk-lib.aws_kms.IKey</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsStringParameter.property.env">env</a></code> | <code>aws-cdk-lib.ResourceEnvironment</code> | The environment this resource belongs to. |
| <code><a href="#cdk-sops-secrets.SopsStringParameter.property.parameterArn">parameterArn</a></code> | <code>string</code> | The ARN of the SSM Parameter resource. |
| <code><a href="#cdk-sops-secrets.SopsStringParameter.property.parameterName">parameterName</a></code> | <code>string</code> | The name of the SSM Parameter resource. |
| <code><a href="#cdk-sops-secrets.SopsStringParameter.property.parameterType">parameterType</a></code> | <code>string</code> | The type of the SSM Parameter resource. |
| <code><a href="#cdk-sops-secrets.SopsStringParameter.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this resource is defined. |
| <code><a href="#cdk-sops-secrets.SopsStringParameter.property.stringValue">stringValue</a></code> | <code>string</code> | The parameter value. |
| <code><a href="#cdk-sops-secrets.SopsStringParameter.property.sync">sync</a></code> | <code><a href="#cdk-sops-secrets.SopsSync">SopsSync</a></code> | *No description.* |

---

##### `node`<sup>Required</sup> <a name="node" id="cdk-sops-secrets.SopsStringParameter.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `encryptionKey`<sup>Required</sup> <a name="encryptionKey" id="cdk-sops-secrets.SopsStringParameter.property.encryptionKey"></a>

```typescript
public readonly encryptionKey: IKey;
```

- *Type:* aws-cdk-lib.aws_kms.IKey

---

##### `env`<sup>Required</sup> <a name="env" id="cdk-sops-secrets.SopsStringParameter.property.env"></a>

```typescript
public readonly env: ResourceEnvironment;
```

- *Type:* aws-cdk-lib.ResourceEnvironment

The environment this resource belongs to.

For resources that are created and managed by the CDK
(generally, those created by creating new class instances like Role, Bucket, etc.),
this is always the same as the environment of the stack they belong to;
however, for imported resources
(those obtained from static methods like fromRoleArn, fromBucketName, etc.),
that might be different than the stack they were imported into.

---

##### `parameterArn`<sup>Required</sup> <a name="parameterArn" id="cdk-sops-secrets.SopsStringParameter.property.parameterArn"></a>

```typescript
public readonly parameterArn: string;
```

- *Type:* string

The ARN of the SSM Parameter resource.

---

##### `parameterName`<sup>Required</sup> <a name="parameterName" id="cdk-sops-secrets.SopsStringParameter.property.parameterName"></a>

```typescript
public readonly parameterName: string;
```

- *Type:* string

The name of the SSM Parameter resource.

---

##### `parameterType`<sup>Required</sup> <a name="parameterType" id="cdk-sops-secrets.SopsStringParameter.property.parameterType"></a>

```typescript
public readonly parameterType: string;
```

- *Type:* string

The type of the SSM Parameter resource.

---

##### `stack`<sup>Required</sup> <a name="stack" id="cdk-sops-secrets.SopsStringParameter.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this resource is defined.

---

##### `stringValue`<sup>Required</sup> <a name="stringValue" id="cdk-sops-secrets.SopsStringParameter.property.stringValue"></a>

```typescript
public readonly stringValue: string;
```

- *Type:* string

The parameter value.

Value must not nest another parameter. Do not use {{}} in the value.

---

##### `sync`<sup>Required</sup> <a name="sync" id="cdk-sops-secrets.SopsStringParameter.property.sync"></a>

```typescript
public readonly sync: SopsSync;
```

- *Type:* <a href="#cdk-sops-secrets.SopsSync">SopsSync</a>

---


### SopsSync <a name="SopsSync" id="cdk-sops-secrets.SopsSync"></a>

The custom resource, that is syncing the content from a sops file to a secret.

#### Initializers <a name="Initializers" id="cdk-sops-secrets.SopsSync.Initializer"></a>

```typescript
import { SopsSync } from 'cdk-sops-secrets'

new SopsSync(scope: Construct, id: string, props: SopsSyncProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSync.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSync.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSync.Initializer.parameter.props">props</a></code> | <code><a href="#cdk-sops-secrets.SopsSyncProps">SopsSyncProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="cdk-sops-secrets.SopsSync.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

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
| <code><a href="#cdk-sops-secrets.SopsSync.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="cdk-sops-secrets.SopsSync.isConstruct"></a>

```typescript
import { SopsSync } from 'cdk-sops-secrets'

SopsSync.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="cdk-sops-secrets.SopsSync.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSync.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#cdk-sops-secrets.SopsSync.property.converToJSON">converToJSON</a></code> | <code>boolean</code> | Was the format converted to json? |
| <code><a href="#cdk-sops-secrets.SopsSync.property.flatten">flatten</a></code> | <code>boolean</code> | Was the structure flattened? |
| <code><a href="#cdk-sops-secrets.SopsSync.property.stringifiedValues">stringifiedValues</a></code> | <code>boolean</code> | Were the values stringified? |
| <code><a href="#cdk-sops-secrets.SopsSync.property.versionId">versionId</a></code> | <code>string</code> | The current versionId of the secret populated via this resource. |

---

##### `node`<sup>Required</sup> <a name="node" id="cdk-sops-secrets.SopsSync.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `converToJSON`<sup>Required</sup> <a name="converToJSON" id="cdk-sops-secrets.SopsSync.property.converToJSON"></a>

```typescript
public readonly converToJSON: boolean;
```

- *Type:* boolean

Was the format converted to json?

---

##### `flatten`<sup>Required</sup> <a name="flatten" id="cdk-sops-secrets.SopsSync.property.flatten"></a>

```typescript
public readonly flatten: boolean;
```

- *Type:* boolean

Was the structure flattened?

---

##### `stringifiedValues`<sup>Required</sup> <a name="stringifiedValues" id="cdk-sops-secrets.SopsSync.property.stringifiedValues"></a>

```typescript
public readonly stringifiedValues: boolean;
```

- *Type:* boolean

Were the values stringified?

---

##### `versionId`<sup>Required</sup> <a name="versionId" id="cdk-sops-secrets.SopsSync.property.versionId"></a>

```typescript
public readonly versionId: string;
```

- *Type:* string

The current versionId of the secret populated via this resource.

---


### SopsSyncProvider <a name="SopsSyncProvider" id="cdk-sops-secrets.SopsSyncProvider"></a>

- *Implements:* aws-cdk-lib.aws_iam.IGrantable

#### Initializers <a name="Initializers" id="cdk-sops-secrets.SopsSyncProvider.Initializer"></a>

```typescript
import { SopsSyncProvider } from 'cdk-sops-secrets'

new SopsSyncProvider(scope: Construct, id?: string)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="cdk-sops-secrets.SopsSyncProvider.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Optional</sup> <a name="id" id="cdk-sops-secrets.SopsSyncProvider.Initializer.parameter.id"></a>

- *Type:* string

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.applyRemovalPolicy">applyRemovalPolicy</a></code> | Apply the given removal policy to this resource. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.addEventSource">addEventSource</a></code> | Adds an event source to this function. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.addEventSourceMapping">addEventSourceMapping</a></code> | Adds an event source that maps to this AWS Lambda function. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.addFunctionUrl">addFunctionUrl</a></code> | Adds a url to this lambda function. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.addPermission">addPermission</a></code> | Adds a permission to the Lambda resource policy. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.addToRolePolicy">addToRolePolicy</a></code> | Adds a statement to the IAM role assumed by the instance. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.configureAsyncInvoke">configureAsyncInvoke</a></code> | Configures options for asynchronous invocation. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.considerWarningOnInvokeFunctionPermissions">considerWarningOnInvokeFunctionPermissions</a></code> | A warning will be added to functions under the following conditions: - permissions that include `lambda:InvokeFunction` are added to the unqualified function. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.grantInvoke">grantInvoke</a></code> | Grant the given identity permissions to invoke this Lambda. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.grantInvokeCompositePrincipal">grantInvokeCompositePrincipal</a></code> | Grant multiple principals the ability to invoke this Lambda via CompositePrincipal. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.grantInvokeUrl">grantInvokeUrl</a></code> | Grant the given identity permissions to invoke this Lambda Function URL. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.metric">metric</a></code> | Return the given named metric for this Function. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.metricDuration">metricDuration</a></code> | How long execution of this Lambda takes. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.metricErrors">metricErrors</a></code> | How many invocations of this Lambda fail. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.metricInvocations">metricInvocations</a></code> | How often this Lambda is invoked. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.metricThrottles">metricThrottles</a></code> | How often this Lambda is throttled. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.addDependency">addDependency</a></code> | Using node.addDependency() does not work on this method as the underlying lambda function is modeled as a singleton across the stack. Use this method instead to declare dependencies. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.addEnvironment">addEnvironment</a></code> | Adds an environment variable to this Lambda function. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.addLayers">addLayers</a></code> | Adds one or more Lambda Layers to this Lambda function. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.dependOn">dependOn</a></code> | The SingletonFunction construct cannot be added as a dependency of another construct using node.addDependency(). Use this method instead to declare this as a dependency of another construct. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.addAgeKey">addAgeKey</a></code> | *No description.* |

---

##### `toString` <a name="toString" id="cdk-sops-secrets.SopsSyncProvider.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="cdk-sops-secrets.SopsSyncProvider.applyRemovalPolicy"></a>

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

###### `policy`<sup>Required</sup> <a name="policy" id="cdk-sops-secrets.SopsSyncProvider.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

##### `addEventSource` <a name="addEventSource" id="cdk-sops-secrets.SopsSyncProvider.addEventSource"></a>

```typescript
public addEventSource(source: IEventSource): void
```

Adds an event source to this function.

Event sources are implemented in the aws-cdk-lib/aws-lambda-event-sources module.

The following example adds an SQS Queue as an event source:
```
import { SqsEventSource } from 'aws-cdk-lib/aws-lambda-event-sources';
myFunction.addEventSource(new SqsEventSource(myQueue));
```

###### `source`<sup>Required</sup> <a name="source" id="cdk-sops-secrets.SopsSyncProvider.addEventSource.parameter.source"></a>

- *Type:* aws-cdk-lib.aws_lambda.IEventSource

---

##### `addEventSourceMapping` <a name="addEventSourceMapping" id="cdk-sops-secrets.SopsSyncProvider.addEventSourceMapping"></a>

```typescript
public addEventSourceMapping(id: string, options: EventSourceMappingOptions): EventSourceMapping
```

Adds an event source that maps to this AWS Lambda function.

###### `id`<sup>Required</sup> <a name="id" id="cdk-sops-secrets.SopsSyncProvider.addEventSourceMapping.parameter.id"></a>

- *Type:* string

---

###### `options`<sup>Required</sup> <a name="options" id="cdk-sops-secrets.SopsSyncProvider.addEventSourceMapping.parameter.options"></a>

- *Type:* aws-cdk-lib.aws_lambda.EventSourceMappingOptions

---

##### `addFunctionUrl` <a name="addFunctionUrl" id="cdk-sops-secrets.SopsSyncProvider.addFunctionUrl"></a>

```typescript
public addFunctionUrl(options?: FunctionUrlOptions): FunctionUrl
```

Adds a url to this lambda function.

###### `options`<sup>Optional</sup> <a name="options" id="cdk-sops-secrets.SopsSyncProvider.addFunctionUrl.parameter.options"></a>

- *Type:* aws-cdk-lib.aws_lambda.FunctionUrlOptions

---

##### `addPermission` <a name="addPermission" id="cdk-sops-secrets.SopsSyncProvider.addPermission"></a>

```typescript
public addPermission(name: string, permission: Permission): void
```

Adds a permission to the Lambda resource policy.

###### `name`<sup>Required</sup> <a name="name" id="cdk-sops-secrets.SopsSyncProvider.addPermission.parameter.name"></a>

- *Type:* string

---

###### `permission`<sup>Required</sup> <a name="permission" id="cdk-sops-secrets.SopsSyncProvider.addPermission.parameter.permission"></a>

- *Type:* aws-cdk-lib.aws_lambda.Permission

---

##### `addToRolePolicy` <a name="addToRolePolicy" id="cdk-sops-secrets.SopsSyncProvider.addToRolePolicy"></a>

```typescript
public addToRolePolicy(statement: PolicyStatement): void
```

Adds a statement to the IAM role assumed by the instance.

###### `statement`<sup>Required</sup> <a name="statement" id="cdk-sops-secrets.SopsSyncProvider.addToRolePolicy.parameter.statement"></a>

- *Type:* aws-cdk-lib.aws_iam.PolicyStatement

---

##### `configureAsyncInvoke` <a name="configureAsyncInvoke" id="cdk-sops-secrets.SopsSyncProvider.configureAsyncInvoke"></a>

```typescript
public configureAsyncInvoke(options: EventInvokeConfigOptions): void
```

Configures options for asynchronous invocation.

###### `options`<sup>Required</sup> <a name="options" id="cdk-sops-secrets.SopsSyncProvider.configureAsyncInvoke.parameter.options"></a>

- *Type:* aws-cdk-lib.aws_lambda.EventInvokeConfigOptions

---

##### `considerWarningOnInvokeFunctionPermissions` <a name="considerWarningOnInvokeFunctionPermissions" id="cdk-sops-secrets.SopsSyncProvider.considerWarningOnInvokeFunctionPermissions"></a>

```typescript
public considerWarningOnInvokeFunctionPermissions(scope: Construct, action: string): void
```

A warning will be added to functions under the following conditions: - permissions that include `lambda:InvokeFunction` are added to the unqualified function.

function.currentVersion is invoked before or after the permission is created.

This applies only to permissions on Lambda functions, not versions or aliases.
This function is overridden as a noOp for QualifiedFunctionBase.

###### `scope`<sup>Required</sup> <a name="scope" id="cdk-sops-secrets.SopsSyncProvider.considerWarningOnInvokeFunctionPermissions.parameter.scope"></a>

- *Type:* constructs.Construct

---

###### `action`<sup>Required</sup> <a name="action" id="cdk-sops-secrets.SopsSyncProvider.considerWarningOnInvokeFunctionPermissions.parameter.action"></a>

- *Type:* string

---

##### `grantInvoke` <a name="grantInvoke" id="cdk-sops-secrets.SopsSyncProvider.grantInvoke"></a>

```typescript
public grantInvoke(grantee: IGrantable): Grant
```

Grant the given identity permissions to invoke this Lambda.

###### `grantee`<sup>Required</sup> <a name="grantee" id="cdk-sops-secrets.SopsSyncProvider.grantInvoke.parameter.grantee"></a>

- *Type:* aws-cdk-lib.aws_iam.IGrantable

---

##### `grantInvokeCompositePrincipal` <a name="grantInvokeCompositePrincipal" id="cdk-sops-secrets.SopsSyncProvider.grantInvokeCompositePrincipal"></a>

```typescript
public grantInvokeCompositePrincipal(compositePrincipal: CompositePrincipal): Grant[]
```

Grant multiple principals the ability to invoke this Lambda via CompositePrincipal.

###### `compositePrincipal`<sup>Required</sup> <a name="compositePrincipal" id="cdk-sops-secrets.SopsSyncProvider.grantInvokeCompositePrincipal.parameter.compositePrincipal"></a>

- *Type:* aws-cdk-lib.aws_iam.CompositePrincipal

---

##### `grantInvokeUrl` <a name="grantInvokeUrl" id="cdk-sops-secrets.SopsSyncProvider.grantInvokeUrl"></a>

```typescript
public grantInvokeUrl(grantee: IGrantable): Grant
```

Grant the given identity permissions to invoke this Lambda Function URL.

###### `grantee`<sup>Required</sup> <a name="grantee" id="cdk-sops-secrets.SopsSyncProvider.grantInvokeUrl.parameter.grantee"></a>

- *Type:* aws-cdk-lib.aws_iam.IGrantable

---

##### `metric` <a name="metric" id="cdk-sops-secrets.SopsSyncProvider.metric"></a>

```typescript
public metric(metricName: string, props?: MetricOptions): Metric
```

Return the given named metric for this Function.

###### `metricName`<sup>Required</sup> <a name="metricName" id="cdk-sops-secrets.SopsSyncProvider.metric.parameter.metricName"></a>

- *Type:* string

---

###### `props`<sup>Optional</sup> <a name="props" id="cdk-sops-secrets.SopsSyncProvider.metric.parameter.props"></a>

- *Type:* aws-cdk-lib.aws_cloudwatch.MetricOptions

---

##### `metricDuration` <a name="metricDuration" id="cdk-sops-secrets.SopsSyncProvider.metricDuration"></a>

```typescript
public metricDuration(props?: MetricOptions): Metric
```

How long execution of this Lambda takes.

Average over 5 minutes

###### `props`<sup>Optional</sup> <a name="props" id="cdk-sops-secrets.SopsSyncProvider.metricDuration.parameter.props"></a>

- *Type:* aws-cdk-lib.aws_cloudwatch.MetricOptions

---

##### `metricErrors` <a name="metricErrors" id="cdk-sops-secrets.SopsSyncProvider.metricErrors"></a>

```typescript
public metricErrors(props?: MetricOptions): Metric
```

How many invocations of this Lambda fail.

Sum over 5 minutes

###### `props`<sup>Optional</sup> <a name="props" id="cdk-sops-secrets.SopsSyncProvider.metricErrors.parameter.props"></a>

- *Type:* aws-cdk-lib.aws_cloudwatch.MetricOptions

---

##### `metricInvocations` <a name="metricInvocations" id="cdk-sops-secrets.SopsSyncProvider.metricInvocations"></a>

```typescript
public metricInvocations(props?: MetricOptions): Metric
```

How often this Lambda is invoked.

Sum over 5 minutes

###### `props`<sup>Optional</sup> <a name="props" id="cdk-sops-secrets.SopsSyncProvider.metricInvocations.parameter.props"></a>

- *Type:* aws-cdk-lib.aws_cloudwatch.MetricOptions

---

##### `metricThrottles` <a name="metricThrottles" id="cdk-sops-secrets.SopsSyncProvider.metricThrottles"></a>

```typescript
public metricThrottles(props?: MetricOptions): Metric
```

How often this Lambda is throttled.

Sum over 5 minutes

###### `props`<sup>Optional</sup> <a name="props" id="cdk-sops-secrets.SopsSyncProvider.metricThrottles.parameter.props"></a>

- *Type:* aws-cdk-lib.aws_cloudwatch.MetricOptions

---

##### `addDependency` <a name="addDependency" id="cdk-sops-secrets.SopsSyncProvider.addDependency"></a>

```typescript
public addDependency(up: IDependable): void
```

Using node.addDependency() does not work on this method as the underlying lambda function is modeled as a singleton across the stack. Use this method instead to declare dependencies.

###### `up`<sup>Required</sup> <a name="up" id="cdk-sops-secrets.SopsSyncProvider.addDependency.parameter.up"></a>

- *Type:* constructs.IDependable

---

##### `addEnvironment` <a name="addEnvironment" id="cdk-sops-secrets.SopsSyncProvider.addEnvironment"></a>

```typescript
public addEnvironment(key: string, value: string, options?: EnvironmentOptions): Function
```

Adds an environment variable to this Lambda function.

If this is a ref to a Lambda function, this operation results in a no-op.

###### `key`<sup>Required</sup> <a name="key" id="cdk-sops-secrets.SopsSyncProvider.addEnvironment.parameter.key"></a>

- *Type:* string

The environment variable key.

---

###### `value`<sup>Required</sup> <a name="value" id="cdk-sops-secrets.SopsSyncProvider.addEnvironment.parameter.value"></a>

- *Type:* string

The environment variable's value.

---

###### `options`<sup>Optional</sup> <a name="options" id="cdk-sops-secrets.SopsSyncProvider.addEnvironment.parameter.options"></a>

- *Type:* aws-cdk-lib.aws_lambda.EnvironmentOptions

Environment variable options.

---

##### `addLayers` <a name="addLayers" id="cdk-sops-secrets.SopsSyncProvider.addLayers"></a>

```typescript
public addLayers(layers: ILayerVersion): void
```

Adds one or more Lambda Layers to this Lambda function.

###### `layers`<sup>Required</sup> <a name="layers" id="cdk-sops-secrets.SopsSyncProvider.addLayers.parameter.layers"></a>

- *Type:* aws-cdk-lib.aws_lambda.ILayerVersion

the layers to be added.

---

##### `dependOn` <a name="dependOn" id="cdk-sops-secrets.SopsSyncProvider.dependOn"></a>

```typescript
public dependOn(down: IConstruct): void
```

The SingletonFunction construct cannot be added as a dependency of another construct using node.addDependency(). Use this method instead to declare this as a dependency of another construct.

###### `down`<sup>Required</sup> <a name="down" id="cdk-sops-secrets.SopsSyncProvider.dependOn.parameter.down"></a>

- *Type:* constructs.IConstruct

---

##### `addAgeKey` <a name="addAgeKey" id="cdk-sops-secrets.SopsSyncProvider.addAgeKey"></a>

```typescript
public addAgeKey(key: SecretValue): void
```

###### `key`<sup>Required</sup> <a name="key" id="cdk-sops-secrets.SopsSyncProvider.addAgeKey.parameter.key"></a>

- *Type:* aws-cdk-lib.SecretValue

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.isOwnedResource">isOwnedResource</a></code> | Returns true if the construct was created by CDK, and false otherwise. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.isResource">isResource</a></code> | Check whether the given construct is a Resource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="cdk-sops-secrets.SopsSyncProvider.isConstruct"></a>

```typescript
import { SopsSyncProvider } from 'cdk-sops-secrets'

SopsSyncProvider.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="cdk-sops-secrets.SopsSyncProvider.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isOwnedResource` <a name="isOwnedResource" id="cdk-sops-secrets.SopsSyncProvider.isOwnedResource"></a>

```typescript
import { SopsSyncProvider } from 'cdk-sops-secrets'

SopsSyncProvider.isOwnedResource(construct: IConstruct)
```

Returns true if the construct was created by CDK, and false otherwise.

###### `construct`<sup>Required</sup> <a name="construct" id="cdk-sops-secrets.SopsSyncProvider.isOwnedResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

##### `isResource` <a name="isResource" id="cdk-sops-secrets.SopsSyncProvider.isResource"></a>

```typescript
import { SopsSyncProvider } from 'cdk-sops-secrets'

SopsSyncProvider.isResource(construct: IConstruct)
```

Check whether the given construct is a Resource.

###### `construct`<sup>Required</sup> <a name="construct" id="cdk-sops-secrets.SopsSyncProvider.isResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.property.env">env</a></code> | <code>aws-cdk-lib.ResourceEnvironment</code> | The environment this resource belongs to. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this resource is defined. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.property.architecture">architecture</a></code> | <code>aws-cdk-lib.aws_lambda.Architecture</code> | The architecture of this Lambda Function. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.property.connections">connections</a></code> | <code>aws-cdk-lib.aws_ec2.Connections</code> | Access the Connections object. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.property.functionArn">functionArn</a></code> | <code>string</code> | The ARN fo the function. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.property.functionName">functionName</a></code> | <code>string</code> | The name of the function. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.property.grantPrincipal">grantPrincipal</a></code> | <code>aws-cdk-lib.aws_iam.IPrincipal</code> | The principal this Lambda Function is running as. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.property.isBoundToVpc">isBoundToVpc</a></code> | <code>boolean</code> | Whether or not this Lambda function was bound to a VPC. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.property.latestVersion">latestVersion</a></code> | <code>aws-cdk-lib.aws_lambda.IVersion</code> | The `$LATEST` version of this function. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.property.permissionsNode">permissionsNode</a></code> | <code>constructs.Node</code> | The construct node where permissions are attached. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.property.resourceArnsForGrantInvoke">resourceArnsForGrantInvoke</a></code> | <code>string[]</code> | The ARN(s) to put into the resource field of the generated IAM policy for grantInvoke(). |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.property.role">role</a></code> | <code>aws-cdk-lib.aws_iam.IRole</code> | The IAM role associated with this function. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.property.currentVersion">currentVersion</a></code> | <code>aws-cdk-lib.aws_lambda.Version</code> | Returns a `lambda.Version` which represents the current version of this singleton Lambda function. A new version will be created every time the function's configuration changes. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.property.logGroup">logGroup</a></code> | <code>aws-cdk-lib.aws_logs.ILogGroup</code> | The LogGroup where the Lambda function's logs are made available. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.property.runtime">runtime</a></code> | <code>aws-cdk-lib.aws_lambda.Runtime</code> | The runtime environment for the Lambda function. |

---

##### `node`<sup>Required</sup> <a name="node" id="cdk-sops-secrets.SopsSyncProvider.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `env`<sup>Required</sup> <a name="env" id="cdk-sops-secrets.SopsSyncProvider.property.env"></a>

```typescript
public readonly env: ResourceEnvironment;
```

- *Type:* aws-cdk-lib.ResourceEnvironment

The environment this resource belongs to.

For resources that are created and managed by the CDK
(generally, those created by creating new class instances like Role, Bucket, etc.),
this is always the same as the environment of the stack they belong to;
however, for imported resources
(those obtained from static methods like fromRoleArn, fromBucketName, etc.),
that might be different than the stack they were imported into.

---

##### `stack`<sup>Required</sup> <a name="stack" id="cdk-sops-secrets.SopsSyncProvider.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this resource is defined.

---

##### `architecture`<sup>Required</sup> <a name="architecture" id="cdk-sops-secrets.SopsSyncProvider.property.architecture"></a>

```typescript
public readonly architecture: Architecture;
```

- *Type:* aws-cdk-lib.aws_lambda.Architecture

The architecture of this Lambda Function.

---

##### `connections`<sup>Required</sup> <a name="connections" id="cdk-sops-secrets.SopsSyncProvider.property.connections"></a>

```typescript
public readonly connections: Connections;
```

- *Type:* aws-cdk-lib.aws_ec2.Connections

Access the Connections object.

Will fail if not a VPC-enabled Lambda Function

---

##### `functionArn`<sup>Required</sup> <a name="functionArn" id="cdk-sops-secrets.SopsSyncProvider.property.functionArn"></a>

```typescript
public readonly functionArn: string;
```

- *Type:* string

The ARN fo the function.

---

##### `functionName`<sup>Required</sup> <a name="functionName" id="cdk-sops-secrets.SopsSyncProvider.property.functionName"></a>

```typescript
public readonly functionName: string;
```

- *Type:* string

The name of the function.

---

##### `grantPrincipal`<sup>Required</sup> <a name="grantPrincipal" id="cdk-sops-secrets.SopsSyncProvider.property.grantPrincipal"></a>

```typescript
public readonly grantPrincipal: IPrincipal;
```

- *Type:* aws-cdk-lib.aws_iam.IPrincipal

The principal this Lambda Function is running as.

---

##### `isBoundToVpc`<sup>Required</sup> <a name="isBoundToVpc" id="cdk-sops-secrets.SopsSyncProvider.property.isBoundToVpc"></a>

```typescript
public readonly isBoundToVpc: boolean;
```

- *Type:* boolean

Whether or not this Lambda function was bound to a VPC.

If this is is `false`, trying to access the `connections` object will fail.

---

##### `latestVersion`<sup>Required</sup> <a name="latestVersion" id="cdk-sops-secrets.SopsSyncProvider.property.latestVersion"></a>

```typescript
public readonly latestVersion: IVersion;
```

- *Type:* aws-cdk-lib.aws_lambda.IVersion

The `$LATEST` version of this function.

Note that this is reference to a non-specific AWS Lambda version, which
means the function this version refers to can return different results in
different invocations.

To obtain a reference to an explicit version which references the current
function configuration, use `lambdaFunction.currentVersion` instead.

---

##### `permissionsNode`<sup>Required</sup> <a name="permissionsNode" id="cdk-sops-secrets.SopsSyncProvider.property.permissionsNode"></a>

```typescript
public readonly permissionsNode: Node;
```

- *Type:* constructs.Node

The construct node where permissions are attached.

---

##### `resourceArnsForGrantInvoke`<sup>Required</sup> <a name="resourceArnsForGrantInvoke" id="cdk-sops-secrets.SopsSyncProvider.property.resourceArnsForGrantInvoke"></a>

```typescript
public readonly resourceArnsForGrantInvoke: string[];
```

- *Type:* string[]

The ARN(s) to put into the resource field of the generated IAM policy for grantInvoke().

---

##### `role`<sup>Optional</sup> <a name="role" id="cdk-sops-secrets.SopsSyncProvider.property.role"></a>

```typescript
public readonly role: IRole;
```

- *Type:* aws-cdk-lib.aws_iam.IRole

The IAM role associated with this function.

Undefined if the function was imported without a role.

---

##### `currentVersion`<sup>Required</sup> <a name="currentVersion" id="cdk-sops-secrets.SopsSyncProvider.property.currentVersion"></a>

```typescript
public readonly currentVersion: Version;
```

- *Type:* aws-cdk-lib.aws_lambda.Version

Returns a `lambda.Version` which represents the current version of this singleton Lambda function. A new version will be created every time the function's configuration changes.

You can specify options for this version using the `currentVersionOptions`
prop when initializing the `lambda.SingletonFunction`.

---

##### `logGroup`<sup>Required</sup> <a name="logGroup" id="cdk-sops-secrets.SopsSyncProvider.property.logGroup"></a>

```typescript
public readonly logGroup: ILogGroup;
```

- *Type:* aws-cdk-lib.aws_logs.ILogGroup

The LogGroup where the Lambda function's logs are made available.

If either `logRetention` is set or this property is called, a CloudFormation custom resource is added to the stack that
pre-creates the log group as part of the stack deployment, if it already doesn't exist, and sets the correct log retention
period (never expire, by default).

Further, if the log group already exists and the `logRetention` is not set, the custom resource will reset the log retention
to never expire even if it was configured with a different value.

---

##### `runtime`<sup>Required</sup> <a name="runtime" id="cdk-sops-secrets.SopsSyncProvider.property.runtime"></a>

```typescript
public readonly runtime: Runtime;
```

- *Type:* aws-cdk-lib.aws_lambda.Runtime

The runtime environment for the Lambda function.

---


## Structs <a name="Structs" id="Structs"></a>

### MultiStringParameterProps <a name="MultiStringParameterProps" id="cdk-sops-secrets.MultiStringParameterProps"></a>

#### Initializer <a name="Initializer" id="cdk-sops-secrets.MultiStringParameterProps.Initializer"></a>

```typescript
import { MultiStringParameterProps } from 'cdk-sops-secrets'

const multiStringParameterProps: MultiStringParameterProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.MultiStringParameterProps.property.convertToJSON">convertToJSON</a></code> | <code>boolean</code> | Should the encrypted sops value should be converted to JSON? |
| <code><a href="#cdk-sops-secrets.MultiStringParameterProps.property.creationType">creationType</a></code> | <code><a href="#cdk-sops-secrets.CreationType">CreationType</a></code> | *No description.* |
| <code><a href="#cdk-sops-secrets.MultiStringParameterProps.property.flatten">flatten</a></code> | <code>boolean</code> | Should the structure be flattened? |
| <code><a href="#cdk-sops-secrets.MultiStringParameterProps.property.flattenSeparator">flattenSeparator</a></code> | <code>string</code> | If the structure should be flattened use the provided separator between keys. |
| <code><a href="#cdk-sops-secrets.MultiStringParameterProps.property.resourceType">resourceType</a></code> | <code><a href="#cdk-sops-secrets.ResourceType">ResourceType</a></code> | *No description.* |
| <code><a href="#cdk-sops-secrets.MultiStringParameterProps.property.sopsAgeKey">sopsAgeKey</a></code> | <code>aws-cdk-lib.SecretValue</code> | The age key that should be used for encryption. |
| <code><a href="#cdk-sops-secrets.MultiStringParameterProps.property.sopsFileFormat">sopsFileFormat</a></code> | <code>string</code> | The format of the sops file. |
| <code><a href="#cdk-sops-secrets.MultiStringParameterProps.property.sopsFilePath">sopsFilePath</a></code> | <code>string</code> | The filepath to the sops file. |
| <code><a href="#cdk-sops-secrets.MultiStringParameterProps.property.sopsKmsKey">sopsKmsKey</a></code> | <code>aws-cdk-lib.aws_kms.IKey[]</code> | The kmsKey used to encrypt the sops file. |
| <code><a href="#cdk-sops-secrets.MultiStringParameterProps.property.sopsProvider">sopsProvider</a></code> | <code><a href="#cdk-sops-secrets.SopsSyncProvider">SopsSyncProvider</a></code> | The custom resource provider to use. |
| <code><a href="#cdk-sops-secrets.MultiStringParameterProps.property.sopsS3Bucket">sopsS3Bucket</a></code> | <code>string</code> | If you want to pass the sops file via s3, you can specify the bucket you can use cfn parameter here Both, sopsS3Bucket and sopsS3Key have to be specified. |
| <code><a href="#cdk-sops-secrets.MultiStringParameterProps.property.sopsS3Key">sopsS3Key</a></code> | <code>string</code> | If you want to pass the sops file via s3, you can specify the key inside the bucket you can use cfn parameter here Both, sopsS3Bucket and sopsS3Key have to be specified. |
| <code><a href="#cdk-sops-secrets.MultiStringParameterProps.property.stringifyValues">stringifyValues</a></code> | <code>boolean</code> | Shall all values be flattened? |
| <code><a href="#cdk-sops-secrets.MultiStringParameterProps.property.uploadType">uploadType</a></code> | <code><a href="#cdk-sops-secrets.UploadType">UploadType</a></code> | How should the secret be passed to the CustomResource? |
| <code><a href="#cdk-sops-secrets.MultiStringParameterProps.property.keySeperator">keySeperator</a></code> | <code>string</code> | *No description.* |

---

##### `convertToJSON`<sup>Optional</sup> <a name="convertToJSON" id="cdk-sops-secrets.MultiStringParameterProps.property.convertToJSON"></a>

```typescript
public readonly convertToJSON: boolean;
```

- *Type:* boolean
- *Default:* true

Should the encrypted sops value should be converted to JSON?

Only JSON can be handled by cloud formations dynamic references.

---

##### `creationType`<sup>Optional</sup> <a name="creationType" id="cdk-sops-secrets.MultiStringParameterProps.property.creationType"></a>

```typescript
public readonly creationType: CreationType;
```

- *Type:* <a href="#cdk-sops-secrets.CreationType">CreationType</a>

---

##### `flatten`<sup>Optional</sup> <a name="flatten" id="cdk-sops-secrets.MultiStringParameterProps.property.flatten"></a>

```typescript
public readonly flatten: boolean;
```

- *Type:* boolean
- *Default:* true

Should the structure be flattened?

The result will be a flat structure and all
object keys will be replaced with the full jsonpath as key.
This is usefull for dynamic references, as those don't support nested objects.

---

##### `flattenSeparator`<sup>Optional</sup> <a name="flattenSeparator" id="cdk-sops-secrets.MultiStringParameterProps.property.flattenSeparator"></a>

```typescript
public readonly flattenSeparator: string;
```

- *Type:* string
- *Default:* '.'

If the structure should be flattened use the provided separator between keys.

---

##### `resourceType`<sup>Optional</sup> <a name="resourceType" id="cdk-sops-secrets.MultiStringParameterProps.property.resourceType"></a>

```typescript
public readonly resourceType: ResourceType;
```

- *Type:* <a href="#cdk-sops-secrets.ResourceType">ResourceType</a>

---

##### `sopsAgeKey`<sup>Optional</sup> <a name="sopsAgeKey" id="cdk-sops-secrets.MultiStringParameterProps.property.sopsAgeKey"></a>

```typescript
public readonly sopsAgeKey: SecretValue;
```

- *Type:* aws-cdk-lib.SecretValue

The age key that should be used for encryption.

---

##### `sopsFileFormat`<sup>Optional</sup> <a name="sopsFileFormat" id="cdk-sops-secrets.MultiStringParameterProps.property.sopsFileFormat"></a>

```typescript
public readonly sopsFileFormat: string;
```

- *Type:* string
- *Default:* The fileformat will be derived from the file ending

The format of the sops file.

---

##### `sopsFilePath`<sup>Optional</sup> <a name="sopsFilePath" id="cdk-sops-secrets.MultiStringParameterProps.property.sopsFilePath"></a>

```typescript
public readonly sopsFilePath: string;
```

- *Type:* string

The filepath to the sops file.

---

##### `sopsKmsKey`<sup>Optional</sup> <a name="sopsKmsKey" id="cdk-sops-secrets.MultiStringParameterProps.property.sopsKmsKey"></a>

```typescript
public readonly sopsKmsKey: IKey[];
```

- *Type:* aws-cdk-lib.aws_kms.IKey[]
- *Default:* The key will be derived from the sops file

The kmsKey used to encrypt the sops file.

Encrypt permissions
will be granted to the custom resource provider.

---

##### `sopsProvider`<sup>Optional</sup> <a name="sopsProvider" id="cdk-sops-secrets.MultiStringParameterProps.property.sopsProvider"></a>

```typescript
public readonly sopsProvider: SopsSyncProvider;
```

- *Type:* <a href="#cdk-sops-secrets.SopsSyncProvider">SopsSyncProvider</a>
- *Default:* A new singleton provider will be created

The custom resource provider to use.

If you don't specify any, a new
provider will be created - or if already exists within this stack - reused.

---

##### `sopsS3Bucket`<sup>Optional</sup> <a name="sopsS3Bucket" id="cdk-sops-secrets.MultiStringParameterProps.property.sopsS3Bucket"></a>

```typescript
public readonly sopsS3Bucket: string;
```

- *Type:* string

If you want to pass the sops file via s3, you can specify the bucket you can use cfn parameter here Both, sopsS3Bucket and sopsS3Key have to be specified.

---

##### `sopsS3Key`<sup>Optional</sup> <a name="sopsS3Key" id="cdk-sops-secrets.MultiStringParameterProps.property.sopsS3Key"></a>

```typescript
public readonly sopsS3Key: string;
```

- *Type:* string

If you want to pass the sops file via s3, you can specify the key inside the bucket you can use cfn parameter here Both, sopsS3Bucket and sopsS3Key have to be specified.

---

##### `stringifyValues`<sup>Optional</sup> <a name="stringifyValues" id="cdk-sops-secrets.MultiStringParameterProps.property.stringifyValues"></a>

```typescript
public readonly stringifyValues: boolean;
```

- *Type:* boolean

Shall all values be flattened?

This is usefull for dynamic references, as there
are lookup errors for certain float types

---

##### `uploadType`<sup>Optional</sup> <a name="uploadType" id="cdk-sops-secrets.MultiStringParameterProps.property.uploadType"></a>

```typescript
public readonly uploadType: UploadType;
```

- *Type:* <a href="#cdk-sops-secrets.UploadType">UploadType</a>
- *Default:* INLINE

How should the secret be passed to the CustomResource?

---

##### `keySeperator`<sup>Optional</sup> <a name="keySeperator" id="cdk-sops-secrets.MultiStringParameterProps.property.keySeperator"></a>

```typescript
public readonly keySeperator: string;
```

- *Type:* string

---

### SopsSecretProps <a name="SopsSecretProps" id="cdk-sops-secrets.SopsSecretProps"></a>

The configuration options of the SopsSecret.

#### Initializer <a name="Initializer" id="cdk-sops-secrets.SopsSecretProps.Initializer"></a>

```typescript
import { SopsSecretProps } from 'cdk-sops-secrets'

const sopsSecretProps: SopsSecretProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.description">description</a></code> | <code>string</code> | An optional, human-friendly description of the secret. |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.encryptionKey">encryptionKey</a></code> | <code>aws-cdk-lib.aws_kms.IKey</code> | The customer-managed encryption key to use for encrypting the secret value. |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.generateSecretString">generateSecretString</a></code> | <code>aws-cdk-lib.aws_secretsmanager.SecretStringGenerator</code> | Configuration for how to generate a secret value. |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.removalPolicy">removalPolicy</a></code> | <code>aws-cdk-lib.RemovalPolicy</code> | Policy to apply when the secret is removed from this stack. |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.replicaRegions">replicaRegions</a></code> | <code>aws-cdk-lib.aws_secretsmanager.ReplicaRegion[]</code> | A list of regions where to replicate this secret. |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.secretName">secretName</a></code> | <code>string</code> | A name for the secret. |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.secretObjectValue">secretObjectValue</a></code> | <code>{[ key: string ]: aws-cdk-lib.SecretValue}</code> | Initial value for a JSON secret. |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.secretStringBeta1">secretStringBeta1</a></code> | <code>aws-cdk-lib.aws_secretsmanager.SecretStringValueBeta1</code> | Initial value for the secret. |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.secretStringValue">secretStringValue</a></code> | <code>aws-cdk-lib.SecretValue</code> | Initial value for the secret. |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.convertToJSON">convertToJSON</a></code> | <code>boolean</code> | Should the encrypted sops value should be converted to JSON? |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.creationType">creationType</a></code> | <code><a href="#cdk-sops-secrets.CreationType">CreationType</a></code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.flatten">flatten</a></code> | <code>boolean</code> | Should the structure be flattened? |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.flattenSeparator">flattenSeparator</a></code> | <code>string</code> | If the structure should be flattened use the provided separator between keys. |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.resourceType">resourceType</a></code> | <code><a href="#cdk-sops-secrets.ResourceType">ResourceType</a></code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.sopsAgeKey">sopsAgeKey</a></code> | <code>aws-cdk-lib.SecretValue</code> | The age key that should be used for encryption. |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.sopsFileFormat">sopsFileFormat</a></code> | <code>string</code> | The format of the sops file. |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.sopsFilePath">sopsFilePath</a></code> | <code>string</code> | The filepath to the sops file. |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.sopsKmsKey">sopsKmsKey</a></code> | <code>aws-cdk-lib.aws_kms.IKey[]</code> | The kmsKey used to encrypt the sops file. |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.sopsProvider">sopsProvider</a></code> | <code><a href="#cdk-sops-secrets.SopsSyncProvider">SopsSyncProvider</a></code> | The custom resource provider to use. |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.sopsS3Bucket">sopsS3Bucket</a></code> | <code>string</code> | If you want to pass the sops file via s3, you can specify the bucket you can use cfn parameter here Both, sopsS3Bucket and sopsS3Key have to be specified. |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.sopsS3Key">sopsS3Key</a></code> | <code>string</code> | If you want to pass the sops file via s3, you can specify the key inside the bucket you can use cfn parameter here Both, sopsS3Bucket and sopsS3Key have to be specified. |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.stringifyValues">stringifyValues</a></code> | <code>boolean</code> | Shall all values be flattened? |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.uploadType">uploadType</a></code> | <code><a href="#cdk-sops-secrets.UploadType">UploadType</a></code> | How should the secret be passed to the CustomResource? |

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

- *Type:* aws-cdk-lib.aws_kms.IKey
- *Default:* A default KMS key for the account and region is used.

The customer-managed encryption key to use for encrypting the secret value.

---

##### `generateSecretString`<sup>Optional</sup> <a name="generateSecretString" id="cdk-sops-secrets.SopsSecretProps.property.generateSecretString"></a>

```typescript
public readonly generateSecretString: SecretStringGenerator;
```

- *Type:* aws-cdk-lib.aws_secretsmanager.SecretStringGenerator
- *Default:* 32 characters with upper-case letters, lower-case letters, punctuation and numbers (at least one from each category), per the default values of ``SecretStringGenerator``.

Configuration for how to generate a secret value.

Only one of `secretString` and `generateSecretString` can be provided.

---

##### `removalPolicy`<sup>Optional</sup> <a name="removalPolicy" id="cdk-sops-secrets.SopsSecretProps.property.removalPolicy"></a>

```typescript
public readonly removalPolicy: RemovalPolicy;
```

- *Type:* aws-cdk-lib.RemovalPolicy
- *Default:* Not set.

Policy to apply when the secret is removed from this stack.

---

##### `replicaRegions`<sup>Optional</sup> <a name="replicaRegions" id="cdk-sops-secrets.SopsSecretProps.property.replicaRegions"></a>

```typescript
public readonly replicaRegions: ReplicaRegion[];
```

- *Type:* aws-cdk-lib.aws_secretsmanager.ReplicaRegion[]
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

##### `secretObjectValue`<sup>Optional</sup> <a name="secretObjectValue" id="cdk-sops-secrets.SopsSecretProps.property.secretObjectValue"></a>

```typescript
public readonly secretObjectValue: {[ key: string ]: SecretValue};
```

- *Type:* {[ key: string ]: aws-cdk-lib.SecretValue}
- *Default:* SecretsManager generates a new secret value.

Initial value for a JSON secret.

**NOTE:** *It is **highly** encouraged to leave this field undefined and allow SecretsManager to create the secret value.
The secret object -- if provided -- will be included in the output of the cdk as part of synthesis,
and will appear in the CloudFormation template in the console. This can be secure(-ish) if that value is merely reference to
another resource (or one of its attributes), but if the value is a plaintext string, it will be visible to anyone with access
to the CloudFormation template (via the AWS Console, SDKs, or CLI).

Specifies a JSON object that you want to encrypt and store in this new version of the secret.
To specify a simple string value instead, use `SecretProps.secretStringValue`

Only one of `secretStringBeta1`, `secretStringValue`, 'secretObjectValue', and `generateSecretString` can be provided.

---

*Example*

```typescript
declare const user: iam.User;
declare const accessKey: iam.AccessKey;
declare const stack: Stack;
new secretsmanager.Secret(stack, 'JSONSecret', {
  secretObjectValue: {
    username: SecretValue.unsafePlainText(user.userName), // intrinsic reference, not exposed as plaintext
    database: SecretValue.unsafePlainText('foo'), // rendered as plain text, but not a secret
    password: accessKey.secretAccessKey, // SecretValue
  },
});
```


##### ~~`secretStringBeta1`~~<sup>Optional</sup> <a name="secretStringBeta1" id="cdk-sops-secrets.SopsSecretProps.property.secretStringBeta1"></a>

- *Deprecated:* Use `secretStringValue` instead.

```typescript
public readonly secretStringBeta1: SecretStringValueBeta1;
```

- *Type:* aws-cdk-lib.aws_secretsmanager.SecretStringValueBeta1
- *Default:* SecretsManager generates a new secret value.

Initial value for the secret.

**NOTE:** *It is **highly** encouraged to leave this field undefined and allow SecretsManager to create the secret value.
The secret string -- if provided -- will be included in the output of the cdk as part of synthesis,
and will appear in the CloudFormation template in the console. This can be secure(-ish) if that value is merely reference to
another resource (or one of its attributes), but if the value is a plaintext string, it will be visible to anyone with access
to the CloudFormation template (via the AWS Console, SDKs, or CLI).

Specifies text data that you want to encrypt and store in this new version of the secret.
May be a simple string value, or a string representation of a JSON structure.

Only one of `secretStringBeta1`, `secretStringValue`, and `generateSecretString` can be provided.

---

##### `secretStringValue`<sup>Optional</sup> <a name="secretStringValue" id="cdk-sops-secrets.SopsSecretProps.property.secretStringValue"></a>

```typescript
public readonly secretStringValue: SecretValue;
```

- *Type:* aws-cdk-lib.SecretValue
- *Default:* SecretsManager generates a new secret value.

Initial value for the secret.

**NOTE:** *It is **highly** encouraged to leave this field undefined and allow SecretsManager to create the secret value.
The secret string -- if provided -- will be included in the output of the cdk as part of synthesis,
and will appear in the CloudFormation template in the console. This can be secure(-ish) if that value is merely reference to
another resource (or one of its attributes), but if the value is a plaintext string, it will be visible to anyone with access
to the CloudFormation template (via the AWS Console, SDKs, or CLI).

Specifies text data that you want to encrypt and store in this new version of the secret.
May be a simple string value. To provide a string representation of JSON structure, use `SecretProps.secretObjectValue` instead.

Only one of `secretStringBeta1`, `secretStringValue`, 'secretObjectValue', and `generateSecretString` can be provided.

---

##### `convertToJSON`<sup>Optional</sup> <a name="convertToJSON" id="cdk-sops-secrets.SopsSecretProps.property.convertToJSON"></a>

```typescript
public readonly convertToJSON: boolean;
```

- *Type:* boolean
- *Default:* true

Should the encrypted sops value should be converted to JSON?

Only JSON can be handled by cloud formations dynamic references.

---

##### `creationType`<sup>Optional</sup> <a name="creationType" id="cdk-sops-secrets.SopsSecretProps.property.creationType"></a>

```typescript
public readonly creationType: CreationType;
```

- *Type:* <a href="#cdk-sops-secrets.CreationType">CreationType</a>

---

##### `flatten`<sup>Optional</sup> <a name="flatten" id="cdk-sops-secrets.SopsSecretProps.property.flatten"></a>

```typescript
public readonly flatten: boolean;
```

- *Type:* boolean
- *Default:* true

Should the structure be flattened?

The result will be a flat structure and all
object keys will be replaced with the full jsonpath as key.
This is usefull for dynamic references, as those don't support nested objects.

---

##### `flattenSeparator`<sup>Optional</sup> <a name="flattenSeparator" id="cdk-sops-secrets.SopsSecretProps.property.flattenSeparator"></a>

```typescript
public readonly flattenSeparator: string;
```

- *Type:* string
- *Default:* '.'

If the structure should be flattened use the provided separator between keys.

---

##### `resourceType`<sup>Optional</sup> <a name="resourceType" id="cdk-sops-secrets.SopsSecretProps.property.resourceType"></a>

```typescript
public readonly resourceType: ResourceType;
```

- *Type:* <a href="#cdk-sops-secrets.ResourceType">ResourceType</a>

---

##### `sopsAgeKey`<sup>Optional</sup> <a name="sopsAgeKey" id="cdk-sops-secrets.SopsSecretProps.property.sopsAgeKey"></a>

```typescript
public readonly sopsAgeKey: SecretValue;
```

- *Type:* aws-cdk-lib.SecretValue

The age key that should be used for encryption.

---

##### `sopsFileFormat`<sup>Optional</sup> <a name="sopsFileFormat" id="cdk-sops-secrets.SopsSecretProps.property.sopsFileFormat"></a>

```typescript
public readonly sopsFileFormat: string;
```

- *Type:* string
- *Default:* The fileformat will be derived from the file ending

The format of the sops file.

---

##### `sopsFilePath`<sup>Optional</sup> <a name="sopsFilePath" id="cdk-sops-secrets.SopsSecretProps.property.sopsFilePath"></a>

```typescript
public readonly sopsFilePath: string;
```

- *Type:* string

The filepath to the sops file.

---

##### `sopsKmsKey`<sup>Optional</sup> <a name="sopsKmsKey" id="cdk-sops-secrets.SopsSecretProps.property.sopsKmsKey"></a>

```typescript
public readonly sopsKmsKey: IKey[];
```

- *Type:* aws-cdk-lib.aws_kms.IKey[]
- *Default:* The key will be derived from the sops file

The kmsKey used to encrypt the sops file.

Encrypt permissions
will be granted to the custom resource provider.

---

##### `sopsProvider`<sup>Optional</sup> <a name="sopsProvider" id="cdk-sops-secrets.SopsSecretProps.property.sopsProvider"></a>

```typescript
public readonly sopsProvider: SopsSyncProvider;
```

- *Type:* <a href="#cdk-sops-secrets.SopsSyncProvider">SopsSyncProvider</a>
- *Default:* A new singleton provider will be created

The custom resource provider to use.

If you don't specify any, a new
provider will be created - or if already exists within this stack - reused.

---

##### `sopsS3Bucket`<sup>Optional</sup> <a name="sopsS3Bucket" id="cdk-sops-secrets.SopsSecretProps.property.sopsS3Bucket"></a>

```typescript
public readonly sopsS3Bucket: string;
```

- *Type:* string

If you want to pass the sops file via s3, you can specify the bucket you can use cfn parameter here Both, sopsS3Bucket and sopsS3Key have to be specified.

---

##### `sopsS3Key`<sup>Optional</sup> <a name="sopsS3Key" id="cdk-sops-secrets.SopsSecretProps.property.sopsS3Key"></a>

```typescript
public readonly sopsS3Key: string;
```

- *Type:* string

If you want to pass the sops file via s3, you can specify the key inside the bucket you can use cfn parameter here Both, sopsS3Bucket and sopsS3Key have to be specified.

---

##### `stringifyValues`<sup>Optional</sup> <a name="stringifyValues" id="cdk-sops-secrets.SopsSecretProps.property.stringifyValues"></a>

```typescript
public readonly stringifyValues: boolean;
```

- *Type:* boolean

Shall all values be flattened?

This is usefull for dynamic references, as there
are lookup errors for certain float types

---

##### `uploadType`<sup>Optional</sup> <a name="uploadType" id="cdk-sops-secrets.SopsSecretProps.property.uploadType"></a>

```typescript
public readonly uploadType: UploadType;
```

- *Type:* <a href="#cdk-sops-secrets.UploadType">UploadType</a>
- *Default:* INLINE

How should the secret be passed to the CustomResource?

---

### SopsStringParameterProps <a name="SopsStringParameterProps" id="cdk-sops-secrets.SopsStringParameterProps"></a>

The configuration options of the StringParameter.

#### Initializer <a name="Initializer" id="cdk-sops-secrets.SopsStringParameterProps.Initializer"></a>

```typescript
import { SopsStringParameterProps } from 'cdk-sops-secrets'

const sopsStringParameterProps: SopsStringParameterProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.SopsStringParameterProps.property.convertToJSON">convertToJSON</a></code> | <code>boolean</code> | Should the encrypted sops value should be converted to JSON? |
| <code><a href="#cdk-sops-secrets.SopsStringParameterProps.property.creationType">creationType</a></code> | <code><a href="#cdk-sops-secrets.CreationType">CreationType</a></code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsStringParameterProps.property.flatten">flatten</a></code> | <code>boolean</code> | Should the structure be flattened? |
| <code><a href="#cdk-sops-secrets.SopsStringParameterProps.property.flattenSeparator">flattenSeparator</a></code> | <code>string</code> | If the structure should be flattened use the provided separator between keys. |
| <code><a href="#cdk-sops-secrets.SopsStringParameterProps.property.resourceType">resourceType</a></code> | <code><a href="#cdk-sops-secrets.ResourceType">ResourceType</a></code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsStringParameterProps.property.sopsAgeKey">sopsAgeKey</a></code> | <code>aws-cdk-lib.SecretValue</code> | The age key that should be used for encryption. |
| <code><a href="#cdk-sops-secrets.SopsStringParameterProps.property.sopsFileFormat">sopsFileFormat</a></code> | <code>string</code> | The format of the sops file. |
| <code><a href="#cdk-sops-secrets.SopsStringParameterProps.property.sopsFilePath">sopsFilePath</a></code> | <code>string</code> | The filepath to the sops file. |
| <code><a href="#cdk-sops-secrets.SopsStringParameterProps.property.sopsKmsKey">sopsKmsKey</a></code> | <code>aws-cdk-lib.aws_kms.IKey[]</code> | The kmsKey used to encrypt the sops file. |
| <code><a href="#cdk-sops-secrets.SopsStringParameterProps.property.sopsProvider">sopsProvider</a></code> | <code><a href="#cdk-sops-secrets.SopsSyncProvider">SopsSyncProvider</a></code> | The custom resource provider to use. |
| <code><a href="#cdk-sops-secrets.SopsStringParameterProps.property.sopsS3Bucket">sopsS3Bucket</a></code> | <code>string</code> | If you want to pass the sops file via s3, you can specify the bucket you can use cfn parameter here Both, sopsS3Bucket and sopsS3Key have to be specified. |
| <code><a href="#cdk-sops-secrets.SopsStringParameterProps.property.sopsS3Key">sopsS3Key</a></code> | <code>string</code> | If you want to pass the sops file via s3, you can specify the key inside the bucket you can use cfn parameter here Both, sopsS3Bucket and sopsS3Key have to be specified. |
| <code><a href="#cdk-sops-secrets.SopsStringParameterProps.property.stringifyValues">stringifyValues</a></code> | <code>boolean</code> | Shall all values be flattened? |
| <code><a href="#cdk-sops-secrets.SopsStringParameterProps.property.uploadType">uploadType</a></code> | <code><a href="#cdk-sops-secrets.UploadType">UploadType</a></code> | How should the secret be passed to the CustomResource? |
| <code><a href="#cdk-sops-secrets.SopsStringParameterProps.property.allowedPattern">allowedPattern</a></code> | <code>string</code> | A regular expression used to validate the parameter value. |
| <code><a href="#cdk-sops-secrets.SopsStringParameterProps.property.description">description</a></code> | <code>string</code> | Information about the parameter that you want to add to the system. |
| <code><a href="#cdk-sops-secrets.SopsStringParameterProps.property.parameterName">parameterName</a></code> | <code>string</code> | The name of the parameter. |
| <code><a href="#cdk-sops-secrets.SopsStringParameterProps.property.simpleName">simpleName</a></code> | <code>boolean</code> | Indicates if the parameter name is a simple name (i.e. does not include "/" separators). |
| <code><a href="#cdk-sops-secrets.SopsStringParameterProps.property.tier">tier</a></code> | <code>aws-cdk-lib.aws_ssm.ParameterTier</code> | The tier of the string parameter. |
| <code><a href="#cdk-sops-secrets.SopsStringParameterProps.property.stringValue">stringValue</a></code> | <code>string</code> | The value of the parameter. |
| <code><a href="#cdk-sops-secrets.SopsStringParameterProps.property.dataType">dataType</a></code> | <code>aws-cdk-lib.aws_ssm.ParameterDataType</code> | The data type of the parameter, such as `text` or `aws:ec2:image`. |
| <code><a href="#cdk-sops-secrets.SopsStringParameterProps.property.type">type</a></code> | <code>aws-cdk-lib.aws_ssm.ParameterType</code> | The type of the string parameter. |
| <code><a href="#cdk-sops-secrets.SopsStringParameterProps.property.encryptionKey">encryptionKey</a></code> | <code>aws-cdk-lib.aws_kms.IKey</code> | *No description.* |

---

##### `convertToJSON`<sup>Optional</sup> <a name="convertToJSON" id="cdk-sops-secrets.SopsStringParameterProps.property.convertToJSON"></a>

```typescript
public readonly convertToJSON: boolean;
```

- *Type:* boolean
- *Default:* true

Should the encrypted sops value should be converted to JSON?

Only JSON can be handled by cloud formations dynamic references.

---

##### `creationType`<sup>Optional</sup> <a name="creationType" id="cdk-sops-secrets.SopsStringParameterProps.property.creationType"></a>

```typescript
public readonly creationType: CreationType;
```

- *Type:* <a href="#cdk-sops-secrets.CreationType">CreationType</a>

---

##### `flatten`<sup>Optional</sup> <a name="flatten" id="cdk-sops-secrets.SopsStringParameterProps.property.flatten"></a>

```typescript
public readonly flatten: boolean;
```

- *Type:* boolean
- *Default:* true

Should the structure be flattened?

The result will be a flat structure and all
object keys will be replaced with the full jsonpath as key.
This is usefull for dynamic references, as those don't support nested objects.

---

##### `flattenSeparator`<sup>Optional</sup> <a name="flattenSeparator" id="cdk-sops-secrets.SopsStringParameterProps.property.flattenSeparator"></a>

```typescript
public readonly flattenSeparator: string;
```

- *Type:* string
- *Default:* '.'

If the structure should be flattened use the provided separator between keys.

---

##### `resourceType`<sup>Optional</sup> <a name="resourceType" id="cdk-sops-secrets.SopsStringParameterProps.property.resourceType"></a>

```typescript
public readonly resourceType: ResourceType;
```

- *Type:* <a href="#cdk-sops-secrets.ResourceType">ResourceType</a>

---

##### `sopsAgeKey`<sup>Optional</sup> <a name="sopsAgeKey" id="cdk-sops-secrets.SopsStringParameterProps.property.sopsAgeKey"></a>

```typescript
public readonly sopsAgeKey: SecretValue;
```

- *Type:* aws-cdk-lib.SecretValue

The age key that should be used for encryption.

---

##### `sopsFileFormat`<sup>Optional</sup> <a name="sopsFileFormat" id="cdk-sops-secrets.SopsStringParameterProps.property.sopsFileFormat"></a>

```typescript
public readonly sopsFileFormat: string;
```

- *Type:* string
- *Default:* The fileformat will be derived from the file ending

The format of the sops file.

---

##### `sopsFilePath`<sup>Optional</sup> <a name="sopsFilePath" id="cdk-sops-secrets.SopsStringParameterProps.property.sopsFilePath"></a>

```typescript
public readonly sopsFilePath: string;
```

- *Type:* string

The filepath to the sops file.

---

##### `sopsKmsKey`<sup>Optional</sup> <a name="sopsKmsKey" id="cdk-sops-secrets.SopsStringParameterProps.property.sopsKmsKey"></a>

```typescript
public readonly sopsKmsKey: IKey[];
```

- *Type:* aws-cdk-lib.aws_kms.IKey[]
- *Default:* The key will be derived from the sops file

The kmsKey used to encrypt the sops file.

Encrypt permissions
will be granted to the custom resource provider.

---

##### `sopsProvider`<sup>Optional</sup> <a name="sopsProvider" id="cdk-sops-secrets.SopsStringParameterProps.property.sopsProvider"></a>

```typescript
public readonly sopsProvider: SopsSyncProvider;
```

- *Type:* <a href="#cdk-sops-secrets.SopsSyncProvider">SopsSyncProvider</a>
- *Default:* A new singleton provider will be created

The custom resource provider to use.

If you don't specify any, a new
provider will be created - or if already exists within this stack - reused.

---

##### `sopsS3Bucket`<sup>Optional</sup> <a name="sopsS3Bucket" id="cdk-sops-secrets.SopsStringParameterProps.property.sopsS3Bucket"></a>

```typescript
public readonly sopsS3Bucket: string;
```

- *Type:* string

If you want to pass the sops file via s3, you can specify the bucket you can use cfn parameter here Both, sopsS3Bucket and sopsS3Key have to be specified.

---

##### `sopsS3Key`<sup>Optional</sup> <a name="sopsS3Key" id="cdk-sops-secrets.SopsStringParameterProps.property.sopsS3Key"></a>

```typescript
public readonly sopsS3Key: string;
```

- *Type:* string

If you want to pass the sops file via s3, you can specify the key inside the bucket you can use cfn parameter here Both, sopsS3Bucket and sopsS3Key have to be specified.

---

##### `stringifyValues`<sup>Optional</sup> <a name="stringifyValues" id="cdk-sops-secrets.SopsStringParameterProps.property.stringifyValues"></a>

```typescript
public readonly stringifyValues: boolean;
```

- *Type:* boolean

Shall all values be flattened?

This is usefull for dynamic references, as there
are lookup errors for certain float types

---

##### `uploadType`<sup>Optional</sup> <a name="uploadType" id="cdk-sops-secrets.SopsStringParameterProps.property.uploadType"></a>

```typescript
public readonly uploadType: UploadType;
```

- *Type:* <a href="#cdk-sops-secrets.UploadType">UploadType</a>
- *Default:* INLINE

How should the secret be passed to the CustomResource?

---

##### `allowedPattern`<sup>Optional</sup> <a name="allowedPattern" id="cdk-sops-secrets.SopsStringParameterProps.property.allowedPattern"></a>

```typescript
public readonly allowedPattern: string;
```

- *Type:* string
- *Default:* no validation is performed

A regular expression used to validate the parameter value.

For example, for String types with values restricted to
numbers, you can specify the following: ``^\d+$``

---

##### `description`<sup>Optional</sup> <a name="description" id="cdk-sops-secrets.SopsStringParameterProps.property.description"></a>

```typescript
public readonly description: string;
```

- *Type:* string
- *Default:* none

Information about the parameter that you want to add to the system.

---

##### `parameterName`<sup>Optional</sup> <a name="parameterName" id="cdk-sops-secrets.SopsStringParameterProps.property.parameterName"></a>

```typescript
public readonly parameterName: string;
```

- *Type:* string
- *Default:* a name will be generated by CloudFormation

The name of the parameter.

---

##### `simpleName`<sup>Optional</sup> <a name="simpleName" id="cdk-sops-secrets.SopsStringParameterProps.property.simpleName"></a>

```typescript
public readonly simpleName: boolean;
```

- *Type:* boolean
- *Default:* auto-detect based on `parameterName`

Indicates if the parameter name is a simple name (i.e. does not include "/" separators).

This is required only if `parameterName` is a token, which means we
are unable to detect if the name is simple or "path-like" for the purpose
of rendering SSM parameter ARNs.

If `parameterName` is not specified, `simpleName` must be `true` (or
undefined) since the name generated by AWS CloudFormation is always a
simple name.

---

##### `tier`<sup>Optional</sup> <a name="tier" id="cdk-sops-secrets.SopsStringParameterProps.property.tier"></a>

```typescript
public readonly tier: ParameterTier;
```

- *Type:* aws-cdk-lib.aws_ssm.ParameterTier
- *Default:* undefined

The tier of the string parameter.

---

##### `stringValue`<sup>Required</sup> <a name="stringValue" id="cdk-sops-secrets.SopsStringParameterProps.property.stringValue"></a>

```typescript
public readonly stringValue: string;
```

- *Type:* string

The value of the parameter.

It may not reference another parameter and ``{{}}`` cannot be used in the value.

---

##### `dataType`<sup>Optional</sup> <a name="dataType" id="cdk-sops-secrets.SopsStringParameterProps.property.dataType"></a>

```typescript
public readonly dataType: ParameterDataType;
```

- *Type:* aws-cdk-lib.aws_ssm.ParameterDataType
- *Default:* ParameterDataType.TEXT

The data type of the parameter, such as `text` or `aws:ec2:image`.

---

##### ~~`type`~~<sup>Optional</sup> <a name="type" id="cdk-sops-secrets.SopsStringParameterProps.property.type"></a>

- *Deprecated:* - type will always be 'String'

```typescript
public readonly type: ParameterType;
```

- *Type:* aws-cdk-lib.aws_ssm.ParameterType
- *Default:* ParameterType.STRING

The type of the string parameter.

---

##### `encryptionKey`<sup>Required</sup> <a name="encryptionKey" id="cdk-sops-secrets.SopsStringParameterProps.property.encryptionKey"></a>

```typescript
public readonly encryptionKey: IKey;
```

- *Type:* aws-cdk-lib.aws_kms.IKey

---

### SopsSyncOptions <a name="SopsSyncOptions" id="cdk-sops-secrets.SopsSyncOptions"></a>

Configuration options for the SopsSync.

#### Initializer <a name="Initializer" id="cdk-sops-secrets.SopsSyncOptions.Initializer"></a>

```typescript
import { SopsSyncOptions } from 'cdk-sops-secrets'

const sopsSyncOptions: SopsSyncOptions = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSyncOptions.property.convertToJSON">convertToJSON</a></code> | <code>boolean</code> | Should the encrypted sops value should be converted to JSON? |
| <code><a href="#cdk-sops-secrets.SopsSyncOptions.property.creationType">creationType</a></code> | <code><a href="#cdk-sops-secrets.CreationType">CreationType</a></code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSyncOptions.property.flatten">flatten</a></code> | <code>boolean</code> | Should the structure be flattened? |
| <code><a href="#cdk-sops-secrets.SopsSyncOptions.property.flattenSeparator">flattenSeparator</a></code> | <code>string</code> | If the structure should be flattened use the provided separator between keys. |
| <code><a href="#cdk-sops-secrets.SopsSyncOptions.property.resourceType">resourceType</a></code> | <code><a href="#cdk-sops-secrets.ResourceType">ResourceType</a></code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSyncOptions.property.sopsAgeKey">sopsAgeKey</a></code> | <code>aws-cdk-lib.SecretValue</code> | The age key that should be used for encryption. |
| <code><a href="#cdk-sops-secrets.SopsSyncOptions.property.sopsFileFormat">sopsFileFormat</a></code> | <code>string</code> | The format of the sops file. |
| <code><a href="#cdk-sops-secrets.SopsSyncOptions.property.sopsFilePath">sopsFilePath</a></code> | <code>string</code> | The filepath to the sops file. |
| <code><a href="#cdk-sops-secrets.SopsSyncOptions.property.sopsKmsKey">sopsKmsKey</a></code> | <code>aws-cdk-lib.aws_kms.IKey[]</code> | The kmsKey used to encrypt the sops file. |
| <code><a href="#cdk-sops-secrets.SopsSyncOptions.property.sopsProvider">sopsProvider</a></code> | <code><a href="#cdk-sops-secrets.SopsSyncProvider">SopsSyncProvider</a></code> | The custom resource provider to use. |
| <code><a href="#cdk-sops-secrets.SopsSyncOptions.property.sopsS3Bucket">sopsS3Bucket</a></code> | <code>string</code> | If you want to pass the sops file via s3, you can specify the bucket you can use cfn parameter here Both, sopsS3Bucket and sopsS3Key have to be specified. |
| <code><a href="#cdk-sops-secrets.SopsSyncOptions.property.sopsS3Key">sopsS3Key</a></code> | <code>string</code> | If you want to pass the sops file via s3, you can specify the key inside the bucket you can use cfn parameter here Both, sopsS3Bucket and sopsS3Key have to be specified. |
| <code><a href="#cdk-sops-secrets.SopsSyncOptions.property.stringifyValues">stringifyValues</a></code> | <code>boolean</code> | Shall all values be flattened? |
| <code><a href="#cdk-sops-secrets.SopsSyncOptions.property.uploadType">uploadType</a></code> | <code><a href="#cdk-sops-secrets.UploadType">UploadType</a></code> | How should the secret be passed to the CustomResource? |

---

##### `convertToJSON`<sup>Optional</sup> <a name="convertToJSON" id="cdk-sops-secrets.SopsSyncOptions.property.convertToJSON"></a>

```typescript
public readonly convertToJSON: boolean;
```

- *Type:* boolean
- *Default:* true

Should the encrypted sops value should be converted to JSON?

Only JSON can be handled by cloud formations dynamic references.

---

##### `creationType`<sup>Optional</sup> <a name="creationType" id="cdk-sops-secrets.SopsSyncOptions.property.creationType"></a>

```typescript
public readonly creationType: CreationType;
```

- *Type:* <a href="#cdk-sops-secrets.CreationType">CreationType</a>

---

##### `flatten`<sup>Optional</sup> <a name="flatten" id="cdk-sops-secrets.SopsSyncOptions.property.flatten"></a>

```typescript
public readonly flatten: boolean;
```

- *Type:* boolean
- *Default:* true

Should the structure be flattened?

The result will be a flat structure and all
object keys will be replaced with the full jsonpath as key.
This is usefull for dynamic references, as those don't support nested objects.

---

##### `flattenSeparator`<sup>Optional</sup> <a name="flattenSeparator" id="cdk-sops-secrets.SopsSyncOptions.property.flattenSeparator"></a>

```typescript
public readonly flattenSeparator: string;
```

- *Type:* string
- *Default:* '.'

If the structure should be flattened use the provided separator between keys.

---

##### `resourceType`<sup>Optional</sup> <a name="resourceType" id="cdk-sops-secrets.SopsSyncOptions.property.resourceType"></a>

```typescript
public readonly resourceType: ResourceType;
```

- *Type:* <a href="#cdk-sops-secrets.ResourceType">ResourceType</a>

---

##### `sopsAgeKey`<sup>Optional</sup> <a name="sopsAgeKey" id="cdk-sops-secrets.SopsSyncOptions.property.sopsAgeKey"></a>

```typescript
public readonly sopsAgeKey: SecretValue;
```

- *Type:* aws-cdk-lib.SecretValue

The age key that should be used for encryption.

---

##### `sopsFileFormat`<sup>Optional</sup> <a name="sopsFileFormat" id="cdk-sops-secrets.SopsSyncOptions.property.sopsFileFormat"></a>

```typescript
public readonly sopsFileFormat: string;
```

- *Type:* string
- *Default:* The fileformat will be derived from the file ending

The format of the sops file.

---

##### `sopsFilePath`<sup>Optional</sup> <a name="sopsFilePath" id="cdk-sops-secrets.SopsSyncOptions.property.sopsFilePath"></a>

```typescript
public readonly sopsFilePath: string;
```

- *Type:* string

The filepath to the sops file.

---

##### `sopsKmsKey`<sup>Optional</sup> <a name="sopsKmsKey" id="cdk-sops-secrets.SopsSyncOptions.property.sopsKmsKey"></a>

```typescript
public readonly sopsKmsKey: IKey[];
```

- *Type:* aws-cdk-lib.aws_kms.IKey[]
- *Default:* The key will be derived from the sops file

The kmsKey used to encrypt the sops file.

Encrypt permissions
will be granted to the custom resource provider.

---

##### `sopsProvider`<sup>Optional</sup> <a name="sopsProvider" id="cdk-sops-secrets.SopsSyncOptions.property.sopsProvider"></a>

```typescript
public readonly sopsProvider: SopsSyncProvider;
```

- *Type:* <a href="#cdk-sops-secrets.SopsSyncProvider">SopsSyncProvider</a>
- *Default:* A new singleton provider will be created

The custom resource provider to use.

If you don't specify any, a new
provider will be created - or if already exists within this stack - reused.

---

##### `sopsS3Bucket`<sup>Optional</sup> <a name="sopsS3Bucket" id="cdk-sops-secrets.SopsSyncOptions.property.sopsS3Bucket"></a>

```typescript
public readonly sopsS3Bucket: string;
```

- *Type:* string

If you want to pass the sops file via s3, you can specify the bucket you can use cfn parameter here Both, sopsS3Bucket and sopsS3Key have to be specified.

---

##### `sopsS3Key`<sup>Optional</sup> <a name="sopsS3Key" id="cdk-sops-secrets.SopsSyncOptions.property.sopsS3Key"></a>

```typescript
public readonly sopsS3Key: string;
```

- *Type:* string

If you want to pass the sops file via s3, you can specify the key inside the bucket you can use cfn parameter here Both, sopsS3Bucket and sopsS3Key have to be specified.

---

##### `stringifyValues`<sup>Optional</sup> <a name="stringifyValues" id="cdk-sops-secrets.SopsSyncOptions.property.stringifyValues"></a>

```typescript
public readonly stringifyValues: boolean;
```

- *Type:* boolean

Shall all values be flattened?

This is usefull for dynamic references, as there
are lookup errors for certain float types

---

##### `uploadType`<sup>Optional</sup> <a name="uploadType" id="cdk-sops-secrets.SopsSyncOptions.property.uploadType"></a>

```typescript
public readonly uploadType: UploadType;
```

- *Type:* <a href="#cdk-sops-secrets.UploadType">UploadType</a>
- *Default:* INLINE

How should the secret be passed to the CustomResource?

---

### SopsSyncProps <a name="SopsSyncProps" id="cdk-sops-secrets.SopsSyncProps"></a>

The configuration options extended by the target Secret.

#### Initializer <a name="Initializer" id="cdk-sops-secrets.SopsSyncProps.Initializer"></a>

```typescript
import { SopsSyncProps } from 'cdk-sops-secrets'

const sopsSyncProps: SopsSyncProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.convertToJSON">convertToJSON</a></code> | <code>boolean</code> | Should the encrypted sops value should be converted to JSON? |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.creationType">creationType</a></code> | <code><a href="#cdk-sops-secrets.CreationType">CreationType</a></code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.flatten">flatten</a></code> | <code>boolean</code> | Should the structure be flattened? |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.flattenSeparator">flattenSeparator</a></code> | <code>string</code> | If the structure should be flattened use the provided separator between keys. |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.resourceType">resourceType</a></code> | <code><a href="#cdk-sops-secrets.ResourceType">ResourceType</a></code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.sopsAgeKey">sopsAgeKey</a></code> | <code>aws-cdk-lib.SecretValue</code> | The age key that should be used for encryption. |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.sopsFileFormat">sopsFileFormat</a></code> | <code>string</code> | The format of the sops file. |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.sopsFilePath">sopsFilePath</a></code> | <code>string</code> | The filepath to the sops file. |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.sopsKmsKey">sopsKmsKey</a></code> | <code>aws-cdk-lib.aws_kms.IKey[]</code> | The kmsKey used to encrypt the sops file. |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.sopsProvider">sopsProvider</a></code> | <code><a href="#cdk-sops-secrets.SopsSyncProvider">SopsSyncProvider</a></code> | The custom resource provider to use. |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.sopsS3Bucket">sopsS3Bucket</a></code> | <code>string</code> | If you want to pass the sops file via s3, you can specify the bucket you can use cfn parameter here Both, sopsS3Bucket and sopsS3Key have to be specified. |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.sopsS3Key">sopsS3Key</a></code> | <code>string</code> | If you want to pass the sops file via s3, you can specify the key inside the bucket you can use cfn parameter here Both, sopsS3Bucket and sopsS3Key have to be specified. |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.stringifyValues">stringifyValues</a></code> | <code>boolean</code> | Shall all values be flattened? |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.uploadType">uploadType</a></code> | <code><a href="#cdk-sops-secrets.UploadType">UploadType</a></code> | How should the secret be passed to the CustomResource? |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.encryptionKey">encryptionKey</a></code> | <code>aws-cdk-lib.aws_kms.IKey</code> | The encryption key used for encrypting the ssm parameter if `parameterName` is set. |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.parameterName">parameterName</a></code> | <code>string</code> | The parameter name. |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.parameterNames">parameterNames</a></code> | <code>string[]</code> | The parameter name. |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.secret">secret</a></code> | <code>aws-cdk-lib.aws_secretsmanager.ISecret</code> | The secret that will be populated with the encrypted sops file content. |

---

##### `convertToJSON`<sup>Optional</sup> <a name="convertToJSON" id="cdk-sops-secrets.SopsSyncProps.property.convertToJSON"></a>

```typescript
public readonly convertToJSON: boolean;
```

- *Type:* boolean
- *Default:* true

Should the encrypted sops value should be converted to JSON?

Only JSON can be handled by cloud formations dynamic references.

---

##### `creationType`<sup>Optional</sup> <a name="creationType" id="cdk-sops-secrets.SopsSyncProps.property.creationType"></a>

```typescript
public readonly creationType: CreationType;
```

- *Type:* <a href="#cdk-sops-secrets.CreationType">CreationType</a>

---

##### `flatten`<sup>Optional</sup> <a name="flatten" id="cdk-sops-secrets.SopsSyncProps.property.flatten"></a>

```typescript
public readonly flatten: boolean;
```

- *Type:* boolean
- *Default:* true

Should the structure be flattened?

The result will be a flat structure and all
object keys will be replaced with the full jsonpath as key.
This is usefull for dynamic references, as those don't support nested objects.

---

##### `flattenSeparator`<sup>Optional</sup> <a name="flattenSeparator" id="cdk-sops-secrets.SopsSyncProps.property.flattenSeparator"></a>

```typescript
public readonly flattenSeparator: string;
```

- *Type:* string
- *Default:* '.'

If the structure should be flattened use the provided separator between keys.

---

##### `resourceType`<sup>Optional</sup> <a name="resourceType" id="cdk-sops-secrets.SopsSyncProps.property.resourceType"></a>

```typescript
public readonly resourceType: ResourceType;
```

- *Type:* <a href="#cdk-sops-secrets.ResourceType">ResourceType</a>

---

##### `sopsAgeKey`<sup>Optional</sup> <a name="sopsAgeKey" id="cdk-sops-secrets.SopsSyncProps.property.sopsAgeKey"></a>

```typescript
public readonly sopsAgeKey: SecretValue;
```

- *Type:* aws-cdk-lib.SecretValue

The age key that should be used for encryption.

---

##### `sopsFileFormat`<sup>Optional</sup> <a name="sopsFileFormat" id="cdk-sops-secrets.SopsSyncProps.property.sopsFileFormat"></a>

```typescript
public readonly sopsFileFormat: string;
```

- *Type:* string
- *Default:* The fileformat will be derived from the file ending

The format of the sops file.

---

##### `sopsFilePath`<sup>Optional</sup> <a name="sopsFilePath" id="cdk-sops-secrets.SopsSyncProps.property.sopsFilePath"></a>

```typescript
public readonly sopsFilePath: string;
```

- *Type:* string

The filepath to the sops file.

---

##### `sopsKmsKey`<sup>Optional</sup> <a name="sopsKmsKey" id="cdk-sops-secrets.SopsSyncProps.property.sopsKmsKey"></a>

```typescript
public readonly sopsKmsKey: IKey[];
```

- *Type:* aws-cdk-lib.aws_kms.IKey[]
- *Default:* The key will be derived from the sops file

The kmsKey used to encrypt the sops file.

Encrypt permissions
will be granted to the custom resource provider.

---

##### `sopsProvider`<sup>Optional</sup> <a name="sopsProvider" id="cdk-sops-secrets.SopsSyncProps.property.sopsProvider"></a>

```typescript
public readonly sopsProvider: SopsSyncProvider;
```

- *Type:* <a href="#cdk-sops-secrets.SopsSyncProvider">SopsSyncProvider</a>
- *Default:* A new singleton provider will be created

The custom resource provider to use.

If you don't specify any, a new
provider will be created - or if already exists within this stack - reused.

---

##### `sopsS3Bucket`<sup>Optional</sup> <a name="sopsS3Bucket" id="cdk-sops-secrets.SopsSyncProps.property.sopsS3Bucket"></a>

```typescript
public readonly sopsS3Bucket: string;
```

- *Type:* string

If you want to pass the sops file via s3, you can specify the bucket you can use cfn parameter here Both, sopsS3Bucket and sopsS3Key have to be specified.

---

##### `sopsS3Key`<sup>Optional</sup> <a name="sopsS3Key" id="cdk-sops-secrets.SopsSyncProps.property.sopsS3Key"></a>

```typescript
public readonly sopsS3Key: string;
```

- *Type:* string

If you want to pass the sops file via s3, you can specify the key inside the bucket you can use cfn parameter here Both, sopsS3Bucket and sopsS3Key have to be specified.

---

##### `stringifyValues`<sup>Optional</sup> <a name="stringifyValues" id="cdk-sops-secrets.SopsSyncProps.property.stringifyValues"></a>

```typescript
public readonly stringifyValues: boolean;
```

- *Type:* boolean

Shall all values be flattened?

This is usefull for dynamic references, as there
are lookup errors for certain float types

---

##### `uploadType`<sup>Optional</sup> <a name="uploadType" id="cdk-sops-secrets.SopsSyncProps.property.uploadType"></a>

```typescript
public readonly uploadType: UploadType;
```

- *Type:* <a href="#cdk-sops-secrets.UploadType">UploadType</a>
- *Default:* INLINE

How should the secret be passed to the CustomResource?

---

##### `encryptionKey`<sup>Optional</sup> <a name="encryptionKey" id="cdk-sops-secrets.SopsSyncProps.property.encryptionKey"></a>

```typescript
public readonly encryptionKey: IKey;
```

- *Type:* aws-cdk-lib.aws_kms.IKey

The encryption key used for encrypting the ssm parameter if `parameterName` is set.

---

##### `parameterName`<sup>Optional</sup> <a name="parameterName" id="cdk-sops-secrets.SopsSyncProps.property.parameterName"></a>

```typescript
public readonly parameterName: string;
```

- *Type:* string

The parameter name.

If set this creates an encrypted SSM Parameter instead of a secret.

---

##### `parameterNames`<sup>Optional</sup> <a name="parameterNames" id="cdk-sops-secrets.SopsSyncProps.property.parameterNames"></a>

```typescript
public readonly parameterNames: string[];
```

- *Type:* string[]

The parameter name.

If set this creates an encrypted SSM Parameter instead of a secret.

---

##### `secret`<sup>Optional</sup> <a name="secret" id="cdk-sops-secrets.SopsSyncProps.property.secret"></a>

```typescript
public readonly secret: ISecret;
```

- *Type:* aws-cdk-lib.aws_secretsmanager.ISecret

The secret that will be populated with the encrypted sops file content.

---



## Enums <a name="Enums" id="Enums"></a>

### CreationType <a name="CreationType" id="cdk-sops-secrets.CreationType"></a>

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-sops-secrets.CreationType.SINGLE">SINGLE</a></code> | Create or update a single secret/parameter. |
| <code><a href="#cdk-sops-secrets.CreationType.MULTI">MULTI</a></code> | Create or update a multiple secrets/parameters by flattening the SOPS file. |

---

##### `SINGLE` <a name="SINGLE" id="cdk-sops-secrets.CreationType.SINGLE"></a>

Create or update a single secret/parameter.

---


##### `MULTI` <a name="MULTI" id="cdk-sops-secrets.CreationType.MULTI"></a>

Create or update a multiple secrets/parameters by flattening the SOPS file.

---


### ResourceType <a name="ResourceType" id="cdk-sops-secrets.ResourceType"></a>

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-sops-secrets.ResourceType.SECRET">SECRET</a></code> | *No description.* |
| <code><a href="#cdk-sops-secrets.ResourceType.PARAMETER">PARAMETER</a></code> | *No description.* |

---

##### `SECRET` <a name="SECRET" id="cdk-sops-secrets.ResourceType.SECRET"></a>

---


##### `PARAMETER` <a name="PARAMETER" id="cdk-sops-secrets.ResourceType.PARAMETER"></a>

---


### UploadType <a name="UploadType" id="cdk-sops-secrets.UploadType"></a>

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-sops-secrets.UploadType.INLINE">INLINE</a></code> | Pass the secret data inline (base64 encoded and compressed). |
| <code><a href="#cdk-sops-secrets.UploadType.ASSET">ASSET</a></code> | Uplaod the secert data as asset. |

---

##### `INLINE` <a name="INLINE" id="cdk-sops-secrets.UploadType.INLINE"></a>

Pass the secret data inline (base64 encoded and compressed).

---


##### `ASSET` <a name="ASSET" id="cdk-sops-secrets.UploadType.ASSET"></a>

Uplaod the secert data as asset.

---

