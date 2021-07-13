.PHONY: build
build:
	env GOOS=linux GOARCH=amd64 go build  -v ./cmd/server

deploy:
	rsync -ravzP server de1035:///home/alexeykasatikov/oncallstats/

_DEFAULT_GOAL := build
