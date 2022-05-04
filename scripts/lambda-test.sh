#!/bin/sh

BASEPATH=$(git rev-parse --show-toplevel)
cd "$BASEPATH/lambda"
UPDATE_SNAPS=true go test -race -coverprofile=coverage.out -covermode=atomic -v