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

helm uninstall  -n kube-system web3eye-minio
kubectl get pvc  -n kube-system | grep web3eye-minio | awk '{print $1}' | xargs -n1 kubectl delete pvc -n kube-system
