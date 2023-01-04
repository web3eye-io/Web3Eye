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

kubectl apply -f $SHELL_FOLDER/01-nfs-client.yaml
kubectl apply -f $SHELL_FOLDER/02-nfs-client-provisioner.yaml
kubectl apply -f $SHELL_FOLDER/03-nfs-client-class.yaml

kubectl patch storageclass course-nfs-storage -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"true"}}}'
