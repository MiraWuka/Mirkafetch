# Makefile for Mirkafetch project

# Variables
BINARY_NAME=Mirkafetch
MAIN_PATH=./cmd/Mirkafetch
BUILD_DIR=./bin
GO_FILES=$(shell find . -name "*.go" -type f)

# Build flags
LDFLAGS=-ldflags "-s -w"
BUILD_FLAGS=-trimpath

# Default target
.PHONY: all
all: clean build

# Build the application
.PHONY: build
build: $(BUILD_DIR)/$(BINARY_NAME)

$(BUILD_DIR)/$(BINARY_NAME): $(GO_FILES)
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build $(BUILD_FLAGS) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

# Build for multiple platforms
.PHONY: build-all
build-all: clean
	@echo "Building for multiple platforms..."
	@mkdir -p $(BUILD_DIR)
	# Linux
	GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)
	GOOS=linux GOARCH=arm64 go build $(BUILD_FLAGS) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 $(MAIN_PATH)
	# macOS
	GOOS=darwin GOARCH=amd64 go build $(BUILD_FLAGS) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_PATH)
	GOOS=darwin GOARCH=arm64 go build $(BUILD_FLAGS) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN_PATH)
	# Windows
	GOOS=windows GOARCH=amd64 go build $(BUILD_FLAGS) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)

# Run the application
.PHONY: run
run: build
	$(BUILD_DIR)/$(BINARY_NAME)

# Install the application
.PHONY: install
install:
	go install $(BUILD_FLAGS) $(LDFLAGS) $(MAIN_PATH)

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@go clean

# Run tests
.PHONY: test
test:
	go test -v ./...

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Format code
.PHONY: fmt
fmt:
	go fmt ./...

# Lint code
.PHONY: lint
lint:
	golangci-lint run

# Vet code
.PHONY: vet
vet:
	go vet ./...

# Security check
.PHONY: security
security:
	gosec ./...

# Development workflow
.PHONY: dev
dev: fmt vet test build

# Full check (format, vet, test, build)
.PHONY: check
check: fmt vet lint test build

# Help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build         - Build the application"
	@echo "  build-all     - Build for multiple platforms"
	@echo "  run           - Build and run the application"
	@echo "  install       - Install the application"
	@echo "  clean         - Clean build artifacts"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage"
	@echo "  fmt           - Format code"
	@echo "  lint          - Lint code"
	@echo "  vet           - Vet code"
	@echo "  security      - Run security checks"
	@echo "  dev           - Development workflow (fmt, vet, test, build)"
	@echo "  check         - Full check (fmt, vet, lint, test, build)"
	@echo "  help          - Show this help message"