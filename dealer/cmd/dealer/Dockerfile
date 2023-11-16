FROM centos:centos7.9.2009

COPY dealer /usr/local/bin/dealer
COPY lotus /usr/local/bin/lotus
COPY ipfs /usr/local/bin/ipfs
COPY Dealer.viper.yaml /etc/Dealer/Dealer.viper.yaml
ENV ENV_DATA_DIR=/data \
    ENV_LOG_DIR=/var/log \
    ENV_LOG_LEVEL=debug \
    ENV_ETH_ENDPOINT="" \
    ENV_INVERSTOR_KEY=""
RUN mkdir $ENV_DATA_DIR

RUN yum update -y
RUN yum install -y hwloc-devel
RUN ln -s /usr/lib/x86_64-linux-gnu/libhwloc.so /usr/lib/x86_64-linux-gnu/libhwloc.so.5
RUN yum install ca-certificates curl -y

WORKDIR /usr/local/bin/

CMD ./dealer run
