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

helm repo add bitnami https://charts.bitnami.com/bitnami
# helm install kafka oci://registry-1.docker.io/bitnamicharts/kafka
helm install -n kube-system -f $SHELL_FOLDER/value.yaml kafka bitnami/kafka
sleep 5
kubectl get pods  -n kube-system | grep kafka
