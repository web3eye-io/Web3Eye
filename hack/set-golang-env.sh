 #!/bin/bash
MY_PATH=`cd $(dirname $0);pwd`
ROOT_PATH=$MY_PATH/../

GOVERSION="1.19.13"
GOTMPENV="/tmp/.golang/$GOVERSION"
GOROOT="$GOTMPENV/goroot"
GOPATH="$GOTMPENV/gopath"
GOBIN="$GOROOT/bin"
PATH="$GOBIN:$PATH:$GOBIN"

go_name=go$GOVERSION

echo "Will change go version to $go_name"

go_tar="$go_name.linux-amd64.tar.gz"
go_tar_url="https://go.dev/dl/$go_tar"

go_data=$GOTMPENV

rm -rf $GOTMPENV

mkdir -p $GOPATH
mkdir -p $GOROOT

export GOROOT=$GOROOT
export GOPATH=$GOPATH
export GOBIN=$GOBIN
export PATH=$PATH

[ -z $GOPROXY ] && export GOPROXY="https://proxy.golang.org,direct"

set +e
rc=`go version | grep "$go_name"`
if [ ! $? -eq 0 -o ! -f $GOROOT/.decompressed ]; then
  rm -rf $GOROOT/.decompressed
  echo "Fetching $go_tar from $go_tar_url, stored to $go_data"
  curl -L $go_tar_url -o $go_data/$go_tar
  tar -xvf $go_data/$go_tar --strip-components 1 -C $GOROOT
  touch $GOROOT/.decompressed
fi
set -e