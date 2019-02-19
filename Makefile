SHELL := bash

.PHONY: all

all: fmt test lint

test:
	go test -coverprofile=coverage.txt -covermode=atomic ./...

lint:
	gometalinter ./... --vendor --exclude="exported.*should have comment.*or be unexported\b"

lint-full:
	gometalinter ./... --vendor

fmt:
	gofmt -w=true -s $$(find . -type f -name '*.go' -not -path "./vendor/*")
	goimports -w=true $$(find . -type f -name '*.go' -not -path "./vendor/*")
