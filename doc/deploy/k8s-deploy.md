# 安装说明

每个小节请阅读完成再操作，以免理解错误上下文意思，同时欢迎提Issue帮助改进。

部署顺序如下：

![deploy steps](https://raw.githubusercontent.com/web3eye-io/Web3Eye/master/doc/picture/deploy-steps.png)

# 机器规划

测试环境主机规划表

| IP           | hostname | 硬件配置                           | 角色                                     | 位置 | 系统       |
|--------------|----------|--------------------------------|------------------------------------------|------|------------|
| 172.16.29.49 |          | CPU:4核  内存：8G  磁盘：50G         | gateway(for Scientific Internet)、Jenkins | IDC  | centos7    |
| 172.16.29.51 | idcnode1 | CPU:8核  内存：16G  磁盘：200G       | k8s-master                               | IDC  | ubuntu20.4 |
| 172.16.29.52 | idcnode2 | CPU:16核  内存：32G  磁盘：100G      | k8s-worker                               | IDC  | ubuntu20.4 |
| 172.16.29.53 | idcnode3 | CPU:16核  内存：32G  磁盘：100G      | k8s-worker                               | IDC  | ubuntu20.4 |
| 172.16.29.54 | idcnode4 | CPU:16核  内存：32G  磁盘：100G+400G | k8s-worker,nfs-server                    | IDC  | ubuntu20.4 |
| 172.23.10.87 | awsnode1 | CPU:8核  内存：16G  磁盘：100G       | k8s-master,k8s-worker                    | AWS  | ubuntu20.4 |

总共部署两套K8s环境，分别是IDC、AWS环境；IDC负责解析链数据、存储链数据、提供搜索服务；而AWS主要为客户和管理员提供可视化服务。

两套环境都是用同一个Jenkins创建Job进行部署，K8s上的组件和服务也都由Jenkins上的Job部署。

可依据高可用需求灵活扩展K8s集群规模

| 集群角色 | TARGET_ENV                                      |
|------|-------------------------------------------------|
| 开发环境 | web3eye-development-idc、web3eye-development-aws |
| 测试环境 | web3eye-testing-idc、web3eye-testing-aws         |
| 生产环境 | web3eye-production-idc、web3eye-production-aws   |

以下的TARGET_ENV使用web3eye-testing-*为例子

# 安装及初始化系统

按照配置正常安装系统即可，若是在虚拟机上安装可考虑用克隆的方式提高安装速度。

**K8s集群内的所有机器**

目标：
- 设置root密码
- 开启root的ssh登录
- 固定IP

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

# 设置hostname
hostnamectl set-hostname <newhostname>
hostname <new-hostname>
```

## 安装V2rayA（可选）

Gateway机器主要为IDC提供统一的网络控制，主要为了更好的科学上网，此处也可选其他方式实现，如果没有科学上网的需求可不要gateway节点。

选择V2rayA主要考虑代理能力强，模式多且设置便捷，而且可用Web操作，在无图形化服务器非常好用。

安装方法：https://v2raya.org/docs/prologue/installation/redhat/

```Shell
curl -Ls https://mirrors.v2raya.org/go.sh | bash
systemctl disable v2ray --now ### Xray 需要替换服务为 xray

yum install wget -y

wget https://github.com/v2rayA/v2rayA/releases/download/v2.2.4/installer_redhat_x64_2.2.4.rpm -O /tmp/v2raya.rpm
rpm -i /tmp/v2raya.rpm

systemctl start v2raya.service
systemctl enable v2raya.service
```

安装完成后导入代理节点即可使用，同时将其他机器的网关设置成Gateway机器的IP，其他机器也能科学上网。

# 安装Jenkins

安装Jenkins此处安装在gateway角色上，一般来说只要不安装在K8s集群的机器上都可以

目标：
- 关闭Selinux和防火墙
- 安装Docker
- 启动Jenkins
- 初始化Jenkins
- 配置Git插件

## 关闭Selinux和防火墙

因为gateway服务器是Centos系统，所以需要设置SeLinux和防火墙。

```Shell
# 临时关闭Selinux
setenforce 0
# 永久关闭Selinux
vi /etc/selinux/config
## 将SELINUX=enforcing改为SELINUX=disabled


# 关闭防火墙
systemctl stop firewalld.service
systemctl disable firewalld.service
```

## 安装Docker

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

systemctl start docker
systemctl enable docker
```

## 启动Jenkins

```Shell
docker run \
  --name jenkins \
  -d \
  -p 18080:8080 \
  -p 50000:50000 \
  -v /opt/share/jenkins:/var/lib/jenkins   \
  -v /sys/fs/cgroup:/sys/fs/cgroup:ro      \
  --tmpfs /tmp:exec --tmpfs /run --tmpfs /run/lock --tmpfs /var/run \
  -v /var/run/docker.sock:/var/run/docker.sock \
  --restart always \
  --privileged coastlinesss/jenkins:0.0.1 /usr/sbin/init
```

## 初始化Jenkins

获取jenkins初始密码

```shell
docker exec -it jenkins cat /var/lib/jenkins/secrets/initialAdminPassword
```

访问jenkins web页面(172.16.29.49:18080)，完成Jenkins初始配置，如添加用户等，在安装插件时可先安装建议插件。

### 配置Git插件

**配置Git** 接受第一次连接（Dashboard > 系统管理 > 全局安全配置 ），找到Git Host Key Verification Configuration选择Accept first connection

**配置Git名称**（Dashboard > 系统管理 > 全局工具配置 ），找到Git 配置Path to Git executable 和 Name 为git

# 提供NFS-SERVER服务

本示例使用NFS作为存储类，也可以替换成其他存储方案。

首先选择一台主机（例子中提供存储的是idcnode4）安装nfs-server并配置一个路径提供NFS服务，后续通过Jenkins Job为K8s设置默认存储类。

nfs-server安装示例：

```Shell
apt update
apt install nfs-kernel-server -y

# 本例子/k8sdata为提供存储的目录
echo '/k8sdata *(rw,async,no_subtree_check,no_root_squash)' >> /etc/exports
systemctl start nfs-kernel-server.service

exportfs -a
```
# Jenkins任务说明

试图及任务说明

| 视图                                                | 编号格式                                        | 任务说明     |
|-----------------------------------------------------|-------------------------------------------------|-------------|
| 00-kubernets                                        | (development\testing\production)-000N-*         | 安装K8s环境  |
| 01-basement                                         | (development\testing\production)-100N-*         | 安装基础组件 |
| 02-build                                            | (development\testing\production)-200N-*         | 构建项目     |
| 03-tag-(testing\production)                         | (testing\production)-300N-*                     | 打tag        |
| 04-release-(feature\development\testing\production) | (feature\development\testing\production)-400N-* | release项目  |
| 05-deploy-(feature\development\testing\production)  | (feature\development\testing\production)-500N-* | 部署项目     |
| 05-domain-(feature\development\testing\production)  | (feature\development\testing\production)-600N-* | 网站上线相关     |


# 安装K8s

## 安装K8s

在Jenkins使用Kubeasz安装K8s，按照主机规划在IDC和AWS分别部署K8S环境。

按照主机规划修改
- hosts
  - 主机规划，比如etcd、kube_master、kube_node
  - [all:vars]下的用户名及密码
  - cluster_dir
- config.yml
  - nfs-provisioner下的配置，主要关注nfs_server及nfs_path

建议job-name： testing-0001-IDC-k8s-cluster、testing-0002-AWS-k8s-cluster

需要在IDC以及AWS环境中各安装一套K8S
## 安装Helm工具

安装Helm和设置默认存储类的Job参数如下

脚本路径：basement/Jenkinsfile 

| 参数名    | testing-0002-install-helm-for-jenkins |
|-----------|---------------------------------------|
| INSTALL   | true                                  |
| UNINSTALL | false                                 |
| TARGET    | helm                                  |

# 安装基础组件

在Jenkins创建对应Job进行环境设置和中间件安装，与其他项目Job设置只是Jenkinsfile文件路径有差异，需要关注。

参数名和值参考给出的表格。

主要目标：
- IDC基础组件
  - milvus
  - redis-cluster
  - minio
  - mysql
- AWS基础组件
  - treafik

## IDC基础组件

安装中间件的Job参数如下，其中TARGET参数为"all"时同时安装所有中间件。

也可以分开安装或卸载。指定TARGET参数为特定组件名称即可。

组件名单：
- milvus
- redis-cluster
- minio
- mysql

脚本路径：basement/Jenkinsfile 

| 参数名    | testing-1002-install-all-basement | testing-1002-uninstall-all-basement |
|-----------|-----------------------------------|-------------------------------------|
| INSTALL   | true                              | false                               |
| UNINSTALL | false                             | true                                |
| TARGET    | all                               | all                                 |
| TARGET_ENV    | web3eye-testing-idc                          |


## AWS基础组件

组件名单：
- traefik

脚本路径：basement/Jenkinsfile 

| 参数名    | testing-1003-install-traefik | testing-1003-uninstall-traefik |
|-----------|------------------------------|--------------------------------|
| INSTALL   | true                         | false                          |
| UNINSTALL | false                        | true                           |
| TARGET    | traefik                      | traefik                        |
| TARGET_ENV    | web3eye-testing-aws                          |

# 项目构建&部署任务

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

配置部署Job时建议按照IDC、AWS分类

脚本路径：Jenkinsfile 

| 参数名         | build  | tag        | release | deploy |
|----------------|--------|------------|---------|--------|
| BRANCH_NAME    | 分支名 | 分支名     | 分支名  | 分支名 |
| BUILD_TARGET   | true   | false      | false   | false  |
| DEPLOY_TARGET  | false  | false      | false   | true   |
| RELEASE_TARGET | false  | false      | true    | true   |
| TAG_PATCH      | false  | true/false | false   | false  |
| TAG_MINOR      | false  | true/false | false   | false  |
| TAG_MAJOR      | false  | true/false | false   | false  |
| AIMPROJECT     | 项目名 | 项目名     | 项目名  | 项目名 |
| TAG_FOR        |    | test/prod  |     |    |
| TARGET_ENV     | 环境名 | 环境名     | 环境名  | 环境名 |

release和deploy的Tag关系说明：

|             | feature  | development | testing              | production           |
|-------------|----------|-------------|----------------------|----------------------|
| BRANCH_NAME | branch   | master      | master               | master               |
| TARGET_ENV  | 任意环境 | development | testing              | production           |
| 最终Tag名   | branch名 | latest      | 奇数版本号(如:0.5.3) | 偶数版本号(如:0.5.2) |

# 网站上线

部署好所有项目后，要上线网站还需要配置TLS域名解析。

# 开始使用

打开Dashboard（AWS的32443端口访问），添加Eth和Sol的Endpoints，再添加同步任务即可。

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