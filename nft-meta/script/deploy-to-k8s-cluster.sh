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

version=latest

if [ "x$1" != "x" ]; then
  version=$1
fi

if [ "$1" == "dev" ]; then
  version=latest
fi

# TODO: support change registry
## For testing or production environment, pass the second variable
# if [[ "x" != "x$2" ]]; then
#   DOCKER_REGISTRY=$2
# fi

service_name=$(
  cd $ROOT_FOLDER
  basename $(pwd)
)

echo "Deploy docker image for $PLATFORM -- $version"

sed -i "s/$service_name:latest/$service_name:$version/g" $ROOT_FOLDER/cmd/$service_name/k8s/02-$service_name.yaml
# sed -i "s/uhub.service.ucloud.cn/$DOCKER_REGISTRY/g" cmd/$service_name/k8s/02-$service_name.yaml

kubectl get deployment | grep $service_name
if [ $? == 0 ]; then
  kubectl apply -k $ROOT_FOLDER/cmd/$service_name/k8s
else
  kubectl replace -k $ROOT_FOLDER/cmd/$service_name/k8s
fi
