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

FOLDER=mysql-cluster
URL=https://github.com/NpoolPlatform/mysql-cluster.git
if [ ! -d "$FOLDER" ] ; then
    git clone -b branch/web3eye $URL $FOLDER
else
    cd "$FOLDER"
    git pull $URL
fi

sed -i 's/consul_register_enable: "true"/consul_register_enable: "false"/g' $SHELL_FOLDER/mysql-cluster/k8s/01-configmap.yaml
sed -i 's/pmm_admin_enable: "true"/pmm_admin_enable: "false"/g' $SHELL_FOLDER/mysql-cluster/k8s/01-configmap.yaml          

export MYSQL_ROOT_PASSWORD="web321eye"
envsubst < $SHELL_FOLDER/mysql-cluster/k8s/secret.yaml | kubectl apply -f -
kubectl apply -f $SHELL_FOLDER/mysql-env.yaml
kubectl apply -k $SHELL_FOLDER/mysql-cluster/k8s/

sleep 5
kubectl get pods -A| grep mysql

