## 安装说明

每个小节请阅读完成再操作，以免理解错误上下文意思，同时欢迎提Issue帮助改进。

## 机器准备

主机规划表

| IP           | hostname | 硬件配置                           | 角色                                     | 位置 | 系统       |
|--------------|----------|--------------------------------|------------------------------------------|------|------------|
| 172.16.29.49 |          | CPU:4核  内存：8G  磁盘：50G         | gateway(for Scientific Internet)、Jenkins | IDC  | centos7    |
| 172.16.29.51 | node1    | CPU:8核  内存：16G  磁盘：200G       | k8s-master,jenkins                       | IDC  | ubuntu20.4 |
| 172.16.29.52 | node2    | CPU:16核  内存：32G  磁盘：100G      | k8s-worker                               | IDC  | ubuntu20.4 |
| 172.16.29.53 | node3    | CPU:16核  内存：32G  磁盘：100G      | k8s-worker                               | IDC  | ubuntu20.4 |
| 172.16.29.54 | node4    | CPU:16核  内存：32G  磁盘：100G+400G | k8s-worker,nfs-server                    | IDC  | ubuntu20.4 |
| 172.23.10.87 | node1    | CPU:8核  内存：16G  磁盘：100G       | k8s-master,k8s-worker                    | AWS  | ubuntu20.4 |

系统：Ubuntu20.04 or Ubuntu22.04

### 安装V2rayA（可选）

Gateway机器主要为IDC提供统一的网络控制，此处也可选其他方式实现，主要为了更好的科学上网，如果没有科学上网的需求可不要gateway节点。

选择V2rayA主要考虑代理能力强，模式多且设置便捷，而且可用Web操作，在无图形化服务器非常好用。

安装方法：https://v2raya.org/docs/prologue/installation/redhat/

```Shell
curl -Ls https://mirrors.v2raya.org/go.sh | sudo bash
sudo systemctl disable v2ray --now ### Xray 需要替换服务为 xray

wget https://github.com/v2rayA/v2rayA/releases/download/v2.2.4/installer_redhat_x64_2.2.4.rpm -O /tmp/v2raya.rpm
rpm -i /tmp/v2raya.rpm

sudo systemctl start v2raya.service
sudo systemctl enable v2raya.service
```

安装完成后导入代理节点即可使用，同时将其他机器的网关设置成Gateway机器的IP，其他机器也能科学上网了。

## 初始化系统配置

按照配置正常安装系统即可，若是在虚拟机上安装可考虑用克隆的方式提高安装速度。

**K8s集群内的所有机器**

- 设置root密码

- 开启root的ssh登录

- 设置gateway（可选）

```Shell
# 切换到root用户
su root
# 设置密码
passwd

#允许root登录
sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/g' /etc/ssh/sshd_config
#允许密码登录
sed -i 's/#PasswordAuthentication yes/PasswordAuthentication yes/g' /etc/ssh/sshd_config
systemctl restart sshd

# 设置gateway，修改gateway配置成安装了V2rayA的机器IP
vim /etc/netplan/00-installer-config.yaml
netplan apply
```

初始化K8s集群机器，在**jenkins所在服务器**执行，这里jenkins和gateway是同一个服务器。

初始化系统配置内容如下，配置好后执行脚本即可

1.按照规划修改IP和hostname

2.Master免密登录其他机器

3.清理已安装的docker

4.安装python3

复制脚本到.sh文件中，并配置后执行

```Shell
#!/bin/bash

# a host info like:
# IP hostname
hosts=(
    172.16.29.51 node1
    172.16.29.52 node2
    172.16.29.53 node3
    172.16.29.54 node4
)

# ssh-keygen in me
# value: true or false
enableSSHKeygen=true

# clear old docker in me
# value: true or false
clearOldDocker=true

if [ $enableSSHKeygen ];then
    echo "start ssh-keygen" 
    ssh-keygen
    echo "end ssh-keygen" 
fi


rem=0
for index in "${!hosts[@]}";   
do   
    if [ $rem != 0 ];then
        rem=0
        continue
    fi
    rem=1

    ssh-copy-id "${hosts[$index]}"
done  


rem=0
for index in "${!hosts[@]}";   
do   
    if [ $rem != 0 ];then
        rem=0
        continue
    fi
    rem=1
    
    if [ $clearOldDocker ];then
        echo "start clear old-docker" 
        # ssh "${hosts[$index]}" "for pkg in docker.io docker-doc docker-compose podman-docker containerd runc; do sudo apt-get remove \$pkg; done"
        echo "end clear old-docker" 
    fi

    index1=`expr $index + 1`
    ssh "${hosts[$index]}" "hostnamectl set-hostname \"${hosts[$index1]}\""
    ssh "${hosts[$index]}" "apt install python3 -y"
done  

```


## 准备Jenkins环境

Jenkins管理K8s集群的生命周期，所以存在独立与K8s集群的机器上。

安装Docker:

```Shell
# 清除旧版本Docker
yum remove -y docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-engine \
                  docker \
                  docker-ce \
                  docker-ce-cli

yum install -y yum-utils
yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo

yum install docker
```

