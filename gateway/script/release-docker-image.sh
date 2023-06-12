#!/usr/bin/env bash
SHELL_FOLDER=$(
  cd "$(dirname "$0")"
  pwd
)
PROJECT_FOLDER=$(
  cd $SHELL_FOLDER/../
  pwd
)

set -o errexit
set -o nounset
set -o pipefail

PLATFORM=linux/amd64
OUTPUT=$PROJECT_FOLDER/output

pkg=github.com/NpoolPlatform/go-service-framework/pkg/version

OS="${PLATFORM%/*}"
ARCH=$(basename "$PLATFORM")

if git_status=$(git status --porcelain --untracked=no 2>/dev/null) && [[ -z "${git_status}" ]]; then
  git_tree_state=clean
fi

service_name=$(
  cd $PROJECT_FOLDER
  basename $(pwd)
)

if [[ "x" != "x$1" ]]; then
  version=$1
fi

## For development environment, pass the second variable
if [[ "xdev" == "x$1" ]]; then
  version=latest
fi

# TODO: should be official registry
# registry=uhub.service.ucloud.cn
registry=""
OrginazeName=coastlinesss
# OrginazeName=web3eye

if [[ "x" != $2 ]]; then
  registry=$2/
fi

echo "Release docker image for $PLATFORM -- $version"

user=$(whoami)
if [ "$user" == "root" ]; then
  docker push ${registry}${OrginazeName}/$service_name:$version
else
  sudo docker push ${registry}${OrginazeName}/$service_name:$version
fi
