#!/bin/bash
echo "GOROOT: $GOROOT / GOPATH: $GOPATH"
BASEPATH=$(git rev-parse --show-toplevel)
cd "$BASEPATH/lambda"
GOOS=linux
GOARCH=amd64
CGO_ENABLED="1"
CGO_FLAGS="-fdebug-prefix-map=$SRC_ROOT=/tmp/go-build"
GOPROXY=https://proxy.golang.org,direct
ZERO_AR_DATE=1
go build -trimpath -buildvcs=false -gcflags '-N -l'
ls -la cdk-sops-secrets
touch -t 202002020000 cdk-sops-secrets
chmod 755 cdk-sops-secrets
shasum cdk-sops-secrets
ls -la cdk-sops-secrets
