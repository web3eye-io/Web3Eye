#!/usr/bin/env bash
SHELL_FOLDER=$(cd "$(dirname "$0")";pwd)
ROOT_FOLDER=$(cd $SHELL_FOLDER/../;pwd)

set -o errexit
set -o nounset
set -o pipefail


$SHELL_FOLDER/build-docker-image.sh

service_name=$(cd $ROOT_FOLDER;basename `pwd`)

version=latest

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


echo "Release docker image for $PLATFORM -- $version"

user=`whoami`
if [ "$user" == "root" ]; then
    docker push ${registry}web3eye/$service_name:$version
else
    sudo docker push ${registry}web3eye/$service_name:$version
fi
