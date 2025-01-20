#!/bin/bash
echo "GOROOT: $GOROOT / GOPATH: $GOPATH"
BASEPATH=$(git rev-parse --show-toplevel)
cd "$BASEPATH/lambda"
export GOOS=linux 
export GOARCH=amd64
export GOPROXY=https://proxy.golang.org,direct
export CGO_ENABLED=0
go build -trimpath -buildvcs=false -tags lambda.norpc -o bootstrap -ldflags="-s -w -buildid="
ls -la bootstrap
sha1sum bootstrap
touch -t 202002020000 bootstrap
chmod 755 bootstrap
ls -la bootstrap
sha1sum bootstrap
