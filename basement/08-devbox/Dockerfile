FROM jrei/systemd-ubuntu:20.04

USER root

# RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list

RUN apt-get update -y
RUN apt-get install git make curl jq wget apt-utils vim net-tools sudo iputils-ping iproute2 openssh-server dnsutils -y

RUN sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/g' /etc/ssh/sshd_config
RUN sed -i 's/#PubkeyAuthentication yes/PubkeyAuthentication yes/g' /etc/ssh/sshd_config
RUN sed -i 's/$AuthorizedKeysFile/AuthorizedKeysFile/g' /etc/ssh/sshd_config
RUN sed -i 's/UsePAM yes/UsePam no/g' /etc/ssh/sshd_config
RUN sed -i 's/#Port 22/Port 22/g' /etc/ssh/sshd_config
RUN echo root:qq | chpasswd
RUN service ssh restart

ARG ALL_PROXY

RUN all_proxy=$ALL_PROXY curl -sL -o /tmp/go1.19.9.linux-amd64.tar.gz https://dl.google.com/go/go1.19.9.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf /tmp/go1.19.9.linux-amd64.tar.gz
# RUN mv /usr/local/go1.19.9 /usr/local/go
RUN cp /usr/local/go/bin/go /usr/bin/go
RUN curl -sSL "https://github.com/fullstorydev/grpcurl/releases/download/v1.8.7/grpcurl_1.8.7_linux_x86_64.tar.gz" | sudo tar -xz -C /usr/local/bin


ENV GOBIN=/usr/bin
ENV GOROOT=/usr/local/go
ENV GOTOOLDIR=/usr/lib/go/pkg/tool/linux_amd64

CMD ["/bin/bash"]
