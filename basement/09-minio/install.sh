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

helm repo add minio https://helm.min.io/
helm install web3eye-minio minio/minio -f $SHELL_FOLDER/value.yaml

sleep 5
kubectl get pods | grep minio
