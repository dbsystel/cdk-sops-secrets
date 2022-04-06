#!/bin/bash

BASEPATH=$(git rev-parse --show-toplevel)
mkdir -p "$BASEPATH/assets"
cd "$BASEPATH/lambda"
dzip "$BASEPATH/assets/cdk-sops-lambda.zip" cdk-sops-secrets
rm -Rf cdk-sops-secrets
shasum "$BASEPATH/assets/cdk-sops-lambda.zip"
ls -la "$BASEPATH/assets/cdk-sops-lambda.zip"