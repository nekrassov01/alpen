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
	go build -ldflags "-X main.Version=$(VERSION) -X main.Revision=$(REVISION)" -o $(BIN)

.PHONY: check
check: test vet golangci-lint govulncheck

.PHONY: deps
deps:
ifndef HAS_LINT
	go install $(BIN_LINT)
endif
ifndef HAS_VULNCHECK
	go install $(BIN_GOVULNCHECK)
endif
ifndef HAS_GOBUMP
	go install $(BIN_GOBUMP)
endif

.PHONY: golangci-lint
golangci-lint:
	golangci-lint run ./... -v --tests

.PHONY: govulncheck
govulncheck: deps
	$(GOBIN)/govulncheck ./...

.PHONY: show-version
show-version: deps
	$(GOBIN)/gobump show -r .

.PHONY: bump
bump: deps
	@gobump up -w .

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go test -race -cover -v ./... -coverprofile=cover.out -covermode=atomic

.PHONY: cover
cover:
	go tool cover -html=cover.out -o cover.html

.PHONY: clean
clean:
	go clean
	rm -f $(BIN)
