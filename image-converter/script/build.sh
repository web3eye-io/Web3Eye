#!/usr/bin/env bash
SHELL_FOLDER=$(cd "$(dirname "$0")";pwd)
ROOT_FOLDER=$(cd $SHELL_FOLDER/../;pwd)

set -o errexit
set -o nounset
set -o pipefail

echo "Nothing to do"
cd $ROOT_FOLDER