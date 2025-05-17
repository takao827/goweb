GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

.PHONY: all
all: gobuild

.PHONY: clean
clean:
	@go clean -cache
	@rm -rf bin

.PHONY: test
test:
	@go test ./...

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: vet
vet:
	@go vet ./...

check: fmt lint vet test

.PHONY: build
build:
	@go build -o ./bin/${GOOS}/${GOARCH}/ ./cmd/goweb

.PHONY: run
run: 
	@go run ./cmd/goweb
