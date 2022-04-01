#!/bin/bash


BASEPATH=$(git rev-parse --show-toplevel)
cd scripts
./clean.sh
./lambda-build.sh
./lambda-zip.sh
