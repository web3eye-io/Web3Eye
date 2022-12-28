#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

VERSION=v1.46.2

URL=github.com/golangci/golangci-lint/cmd/golangci-lint@$VERSION

if [[ ! -f .golangci.yml ]]; then
    echo 'ERROR: missing .golangci.yml in repo root' >&2
    exit 1
fi

if ! command -v gofumpt; then
    go install mvdan.cc/gofumpt@v0.3.1
fi

if ! command -v golangci-lint; then
    export CGO_ENABLED=0
    go install $URL
    PATH=$PATH:bin
fi

golangci-lint version
golangci-lint linters
golangci-lint run "$@"
