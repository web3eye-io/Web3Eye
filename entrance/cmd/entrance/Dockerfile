FROM centos:centos7.9.2009

COPY entrance /usr/local/bin/entrance
COPY Entrance.viper.yaml /etc/Entrance/Entrance.viper.yaml
ENV ENV_DATA_DIR=/data \
    ENV_LOG_DIR=/var/log \
    ENV_LOG_LEVEL=debug \
    ENV_ETH_ENDPOINT="" \
    ENV_INVERSTOR_KEY=""
RUN mkdir $ENV_DATA_DIR

RUN yum update -y
RUN yum install ca-certificates curl -y

WORKDIR /usr/local/bin/

CMD ./entrance run