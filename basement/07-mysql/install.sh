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

cd $SHELL_FOLDER

git clone -b branch/web3eye https://github.com/NpoolPlatform/mysql-cluster.git
sed -i 's/consul_register_enable: "true"/consul_register_enable: "false"/g' $SHELL_FOLDER/mysql-cluster/k8s/01-configmap.yaml
sed -i 's/pmm_admin_enable: "true"/pmm_admin_enable: "false"/g' $SHELL_FOLDER/mysql-cluster/k8s/01-configmap.yaml          

kubectl apply -f $SHELL_FOLDER/mysql-cluster/mysql-env.yaml
export MYSQL_ROOT_PASSWORD="web321eye"
envsubst < $SHELL_FOLDER/mysql-cluster/k8s/secret.yaml | kubectl apply -f -
kubectl apply -k $SHELL_FOLDER/mysql-cluster/k8s/

sleep 5
kubectl get pods | grep mysql
