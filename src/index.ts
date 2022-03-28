import { IGrantable, IPrincipal } from '@aws-cdk/aws-iam';
import { Code, Function, Runtime } from '@aws-cdk/aws-lambda';
import { Construct } from '@aws-cdk/core';
import { Provider } from '@aws-cdk/custom-resources';

export interface SopsSecretProps {

}
export class SopsSecrets extends Construct {
  readonly props: SopsSecretProps;

  public constructor(scope: Construct, id: string, props: SopsSecretProps) {
    super(scope, id);
    this.props = props;
  }

}
export interface SopsSecretsProviderProps {

}
export class SopsSecretsProvider extends Construct implements IGrantable {

  readonly grantPrincipal: IPrincipal;
  readonly props: SopsSecretsProviderProps;

  public constructor(scope: Construct, id: string, props: SopsSecretsProviderProps ) {
    super(scope, id);
    this.props = props;
    const lambda = new Function(this, 'Function', {
      code: Code.fromAsset('../assets/cdk-sops-lambda.zip'),
      runtime: Runtime.GO_1_X,
      handler: 'cdk-sops-lambda',
    });
    this.grantPrincipal = lambda.role!;
    new Provider(this, 'Resource', {
      onEventHandler: lambda,
    });
  }
}