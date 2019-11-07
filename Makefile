
all: build

build: fmt
	go build

setup: dep
	@go get golang.org/x/lint/golint
	@go get golang.org/x/tools/cmd/goimports

dep:
	@dep ensure

test: build
	go test

.PHONY: all setup dep
