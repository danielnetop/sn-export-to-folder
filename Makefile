SHELL := bash

LINTER_VERSION=v1.51.2

BINARY = sn-export-folder
ifeq ($(OS),Windows_NT)
	BINARY := $(BINARY).exe
endif

.PHONY: test run

## default: linting, test and build the application binary
default: build

## mod: run go mod vendor -v
mod:
	go mod vendor -v

## tidy: run go mod tidy -v
tidy:
	go mod tidy -v

## mod-tidy: runs go mod and tidy together
mod-tidy: mod tidy

## lint: download/install golangci-lint and analyse the source code with the configuration in .golangci.yml
lint: #get-linter
	./bin/golangci-lint run --timeout=5m

## lint-fix: download/install golangci-lint, analyse and fix the source code with the configuration in .golangci.yml
lint-fix: get-linter
	./bin/golangci-lint run --timeout=5m --fix

get-linter:
	command -v ./bin/golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s $(LINTER_VERSION)

## build: test and build the application binary
build: clean test
	go build -o ./$(BINARY) ./...

## format: run go format
format:
	go fmt ./...

## test: run mod, tidy, linters and run unit tests
test: mod-tidy lint test-unit

## test-unit: run unit tests
test-unit: format
	go test ./...

## cover-report: run coverage and show html report
cover-report: cover
	go tool cover -html=coverage.out

## cover: run tests with coverage report.
cover:
	go test -v -coverprofile=coverage.out -coverpkg=./... ./...
	go tool cover -func coverage.out

## run: runs the service locally with go run using the config defined on the Makefile
run:
	go run ./...

## clean: remove exported folders, built binary and coverage
clean:
	rm -f $(BINARY)
	rm -rf exported
	rm -f coverage.out

## help: prints this help message
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
