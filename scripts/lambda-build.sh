#!/bin/bash
echo "GOROOT: $GOROOT / GOPATH: $GOPATH"
BASEPATH=$(git rev-parse --show-toplevel)
cd "$BASEPATH/lambda"
export GOOS=linux 
export GOARCH=amd64
export GOPROXY=https://proxy.golang.org,direct
export CGO_ENABLED=0
go build -trimpath -buildvcs=false -ldflags="-s -w -buildid="
ls -la cdk-sops-secrets
shasum cdk-sops-secrets
touch -t 202002020000 cdk-sops-secrets
chmod 755 cdk-sops-secrets
ls -la cdk-sops-secrets
shasum cdk-sops-secrets
