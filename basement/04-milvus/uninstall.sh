#!/bin/bash
# install docs
# https://milvus.io/docs/install_cluster-helm.md

helm uninstall -n kube-system milvus
kubectl get pods -n kube-system| grep milvus
kubectl get pvc -n kube-system| grep milvus | awk '{print $1}' | xargs -n1 kubectl delete pvc -n kube-system
