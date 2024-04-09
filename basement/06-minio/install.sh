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

helm repo add minio https://charts.min.io
helm install -n kube-system web3eye-minio minio/minio -f $SHELL_FOLDER/value.yaml --set mode=standalone,image.tag=RELEASE.2022-04-16T04-26-02Z
# helm repo add minio https://charts.min.io
# kubectl create ns ame
# helm install --set accessKey=admin,secretKey=12341234,rootUser=admin,rootPassword=12341234,mode=distributed,replicas=4,service.type=NodePort,persistence.storageClass=nfs-storage,persistence.size=500Gi,resources.requests.memory=4Gi -name minio minio/minio --debug --wait --timeout 10m
sleep 5
kubectl get pods -n kube-system | grep web3eye-minio