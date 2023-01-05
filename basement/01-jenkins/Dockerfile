FROM jenkins/jenkins:centos7

USER root


RUN curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo

RUN all_proxy= yum update -y
RUN yum install deltarpm -y
RUN yum install gcc -y
RUN yum install vim -y
RUN yum install curl -y
RUN yum install gnupg2 -y
RUN yum install iptables -y
RUN yum install epel-release -y
RUN yum install java-11-openjdk-devel -y
RUN yum install initscripts -y
RUN yum install gettext -y
RUN yum install unzip -y
RUN yum install gcc-c++ -y
RUN yum install make -y
RUN yum install nodejs -y
RUN yum install ansible -y
RUN yum install git -y

RUN curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash

RUN sed -i 's/#host_key_checking = False/host_key_checking = False/g' /etc/ansible/ansible.cfg

