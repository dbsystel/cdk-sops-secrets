# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

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
| <code><a href="#cdk-sops-secrets.SopsSecret.currentVersionId">currentVersionId</a></code> | Returns the current versionId that was created via the SopsSync. |
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

Returns the current versionId that was created via the SopsSync.

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
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.addPermission">addPermission</a></code> | Adds a permission to the Lambda resource policy. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.addToRolePolicy">addToRolePolicy</a></code> | Adds a statement to the IAM role assumed by the instance. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.configureAsyncInvoke">configureAsyncInvoke</a></code> | Configures options for asynchronous invocation. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.grantInvoke">grantInvoke</a></code> | Grant the given identity permissions to invoke this Lambda. |
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

Event sources are implemented in the @aws-cdk/aws-lambda-event-sources module.

The following example adds an SQS Queue as an event source:
```
import { SqsEventSource } from '@aws-cdk/aws-lambda-event-sources';
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

##### `grantInvoke` <a name="grantInvoke" id="cdk-sops-secrets.SopsSyncProvider.grantInvoke"></a>

```typescript
public grantInvoke(grantee: IGrantable): Grant
```

Grant the given identity permissions to invoke this Lambda.

###### `grantee`<sup>Required</sup> <a name="grantee" id="cdk-sops-secrets.SopsSyncProvider.grantInvoke.parameter.grantee"></a>

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
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.property.connections">connections</a></code> | <code>aws-cdk-lib.aws_ec2.Connections</code> | Access the Connections object. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.property.functionArn">functionArn</a></code> | <code>string</code> | The ARN fo the function. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.property.functionName">functionName</a></code> | <code>string</code> | The name of the function. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.property.grantPrincipal">grantPrincipal</a></code> | <code>aws-cdk-lib.aws_iam.IPrincipal</code> | The principal this Lambda Function is running as. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.property.isBoundToVpc">isBoundToVpc</a></code> | <code>boolean</code> | Whether or not this Lambda function was bound to a VPC. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.property.latestVersion">latestVersion</a></code> | <code>aws-cdk-lib.aws_lambda.IVersion</code> | The `$LATEST` version of this function. |
| <code><a href="#cdk-sops-secrets.SopsSyncProvider.property.permissionsNode">permissionsNode</a></code> | <code>constructs.Node</code> | The construct node where permissions are attached. |
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
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.convertToJSON">convertToJSON</a></code> | <code>boolean</code> | Should the encrypted sops value should be converted to JSON? |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.flatten">flatten</a></code> | <code>boolean</code> | Should the structure be flattened? |
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

##### `convertToJSON`<sup>Optional</sup> <a name="convertToJSON" id="cdk-sops-secrets.SopsSecretProps.property.convertToJSON"></a>

```typescript
public readonly convertToJSON: boolean;
```

- *Type:* boolean
- *Default:* true

Should the encrypted sops value should be converted to JSON?

Only JSON can be handled by cloud formations dynamic references.

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
| <code><a href="#cdk-sops-secrets.SopsSyncOptions.property.flatten">flatten</a></code> | <code>boolean</code> | Should the structure be flattened? |
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
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.flatten">flatten</a></code> | <code>boolean</code> | Should the structure be flattened? |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.sopsAgeKey">sopsAgeKey</a></code> | <code>aws-cdk-lib.SecretValue</code> | The age key that should be used for encryption. |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.sopsFileFormat">sopsFileFormat</a></code> | <code>string</code> | The format of the sops file. |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.sopsFilePath">sopsFilePath</a></code> | <code>string</code> | The filepath to the sops file. |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.sopsKmsKey">sopsKmsKey</a></code> | <code>aws-cdk-lib.aws_kms.IKey[]</code> | The kmsKey used to encrypt the sops file. |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.sopsProvider">sopsProvider</a></code> | <code><a href="#cdk-sops-secrets.SopsSyncProvider">SopsSyncProvider</a></code> | The custom resource provider to use. |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.sopsS3Bucket">sopsS3Bucket</a></code> | <code>string</code> | If you want to pass the sops file via s3, you can specify the bucket you can use cfn parameter here Both, sopsS3Bucket and sopsS3Key have to be specified. |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.sopsS3Key">sopsS3Key</a></code> | <code>string</code> | If you want to pass the sops file via s3, you can specify the key inside the bucket you can use cfn parameter here Both, sopsS3Bucket and sopsS3Key have to be specified. |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.stringifyValues">stringifyValues</a></code> | <code>boolean</code> | Shall all values be flattened? |
| <code><a href="#cdk-sops-secrets.SopsSyncProps.property.uploadType">uploadType</a></code> | <code><a href="#cdk-sops-secrets.UploadType">UploadType</a></code> | How should the secret be passed to the CustomResource? |
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

##### `secret`<sup>Required</sup> <a name="secret" id="cdk-sops-secrets.SopsSyncProps.property.secret"></a>

```typescript
public readonly secret: ISecret;
```

- *Type:* aws-cdk-lib.aws_secretsmanager.ISecret

The secret that will be populated with the encrypted sops file content.

---



## Enums <a name="Enums" id="Enums"></a>

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

