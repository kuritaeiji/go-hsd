.DEFAULT_GOAL := test

tidy:
	go mod tidy
.PHONY: tidy

fmt: tidy
	goimports -w -l hsd.go hsd_test.go
.PHONY: fmt

lint: fmt
	golangci-lint run hsd.go hsd_test.go
.PHONY: lint

test: lint
	go test -v -cover -coverprofile=cover.out ./...
	go tool cover -html=cover.out -o cover.html
.PHONY: test