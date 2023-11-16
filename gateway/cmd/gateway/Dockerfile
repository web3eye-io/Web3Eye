FROM centos:centos7.9.2009

COPY gateway /usr/local/bin/gateway
COPY Gateway.viper.yaml /etc/Gateway/Gateway.viper.yaml
ENV ENV_DATA_DIR=/data \
    ENV_LOG_DIR=/var/log \
    ENV_LOG_LEVEL=debug \
    ENV_ETH_ENDPOINT="" \
    ENV_INVERSTOR_KEY=""
RUN mkdir $ENV_DATA_DIR

RUN yum update -y
RUN yum install ca-certificates curl -y

WORKDIR /usr/local/bin/

CMD ./gateway run