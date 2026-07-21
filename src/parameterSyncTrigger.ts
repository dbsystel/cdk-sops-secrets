import { createHash } from 'crypto';
import { CfnResource, Lazy, Stack } from 'aws-cdk-lib';
import { CfnParameter, StringParameter } from 'aws-cdk-lib/aws-ssm';
import { SopsSync } from './SopsSync';

function getParameterResource(parameter: StringParameter): CfnParameter {
  const resource = parameter.node.defaultChild;
  if (!CfnParameter.isCfnParameter(resource)) {
    throw new Error(
      'Expected StringParameter default child to be a CfnParameter',
    );
  }
  return resource;
}

function hashSyncTriggerSource(source: unknown): string {
  return createHash('sha256').update(JSON.stringify(source)).digest('hex');
}

function stableName(source: Record<string, unknown>): string {
  if (typeof source.Name === 'string') {
    return source.Name;
  }

  return JSON.stringify(source.Name) ?? '';
}

function parameterSyncTriggerSource(
  parameter: StringParameter,
): Record<string, unknown> {
  const resource = getParameterResource(parameter);
  const resolvedTags: unknown = Stack.of(parameter).resolve(
    resource.tags.renderedTags,
  );

  return {
    Description: resource.description,
    Name: resource.name,
    Tags: resolvedTags,
    Tier: resource.tier,
    Type: resource.type,
    Value: resource.value,
  };
}

export function parameterSyncTrigger(parameter: StringParameter): string {
  return Lazy.string({
    produce: () => hashSyncTriggerSource(parameterSyncTriggerSource(parameter)),
  });
}

export function parameterGroupSyncTrigger(
  parameters: StringParameter[],
): string {
  return Lazy.string({
    produce: () =>
      hashSyncTriggerSource(
        parameters
          .map((parameter) => parameterSyncTriggerSource(parameter))
          .sort((left, right) =>
            stableName(left).localeCompare(stableName(right)),
          ),
      ),
  });
}

function getSyncResource(sync: SopsSync): CfnResource {
  const resource = sync.node.tryFindChild('Resource');
  if (resource === undefined) {
    throw new Error('Expected SopsSync to contain a Resource child');
  }

  const cfnResource = resource.node.defaultChild;
  if (!CfnResource.isCfnResource(cfnResource)) {
    throw new Error(
      'Expected SopsSync Resource default child to be a CfnResource',
    );
  }

  return cfnResource;
}

export function addSyncTrigger(sync: SopsSync, syncTrigger: string): void {
  const cfnResource = getSyncResource(sync);

  cfnResource.addPropertyOverride('SyncTrigger', syncTrigger);
}

export function addSyncDependency(
  sync: SopsSync,
  dependency: StringParameter,
): void {
  getSyncResource(sync).addDependency(getParameterResource(dependency));
}
