FROM debian:stable-slim

COPY block-etl /usr/local/bin/block-etl
COPY BlockETL.viper.yaml /etc/BlockETL/BlockETL.viper.yaml
ENV ENV_DATA_DIR=/data \
    ENV_LOG_DIR=/var/log \
    ENV_LOG_LEVEL=debug \
    ENV_ETH_ENDPOINT="" \
    ENV_INVERSTOR_KEY=""
RUN mkdir $ENV_DATA_DIR
RUN apt update -y
RUN apt install -y --no-install-recommends ca-certificates curl

WORKDIR /usr/local/bin/

CMD ./block-etl run