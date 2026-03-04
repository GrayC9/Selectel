.PHONY: all build test cover lint plugin-build vet

all: vet test build

build:
	go build ./...

test:
	go test ./... -v -race

cover:
	go test ./... -coverprofile=coverage.out -coverpkg=./... && go tool cover -html=coverage.out

lint:
	golangci-lint run ./...

plugin-build:
	go build -tags plugin -buildmode=plugin -o loglint.so ./plugin/

vet:
	go vet ./...
