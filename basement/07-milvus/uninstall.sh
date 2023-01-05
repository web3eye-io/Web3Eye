#!/bin/bash
# install docs
# https://milvus.io/docs/install_cluster-helm.md

helm uninstall milvus
kubectl get pods | grep milvus
kubectl get pvc | grep milvus | awk '{print $1}' | xargs -n1 kubectl delete pvc
