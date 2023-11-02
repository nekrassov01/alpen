BIN := alpen
ifeq ($(OS),Windows_NT)
BIN := $(BIN).exe
endif

GOBIN ?= $(shell go env GOPATH)/bin
VERSION := $$(make -s show-version)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := "-s -w -X main.Version=$(VERSION) -X main.Revision=$(REVISION)"

HAS_LINT := $(shell command -v $(GOBIN)/golangci-lint 2> /dev/null)
HAS_VULNCHECK := $(shell command -v $(GOBIN)/govulncheck 2> /dev/null)
HAS_GOBUMP := $(shell command -v $(GOBIN)/gobump 2> /dev/null)

BIN_LINT := github.com/golangci/golangci-lint/cmd/golangci-lint@latest
BIN_GOVULNCHECK := golang.org/x/vuln/cmd/govulncheck@latest
BIN_GOBUMP := github.com/x-motemen/gobump/cmd/gobump@latest

export GO111MODULE=on

.PHONY: build
build: clean
	go mod tidy
	go build -ldflags "-X main.Version=$(VERSION) -X main.Revision=$(REVISION)" -o $(BIN) .

.PHONY: put
put: build
	cp $(BIN) $(GOBIN)/$(BIN)

.PHONY: check
check: test cover golangci-lint govulncheck

.PHONY: deps
deps: deps-lint deps-govulncheck deps-gobump

.PHONY: deps-lint
deps-lint:
ifndef HAS_LINT
	go install $(BIN_LINT)
endif

.PHONY: deps-govulncheck
deps-govulncheck:
ifndef HAS_VULNCHECK
	go install $(BIN_GOVULNCHECK)
endif

.PHONY: deps-gobump
deps-gobump:
ifndef HAS_GOBUMP
	go install $(BIN_GOBUMP)
endif

.PHONY: test
test:
	go test ./... -v -cover -coverprofile=cover.out

.PHONY: cover
cover:
	go tool cover -html=cover.out -o cover.html

.PHONY: golangci-lint
golangci-lint: deps-lint
	golangci-lint run ./... -v

.PHONY: govulncheck
govulncheck: deps-govulncheck
	$(GOBIN)/govulncheck -test ./...

.PHONY: show-version
show-version: deps-gobump
	$(GOBIN)/gobump show -r .

.PHONY: check-git
ifneq ($(shell git status --porcelain),)
	$(error git workspace is dirty)
endif
ifneq ($(shell git rev-parse --abbrev-ref HEAD),main)
	$(error current branch is not main)
endif

.PHONY: publish
publish: deps-gobump check-git
	$(GOBIN)/gobump up -w .
	git commit -am "bump up version to $(VERSION)"
	git push origin main

.PHONY: release
release: check-git
	git tag "v$(VERSION)"
	git push origin "refs/tags/v$(VERSION)"

.PHONY: clean
clean:
	go clean
	rm -f $(BIN) cover.out cover.html
