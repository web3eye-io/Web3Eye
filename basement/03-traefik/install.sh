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

kubectl apply -k $SHELL_FOLDER

sleep 5
kubectl get pods -n kube-system | grep traefik
