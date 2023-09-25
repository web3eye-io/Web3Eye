## 安装说明

每个小节请阅读完成再操作，以免理解错误上下文意思，同时欢迎提Issue帮助改进。

## 机器准备

主机规划表

| IP           | hostname | 硬件配置                           | 角色                             | 位置 |
|--------------|----------|--------------------------------|----------------------------------|------|
| 172.23.10.29 |          | CPU:4核  内存：8G  磁盘：50G         | gateway(for Scientific Internet) | IDC  |
| 172.23.10.31 | node1    | CPU:8核  内存：16G  磁盘：200G       | k8s-master,jenkins               | IDC  |
| 172.23.10.32 | node2    | CPU:16核  内存：32G  磁盘：100G      | k8s-worker                       | IDC  |
| 172.23.10.33 | node3    | CPU:16核  内存：32G  磁盘：100G      | k8s-worker                       | IDC  |
| 172.23.10.34 | node4    | CPU:16核  内存：32G  磁盘：100G+400G | k8s-worker,nfs-server            | IDC  |
| 172.16.10.85 | node1    | CPU:8核  内存：16G  磁盘：100G       | k8s-master,k8s-worker            | AWS  |

系统：Ubuntu20.04 or Ubuntu22.04

### 安装V2rayA

Gateway机器主要为IDC提供统一的网络控制，此处也可选其他方式实现，主要为了更好的科学上网，如果没有科学上网的需求可不要gateway节点。

选择V2rayA主要考虑代理能力强，模式多且设置便捷，而且可用Web操作，在无图形化服务器非常好用。

安装方法：https://v2raya.org/docs/prologue/installation/debian/

建议直接下载deb文件进行安装。

安装完成后导入代理节点即可使用，同时将其他机器的网关设置成Gateway机器的IP，其他机器也能科学上网了。

## 初始化系统配置

按照配置正常安装系统即可，若是在虚拟机上安装可考虑用克隆的方式提高安装速度。

设置root密码、开启root的ssh登录、设置gateway（可选）

```Shell
# 切换到root用户
su root
# 设置密码
passwd

PermitRootLogin yes #允许root登录
PasswordAuthentication yes #允许密码登录
sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/g' /etc/ssh/sshd_config
sed -i 's/#PasswordAuthentication yes/PasswordAuthentication yes/g' /etc/ssh/sshd_config
systemctl restart sshd

# 设置gateway
vim /etc/netplan/00-installer-config.yaml
## 修改gateway配置成安装了V2rayA的机器IP
netplan apply
```

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
        ssh "${arr[$index]}" "for pkg in docker.io docker-doc docker-compose podman-docker containerd runc; do sudo apt-get remove \$pkg; done"
        echo "end clear old-docker" 
    fi

    index1=`expr $index + 1`
    ssh "${arr[$index]}" "hostnamectl set-hostname \"\${hosts[\$index1]}\""
    ssh "${arr[$index]}" "apt install python3 -y"
done  

```

### 安装K8s

安装K8s集群(版本为1.24)，可选择kubeasz快速安装(项目链接:<https://github.com/easzlab/kubeasz>)。另外集群中主机名不能重复，否则k8s网络可能会出现问题。

完整的安装指导如下链接，也可参考浓缩版安装指导

安装方法：https://github.com/easzlab/kubeasz/blob/master/docs/setup/00-planning_and_overall_intro.md

#### 浓缩版安装指导

提供快速安装步骤，如有问题请参考官方文档

```Shell
# 下载指定版本工具
export release=3.5.0
wget https://github.com/easzlab/kubeasz/releases/download/${release}/ezdown
chmod +x ./ezdown
./ezdown -D -m standard
./ezdown -S

# 在docker中启动工具
docker exec -it kubeasz ezctl new k8s-01

# 需要进行配置，主机配置示例在下方
vim /etc/kubeasz/clusters/k8s-01/hosts 

# 配置环境变量
echo "alias dk='docker exec -it kubeasz'" >> /etc/profile
source /etc/profile

# 安装
dk ezctl setup k8s-01 all
```

主机配置示例（配置时请在官方提供的配置文件内编辑，此处之给出主机规划部分的配置）：

```ini
# 'etcd' cluster should have odd member(s) (1,3,5,...)
[etcd]
172.16.29.51

# master node(s)
[kube_master]
172.16.29.51

# work node(s)
[kube_node]
172.16.29.52
172.16.29.53
172.16.29.54
```

安装完成后把/etc/kubeasz/bin添加到PATH环境变量。
```Shell
echo "export PATH=\$PATH:/etc/kubeasz/bin" >> /etc/profile
source /etc/profile

