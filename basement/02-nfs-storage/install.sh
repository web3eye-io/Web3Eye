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

helm repo add nfs-subdir-external-provisioner https://kubernetes-sigs.github.io/nfs-subdir-external-provisioner
helm repo update

helm install -n kube-system default-nfs-provisioner \
    nfs-subdir-external-provisioner/nfs-subdir-external-provisioner \
    -f $SHELL_FOLDER/value.yaml
kubectl get pods -n kube-system | grep nfs
kubectl get storageclass -A
