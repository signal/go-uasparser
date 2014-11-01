PWD = $(shell pwd)
.PHONY: env deps test-deps test benchmark clean

# environment

GOPATH = $(PWD)/.go:$(PWD)
GO     = GOPATH=$(GOPATH) go
DEPS   = $(shell go list -f '{{range .Imports}}{{.}} {{end}}' ./...)

.go:
	mkdir -p $(PWD)/.go/{bin,pkg,src}

env: .go

# deps

deps: env
	@echo "# Install deps"
	@for dep in $(DEPS); do $(GO) get -d ${dep}; done

# test deps

tmp/uas-data:
	@echo "# Download integration test data"
	@mkdir -p tmp/uas-data
	@$(PWD)/bin/uas-download-testing-samples.sh

test-deps: tmp/uas-data

# commands

test: deps test-deps
	$(GO) test

benchmark: deps
	$(GO) test -v -bench=. -run='Benchmark*'

clean: env
	rm -fr $(PWD)/.go $(PWD)/tmp/uas-data

