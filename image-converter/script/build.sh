#!/usr/bin/env bash
SHELL_FOLDER=$(cd "$(dirname "$0")";pwd)
ROOT_FOLDER=$(cd $SHELL_FOLDER/../;pwd)
PROJECT_FOLDER=$(cd $ROOT_FOLDER/../;pwd)

set -o errexit
set -o nounset
set -o pipefail

configFile=$PROJECT_FOLDER/config/config.toml
myConfigPath=$ROOT_FOLDER/pkg/utils/
echo "sync config.toml from $configFile to $myConfigPath"
cp $configFile $myConfigPath

cd $ROOT_FOLDER