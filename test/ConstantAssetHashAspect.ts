import * as cdk from 'aws-cdk-lib';
import { IConstruct } from 'constructs';

/**
 * Aspect that ensures constant asset IDs by overriding asset hashes.
 * This is useful for integration tests to prevent snapshot changes due to asset hash variations.
 */
export class ConstantAssetHashAspect implements cdk.IAspect {
  public visit(node: IConstruct): void {
    // Check if the node is a Lambda Function
    if (node instanceof cdk.aws_lambda.Function) {
      const cfnFunction = node.node.defaultChild as cdk.aws_lambda.CfnFunction;
      if (cfnFunction && cfnFunction.code) {
        // Override the asset hash to a constant value
        const code = cfnFunction.code as any;
        if (code.s3Bucket || code.s3Key) {
          // Set a constant hash based on the function ID
          const constantHash = this.generateConstantHash(node.node.id);

          // Override asset parameters
          cfnFunction.addPropertyOverride(
            'Code.S3Key',
            `asset.${constantHash}.zip`,
          );
        }
      }
    }

    // Check if the node is an Asset
    if (node instanceof cdk.aws_s3_assets.Asset) {
      // Override the asset hash
      const cfnResource = node.node.defaultChild as cdk.CfnResource;
      if (cfnResource) {
        const constantHash = this.generateConstantHash(node.node.id);
        cfnResource.addPropertyOverride('AssetHash', constantHash);
      }
    }
  }

  private generateConstantHash(id: string): string {
    // Generate a deterministic hash based on the construct ID
    // Using a simple string-to-hex conversion for reproducibility
    let hash = 0;
    for (let i = 0; i < id.length; i++) {
      const char = id.charCodeAt(i);
      // eslint-disable-next-line no-bitwise
      hash = (hash << 5) - hash + char;
      // eslint-disable-next-line no-bitwise
      hash = hash & hash; // Convert to 32bit integer
    }
    // Return a hex string padded to 64 characters (SHA256 length)
    return Math.abs(hash).toString(16).padStart(64, '0');
  }
}
