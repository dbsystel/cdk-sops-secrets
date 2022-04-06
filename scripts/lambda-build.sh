#!/bin/bash
echo "GOROOT: $GOROOT / GOPATH: $GOPATH"
BASEPATH=$(git rev-parse --show-toplevel)
cd "$BASEPATH/lambda"
GOOS=linux
GOARCH=amd64
GOPROXY=https://proxy.golang.org,direct
go build -trimpath -buildvcs=false -ldflags="-s -w -buildid="
ls -la cdk-sops-secrets
shasum cdk-sops-secrets
touch -t 202002020000 cdk-sops-secrets
chmod 755 cdk-sops-secrets
ls -la cdk-sops-secrets
shasum cdk-sops-secrets

