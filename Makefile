SHELL := /bin/bash
all: build_test test build install
build_test:
	go build -o tests/bin/augoment
test:
	go test -v -cover -race ./...
build:
	go build -o /usr/local/bin/augoment
install:
	go install
