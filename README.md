# Sops Secrets Manager Construct Library

This construct library provides replacement for CDK Secrets, with extended functionality for mozilla/sops.

Using this library it is possible to populate Secrets with values from a mozilla/sops file without additional scripts and steps in the CI stage. Thereby transformations like JSON conversion of YAML files and transformation into a flat, JSONPath like structure can be done.

Secrets filled in this way can be used immediately within the CloudFormation stack.

This way, secrets can be securely stored in git repositories and easily synchronized into AWS SecretsManager Secrets.
