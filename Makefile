# Variables
BINARY_NAME := hash
BUILD_DIR := build
GIT_COMMIT := $(shell git rev-parse HEAD)
LDFLAGS := -ldflags="-X main.version=$(GIT_COMMIT)"

.PHONY: all
all: test format build

.PHONY: test
test:
	go test -v ./...

.PHONY: format
format:
	go fmt ./...

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

.PHONY: build
build: build-linux build-darwin-amd64 build-darwin-arm64 build-windows

.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)_linux_amd64

.PHONY: build-darwin-amd64
build-darwin-amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)_darwin_amd64

.PHONY: build-darwin-arm64
build-darwin-arm64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)_darwin_arm64

.PHONY: build-windows
build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)_windows_amd64.exe

