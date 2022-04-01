#!/bin/bash

BASEPATH=$(git rev-parse --show-toplevel)

$BASEPATH/scripts/clean.sh
docker run -v $BASEPATH:/__w/cdk-sops-secrets/cdk-sops-secrets --platform linux/amd64 golang:1.18-buster /bin/bash -c "cd /__w/cdk-sops-secrets/cdk-sops-secrets && scripts/clean.sh && scripts/lambda-build.sh" -m
x86_64
$BASEPATH/scripts/lambda-zip.sh
npx projen integ:secret:snapshot