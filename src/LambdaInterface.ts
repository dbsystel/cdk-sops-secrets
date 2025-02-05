/* eslint-disable */
/**
 * This file was automatically generated by json-schema-to-typescript.
 * DO NOT MODIFY IT BY HAND. Instead, modify the source JSONSchema file,
 * and run json-schema-to-typescript to regenerate this file.
 */

export interface SopsSyncResourcePropertys {
  ResourceType: "SECRET" | "PARAMETER_MULTI" | "PARAMETER";
  Format: "json" | "yaml" | "dotenv" | "binary";
  OutputFormat: "json" | "yaml" | "dotenv" | "binary";
  SecretARN?: string;
  ParameterName?: string;
  ParameterKeyPrefix?: string;
  EncryptionKey?: string;
  SopsS3File?: SopsS3File;
  SopsInline?: SopsInline;
  Flatten: boolean;
  StringifyValues: boolean;
  FlattenSeparator: string;
}
export interface SopsS3File {
  Bucket: string;
  Key: string;
}
export interface SopsInline {
  Content: string;
  Hash: string;
}
