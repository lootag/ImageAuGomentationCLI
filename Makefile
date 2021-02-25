SHELL := /bin/bash
all: build_test build install
build_test:
	go build -o tests/bin/augoment
build:
	go build -o /usr/local/bin/augoment
	source ~/.bashrc
install:
	go install
