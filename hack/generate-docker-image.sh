#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

PLATFORM=linux/amd64
OUTPUT=./output

pkg=github.com/NpoolPlatform/go-service-framework/pkg/version

OS="${PLATFORM%/*}"
ARCH=$(basename "$PLATFORM")

if git_status=$(git status --porcelain --untracked=no 2>/dev/null) && [[ -z "${git_status}" ]]; then
    git_tree_state=clean
fi

set +e
version=`git describe --tags --abbrev=0`
if [ ! $? -eq 0 ]; then
    version=latest
fi
set -e

service_name=$1
## For development environment, pass the second variable
if [ "xdevelopment" == "x$2" ]; then
  version=latest
fi

registry=uhub.service.ucloud.cn

if [ "x" != $3 ]; then
  registry=$3
fi

echo "Generate docker image for $PLATFORM -- $version"
if [ ! -f $OUTPUT/$PLATFORM/$service_name ]; then
    echo "Run 'make $service_name' before you generate its image"
    exit 1
fi

mkdir -p $OUTPUT/.${service_name}.tmp
cp ./cmd/$service_name/Dockerfile $OUTPUT/.${service_name}.tmp
cp ./cmd/$service_name/*.yaml $OUTPUT/.${service_name}.tmp
cp $OUTPUT/$PLATFORM/$service_name $OUTPUT/.${service_name}.tmp
cd $OUTPUT/.${service_name}.tmp

user=`whoami`
if [ "$user" == "root" ]; then
    docker build -t $registry/entropypool/$service_name:$version .
else
    sudo docker build -t $registry/entropypool/$service_name:$version .
fi
