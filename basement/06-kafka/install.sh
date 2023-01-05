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

helm repo add bitnami https://charts.bitnami.com/bitnami
helm install -f 01-kafka.yaml kafka bitnami/kafka
kubectl get pods | grep kafka
