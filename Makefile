.PHONY: build
build:
	go build -v ./cmd/server

_DEFAULT_GOAL := build