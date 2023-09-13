#!/bin/sh

BASEPATH=$(git rev-parse --show-toplevel)
mkdir -p "$BASEPATH/assets"
cd "$BASEPATH/lambda"
touch -t 202002020000 bootstrap
chmod 755 bootstrap
ls -la bootstrap
sha1sum bootstrap
zip -X9om "$BASEPATH/assets/cdk-sops-lambda.zip" bootstrap
sha1sum "$BASEPATH/assets/cdk-sops-lambda.zip"
ls -la "$BASEPATH/assets/cdk-sops-lambda.zip"