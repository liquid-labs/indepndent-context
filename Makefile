.DELETE_ON_ERROR:

SHELL:=bash

default: all

PHONY_TARGETS:=all default lint qa

build: main.go
	go build -o indepndent-context main.go

lint:
	golangci-lint run

qa: lint