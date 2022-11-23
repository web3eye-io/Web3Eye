#!/usr/bin/env bash

set -euo pipefail

# Default timeout is 1800s
TEST_TIMEOUT=1800

for arg in "$@"
do
    case $arg in
        -t=*|--timeout=*)
        TEST_TIMEOUT="${arg#*=}"
        shift
        ;;
        -t|--timeout)
        TEST_TIMEOUT="$2"
        shift
        shift
    esac
done

REPO_ROOT=$(git rev-parse --show-toplevel)
cd "${REPO_ROOT}"

GO111MODULE=on go test -v -timeout="${TEST_TIMEOUT}s" -count=1 -cover -coverprofile coverage.out ./cmd/...
go tool cover -html coverage.out -o coverage.html
