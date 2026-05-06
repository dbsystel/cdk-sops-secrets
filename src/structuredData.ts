import * as fs from 'fs';
import * as YAML from 'yaml';

export interface JSONObject {
  [key: string]: unknown;
}

export type StructuredFileFormat = 'json' | 'yaml' | 'dotenv';

function flattenJSON(
  data: JSONObject,
  parentKey: string = '',
  result: JSONObject = {},
  keySeparator = '',
): JSONObject {
  for (const key of Object.keys(data)) {
    const value = data[key];
    const newKey = parentKey ? `${parentKey}${keySeparator}${key}` : key;

    if (Array.isArray(value)) {
      value.forEach((item, index) => {
        const arrayKey = `${newKey}[${index}]`;
        if (item !== null && typeof item === 'object' && !Array.isArray(item)) {
          flattenJSON(item as JSONObject, arrayKey, result, keySeparator);
        } else {
          result[arrayKey] = item;
        }
      });
    } else if (
      value !== null &&
      typeof value === 'object' &&
      !Array.isArray(value)
    ) {
      flattenJSON(value as JSONObject, newKey, result, keySeparator);
    } else {
      result[newKey] = value;
    }
  }
  return result;
}

function ensureObject(data: unknown, filePath: string): JSONObject {
  if (data === null || typeof data !== 'object' || Array.isArray(data)) {
    throw new Error(`Expected structured object content in ${filePath}`);
  }
  return data as JSONObject;
}

function parseDotenv(content: string): JSONObject {
  const result: JSONObject = {};

  for (const line of content.split('\n')) {
    if (line !== '' && !line.startsWith('#')) {
      const parts = line.split('=', 2);
      if (parts.length === 2) {
        const key = parts[0].trim();
        const value = parts[1].trim();
        result[key] = value;
      }
    }
  }

  return result;
}

export function inferStructuredFileFormat(
  sopsFilePath: string,
): StructuredFileFormat | undefined {
  const extension = sopsFilePath.split('.').pop();

  switch (extension) {
    case 'json':
      return 'json';
    case 'yaml':
    case 'yml':
      return 'yaml';
    case 'dotenv':
    case 'env':
      return 'dotenv';
    default:
      return undefined;
  }
}

export function parseStructuredFile(
  sopsFilePath: string,
  fileFormat?: StructuredFileFormat,
): JSONObject {
  const resolvedFormat = fileFormat ?? inferStructuredFileFormat(sopsFilePath);
  if (resolvedFormat === undefined) {
    throw new Error(
      `Unsupported structured file format for ${sopsFilePath}. Supported formats: json, yaml, dotenv`,
    );
  }

  const content = fs.readFileSync(sopsFilePath, 'utf-8');

  switch (resolvedFormat) {
    case 'json':
      return ensureObject(JSON.parse(content), sopsFilePath);
    case 'yaml':
      return ensureObject(YAML.parse(content), sopsFilePath);
    case 'dotenv':
      return parseDotenv(content);
  }
}

export function flattenStructuredFile(
  sopsFilePath: string,
  keySeparator: string,
  fileFormat?: StructuredFileFormat,
): JSONObject {
  return flattenJSON(
    parseStructuredFile(sopsFilePath, fileFormat),
    '',
    {},
    keySeparator,
  );
}

export function flattenStructuredFileToStringMap(
  sopsFilePath: string,
  keySeparator: string,
  fileFormat?: StructuredFileFormat,
): Record<string, string> {
  const flat = flattenStructuredFile(sopsFilePath, keySeparator, fileFormat);

  return Object.fromEntries(
    Object.entries(flat).map(([key, value]) => [
      key,
      value === null || value === undefined ? '' : String(value),
    ]),
  );
}
