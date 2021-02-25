SHELL := /bin/bash
all: build_test build install
build_test:
	go build -o tests/bin/augoment
build:
	go build -o augoment
	source ~/.bashrc
install:
	go install
