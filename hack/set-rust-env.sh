 #!/bin/bash

set +e
rc=`rustc --version | grep rustc`
if [ $? -eq 0 ]; then
    exit 0
fi
set -e

curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh