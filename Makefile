
all:

setup:
	@go get github.com/golang/lint/golint
	@go get golang.org/x/tools/cmd/goimports
	@go get github.com/Masterminds/glide

test: setup
	go test $$(glide novendor)

lint: setup
	go vet $$(glide novendor)
	for pkg in $$(glide novendor -x); do \
		golint -set_exit_status $$pkg || exit $$?; \
	done


fmt: setup
	goimports -w $$(glide nv -x)   

.PHONY: all
