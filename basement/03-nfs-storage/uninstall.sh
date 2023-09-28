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

helm uninstall -n kube-system default-nfs-provisioner
kubectl get pods -n kube-system | grep nfs
kubectl get storageclass -A
