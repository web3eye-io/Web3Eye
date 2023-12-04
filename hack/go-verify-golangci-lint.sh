#!/usr/bin/env bas
MY_PATH=`cd $(dirname $0);pwd`
ROOT_PATH=$MY_PATH/../
LINT_BIN=${ROOT_PATH}/bin

set -o errexit
set -o nounset
set -o pipefail

VERSION_NUM=1.48.0
VERSION=v${VERSION_NUM}
URL_BASE=https://raw.githubusercontent.com/golangci/golangci-lint
URL=$URL_BASE/$VERSION/install.sh

if [[ ! -f .golangci.yml ]]; then
    echo 'ERROR: missing .golangci.yml in repo root' >&2
    exit 1
fi

if ! command -v gofumpt; then
    go install mvdan.cc/gofumpt@v0.3.1
fi

PATH=$LINT_BIN:$PATH
set +e
rc=`golangci-lint version | grep $VERSION_NUM`
set -e
if [ ! $? -eq 0 ]; then
  curl -sfL $URL | sh -s $VERSION -b $LINT_BIN
fi

golangci-lint version
golangci-lint linters
golangci-lint run "$@"
