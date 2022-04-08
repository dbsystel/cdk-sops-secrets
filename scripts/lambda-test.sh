#!/bin/sh

BASEPATH=$(git rev-parse --show-toplevel)
cd "$BASEPATH/lambda"
UPDATE_SNAPS=true go test -v  