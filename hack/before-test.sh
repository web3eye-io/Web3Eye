#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

PLATFORM=linux/amd64
OUTPUT=./output

mkdir -p $OUTPUT/$PLATFORM
for service_name in $(ls $(pwd)/cmd); do
    cp $(pwd)/cmd/$service_name/*.viper.yaml $OUTPUT/$PLATFORM
    cd $OUTPUT/$PLATFORM
    ./$service_name run | grep error &
done

sleep 2
