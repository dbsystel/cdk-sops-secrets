#!/bin/bash

BASEPATH=$(git rev-parse --show-toplevel)
cd "$BASEPATH/lambda"
export GOOS=linux
export GOARCH=amd64
go build
chmod a+x cdk-sops-secrets
zip cdk-sops-secrets.zip cdk-sops-secrets
rm -f cdk-sops-secrets
cd ..
mkdir -p assets
mv lambda/cdk-sops-secrets.zip assets/cdk-sops-secrets.zip