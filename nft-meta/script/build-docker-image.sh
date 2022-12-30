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
echo $version "version"
if [ ! $? -eq 0 ]; then
    version=latest
fi
set -e
echo $version "version"
service_name=$(
    cd $ROOT_FOLDER
    basename $(pwd)
)

## For development environment, pass the second variable
if [[ ${!1-x} == x || "xdevelopment" == "x$1" ]]; then
    version=latest
fi
echo $version ${!1-x} $1 $2 "version"
# TODO: should be official registry
# registry=uhub.service.ucloud.cn
registry=""
OrginazeName=coastlinesss
# OrginazeName=web3eye

if [[ ${!2-x} != x && "x" != $2 ]]; then
    registry=$2/
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
cp $service_source $output_d
cd $output_d

user=$(whoami)
if [ "$user" == "root" ]; then
    docker build -t ${registry}${OrginazeName}/$service_name:$version .
else
    sudo docker build -t ${registry}${OrginazeName}/$service_name:$version .
fi
