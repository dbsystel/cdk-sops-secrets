#!/bin/bash

BASEPATH=$(git rev-parse --show-toplevel)

$BASEPATH/scripts/clean.sh
docker run -v $BASEPATH:/__w/cdk-sops-secrets/cdk-sops-secrets golang:1.17.8-buster /bin/bash -c "cd /__w/cdk-sops-secrets/cdk-sops-secrets && scripts/clean.sh && scripts/lambda-build.sh"
$BASEPATH/scripts/lambda-zip.sh
