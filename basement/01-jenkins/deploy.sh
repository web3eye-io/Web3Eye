#!/bin/bash
SHELL_FOLDER=$(
    cd "$(dirname "$0")"
    pwd
)
PROJECT_FOLDER=$(
    cd $SHELL_FOLDER/../
    pwd
)



OrginazeName=web3eye

user=$(whoami)
service_name=jenkins
version=latest
if [ "$user" == "root" ]; then
    docker run \
        -u 0\
        --name jenkins \
        -d \
        --privileged \
        -p 18080:8080 \
        -p 50000:50000 \
        -v /opt/share/jenkins:/var/lib/jenkins   \
        -v /sys/fs/cgroup:/sys/fs/cgroup:ro      \
        --tmpfs /tmp:exec --tmpfs /run --tmpfs /run/lock --tmpfs /var/run \
        -v /var/run/docker.sock:/var/run/docker.sock  \
        -v /root/.kube:/root/.kube  \
        ${registry}/${OrginazeName}/$service_name:$version
    sleep 5
    docker exec -it jenkins cat /var/jenkins_home/secrets/initialAdminPassword
else
    sudo docker run \
        -u 0\
        --name jenkins \
        -d \
        --privileged \
        -p 18080:8080 \
        -p 50000:50000 \
        -v /opt/share/jenkins:/var/lib/jenkins   \
        -v /sys/fs/cgroup:/sys/fs/cgroup:ro      \
        --tmpfs /tmp:exec --tmpfs /run --tmpfs /run/lock --tmpfs /var/run \
        -v /var/run/docker.sock:/var/run/docker.sock  \
        -v /root/.kube:/root/.kube  \
        ${registry}/${OrginazeName}/$service_name:$version
    sleep 5
    sudo docker exec -it jenkins cat /var/jenkins_home/secrets/initialAdminPassword
fi
