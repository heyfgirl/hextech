.PHONY: all
all: vet lint test build

# Go 版本检查
GO_VERSION := $(shell go version | cut -d' ' -f3 | cut -d'.' -f2)
MIN_GO_VERSION := 16

# 检查工具是否安装
HAS_GOLANGCI_LINT := $(shell which golangci-lint 2>/dev/null)
HAS_SWAG := $(shell which swag 2>/dev/null)

.PHONY: check-go-version
check-go-version:
	@if [ $(GO_VERSION) -lt $(MIN_GO_VERSION) ]; then \
		echo "Required Go version >= 1.$(MIN_GO_VERSION). Current version: $(GO_VERSION)"; \
		exit 1; \
	fi

.PHONY: install-tools
install-tools:
	@echo "Installing required tools..."
	@if [ -z "$(HAS_GOLANGCI_LINT)" ]; then \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	@if [ -z "$(HAS_SWAG)" ]; then \
		go install github.com/swaggo/swag/cmd/swag@latest; \
	fi

.PHONY: vet
vet:
	@echo "Running go vet..."
	@go vet ./...

.PHONY: lint
lint:
	@echo "Running golangci-lint..."
	@if [ -z "$(HAS_GOLANGCI_LINT)" ]; then \
		make install-tools; \
	fi
	@golangci-lint run ./...

.PHONY: test
test:
	@echo "Running tests..."
	@go test -v -race -cover ./...

.PHONY: docs
docs:
	@echo "Generating Swagger documentation..."
	@if [ -z "$(HAS_SWAG)" ]; then \
		make install-tools; \
	fi
	@swag init --parseDependency --parseInternal

.PHONY: build
build:
	@echo "Building application..."
	@go build -o main .

.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -f main
	@rm -rf docs/docs.go docs/swagger.json docs/swagger.yaml

.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make all          - Run all checks and build"
	@echo "  make vet         - Run go vet"
	@echo "  make lint        - Run golangci-lint"
	@echo "  make test        - Run tests"
	@echo "  make docs        - Generate Swagger documentation"
	@echo "  make build       - Build the application"
	@echo "  make clean       - Clean up build artifacts" 