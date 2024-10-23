# Copyright 2018, Boon Logic Inc

.PHONY: examples test format format-check clean generate-client docs go-check

# load top-level makefile variables
TOP?=$(shell cd .. && git rev-parse --show-toplevel)
CWD=$(shell pwd)
-include $(TOP)/mk/base.mk

$(info GOPATH=$(GOPATH))
$(info INSTALL_ROOT=$(INSTALL_ROOT))


format: go-check ## Run the formatter on go code
	go fmt ./...

examples_v1:
	go build examples/v1/connect.go && \
	go build examples/v1/full-example.go && \
	go build examples/v1/pretrain.go && \
	go build examples/v1/stream-advanced.go

examples_v2:
	go build examples/v2/connect.go && \
	go build examples/v2/full-example.go && \
	go build examples/v2/pretrain.go && \
	go build examples/v2/stream-advanced.go

examples: examples_v1 examples_v2

format-check: format ## Run the formatter and perform diff (for pipeline)
	git diff --exit-code; if [ $$? -ne 0 ]; then echo "format-check failed"; exit 1; fi; \
	echo "*** format-check passed"

clean: ## clean up go cache and modcache
	go clean -modcache -cache

generate_v1: ## generate amber swagger client code based on json schema file
	bin/swagger generate client --keep-spec-order -t v1 -f amber-api.json

generate_v2: ## generate amber swagger client code based on json schema file
	bin/swagger generate client --keep-spec-order -t v2 -f amber-api-v2.json

docs: go-check ## generate documentation
	./generate-docs.sh

# test-v1, test-v1next, test-dev, test-qa, test-aoc, test-oap
# add additional .license files in test directory to expand / customize tests
test-v1-%: go-check
	AMBER_TEST_LICENSE_ID=$* go test v1/sdk.go v1/sdk_test.go -timeout 30m -v -coverprofile .coverage.out .
 
test-v2-%: go-check
	AMBER_TEST_LICENSE_ID=$* go test v2/sdk.go v2/sdk_test.go -timeout 30m -v -coverprofile .coverage.out .

go-check:
ifndef GOPATH
	$(error GOPATH is undefined)
endif
