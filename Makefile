.PHONY: build, test

build:
	go build -o bin/apiserver ./cmd/apiserver

test:
	go test -v ./...

.DEFAULT_GOAL := build
