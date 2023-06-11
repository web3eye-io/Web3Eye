FROM debian:stable-slim

COPY gen-car /usr/local/bin/gen-car
COPY GenCar.viper.yaml /etc/GenCar/GenCar.viper.yaml
ENV ENV_DATA_DIR=/data \
    ENV_LOG_DIR=/var/log \
    ENV_LOG_LEVEL=debug \
    ENV_ETH_ENDPOINT="" \
    ENV_INVERSTOR_KEY=""
RUN mkdir $ENV_DATA_DIR

RUN sed -i 's#http://deb.debian.org#http://mirrors.163.com#g' /etc/apt/sources.list

RUN apt update -y
RUN apt install -y --no-install-recommends ca-certificates curl

WORKDIR /usr/local/bin/

CMD ./gen-car run