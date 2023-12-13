#!/bin/bash

# cd to the root path
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd -P)"
cd "${ROOT}"

set -o errexit
set -o nounset
set -o pipefail

es_projects=(webui dashboard)

for proj in ${es_projects[@]}
do

cd ${ROOT}/$proj
if [[ ! -f .eslintrc.js ]]; then
    echo 'ERROR: missing .eslintrc.js in project root' >&2
    exit 1
fi

SUDO="sudo"
user=$(whoami)
if [ "$user" == "root" ]; then
    SUDO=""
fi

if ! command -v n; then
    $SUDO npm install -g n -y
fi

if ! command -v yarn; then
    $SUDO npm install -g yarn -y
    
fi

PATH=/usr/local/bin:$PATH
$SUDO npm config set registry https://registry.npm.taobao.org
$SUDO n v16.14.0
$SUDO npm install @typescript-eslint/eslint-plugin --save-dev
$SUDO yarn lint

done



