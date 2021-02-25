SHELL := /bin/bash
all: build_test test build install
build_test:
	go build -o tests/bin/augoment
build:
	sudo go build -o /usr/local/bin/augoment
	source ~/.bashrc
install:
	go install
