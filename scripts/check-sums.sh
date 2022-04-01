#!/bin/bash

BASEPATH=$(git rev-parse --show-toplevel)

shasum "$BASEPATH/test-secrets/sopsfile.enc-age.json"