FROM jrei/systemd-ubuntu:20.04

USER root

RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list

RUN apt-get update -y
RUN apt-get install git make curl jq wget apt-utils vim net-tools sudo iputils-ping iproute2 openssh-server nfs-common dnsutils -y

RUN sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/g' /etc/ssh/sshd_config
RUN sed -i 's/#PubkeyAuthentication yes/PubkeyAuthentication yes/g' /etc/ssh/sshd_config
RUN sed -i 's/$AuthorizedKeysFile/AuthorizedKeysFile/g' /etc/ssh/sshd_config
RUN sed -i 's/UsePAM yes/UsePam no/g' /etc/ssh/sshd_config
RUN sed -i 's/#Port 22/Port 22/g' /etc/ssh/sshd_config
RUN echo root:qq | chpasswd
RUN service ssh restart

ARG ALL_PROXY
RUN apt-get install golang-go -y

RUN all_proxy=$ALL_PROXY curl -sL -o /tmp/go1.17.8.tar.gz https://github.com/golang/go/archive/refs/tags/go1.17.8.tar.gz
RUN mkdir -p /usr/local/go
RUN cd /usr/local/go; tar xvvf /tmp/go1.17.8.tar.gz
RUN mv /usr/local/go/go-go1.17.8/* /usr/local/go
RUN rm -rf /usr/local/go/go-go1.17.8
RUN cd /usr/local/go/src; ./all.bash
RUN cp /usr/local/go/bin/go /usr/bin/go


ENV GOBIN=/usr/bin
ENV GOROOT=/usr/local/go
ENV GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64

CMD ["/bin/bash"]