kubectl get node -A
```

还需要在Master节点安装Helm(安装介绍<https://helm.sh/docs/intro/install/>)。

### 配置nfs为默认存储类

本示例使用NFS作为存储类，也可以替换成其他存储方案。

在k8s集群中的一台服务器上安装nfs客户端（可以选择k8s的master节点，主要是为了找一台方便的节点完成主要的安装操作）。

首先选择一台主机安装nfs-server并配置一个路径提供NFS服务。

在k8s集群的master机器上把web3eye-io/Web3Eye项目clone到服务器并配置NFS。

```shell
git clone https://github.com/web3eye-io/Web3Eye.git
cd Web3Eye/basement
cat 02-nfs-storage/value.yaml
```

主要关注server和path，修改成NFS服务的地址和路径即可

```yaml
nfs:
  server: 172.23.10.83
  path: /data/k8s_storage
```

确认好配置后执行install.sh

```shell
bash install.sh
```

脚本运行完毕后须看到以下输出结果：

```shell
root@k8s-master:~/Web3Eye/basement/02-nfs-storage# kubectl get pods  -o wide  | egrep nfs 
default-nfs-provisioner-nfs-subdir-external-provisioner-7fx2pwz   1/1     Running     6 (5h21m ago)   19h     172.20.144.99    172.23.10.33   <none>           <none>
```

## 安装Jenkins及配置Jenkins环境

### 使用docker直接起一个jenkins

需要按照实际情况配置端口映射关系以及文件映射关系，这里需要明确好docker.sock和.kube配置的路径。

这里需要注意.kube中的kube-apiserver要指向的docker能访问的IP不能指向127.0.0.1

示例如下：

```shell
root@k8s-master:~/.kube# cat config
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURsRENDQW55Z0F3SUJBZ0lVRXFXVjZHQ2JkWVZHMWhsVmRZYkhyMTRpcnNrd0RRWUpLb1pJaHZjTkFRRUwKQlFBd1lURUxNQWtHQTFVRUJoTUNRMDR4RVRBUEJnTlZCQWdUQ0VoaGJtZGFhRzkxTVFzd0NRWURWUVFIRXdKWQpVekVNTUFvR0ExVUVDaE1EYXpoek1ROHdEUVlEVlFRTEV3WlRlWE4wWlcweEV6QVJCZ05WQkFNVENtdDFZbVZ5CmJtVjBaWE13SUJjTk1qTXdNVEV3TVRBek9EQXdXaGdQTWpFeU1qRXlNVGN4TURNNE1EQmFNR0V4Q3pBSkJnTlYKQkFZVEFrTk9NUkV3RHdZRFZRUUlFd2hJWVc1bldtaHZkVEVMTUFrR0ExVUVCeE1DV0ZNeEREQUtCZ05WQkFvVApBMnM0Y3pFUE1BMEdBMVVFQ3hNR1UzbHpkR1Z0TVJNd0VRWURWUVFERXdwcmRXSmxjbTVsZEdWek1JSUJJakFOCkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQTA1MHFDN0hYZzZGQWFKN3Z1ZUJvQlRENEN5aU8KSjJiYitUZ2V3NDN6ekNtRmdSZ1ZUdmxsejJjd0cxWGNCdnBTMzlsdEd3TGRUSEhGNHBuNFhVOTlnKzJGaEhqbgpMZDE5UzZXZFBMRk5PLy9maU0vNGx2enYzN21zUFhhQVNZTHRidjQwV0xuSmYwTDhzeUxSV1VkMkp3Nm1VMTVWCjVEeEU1RG9WYS9Ib3lGSlhnc0xrc0hJQmt0T1QxM0FUZ09nL1V3STdGcGt4aXFhelYzUjBUNU5WcXEzaW1JS20KYlFaYUNVdFdtRHVXTk5uOGJ2Vno0U21vb0N4ejUybDVUdEZXN0E4SzRKK0JXRDhYcFV1dkIvRy9IZVhESFpQKwpIeVJSTDZNemVLYVAwaVNSOVlIbUwyZklZc0hrb0ozcHhJWXFjcStvR2pGUzVaZUZRa01kVS9EU2xRSURBUUFCCm8wSXdRREFPQmdOVkhROEJBZjhFQkFNQ0FRWXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QWRCZ05WSFE0RUZnUVUKaU9yeUUzWXlqdFhFY3JYNDZnS0V2OENYVTZRd0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQkFGaDBlb0N4d2VaTQpzang2b0FXR280Z0JNSEpxbW43YmUybENmc21CMjJBYndsRWVpWjJRWkJ4WWRERHJGTXNHc1Z3RjVxNjZiT3QxCm5xWjBTVS9CY0dISGFnTU9vWmwxRzVyRVJFempoeUJpNWsrajZ2NG5LTWxnSXdoaWw0Qzg5UUhyM0I3QkhYVEMKQTd1dCt1YVYxenVFVitsZGEvYVJGTllOVGZ0d0dSdDFQeWMvL2dVVm50QzlabU9ISHlPU280a1ZiMFNKVk5mZgpNTEdBWGJGdEp5ekc0OUtjVDBQWFVvNjZhTTlCMHZZYWZzTWs1OTR5OGhsRmJiUnJ5elpFTENNZ
    server: https://172.23.10.31:6443
  name: cluster1
