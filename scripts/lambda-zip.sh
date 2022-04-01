#!/bin/bash

BASEPATH=$(git rev-parse --show-toplevel)
mkdir -p "$BASEPATH/assets"
cd "$BASEPATH/lambda"
zip -9om "$BASEPATH/assets/cdk-sops-lambda.zip" cdk-sops-secrets
shasum "$BASEPATH/assets/cdk-sops-lambda.zip"
ls -la "$BASEPATH/assets/cdk-sops-lambda.zip"