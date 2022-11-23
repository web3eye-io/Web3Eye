#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd -P)

# Default timeout is 1800s
TEST_TIMEOUT=${TIMEOUT:-1800}

# Write go test artifacts here
ARTIFACTS=${ARTIFACTS:-"${REPO_ROOT}/tmp"}
pkg=github.com/NpoolPlatform/go-service-framework/pkg/version

for arg in "$@"
do
    case $arg in
        -t=*|--timeout=*)
        TEST_TIMEOUT="${arg#*=}"
        shift
        ;;
        -t|--timeout)
        TEST_TIMEOUT="$2"
        shift
        shift
    esac
done

cd "${REPO_ROOT}"

mkdir -p "${ARTIFACTS}"

go_test_flags=(
    -v
    -count=1
    -timeout="${TEST_TIMEOUT}s"
    -cover -coverprofile "${ARTIFACTS}/coverage.out"
)

if git_status=$(git status --porcelain --untracked=no 2>/dev/null) && [[ -z "${git_status}" ]]; then
    git_tree_state=clean
fi

git_branch=`git rev-parse --abbrev-ref HEAD`
set +e
version=`git describe --tags --abbrev=0`
if [ ! $? -eq 0 ]; then
    version=$git_branch
fi
set -e

compile_date=`date -u +'%Y-%m-%dT%H:%M:%SZ'`
git_revision=`git rev-parse HEAD 2>/dev/null || echo unknow`

go test -ldflags "-s -w -X $pkg.buildDate=${compile_date} \
        -X $pkg.gitCommit=${git_revision} \
        -X $pkg.gitVersion=${version}     \
        -X $pkg.gitBranch=${git_branch}"  \
        ./... -coverprofile ${ARTIFACTS}/coverage.out

go tool cover -html "${ARTIFACTS}/coverage.out" -o "${ARTIFACTS}/coverage.html"
