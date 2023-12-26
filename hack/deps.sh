#!/usr/bin/env bash
MY_PATH=`cd $(dirname $0);pwd`
ROOT_PATH=$MY_PATH/../
source $MY_PATH/set-golang-env.sh

set -o errexit
set -o nounset
set -o pipefail

go env

cd $ROOT_PATH
go get -x github.com/ugorji/go/codec@v1.2.7
go get -d ./...
go mod tidy

go env