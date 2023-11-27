#!/bin/bash
# install docs
# https://milvus.io/docs/install_cluster-helm.md
SHELL_FOLDER=$(
    cd "$(dirname "$0")"
    pwd
)
env |grep proxy
env |grep PROXY

MILVUS_CHART_VERSION=4.0.31

helm repo add milvus https://milvus-io.github.io/milvus-helm/
helm pull milvus/milvus --version $MILVUS_CHART_VERSION

helm install milvus -n kube-system ./milvus-$MILVUS_CHART_VERSION.tgz 
# have too many part,so wait more time
sleep 20
kubectl get pods -n kube-system | grep milvus

# helm install milvus milvus/milvus --set cpu=1
# helm install milvus milvus/milvus --set cluster.enabled=false --set etcd.replicaCount=1 --set minio.mode=standalone --set pulsar.enabled=false --set standalone.resources.limits.cpu: 3
