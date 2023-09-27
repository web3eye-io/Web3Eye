#!/bin/bash
MY_PATH=`cd $(dirname $0);pwd`
ROOT_PATH=$MY_PATH/../
go_name=go$GOVERSION

set +e
rc=`go version | grep $go_name`
if [ $? -eq 0 ]; then
    return
fi
set -e

echo "Will change go version to $go_name"

go_tar="$go_name.linux-amd64.tar.gz"
go_tar_url="https://go.dev/dl/$go_tar"

go_data=$GOTMPENV

mkdir -p $GOPATH
mkdir -p $GOROOT

[ -z $GOPROXY ] && export GOPROXY="https://proxy.golang.org,direct"

shopt -s expand_aliases
alias go="$go_root/bin/go"

set +e
rc=`go version | grep $go_name`
if [ ! $? -eq 0 ]; then
  set -e
  echo "Fetching $go_tar from $go_tar_url, stored to $go_data"
  curl -L $go_tar_url -o $go_data/$go_tar
  tar -zxvf $go_data/$go_tar --strip-components 1 -C $go_root
fi
set -e
