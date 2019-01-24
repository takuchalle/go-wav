
all: build

build: dep fmt
	go build

setup:
	@go get github.com/golang/lint/golint
	@go get golang.org/x/tools/cmd/goimports
	@go get github.com/Masterminds/glide

dep:
	@dep ensure

test: build
	go test $$(glide novendor)

lint:
	go vet $$(glide novendor)
	for pkg in $$(glide novendor -x); do \
		golint -set_exit_status $$pkg || exit $$?; \
	done


fmt:
	goimports -w $$(glide nv -x)   

.PHONY: all setup fmt test lint dep
