#!/usr/bin/env bash
MY_PATH=`cd $(dirname $0);pwd`
ROOT_PATH=$MY_PATH/../
set -o errexit
set -o nounset
set -o pipefail

cd $ROOT_PATH
go get -x github.com/ugorji/go/codec@v1.2.7
go get -d ./...
go mod tidy