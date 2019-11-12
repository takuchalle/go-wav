
all: build

build:
	go build

test: build
	go test

.PHONY: all build test
