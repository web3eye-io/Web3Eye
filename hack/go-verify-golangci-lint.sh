#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

VERSION=v1.48.0
URL_BASE=https://raw.githubusercontent.com/golangci/golangci-lint
URL=$URL_BASE/$VERSION/install.sh

if [[ ! -f .golangci.yml ]]; then
    echo 'ERROR: missing .golangci.yml in repo root' >&2
    exit 1
fi

if ! command -v gofumpt; then
    go install mvdan.cc/gofumpt@v0.3.1
fi

if ! command -v golangci-lint; then
    curl -sfL $URL | sh -s $VERSION
    PATH=$PATH:bin
fi

golangci-lint version
golangci-lint linters
golangci-lint run "$@"
