#!/bin/bash
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

helm uninstall kafka
kubectl get pods | grep kafka
kubectl get pvc | grep kafka | awk '{print $1}' | xargs -n1 kubectl delete pvc
