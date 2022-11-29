#!/usr/bin/env bash
SHELL_FOLDER=$(cd "$(dirname "$0")";pwd)
ROOT_FOLDER=$(cd $SHELL_FOLDER/../;pwd)
set -o errexit
set -o nounset
set -o pipefail

$SHELL_FOLDER/build.sh

set +e
version=`git describe --tags --abbrev=0`
if [ ! $? -eq 0 ]; then
    version=latest
fi
set -e

service_name=$(cd $ROOT_FOLDER;basename `pwd`)

## For development environment, pass the second variable
if [[ ${!1-x} == x || "xdevelopment" == "x$1" ]]; then
  version=latest
fi

# TODO: should be official registry
# registry=uhub.service.ucloud.cn
registry=""

if [[ ${!2-x} != x  && "x" != $2 ]]; then
  registry=$2/
fi

cd $ROOT_FOLDER

user=`whoami`
if [ "$user" == "root" ]; then
    docker build -t ${registry}web3eye/$service_name:$version .
else
    sudo docker build -t ${registry}web3eye/$service_name:$version .
fi
