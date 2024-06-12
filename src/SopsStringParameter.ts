import { Grant, IGrantable } from 'aws-cdk-lib/aws-iam';
import { IKey } from 'aws-cdk-lib/aws-kms';
import {
  IStringParameter,
  StringParameter,
  StringParameterProps,
} from 'aws-cdk-lib/aws-ssm';
import { RemovalPolicy, ResourceEnvironment, Stack } from 'aws-cdk-lib/core';
import { Construct } from 'constructs';
import { SopsSync, SopsSyncOptions } from './SopsSync';

/**
 * The configuration options of the StringParameter
 */
export interface SopsStringParameterProps
  extends SopsSyncOptions,
    StringParameterProps {
  readonly encryptionKey: IKey;
}

/**
 * A drop in replacement for the normal String Parameter, that is populated with the encrypted
 * content of the given sops file.
 */
export class SopsStringParameter extends Construct implements IStringParameter {
  private readonly parameter: StringParameter;
  readonly sync: SopsSync;
  readonly encryptionKey: IKey;
  readonly stack: Stack;
  readonly env: ResourceEnvironment;
  readonly parameterArn: string;
  readonly parameterName: string;
  readonly parameterType: string;
  readonly stringValue: string;

  public constructor(
    scope: Construct,
    id: string,
    props: SopsStringParameterProps,
  ) {
    super(scope, id);

    this.encryptionKey = props.encryptionKey;
    this.stack = Stack.of(scope);
    this.env = {
      account: this.stack.account,
      region: this.stack.region,
    };

    this.parameter = new StringParameter(this, 'Resource', {
      ...(props as StringParameterProps),
      stringValue: ' ',
    });
    this.parameterArn = this.parameter.parameterArn;
    this.parameterName = this.parameter.parameterName;
    this.parameterType = this.parameter.parameterType;
    this.stringValue = this.parameter.stringValue;

    this.sync = new SopsSync(this, 'SopsSync', {
      encryptionKey: this.parameter.encryptionKey,
      parameterName: this.parameter.parameterName,
      ...(props as SopsSyncOptions),
    });
  }
  grantRead(grantee: IGrantable): Grant {
    return this.parameter.grantRead(grantee);
  }
  grantWrite(grantee: IGrantable): Grant {
    return this.parameter.grantWrite(grantee);
  }
  applyRemovalPolicy(policy: RemovalPolicy): void {
    this.parameter.applyRemovalPolicy(policy);
  }
}
