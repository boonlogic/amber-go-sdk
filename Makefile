# Copyright 2018, Boon Logic Inc

.PHONY: test format format-check clean generate-client docs go-check

format: go-check ## Run the formatter on go code
	go fmt ./...

format-check: format ## Run the formatter and perform diff (for pipeline)
	git diff --exit-code; if [ $$? -ne 0 ]; then echo "format-check failed"; exit 1; fi; \
	echo "*** format-check passed"

clean: ## clean up go cache and modcache
	go clean -modcache -cache

generate-client: ## generate amber swagger client code based on json schema file
	bin/swagger generate client --keep-spec-order -f swagger.json

docs: go-check ## generate documentation
	 $(GOPATH)/bin/gomarkdoc --output docs/functions.md . ./models

go-check:
ifndef GOPATH
	$(error GOPATH is undefined)
endif

