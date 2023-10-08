 #!/bin/bash
MY_PATH=`cd $(dirname $0);pwd`
ROOT_PATH=$MY_PATH/../
node_name=v$NODEVERSION

set +e
rc=`node -v | grep $node_name`
if [ $? -eq 0 ]; then
    exit 0
fi
set -e

echo "Will change node version to $node_name"

node_tar="node-$node_name-linux-x64.tar.xz"
node_tar_url="https://nodejs.org/dist/$node_name/$node_tar"

node_data=$NODETMPENV

mkdir -p $NODEHOME

shopt -s expand_aliases
alias NODE="$NODEHOME/bin/node"

set +e
rc=`node -v | grep $node_name`
if [ ! $? -eq 0 ]; then
  set -e
  echo "Fetching $node_tar from $node_tar_url, stored to $node_data"
  curl -L $node_tar_url -o $node_data/$node_tar
  tar -xvf $node_data/$node_tar --strip-components 1 -C $NODEHOME
fi

set -e