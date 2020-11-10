SHELL := /bin/bash
all: build_test test build install
build_test:
	go build -o tests/bin/augoment
test:
	go test ./tests
build:
	go build -o ../bin/augoment
install:
	sudo mv ../bin/augoment /usr/local/bin
	source ~/.bashrc
