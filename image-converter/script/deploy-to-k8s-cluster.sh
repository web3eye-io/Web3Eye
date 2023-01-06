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

if git_status=$(git status --porcelain --untracked=no 2>/dev/null) && [[ -z "${git_status}" ]]; then
  git_tree_state=clean
fi

version=latest

## For testing or production environment, pass the second variable
if [ "x" != "x$2" ]; then
  version=$2
fi

service_name=$(
  cd $ROOT_FOLDER
  basename $(pwd)
)

kubectl apply -k $ROOT_FOLDER/k8s
