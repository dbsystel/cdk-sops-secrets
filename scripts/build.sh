#!/bin/bash


BASEPATH=$(git rev-parse --show-toplevel)
$BASEPATH/scripts/clean.sh
$BASEPATH/scripts/lambda-build.sh
$BASEPATH/scripts/lambda-zip.sh
