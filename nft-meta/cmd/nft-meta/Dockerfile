FROM centos:centos7.9.2009

COPY nft-meta /usr/local/bin/nft-meta
COPY NFTMeta.viper.yaml /etc/NFTMeta/NFTMeta.viper.yaml
ENV ENV_DATA_DIR=/data \
    ENV_LOG_DIR=/var/log \
    ENV_LOG_LEVEL=debug \
    ENV_ETH_ENDPOINT="" \
    ENV_INVERSTOR_KEY=""
RUN mkdir $ENV_DATA_DIR

RUN yum install ca-certificates curl -y

WORKDIR /usr/local/bin/

CMD ./nft-meta run