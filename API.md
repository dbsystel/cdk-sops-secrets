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
| <code><a href="#cdk-sops-secrets.SopsSecrets.property.props">props</a></code> | <code><a href="#cdk-sops-secrets.SopsSecretProps">SopsSecretProps</a></code> | *No description.* |

---

##### `node`<sup>Required</sup> <a name="node" id="cdk-sops-secrets.SopsSecrets.property.node"></a>

```typescript
public readonly node: ConstructNode;
```

- *Type:* @aws-cdk/core.ConstructNode

The construct tree node associated with this construct.

---

##### `props`<sup>Required</sup> <a name="props" id="cdk-sops-secrets.SopsSecrets.property.props"></a>

```typescript
public readonly props: SopsSecretProps;
```

- *Type:* <a href="#cdk-sops-secrets.SopsSecretProps">SopsSecretProps</a>

---


### SopsSecretsProvider <a name="SopsSecretsProvider" id="cdk-sops-secrets.SopsSecretsProvider"></a>

- *Implements:* @aws-cdk/aws-iam.IGrantable

#### Initializers <a name="Initializers" id="cdk-sops-secrets.SopsSecretsProvider.Initializer"></a>

```typescript
import { SopsSecretsProvider } from 'cdk-sops-secrets'

new SopsSecretsProvider(scope: Construct, id: string, props: SopsSecretsProviderProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSecretsProvider.Initializer.parameter.scope">scope</a></code> | <code>@aws-cdk/core.Construct</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSecretsProvider.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk-sops-secrets.SopsSecretsProvider.Initializer.parameter.props">props</a></code> | <code><a href="#cdk-sops-secrets.SopsSecretsProviderProps">SopsSecretsProviderProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="cdk-sops-secrets.SopsSecretsProvider.Initializer.parameter.scope"></a>

- *Type:* @aws-cdk/core.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="cdk-sops-secrets.SopsSecretsProvider.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="cdk-sops-secrets.SopsSecretsProvider.Initializer.parameter.props"></a>

- *Type:* <a href="#cdk-sops-secrets.SopsSecretsProviderProps">SopsSecretsProviderProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSecretsProvider.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="cdk-sops-secrets.SopsSecretsProvider.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSecretsProvider.isConstruct">isConstruct</a></code> | Return whether the given object is a Construct. |

---

##### `isConstruct` <a name="isConstruct" id="cdk-sops-secrets.SopsSecretsProvider.isConstruct"></a>

```typescript
import { SopsSecretsProvider } from 'cdk-sops-secrets'

SopsSecretsProvider.isConstruct(x: any)
```

Return whether the given object is a Construct.

###### `x`<sup>Required</sup> <a name="x" id="cdk-sops-secrets.SopsSecretsProvider.isConstruct.parameter.x"></a>

- *Type:* any

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-sops-secrets.SopsSecretsProvider.property.node">node</a></code> | <code>@aws-cdk/core.ConstructNode</code> | The construct tree node associated with this construct. |
| <code><a href="#cdk-sops-secrets.SopsSecretsProvider.property.grantPrincipal">grantPrincipal</a></code> | <code>@aws-cdk/aws-iam.IPrincipal</code> | The principal to grant permissions to. |
| <code><a href="#cdk-sops-secrets.SopsSecretsProvider.property.props">props</a></code> | <code><a href="#cdk-sops-secrets.SopsSecretsProviderProps">SopsSecretsProviderProps</a></code> | *No description.* |

---

##### `node`<sup>Required</sup> <a name="node" id="cdk-sops-secrets.SopsSecretsProvider.property.node"></a>

```typescript
public readonly node: ConstructNode;
```

- *Type:* @aws-cdk/core.ConstructNode

The construct tree node associated with this construct.

---

##### `grantPrincipal`<sup>Required</sup> <a name="grantPrincipal" id="cdk-sops-secrets.SopsSecretsProvider.property.grantPrincipal"></a>

```typescript
public readonly grantPrincipal: IPrincipal;
```

- *Type:* @aws-cdk/aws-iam.IPrincipal

The principal to grant permissions to.

---

##### `props`<sup>Required</sup> <a name="props" id="cdk-sops-secrets.SopsSecretsProvider.property.props"></a>

```typescript
public readonly props: SopsSecretsProviderProps;
```

- *Type:* <a href="#cdk-sops-secrets.SopsSecretsProviderProps">SopsSecretsProviderProps</a>

---


## Structs <a name="Structs" id="Structs"></a>

### SopsSecretProps <a name="SopsSecretProps" id="cdk-sops-secrets.SopsSecretProps"></a>

#### Initializer <a name="Initializer" id="cdk-sops-secrets.SopsSecretProps.Initializer"></a>

```typescript
import { SopsSecretProps } from 'cdk-sops-secrets'

const sopsSecretProps: SopsSecretProps = { ... }
```


### SopsSecretsProviderProps <a name="SopsSecretsProviderProps" id="cdk-sops-secrets.SopsSecretsProviderProps"></a>

#### Initializer <a name="Initializer" id="cdk-sops-secrets.SopsSecretsProviderProps.Initializer"></a>

```typescript
import { SopsSecretsProviderProps } from 'cdk-sops-secrets'

const sopsSecretsProviderProps: SopsSecretsProviderProps = { ... }
```




