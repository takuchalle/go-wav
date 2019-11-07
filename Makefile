
all: build

build: dep
	go build

setup:
	@go get golang.org/x/lint/golint
	@go get golang.org/x/tools/cmd/goimports

dep:
	@dep ensure

test: build
	go test

.PHONY: all setup dep
