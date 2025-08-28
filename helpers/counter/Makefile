.PHONY: all clean tests lint build format coverage benchmarks

all: build tests lint

# Run tests with coverage
tests:
	@echo "Running tests with coverage..."
	go test -v -race -covermode=atomic -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out || true
	@if command -v go tool cover >/dev/null 2>&1; then \
		go tool cover -html=coverage.out -o coverage.html; \
	fi

# Run benchmarks
benchmarks:
	@echo "Running benchmarks..."
	go test -run=^$$ -bench=. -benchmem ./...

# Build the package
build:
	@echo "Building package..."
	go build ./...

# Format code: fmt, imports (if available), and gofmt -s
format:
	@echo "Formatting code..."
	@gofmt -s -w .
	@if command -v goimports >/dev/null 2>&1; then \
		goimports -w .; \
	else \
		echo "goimports not installed, skipping import reordering"; \
	fi
	@if command -v golines >/dev/null 2>&1; then \
		golines -w .; \
	else \
		echo "golines not installed, skipping line wrapping"; \
	fi

# Lint code if golangci-lint is available
lint:
	@echo "Linting code..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed, skipping lint"; \
	fi

# Generate coverage report (HTML)
coverage: tests
	@go tool cover -html=coverage.out

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@find . -type f -name "*.test" -delete
	@rm -f coverage.out coverage.html
	@rm -rf tmp
