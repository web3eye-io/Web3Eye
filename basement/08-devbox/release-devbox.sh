#!/bin/bash
SHELL_DIR=$(
    cd "$(dirname $0)"
    pwd
)

docker build -t coastlinesss/development-box:latest .
docker push coastlinesss/development-box:latest
