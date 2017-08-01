# Root directory of the project (absolute path).
ROOTDIR=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))

PROJECT_ROOT=github.com/tossmilestone/swarmkit

PACKAGES=$(shell go list ./... | grep -v /vendor/)

# Race detector is only supported on amd64.
RACE := $(shell test $$(go env GOARCH) != "amd64" || (echo "-race"))

.PHONY: all check ci setup lint test

all: check test

check: lint

ci: check test

setup: ## install dependencies
	@go get -u github.com/golang/lint/golint

lint: ## run go lint
	@test -z "$$(golint ./... | grep -v vendor/ | tee /dev/stderr)"

test: ## run tests
	@go test -parallel 8 ${RACE} ${PACKAGES}
