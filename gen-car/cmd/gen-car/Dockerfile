FROM debian:stable-slim

COPY gen-car /usr/local/bin/gen-car
COPY GenCar.viper.yaml /etc/GenCar/GenCar.viper.yaml
ENV ENV_DATA_DIR=/data \
    ENV_LOG_DIR=/var/log \
    ENV_LOG_LEVEL=debug \
    ENV_ETH_ENDPOINT="" \
    ENV_INVERSTOR_KEY=""
RUN mkdir $ENV_DATA_DIR

RUN apt-get check
RUN apt-get update -y
RUN apt-get install -y libhwloc-dev
RUN ln -s /usr/lib/x86_64-linux-gnu/libhwloc.so /usr/lib/x86_64-linux-gnu/libhwloc.so.5
RUN apt-get install -y --no-install-recommends ca-certificates curl

WORKDIR /usr/local/bin/

CMD ./gen-car run