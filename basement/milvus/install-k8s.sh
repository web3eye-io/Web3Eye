#!/bin/bash
# install docs
# https://milvus.io/docs/install_cluster-milvusoperator.md

# Install cert-manager
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.5.3/cert-manager.yaml

# Install Milvus Operator
# kubectl apply -f https://raw.githubusercontent.com/milvus-io/milvus-operator/main/deploy/manifests/deployment.yaml
helm install milvus-operator \
  -n milvus-operator --create-namespace \
  --wait --wait-for-jobs \
  https://github.com/milvus-io/milvus-operator/releases/download/v0.7.0/milvus-operator-0.7.0.tgz
kubectl get pods -n milvus-operator

# Install a Milvus cluster
kubectl apply -f https://raw.githubusercontent.com/milvus-io/milvus-operator/main/config/samples/milvus_cluster_default.yaml
# minimum install 
# kubectl apply -f https://github.com/milvus-io/milvus-operator/blob/main/config/samples/milvus_cluster_minimum.yaml
kubectl get milvus my-release -o yaml
kubectl get pods
