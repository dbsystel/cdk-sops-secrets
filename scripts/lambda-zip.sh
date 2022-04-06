#!/bin/sh

BASEPATH=$(git rev-parse --show-toplevel)
mkdir -p "$BASEPATH/assets"
cd "$BASEPATH/lambda"
touch -t 202002020000 cdk-sops-secrets
chmod 755 cdk-sops-secrets
ls -la cdk-sops-secrets
sha1sum cdk-sops-secrets
zip -X9om "$BASEPATH/assets/cdk-sops-lambda.zip" cdk-sops-secrets
sha1sum "$BASEPATH/assets/cdk-sops-lambda.zip"
ls -la "$BASEPATH/assets/cdk-sops-lambda.zip"