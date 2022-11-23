REPO_ROOT	:= $(shell git rev-parse --show-toplevel)

.DEFAULT_GOAL:=help
SHELL:=/usr/bin/env bash

COLOR:=\\033[36m
NOCOLOR:=\\033[0m
GITREPO=$(shell git remote -v | grep fetch | awk '{print $$2}' | sed 's/\.git//g' | sed 's/https:\/\///g')
SUBCMDS=$(wildcard cmd/*)
SERVICES=$(SUBCMDS:cmd/%=%)
SERVICEIMAGES=$(SERVICES:%=%-image)
SERVICEIMAGERELEASES=$(SERVICES:%=%-release)
SERVICEK8SDEPLOYS=$(SERVICES:%=%-k8s-deploy)

##@ init project
init:
	cp -f .githooks/* .git/hooks

go.mod:
	go mod init ${GITREPO}
	go mod tidy -compat=1.17

deps:
	go get -d ./...
	go mod tidy -compat=1.17

##@ Verify

.PHONY: add-verify-hook verify verify-build verify-golangci-lint verify-go-mod verify-shellcheck verify-spelling all

add-verify-hook: ## Adds verify scripts to git pre-commit hooks.
# Note: The pre-commit hooks can be bypassed by using the flag --no-verify when
# performing a git commit.
	git config --local core.hooksPath "${REPO_ROOT}/.githooks"

# TODO(lint): Uncomment verify-shellcheck once we finish shellchecking the repo.
verify: go.mod verify-build verify-golangci-lint verify-go-mod #verify-shellcheck ## Runs verification scripts to ensure correct execution
	${REPO_ROOT}/hack/verify.sh

verify-build: ## Builds the project for a chosen set of platforms
	${REPO_ROOT}/hack/verify-build.sh ...

verify-go-mod: ## Runs the go module linter
	${REPO_ROOT}/hack/verify-go-mod.sh

verify-golangci-lint: ## Runs all golang linters
	${REPO_ROOT}/hack/verify-golangci-lint.sh

verify-shellcheck: ## Runs shellcheck
	${REPO_ROOT}/hack/verify-shellcheck.sh

verify-spelling: ## Verifies spelling.
	${REPO_ROOT}/hack/verify-spelling.sh

gen-ent:
	go get entgo.io/ent/cmd/ent
	go run entgo.io/ent/cmd/ent generate --feature entql,sql/upsert,privacy,schema/snapshot,sql/modifier ./pkg/db/ent/schema

all: verify-build

${SERVICES}:
	${REPO_ROOT}/hack/verify-build.sh $@

${SERVICEIMAGES}:
	${REPO_ROOT}/hack/generate-docker-image.sh $(@:%-image=%) $(DEVELOPMENT) $(DOCKER_REGISTRY)

${SERVICEIMAGERELEASES}:
	${REPO_ROOT}/hack/release-docker-image.sh $(@:%-release=%) $(TAG) $(DOCKER_REGISTRY)

${SERVICEK8SDEPLOYS}:
	${REPO_ROOT}/hack/deploy-to-k8s-cluster.sh $(@:%-k8s-deploy=%) $(TAG)

generate-docker-images: ${SERVICES} ${SERVICEIMAGES}
release-docker-images: ${generate-docker-images} ${SERVICEIMAGERELEASES}
deploy-to-k8s-cluster: ${SERVICEK8SDEPLOYS}

##@ Tests

.PHONY: test test-go-unit test-go-integration

before-test: verify-build
	${REPO_ROOT}/hack/before-test.sh

test: verify-build test-go-unit ## Runs unit tests
test-verbose:
	VERBOSE=1 make test

after-test:
	${REPO_ROOT}/hack/after-test.sh

test-go-unit: ## Runs Golang unit tests
	${REPO_ROOT}/hack/test-go.sh


##@ Helpers

.PHONY: help

help:  ## Display this help
	@awk \
		-v "col=${COLOR}" -v "nocol=${NOCOLOR}" \
		' \
			BEGIN { \
				FS = ":.*##" ; \
				printf "\nUsage:\n  make %s<target>%s\n", col, nocol \
			} \
			/^[a-zA-Z_-]+:.*?##/ { \
				printf "  %s%-15s%s %s\n", col, $$1, nocol, $$2 \
			} \
			/^##@/ { \
				printf "\n%s%s%s\n", col, substr($$0, 5), nocol \
			} \
		' $(MAKEFILE_LIST)
