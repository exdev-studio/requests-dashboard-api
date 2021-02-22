.PHONY: build

build:
	go build -o bin/apiserver ./cmd/apiserver

.DEFAULT_GOAL := build
