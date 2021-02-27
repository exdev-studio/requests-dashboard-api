.PHONY: build, test

build:
	go build -o bin/apiserver ./cmd/apiserver

test:
	go test ./...

.DEFAULT_GOAL := build
