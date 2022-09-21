#!/bin/bash

# Exit immediately on failure.
set -e

# Load test parameters.
source scripts/ci/test_parameters.sh

# Disable Docker tests for 386.
export MUTAGEN_TEST_DOCKER="false"

# Perform a local-only build so that we have an agent bundle for testing.
GOARCH=386 go run scripts/build.go --mode=local

# Run tests.
GOARCH=386 go test -p 1 ./pkg/...

# Run tests on SSPL code.
GOARCH=386 go test -v -tags sspl ./sspl/...
