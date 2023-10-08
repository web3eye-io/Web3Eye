#!/usr/bin/env bash
SHELL_FOLDER=$(
    cd "$(dirname "$0")"
    pwd
)
PROJECT_FOLDER=$(
    cd $SHELL_FOLDER/../
    pwd
)

set -o errexit
set -o nounset
set -o pipefail

PLATFORMS=(
    linux/amd64
    # windows/amd64
    # darwin/amd64
)
OUTPUT=$PROJECT_FOLDER/output

function GetOnnxFile(){
    ONNX_URL="https://cyber-tracer-public.s3.ap-southeast-1.amazonaws.com/resnet50_v2.onnx"
    ONNX_FILE=$PROJECT_FOLDER/model/resnet50_v2.onnx
    ONNX_MD5SUM="7fae99a20c1b51aeb5bb2c04b6613642"
    md5sum=""
    if [ -f "$ONNX_FILE" ]; then
        md5sum=$(md5sum $ONNX_FILE | awk '{ print $1 }')
    fi

    if [ "$md5sum" != "$ONNX_MD5SUM" ]; then
        curl $ONNX_URL -o $ONNX_FILE
    fi
}

GetOnnxFile

pkg=github.com/NpoolPlatform/go-service-framework/pkg/version
    service_name=$(
        cd $PROJECT_FOLDER
        basename $(pwd)
    )


for PLATFORM in "${PLATFORMS[@]}"; do
    OS="${PLATFORM%/*}"
    ARCH=$(basename "$PLATFORM")

    if git_status=$(git status --porcelain --untracked=no 2>/dev/null) && [[ -z "${git_status}" ]]; then
        git_tree_state=clean
    fi

    git_branch=$(git rev-parse --abbrev-ref HEAD)
    set +e
    version=$(git describe --tags --abbrev=0)
    if [ ! $? -eq 0 ]; then
        version=$git_branch
    fi
    set -e

    compile_date=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
    git_revision=$(git rev-parse HEAD 2>/dev/null || echo unknow)

    echo "Building project for $PLATFORM -- $version $compile_date $git_revision"
    GOOS=${OS} GOARCH=${ARCH} go build -v -ldflags "-s -w \
        -X $pkg.buildDate=${compile_date} \
        -X $pkg.gitCommit=${git_revision} \
        -X $pkg.gitVersion=${version}     \
        -X $pkg.gitBranch=${git_branch}" \
        -o "${OUTPUT}/${OS}/${ARCH}/" "$PROJECT_FOLDER/cmd/$service_name"
done
