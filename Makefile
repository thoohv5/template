include Makefile.common

VERSION = $(shell git describe --tags --always)

.PHONY: generate
# generate
generate:
	@echo ">> generating code"
	@$(GO) generate ./...
	@$(MAKE) format


.PHONY: build
# build
build:
	mkdir -p build/ && $(GO) build -ldflags "-X main.Version=$(VERSION)" -o build/ ./...


.PHONY: run
# run
run:
	@$(MAKE) generate
	@$(MAKE) common-config
	@$(GO) run $(PREFIX)/cmd/server -conf $(PREFIX)/configs/config.toml


