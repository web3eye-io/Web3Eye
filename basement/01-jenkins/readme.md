docker pull jenkins/jenkins:centos7

docker run \
  --name jenkins \
  -d \
  -p 8080:8080 \
  -p 50000:50000 \
  <!-- -v jenkins-data:/var/jenkins_home \ -->
jenkins/jenkins

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
  jenkins/jenkins:centos7

docker exec -it jenkins /bin/bash

docker exec -it jenkins cat /var/jenkins_home/secrets/initialAdminPassword

<!-- 部署完jenkins需要做的步骤 -->
/root/.kube 当中的配置 IP 要指向kube-apiserver的地址不能指向127.0.0.1
安装Go插件
安装Kubectl插件
配置Git 接受第一次连接（Dashboard > Manage Jenkins > Configure Global Security）
安装docker-ce（<https://yeasy.gitbook.io/docker_practice/install/centos>）
登录Docker（docker login）
git config --global user.email "670884108@qq.com"
git config --global user.name "Greapefurit"

设置代理（不知道能不能用到部署步骤，部署步骤需要连接k8s,可能会连接不到本地地址）
Dashboard > 系统管理 > Configure system 设置环境变量 设置环境变量即可

<!-- 还需处理的问题 -->
1.npm需要安装的命令，考虑用jenkins插件安装，在设置nodejs插件时安装
jenkins中还没有nodejs的环境，安装了node记得一起安装yarn命令
安装n
安装yarn
安装eslint

2.需要安装docker