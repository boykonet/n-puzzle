BINARY_NAME=n_puzzle
SRC_DIR=.
# Gets the version info
VERSION=$(shell git describe --tags --always --dirty)
BUILD_DIR=.
LDFLAGS=-ldflags "-X main.version=$(VERSION)"
export GO111MODULE=on

deps:
	@echo "Installing dependencies..."
	@go mod tidy
	@go mod download

clean:
	@echo "Cleaning up..."
	@rm -rf $(BINARY_NAME)
	@go clean
	@echo "Clean complete"

build-linux:
	@echo "Building for Linux..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)
	@echo "Linux build complete: $(BUILD_DIR)/$(BINARY_NAME)"

build-macos:
	@echo "Building for macOS..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)
	@echo "macOS build complete: $(BUILD_DIR)/$(BINARY_NAME)"

.PHONY: build clean deps build-linux build-macos