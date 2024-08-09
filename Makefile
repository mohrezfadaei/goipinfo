BIN_DIR=./build/bin
BINARY_NAME=goipinfo
BINARY_PATH=$(BIN_DIR)/$(BINARY_NAME)

build:
		go build -o $(BINARY_PATH) .

run: build
		./$(BINARY_PATH)

clean:
		rm $(BINARY_PATH)

test:
		go test -v ./...

fmt:
		go fmt ./...

lint:
		golangci-lint run

deps:
		go mod tidy

build-linux:
		GOOS=linux GOARCH=amd64 go build -o $(BINARY_PATH)-linux-amd64 .
		GOOS=linux GOARCH=arm64 go build -o $(BINARY_PATH)-linux-arm64 .

.PHONY: build run clean test fmt lint deps build-linux
