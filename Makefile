# Copyright 2018, Boon Logic Inc

.PHONY: examples test format format-check clean generate-client docs go-check

# load top-level makefile variables
TOP?=$(shell cd .. && git rev-parse --show-toplevel)
CWD=$(shell pwd)
-include $(TOP)/mk/base.mk

format: go-check ## Run the formatter on go code
	go fmt ./...

examples:
	go build examples/connect.go && \
	go build examples/full-example.go && \
	go build examples/pretrain.go && \
	go build examples/stream-advanced.go

format-check: format ## Run the formatter and perform diff (for pipeline)
	git diff --exit-code; if [ $$? -ne 0 ]; then echo "format-check failed"; exit 1; fi; \
	echo "*** format-check passed"

clean: ## clean up go cache and modcache
	go clean -modcache -cache

generate: ## generate amber swagger client code based on json schema file
	bin/swagger generate client --keep-spec-order -f swagger.json

docs: go-check ## generate documentation
	./generate-docs.sh

# test-v1, test-v1next, test-dev, test-qa, test-aoc, test-oap
# add additional .license files in test directory to expand / customize tests
test-%: go-check
	AMBER_TEST_LICENSE_ID=$* go test -timeout 30m -v -coverprofile .coverage.out .

go-check:
ifndef GOPATH
	$(error GOPATH is undefined)
endif
