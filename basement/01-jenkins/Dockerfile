FROM jenkins/jenkins:centos7

USER root


# RUN curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo

RUN all_proxy= yum update -y
RUN yum install deltarpm gcc vim curl gnupg2 iptables epel-release java-11-openjdk-devel -y
RUN yum install initscripts gettext unzip gcc-c++ make nodejs npm ansible git wget  -y

RUN yum remove docker \
            docker-client \
            docker-client-latest \
            docker-common \
            docker-latest \
            docker-latest-logrotate \
            docker-logrotate \
            docker-engine

# for build gen-car ,because import something from filecoin-project
RUN yum install -y gcc hwloc* ocl-icd*

RUN yum install -y yum-utils
RUN yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo
RUN yum install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

RUN curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
RUN curl -L "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" -o /usr/bin/kubectl
RUN chmod +x /usr/bin/kubectl

RUN sed -i 's/#host_key_checking = False/host_key_checking = False/g' /etc/ansible/ansible.cfg

