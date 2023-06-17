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

helm uninstall redis-cluster
kubectl get pvc | grep redis-cluster | awk '{print $1}' | xargs -n1 kubectl delete pvc
