#!/bin/sh

BASEPATH=$(git rev-parse --show-toplevel)
cd "$BASEPATH/lambda"
mkdir -p "$BASEPATH/coverage"
go test -race -coverprofile=../coverage/coverage.out -covermode=atomic -v