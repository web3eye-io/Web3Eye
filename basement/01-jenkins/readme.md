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

<!-- 构建jenkins已安装工具 -->
1.docker
2.kubectl
3.helm
4.nodejs npm

<!-- 部署完jenkins需要做的步骤 -->
1.确认kube-apiserver的地址
/root/.kube 当中的配置 IP 要指向kube-apiserver的地址不能指向127.0.0.1

2.安装Go插件
3.配置Git 接受第一次连接（Dashboard > Manage Jenkins > Configure Global Security）
4.登录Docker（docker login）

5.设置代理（不知道能不能用到部署步骤，部署步骤需要连接k8s,可能会连接不到本地地址）
Dashboard > 系统管理 > Configure system 设置环境变量 设置环境变量即可

6.添加凭证web3eye-git-ssh-private-key，保证jenkins能操作仓库
<!-- 还需处理的问题 -->
1.npm需要安装的命令，考虑用jenkins插件安装，在设置nodejs插件时安装
jenkins中还没有nodejs的环境，安装了node记得一起安装yarn命令
安装n
安装yarn
安装eslint

