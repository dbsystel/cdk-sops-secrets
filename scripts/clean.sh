#!/bin/bash

BASEPATH=$(git rev-parse --show-toplevel)

rm -Rf "$BASEPATH/assets"
rm -f "$BASEPATH/lambda/cdk-sops-secrets"