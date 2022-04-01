#!/bin/bash

BASEPATH=$(git rev-parse --show-toplevel)
cd "$BASEPATH/lambda"
GOOS=linux GOARCH=amd64 GOFLAGS="-trimpath" GOPROXY=https://proxy.golang.org,direct go build
ls -la cdk-sops-secrets
touch -t 202002020000 cdk-sops-secrets
chmod 755 cdk-sops-secrets
shasum cdk-sops-secrets
ls -la cdk-sops-secrets
