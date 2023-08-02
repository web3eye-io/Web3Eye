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
helm install -n kube-system redis-cluster bitnami/redis-cluster --version 8.6.12

sleep 5
kubectl get pods -n kube-system | grep redis
