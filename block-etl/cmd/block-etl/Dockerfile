FROM jrei/systemd-centos:7

COPY block-etl /usr/local/bin/block-etl
COPY BlockETL.viper.yaml /etc/BlockETL/BlockETL.viper.yaml
ENV ENV_DATA_DIR=/data \
    ENV_LOG_DIR=/var/log \
    ENV_LOG_LEVEL=debug \
    ENV_ETH_ENDPOINT="" \
    ENV_INVERSTOR_KEY=""
RUN mkdir $ENV_DATA_DIR

RUN yum update -y
RUN yum install ca-certificates curl -y

WORKDIR /usr/local/bin/

CMD ./block-etl run