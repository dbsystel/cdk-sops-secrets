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

---

##### `toString` <a name="toString" id="cdk-sops-secrets.SopsSecrets.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSecrets.isConstruct">isConstruct</a></code> | Return whether the given object is a Construct. |

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

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSecrets.property.node">node</a></code> | <code>@aws-cdk/core.ConstructNode</code> | The construct tree node associated with this construct. |

---

##### `node`<sup>Required</sup> <a name="node" id="cdk-sops-secrets.SopsSecrets.property.node"></a>

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
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.secret">secret</a></code> | <code>@aws-cdk/aws-secretsmanager.ISecret</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.sopsFilePath">sopsFilePath</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.sopsKmsKey">sopsKmsKey</a></code> | <code>@aws-cdk/aws-kms.IKey</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.provider">provider</a></code> | <code>@aws-cdk/aws-lambda.IFunction</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSecretProps.property.sopsFileFormat">sopsFileFormat</a></code> | <code>string</code> | *No description.* |

---

##### `secret`<sup>Required</sup> <a name="secret" id="cdk-sops-secrets.SopsSecretProps.property.secret"></a>

```typescript
public readonly secret: ISecret;
```

- *Type:* @aws-cdk/aws-secretsmanager.ISecret

---

##### `sopsFilePath`<sup>Required</sup> <a name="sopsFilePath" id="cdk-sops-secrets.SopsSecretProps.property.sopsFilePath"></a>

```typescript
public readonly sopsFilePath: string;
```

- *Type:* string

---

##### `sopsKmsKey`<sup>Required</sup> <a name="sopsKmsKey" id="cdk-sops-secrets.SopsSecretProps.property.sopsKmsKey"></a>

```typescript
public readonly sopsKmsKey: IKey;
```

- *Type:* @aws-cdk/aws-kms.IKey

---

##### `provider`<sup>Optional</sup> <a name="provider" id="cdk-sops-secrets.SopsSecretProps.property.provider"></a>

```typescript
public readonly provider: IFunction;
```

- *Type:* @aws-cdk/aws-lambda.IFunction

---

##### `sopsFileFormat`<sup>Optional</sup> <a name="sopsFileFormat" id="cdk-sops-secrets.SopsSecretProps.property.sopsFileFormat"></a>

```typescript
public readonly sopsFileFormat: string;
```

- *Type:* string

---



