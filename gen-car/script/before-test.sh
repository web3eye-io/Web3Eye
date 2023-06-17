#!/usr/bin/env bash
SHELL_FOLDER=$(cd "$(dirname "$0")";pwd)
PROJECT_FOLDER=$(cd $SHELL_FOLDER/../;pwd)

cd $PROJECT_FOLDER

set -o errexit
set -o nounset
set -o pipefail

PLATFORM=linux/amd64
OUTPUT=./output

mkdir -p $OUTPUT/$PLATFORM
for service_name in $(ls $(pwd)/cmd); do
    cd $OUTPUT/$PLATFORM
    ./$service_name run | grep error &
done

sleep 2
