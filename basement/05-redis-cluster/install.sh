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
helm install redis-cluster --set redis.password="default" bitnami/redis-cluster
kubectl get pods | grep redis

# TODO: bitnami安装redis默认没生成secret，需要手动生成一个，提供密码
# helm install redis --set auth.password="q123" bitnami/redis
