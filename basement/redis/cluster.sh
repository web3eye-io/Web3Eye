#!/bin/bash
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install redis-cluster bitnami/redis-cluster

helm install redis --set auth.password="q123" bitnami/redis
