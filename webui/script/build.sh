#!/usr/bin/env bash
SHELL_FOLDER=$(
    cd "$(dirname "$0")"
    pwd
)
ROOT_FOLDER=$(
    cd $SHELL_FOLDER/../
    pwd
)

set -o errexit
set -o nounset
set -o pipefail

cd $ROOT_FOLDER

if git_status=$(git status --porcelain --untracked=no 2>/dev/null) && [[ -z "${git_status}" ]]; then
    git_tree_state=clean
fi

git_branch=$(git rev-parse --abbrev-ref HEAD)
set +e

if [ ! $? -eq 0 ]; then
    version=$git_branch
fi
set -e

git_revision=$(git rev-parse HEAD 2>/dev/null || echo unknow)

tag=""
if [ "xlatest" == "x$1" ] || ["x" == "x$1" ]; then
    tag=$(git describe --tags --abbrev=0)
else
    tag=$1
fi

major=$(echo $1 | awk -F '.' '{ print $1 }')
minor=$(echo $1 | awk -F '.' '{ print $2 }')
patch=$(echo $1 | awk -F '.' '{ print $3 }')
version=$major.$minor.$patch

sed -ri "s#\"version(.*)#\"version\": \"$tag\",#" package.json

PATH=/usr/local/bin:$PATH:./node_modules/@quasar/app/bin command quasar
rc=$?
set -e
if [ ! $rc -eq 0 ]; then
    n v16.14.0
    cd ~
    PATH=/usr/local/bin:$PATH npm install --global --registry https://registry.npm.taobao.org yarn
    PATH=/usr/local/bin:$PATH yarn add global quasar-cli@latest
fi

cd $ROOT_FOLDER
PATH=/usr/local/bin:$PATH:./node_modules/@quasar/app/bin yarn install --registry https://registry.npm.taobao.org/
PATH=/usr/local/bin:$PATH:./node_modules/@quasar/app/bin quasar build
