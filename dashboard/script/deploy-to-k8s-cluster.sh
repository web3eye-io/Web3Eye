#!/bin/bash
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

if git_status=$(git status --porcelain --untracked=no 2>/dev/null) && [[ -z "${git_status}" ]]; then
  git_tree_state=clean
fi

if [[ "x" == "x$1" ]]; then
  version=latest
else
  version=$1
fi

# TODO: support change registry
## For testing or production environment, pass the second variable
# if [[ "x" != "x$2" ]]; then
#   DOCKER_REGISTRY=$2
# fi

service_name=$(
  cd $PROJECT_FOLDER
  basename $(pwd)
)

sed -i "s/$service_name:latest/$service_name:$version/g" $PROJECT_FOLDER/k8s/01-$service_name.yaml
sed -i "s/uhub.service.ucloud.cn/$DOCKER_REGISTRY/g" $PROJECT_FOLDER/k8s/01-$service_name.yaml
sed -i "s/web3eye.webui.io/$CLOUD_ROOT_DOMAIN/g" $PROJECT_FOLDER/k8s/02-ingress-vpn.yaml
sed -i "s/web3eye-webui-io/$CLOUD_CERT_NAME/g" $PROJECT_FOLDER/k8s/02-ingress-vpn.yaml

set +e

# check have deployment
kubectl get deployment | grep $service_name
if [ $? == 0 ]; then
  kubectl replace -k $PROJECT_FOLDER/k8s
else
  kubectl apply -k $PROJECT_FOLDER/k8s
fi