启动Jenkins

```Shell
docker run -d -p 9090:8080 -p 60000:50000 \
    --name jenkins-centos-7 \
    -v /opt/share/jenkins:/var/lib/jenkins   \
    -v /sys/fs/cgroup:/sys/fs/cgroup:ro      \
    --tmpfs /tmp:exec --tmpfs /run --tmpfs /run/lock --tmpfs /var/run \
    -v /var/run/docker.sock:/var/run/docker.sock \
    --restart always \
    --privileged uhub.service.ucloud.cn/entropypool/jenkins:centos-7 /usr/sbin/init

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
  -v /root/.kube:/root/.kube  \
  coastlinesss/jenkins
```

获取jenkins初始密码

```shell
docker exec -it jenkins cat /var/jenkins_home/secrets/initialAdminPassword
```

访问jenkins web页面(172.16.29.51:18080)，完成Jenkins初始配置，如添加用户等，在安装插件时可先安装建议插件。

### 配置Jenkins环境

**配置Git** 接受第一次连接（Dashboard > 系统管理 > 全局安全配置 ），找到Git Host Key Verification Configuration选择Accept first connection

**配置Git名称**（Dashboard > 系统管理 > 全局工具配置 ），找到Git 配置Path to Git executable 和 Name 为git

### 安装K8s

在Jenkins使用Kubeasz安装K8s

建议job-name： 0001-set-testing-k8s-cluster

### 安装NFS-server提供存储

本示例使用NFS作为存储类，也可以替换成其他存储方案。

首先选择一台主机安装nfs-server并配置一个路径提供NFS服务，后续通过Jenkins Job为K8s设置默认存储类。

nfs-server安装示例：

```Shell
apt update
apt install nfs-kernel-server -y

# 本例子/scratch为提供存储的目录
echo '/scratch *(rw,async,no_subtree_check,no_root_squash)' >> /etc/exports
systemctl start nfs-kernel-server.service

exportfs -a
```

## 初始化K8s以及安装基础组件

在Jenkins创建对应Job进行环境设置和中间件安装，与其他项目Job设置只是Jenkinsfile文件路径有差异，需要关注。

参数名和值参考给出的表格。

主要目标：

1.初始化K8s
    1.1 安装Helm工具
    1.2 设置默认存储类
2.安装基础组件
    - milvus
    - redis-cluster
    - minio
    - mysql
    - treafik

### 初始化K8s

安装Helm和设置默认存储类的Job参数如下

脚本路径：basement/Jenkinsfile 

| 参数名    | 0002-install-helm-for-jenkins | 0003-set-k8s-nfs-storage |
|-----------|-------------------------------|--------------------------|
| INSTALL   | true                          | true                     |
| UNINSTALL | false                         | false                    |
| TARGET    | helm                          | nfs-storage              |

### 安装初始化组件

安装中间件的Job参数如下，其中TARGET参数为"all"时同时安装所有中间件。

也可以分开安装或卸载。指定TARGET参数为特定组件名称即可。

组件名单：
- milvus
- redis-cluster
- minio
- mysql
- treafik

脚本路径：basement/Jenkinsfile 

| 参数名    | 1001-install-all-basement | 1001-uninstall-all-basement |
|-----------|---------------------------|-----------------------------|
| INSTALL   | true                      | false                       |
| UNINSTALL | false                     | true                        |
| TARGET    | all                       | all                         |

## 项目构建&部署任务

项目名清单：
- IDC
  - nft-meta
  - ranker
  - transform
  - gateway
  - block-etl
- AWS
  - transform
  - cloud-proxy
  - entrance
  - webui
  - dashboard

脚本路径：Jenkinsfile 

| 参数名         | build  | tag           | release       | deploy        |
|----------------|--------|---------------|---------------|---------------|
| BRANCH_NAME    | 分支名 | 分支名        | 分支名        | 分支名        |
| BUILD_TARGET   | true   | false         | false         | false         |
| DEPLOY_TARGET  | false  | false         | false         | true          |
| RELEASE_TARGET | false  | false         | true          | true          |
| TAG_PATCH      | false  | true/false    | false         | false         |
| TAG_MINOR      | false  | true/false    | false         | false         |
| TAG_MAJOR      | false  | true/false    | false         | false         |
| AIMPROJECT     | 项目名 | 项目名        | 项目名        | 项目名        |
| TAG_FOR        | none   | dev/test/prod | dev/test/prod | dev/test/prod |
| TARGET_ENV     | 环境名 | 环境名        | 环境名        | 环境名        |

## 配置说明

所有配置都在config/config.toml中，如果想修改有两种途径：

1.修改config/config.toml重新编译打包成Docker镜像

2.通过设置环境变量即可，比如k8s中可设置configMap、export变量

config.toml -> environment 转换规则

例：

```toml
path="/uu/ii"
port=50515
project-name="Web3Eye"

[mysql]
host="mysql"
port=3306
max-connect=100

log-dir="/var/log"
```

```shell
path=/uu/ii
port=50515
project_name=Web3Eye

mysql_host=mysql
mysql_port=3306
mysql_max_connect=100

log_dir=/var/log
```