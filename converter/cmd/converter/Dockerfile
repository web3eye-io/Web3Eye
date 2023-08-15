FROM bitnami/python:3.8

COPY converter /usr/local/bin/converter
COPY model /pycode
ENV ENV_DATA_DIR=/data \
    ENV_LOG_DIR=/var/log \
    ENV_LOG_LEVEL=debug \
    ENV_ETH_ENDPOINT="" \
    ENV_INVERSTOR_KEY=""
RUN mkdir $ENV_DATA_DIR

RUN apt update -y
RUN apt install -y --no-install-recommends ca-certificates curl

WORKDIR /pycode
RUN pip install --no-cache-dir -r requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple

WORKDIR /usr/local/bin/
CMD ./converter run