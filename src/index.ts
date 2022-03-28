import { IGrantable, IPrincipal } from '@aws-cdk/aws-iam';
import { Construct, CustomResource, CustomResourceProvider } from '@aws-cdk/core';
import { Code, Function } from '@aws-cdk/aws-lambda';
import * as cr from '@aws-cdk/custom-resources';
import { Provider } from '@aws-cdk/custom-resources';

export interface SopsSecretProps {

}
export class SopsSecrets extends Construct {

  public constructor(scope: Construct, id: string, props: SopsSecretProps) {
    super(scope, id);
    new CustomResource(this, 'Resource', {

    })
  }

}
export interface SopsSecretsProviderProps {

}
export class SopsSecretsProvider extends Construct implements IGrantable {
  
  private grantPrincipal: IPrincipal;

  public constructor(scope: Construct, id: string, props: SopsSecretsProviderProps ) {
    super(scope, id);
    const lambda = new Function(this, 'Function', {
      code: Code.fromAsset('../assets/lambdasource.zip'),
      
    })
    const cr = new Provider(this, 'Resource', {
      onEventHandler:
    })

    })
  }
  
}