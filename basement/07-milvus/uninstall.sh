#!/bin/bash
# install docs
# https://milvus.io/docs/install_cluster-helm.md

helm uninstall milvus
kubectl get pods | grep milvus
