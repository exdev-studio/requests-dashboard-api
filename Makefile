.PHONY: build test install

build:
	go build -o bin/apiserver ./cmd/apiserver

test:
	go test ./...

install:
	go install ./...

.DEFAULT_GOAL := build
