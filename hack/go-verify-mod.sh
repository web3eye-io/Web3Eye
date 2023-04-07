#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

go mod tidy -compat=1.17
git diff --exit-code go.*
