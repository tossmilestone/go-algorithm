# Root directory of the project (absolute path).
ROOTDIR=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))

PROJECT_ROOT=github.com/tossmilestone/swarmkit

# Race detector is only supported on amd64.
RACE := $(shell test $$(go env GOARCH) != "amd64" || (echo "-race"))

.PHONY test

all: check test

check: lint

ci: check

lint: ## run go lint
	@test -z "$$(golint ./... | tee /dev/stderr)"

test: ## run tests, except integration tests
	@go test -parallel 8 ${RACE}
