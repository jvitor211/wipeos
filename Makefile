.PHONY: help build test clean lint install dev release

# Default target
help: ## Show this help message
	@echo "Available targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Development
dev: ## Run the application in development mode
	go run main.go

build: ## Build the binary
	go build -ldflags="-s -w" -o bin/wipeOs main.go

test: ## Run tests
	go test -v ./...

test-coverage: ## Run tests with coverage
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

lint: ## Run linter
	golangci-lint run

clean: ## Clean build artifacts
	rm -rf bin/ dist/ coverage.out coverage.html

# Dependencies
deps: ## Download dependencies
	go mod download
	go mod verify

deps-update: ## Update dependencies
	go get -u ./...
	go mod tidy

# Installation
install: ## Install the binary
	go install

install-tools: ## Install development tools
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/goreleaser/goreleaser@latest

# Release
release-snapshot: ## Create a snapshot release
	goreleaser build --snapshot --clean

release: ## Create a full release
	goreleaser release --clean

release-dry: ## Dry run release
	goreleaser release --dry-run

# Docker
docker-build: ## Build Docker image
	docker build -t wipeOs .

docker-run: ## Run in Docker container
	docker run --rm -it wipeOs

# CI/CD helpers
ci-test: ## Run tests for CI
	go test -race -coverprofile=coverage.out -covermode=atomic ./...

ci-lint: ## Run linter for CI
	golangci-lint run --out-format=github-actions

# Format code
fmt: ## Format Go code
	go fmt ./...
	goimports -w .

# Security
security-scan: ## Run security scan
	gosec ./...

# Benchmarks
bench: ## Run benchmarks
	go test -bench=. -benchmem ./...

# Check
check: fmt lint test ## Run all checks (format, lint, test) 