```

运行jenkins容器

```shell
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
  coastlinesss/jenkins
```

### 初始化Jenkins

```shell
docker exec -it jenkins cat /var/jenkins_home/secrets/initialAdminPassword
```

访问jenkins web页面(172.23.10.31:18080)，完成Jenkins初始配置，如添加用户等，在安装插件时可先安装建议插件。

## 安装依赖组件

### 配置Jenkins环境

**安装Git插件**（Dashboard > 系统管理 > 插件管理 > Available plugins > 搜索Git并安装）

**配置Git** 接受第一次连接（Dashboard > 系统管理 > 全局安全配置 ），找到Git Host Key Verification Configuration选择Accept first connection

**配置Git名称**（Dashboard > 系统管理 > 全局工具配置 ），找到Git 配置Path to Git executable 和 Name 为git

**安装Go插件**（Dashboard > 系统管理 > 插件管理 > Available plugins > 搜索Go并安装）

**配置Go插件**（Dashboard > 系统管理 > 全局工具配置 > 找到Go）,设置别名为go, 安装一个Go 1.19

### 新建安装任务

在Jenkins中新建安装组件的任务（即job）

创建item: install_components

选择Pipeline类型

**勾选GitHub项目：**  
    项目URL：<https://github.com/web3eye-io/Web3Eye.git/>  

**勾选This project is parameterized：**  
    增加三个字符参数分别为：  
        名称：INSTALL 默认值：true  
        名称：UNINSTALL 默认值：false  
        名称：TARGET 默认值：all  

流水线中选择Pileline script from SCM  
    SCM:Git  
        Repositories:  
            Repository URL: <https://github.com/web3eye-io/Web3Eye.git/>  
            Credentials: 配置一个Git的凭证，可选择SSH Username with private key或Username with password  
        Branches to build:  
            指定分支：*/master  
    脚本路径：basement/Jenkinsfile  

选择保存

新建视图basement，选择列表视图(install_components)。

### 执行安装任务

Dashboard > basement > install_components

选择**Build with Parameters**,点击 **开始构建**

观察构建过程，全部完成后组件就安装成功了

登录服务器查看安装成功的组件

```shell
root@k8s-master:/home/test# kubectl get pod
NAME                                                              READY   STATUS      RESTARTS        AGE
default-nfs-provisioner-nfs-subdir-external-provisioner-57czz2m   1/1     Running     0               4d17h
development-box-0                                                 1/1     Running     0               17h
kafka-0                                                           1/1     Running     0               3d22h
kafka-1                                                           1/1     Running     0               3d22h
kafka-2                                                           0/1     Pending     0               3d22h
kafka-zookeeper-0                                                 1/1     Running     0               3d22h
kafka-zookeeper-1                                                 1/1     Running     0               3d22h
kafka-zookeeper-2                                                 1/1     Running     0               3d22h
milvus-datacoord-5f7b497444-28k8m                                 1/1     Running     2 (3d18h ago)   3d23h
milvus-datanode-684c8d4986-gnpzs                                  1/1     Running     2 (3d18h ago)   3d23h
milvus-etcd-0                                                     1/1     Running     1 (3d18h ago)   3d23h
milvus-etcd-1                                                     1/1     Running     3 (3d18h ago)   3d23h
milvus-etcd-2                                                     1/1     Running     1 (3d18h ago)   3d23h
milvus-indexcoord-7df986464d-sqlzq                                1/1     Running     2 (3d18h ago)   3d23h
milvus-indexnode-6b6c7f7797-mpxss                                 1/1     Running     1 (3d18h ago)   3d23h
milvus-minio-0                                                    1/1     Running     0               3d23h
milvus-minio-1                                                    1/1     Running     0               3d23h
milvus-minio-2                                                    1/1     Running     0               3d23h
milvus-minio-3                                                    1/1     Running     0               3d23h
milvus-proxy-645fbb45f4-ntw94                                     1/1     Running     2 (3d18h ago)   3d23h
milvus-pulsar-bookie-0                                            1/1     Running     0               3d23h
milvus-pulsar-bookie-1                                            1/1     Running     0               3d23h
milvus-pulsar-bookie-2                                            1/1     Running     0               3d23h
milvus-pulsar-bookie-init-jsjbf                                   0/1     Completed   0               3d23h
milvus-pulsar-broker-0                                            1/1     Running     1 (3d18h ago)   3d23h
milvus-pulsar-proxy-0                                             1/1     Running     0               3d23h
milvus-pulsar-pulsar-init-v9bfg                                   0/1     Completed   0               3d23h
milvus-pulsar-recovery-0                                          1/1     Running     0               3d23h
milvus-pulsar-zookeeper-0                                         1/1     Running     0               3d23h
milvus-pulsar-zookeeper-1                                         1/1     Running     0               3d23h
milvus-pulsar-zookeeper-2                                         1/1     Running     0               3d23h
milvus-querycoord-6778454959-v6zks                                1/1     Running     2 (3d18h ago)   3d23h
milvus-querynode-569c9db6ff-w2968                                 1/1     Running     1 (3d18h ago)   3d23h
milvus-rootcoord-57c9dbfcd9-whttz                                 1/1     Running     2 (3d18h ago)   3d23h
mysql-0                                                           1/1     Running     0               3d23h
redis-cluster-0                                                   1/1     Running     1 (12h ago)     12h
redis-cluster-1                                                   1/1     Running     1 (12h ago)     12h
redis-cluster-2                                                   1/1     Running     1 (12h ago)     12h
redis-cluster-3                                                   1/1     Running     1 (12h ago)     12h
redis-cluster-4                                                   1/1     Running     1 (12h ago)     12h
redis-cluster-5                                                   1/1     Running     1 (12h ago)     12h
traefik-4f9vc                                                     1/1     Running     0               3d23h
traefik-9fxc4                                                     1/1     Running     0               3d23h
traefik-9lxvl                                                     1/1     Running     0               3d23h
whoami-58b8d4f6f6-cklq5                                           1/1     Running     0               3d23h
whoami-58b8d4f6f6-sh2cc                                           1/1     Running     0               3d23h
```

## 部署项目

### 创建视图和任务

新建deploy-dev视图，新建部署项目的任务

参考 安装依赖组件 中的新建任务(可直接克隆)，除了参数化构建过程中的参数不一样以及最后一步SCM中的脚本路径为Jenkinsfile外，其他配置都一致。

参数化构建过程中的Jenkinsfile任务参数矩阵，选择[项目构建&部署任务](#001)中的d-dev取值，根据AIMPROJECT的三个取值创建成三个不同的部署任务。

![部署任务视图](doc/picture/jenkins-deploy-dev.jpg)

### 部署项目

依次参数化构建，建议部署顺序：nft-meta、block-etl、converter

构建完成后访问k8s-master-IP:80/api/nft-meta/可访问项目测试页面

## CICD

### Jenkins任务参数矩阵

#### 安装/卸载组件任务

| 参数名    | install | uninstall |
|-----------|---------|-----------|
| INSTALL   | true    | false     |
| UNINSTALL | false   | true      |
| TARGET    | all     | all       |

TARGET可选值：all、traefik、milvus、redis-cluster、kafka、mysql

#### 项目构建&部署任务

<p id="001">
表头中 b-代表build、r-代表release、d-代表deploy
</p>

| 参数名         | build  | tag        | r-dev  | r-test/prod | d-dev  | d-test/prod |
|----------------|--------|------------|--------|-------------|--------|-------------|
| BRANCH_NAME    | 分支名 | 分支名     | 分支名 | none        | none   | none        |
| BUILD_TARGET   | true   | false      | false  | false       | false  | false       |
| DEPLOY_TARGET  | false  | false      | false  | false       | true   | true        |
| RELEASE_TARGET | false  | false      | true   | true        | false  | false       |
| TAG_PATCH      | false  | true/false | false  | false       | false  | false       |
| TAG_MINOR      | false  | true/false | false  | false       | false  | false       |
| TAG_MAJOR      | false  | true/false | false  | false       | false  | false       |
| AIMPROJECT     | 项目名 | none       | 项目名 | 项目名      | 项目名 | 项目名      |
| TAG_FOR        | none   | test/prod  | dev    | test/prod   | none   | none        |
| TARGET_ENV     | none   | none       | none   | none        | dev    | test/prod   |

参数说明：

AIMPROJECT指定的项目名根据项目选择：nft-meta、block-etl、converter、webui

BRANCH_NAME指定的分支名默认为master，除了prod之外其他可按需指定分支名称

## 配置

所有配置都在config/config.toml中，如果想修改有两种途径：  
1.修改config/config.toml重新编译打包成Docker镜像  
2.通过设置环境变量即可，在k8s中可设置configMap  

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