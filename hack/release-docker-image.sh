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

version=latest

## For testing or production environment, pass the second variable
if [ "x" != "x$2" ]; then
  version=$2
fi

registry=uhub.service.ucloud.cn

if [ "x" != $3 ]; then
  registry=$3
fi

service_name=$1

echo "Release docker image for $PLATFORM -- $version"

user=`whoami`
if [ "$user" == "root" ]; then
    docker push $registry/entropypool/$service_name:$version
else
    sudo docker push $registry/entropypool/$service_name:$version
fi
