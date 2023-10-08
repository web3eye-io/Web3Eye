#!/bin/bash
# install docs
# https://helm.sh/docs/intro/install/
SHELL_FOLDER=$(
    cd "$(dirname "$0")"
    pwd
)

curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh