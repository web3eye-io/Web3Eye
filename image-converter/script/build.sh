#!/usr/bin/env bash
SHELL_FOLDER=$(cd "$(dirname "$0")";pwd)
PROJECT_FOLDER=$(cd $SHELL_FOLDER/../;pwd)
ROOT_FOLDER=$(cd $PROJECT_FOLDER/../;pwd)

set -o errexit
set -o nounset
set -o pipefail

configFile=$ROOT_FOLDER/config/config.toml
myConfigPath=$PROJECT_FOLDER/pkg/utils/
echo "sync config.toml from $configFile to $myConfigPath"
cp $configFile $myConfigPath
cp $configFile $PROJECT_FOLDER

cd $PROJECT_FOLDER