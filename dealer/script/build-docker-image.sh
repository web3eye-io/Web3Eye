#!/usr/bin/env bash
SHELL_FOLDER=$(
    cd "$(dirname "$0")"
    pwd
)
ROOT_FOLDER=$(
    cd $SHELL_FOLDER/../
    pwd
)
set -o errexit
set -o nounset
set -o pipefail

PLATFORM=linux/amd64
OUTPUT=$ROOT_FOLDER/output

pkg=github.com/NpoolPlatform/go-service-framework/pkg/version

OS="${PLATFORM%/*}"
ARCH=$(basename "$PLATFORM")

if git_status=$(git status --porcelain --untracked=no 2>/dev/null) && [[ -z "${git_status}" ]]; then
    git_tree_state=clean
fi

set +e
version=$(git describe --tags --abbrev=0)

if [ ! $? -eq 0 ]; then
    version=latest
fi
set -e

service_name=$(
    cd $ROOT_FOLDER
    basename $(pwd)
)

if [[ "x" != "x$1" ]]; then
  version=$1
fi

## For development environment, pass the second variable
if [[ "xdev" == "x$1" ]]; then
    version=latest
fi
registry=uhub.service.ucloud.cn
OrginazeName=web3eye

if [[ "x" != $2 ]]; then
    registry=$2
fi

service_source=$OUTPUT/$PLATFORM/$service_name

echo "Generate docker image for $PLATFORM -- $version"
if [ ! -f $service_source ]; then
    echo "Run 'make $service_name' before you generate its image"
    exit 1
fi

output_d=$OUTPUT/.${service_name}.tmp
config_d=$ROOT_FOLDER/cmd/$service_name

mkdir -p $output_d
cp $config_d/Dockerfile $output_d
cp $config_d/*.yaml $output_d || echo "have no yaml files"
cp $service_source $OUTPUT/$PLATFORM/lotus $OUTPUT/$PLATFORM/ipfs $output_d
cd $output_d

user=$(whoami)
if [ "$user" == "root" ]; then
    docker build -t ${registry}/${OrginazeName}/$service_name:$version .
else
    sudo docker build -t ${registry}/${OrginazeName}/$service_name:$version .
fi
