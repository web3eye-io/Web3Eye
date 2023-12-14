FROM centos:centos7.9.2009

COPY gen-car /usr/local/bin/gen-car
COPY GenCar.viper.yaml /etc/GenCar/GenCar.viper.yaml
ENV ENV_DATA_DIR=/data \
    ENV_LOG_DIR=/var/log \
    ENV_LOG_LEVEL=debug \
    ENV_ETH_ENDPOINT="" \
    ENV_INVERSTOR_KEY=""
RUN mkdir $ENV_DATA_DIR

RUN yum install -y hwloc-devel
RUN ln -s /usr/lib/x86_64-linux-gnu/libhwloc.so /usr/lib/x86_64-linux-gnu/libhwloc.so.5
RUN yum install ca-certificates curl -y

WORKDIR /usr/local/bin/

CMD ./gen-car run