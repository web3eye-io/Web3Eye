#/bin/bash
UNZIP := $(shell which unzip)

ifeq ($(GOBIN),)
	GOBIN=$(shell dirname $(shell which go))
endif

install: pre tool

pre:
ifndef UNZIP
	$(error "unzip not exist, please run 'apt install unzip'")
endif

tool:
	curl -sL -o /tmp/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.18.1/protoc-3.18.1-linux-x86_64.zip && \
	${UNZIP} -u /tmp/protoc.zip -d /tmp/protoc && \
	mv /tmp/protoc/bin/protoc ${GOBIN} && \
	go mod tidy -compat=1.19 && \
	go get -d google.golang.org/grpc && \
	go get -d github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc && \
	go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc \
    github.com/favadi/protoc-go-inject-tag
