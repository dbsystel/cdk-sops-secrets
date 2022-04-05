# 🔑 Sops Secrets Construct Library


[![npm](https://img.shields.io/npm/v/cdk-sops-secrets.svg)](https://www.npmjs.com/package/cdk-sops-secrets)&nbsp;&nbsp;
[![Known Vulnerabilities](https://snyk.io/test/github/markussiebert/cdk-sops-secrets/badge.svg)](https://snyk.io/test/github/markussiebert/cdk-sops-secrets)&nbsp;&nbsp;
[![release](https://github.com/markussiebert/cdk-sops-secrets/actions/workflows/release.yml/badge.svg)](https://github.com/markussiebert/cdk-sops-secrets/actions/workflows/release.yml)

This construct library provides a replacement for CDK SecretsManager Secrets, with extended functionality for mozilla/sops.

Using this library it is possible to populate Secrets with values from a mozilla/sops file without additional scripts and steps in the CI stage. Thereby transformations like JSON conversion of YAML files and transformation into a flat, JSONPath like structure can be done.

Secrets filled in this way can be used immediately within the CloudFormation stacks and dynamic references.

This way, secrets can be securely stored in git repositories and easily synchronized into AWS SecretsManager Secrets.

# Motivation

This project was created to solve a recurring problem of syncing mozilla/sops secrets into AWS SecretsManager in a convenient, secure way.
 
Apart from that, the project was for learning:
- Write a golang lambda
- Writing unit tests incl. mocks in golang
- Reproducible builds of golang binaries (byte-by-byte identical)
- Build reproducible zips (byte-by-byte identical)
- Release an npm package
- Setting up projects with projen
- CI/CD with github actions
- CDK unit and integration tests