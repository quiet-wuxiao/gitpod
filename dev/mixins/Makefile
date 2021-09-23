BIN_DIR?=$(shell pwd)/tmp/bin
JB_BIN=$(BIN_DIR)/jb
JSONNET_BIN=$(BIN_DIR)/jsonnet
JSONNETFMT_BIN=$(BIN_DIR)/jsonnetfmt
TOOLING=$(JSONNETFMT_BIN) $(JSONNET_BIN) $(JB_BIN)

JSONNET_FMT := $(JSONNETFMT_BIN) -n 2 --max-blank-lines 2 --string-style s --comment-style s

all: setup-workspace fmt lint

.PHONY: clean
clean:
    # Delete files marked in .gitignore
	git clean -Xfd .

.PHONY: setup-workspace
setup-workspace:
	go install github.com/jsonnet-bundler/jsonnet-bundler/cmd/jb@master
	go install github.com/google/go-jsonnet/cmd/jsonnet@master
	go install github.com/google/go-jsonnet/cmd/jsonnetfmt@master
	GO111MODULE=on go install github.com/prometheus/prometheus/cmd/promtool@release-2.26
	export PATH=$(PATH):$(PWD)/tmp/bin

.PHONY: fmt
fmt: $(JSONNETFMT_BIN)
	find . -name 'vendor' -prune -o -name '*.libsonnet' -print -o -name '*.jsonnet' -print | \
		xargs -n 1 -- $(JSONNET_FMT) -i

.PHONY: lint
lint: $(JSONNETFMT_BIN)
	find . -name 'vendor' -prune -o -name '*.libsonnet' -print -o -name '*.jsonnet' -print | \
		while read f; do \
			$(JSONNET_FMT) "$$f" | diff -u "$$f" -; \
		done

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

$(TOOLING): $(BIN_DIR)
	@echo Installing tools from tools.go
	cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go build -modfile=go.mod -o $(BIN_DIR) %