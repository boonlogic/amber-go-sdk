# Copyright 2018, Boon Logic Inc

#build: go-check update-template ## build the core server file
#	go build -a -modcacherw -o $(INSTALL_ROOT)/bin/expert-server$(EXT) -work $(STATIC_BUILD_OPTS) cmd/nano-server/*.go && \
#	cp swagger/amber/amber-api.json $(INSTALL_ROOT)/static && \
#	cp swagger/expert/nano-api.json $(INSTALL_ROOT)/static && \
#	cp start-nano $(INSTALL_ROOT)/bin && \
#	mkdir -p $(INSTALL_ROOT)/security && \
#	cp security/server.crt security/server.key $(INSTALL_ROOT)/security && \
#	cp bin/install $(INSTALL_ROOT)/bin

#install: build ## install expert server (same as build)

#test: ## run all unit tests and compile coverage results
#	mkdir -p $(INSTALL_ROOT)/mnt/amber && \
#	BOON_AMBER_STORE=$(INSTALL_ROOT)/mnt/amber bin/test-runner

format: ## Run the formatter on go code
	cd $(GOPATH)/src/$(DEVPATH)/expert/expert-api && go fmt ./...

format-check: format ## Run the formatter and perform diff (for pipeline)
	git diff --exit-code; if [ $$? -ne 0 ]; then echo "format-check failed"; exit 1; fi; \
	echo "*** format-check passed"

clean: ## clean up go cache and modcache
	go clean -modcache -cache

generate-client: ## generate amber swagger client code based on json schema file
	bin/swagger generate client --keep-spec-order -f swagger.json

