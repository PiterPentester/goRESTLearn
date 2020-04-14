.PHONY: build
build:
	go build  -v ./cmd/apiserver

.PHONY: build
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build
