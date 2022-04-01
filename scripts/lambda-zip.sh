#!/bin/sh

BASEPATH=$(git rev-parse --show-toplevel)
mkdir -p "$BASEPATH/assets"
cd "$BASEPATH/lambda"
zip -X9om "$BASEPATH/assets/cdk-sops-lambda.zip" cdk-sops-secrets
sha1sum "$BASEPATH/assets/cdk-sops-lambda.zip"
ls -la "$BASEPATH/assets/cdk-sops-lambda.zip"