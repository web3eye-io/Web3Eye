FROM jrei/systemd-centos:7

USER root

ARG ALL_PROXY

# RUN curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo

RUN all_proxy= yum update -y
RUN yum install deltarpm -y
RUN yum install gcc -y
RUN yum install vim -y
RUN yum install curl -y
RUN yum install gnupg2 -y
RUN yum install iptables -y
RUN yum install epel-release -y
RUN yum install fontconfig java-11-openjdk-devel java-11-openjdk wget -y
RUN yum install initscripts -y
RUN yum install gettext -y
RUN yum install unzip -y
RUN yum install gcc-c++ -y
RUN yum install make -y
RUN yum install nodejs -y
RUN yum install ansible -y
RUN yum install git -y

RUN sed -i 's/#host_key_checking = False/host_key_checking = False/g' /etc/ansible/ansible.cfg

RUN wget -O /etc/yum.repos.d/jenkins.repo https://pkg.jenkins.io/redhat-stable/jenkins.repo
RUN rpm --import https://pkg.jenkins.io/redhat-stable/jenkins.io-2023.key
RUN yum install jenkins -y

RUN mv /etc/rc.d/init.d/jenkins /usr/local/bin/jenkins-inner
RUN echo "#!/bin/sh" > /etc/rc.d/init.d/jenkins
RUN echo >> /etc/rc.d/init.d/jenkins
RUN echo "# chkconfig: - 85 15" >> /etc/rc.d/init.d/jenkins
RUN echo "# description: Jenkins CI server" >> /etc/rc.d/init.d/jenkins
RUN echo >> /etc/rc.d/init.d/jenkins
RUN echo "export all_proxy=$ALL_PROXY" >> /etc/rc.d/init.d/jenkins
RUN echo >> /etc/rc.d/init.d/jenkins
RUN echo '. /usr/local/bin/jenkins-inner' >> /etc/rc.d/init.d/jenkins

RUN chmod a+x /etc/rc.d/init.d/jenkins
RUN chmod a+x /etc/rc.d/init.d/functions
RUN systemctl enable jenkins

RUN sed -i 's/User=jenkins/User=root/g' /usr/lib/systemd/system/jenkins.service
RUN sed -i 's/Group=jenkins/Group=root/g' /usr/lib/systemd/system/jenkins.service