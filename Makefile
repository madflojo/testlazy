.PHONY: all clean tests lint build format benchmarks

all: build tests lint

COMPONENTS = things/testurl helpers/counter

# Run tests for all modules
tests:
	@echo "Running tests for all modules..."
	@for dir in $(COMPONENTS); do \
		$(MAKE) -C $$dir tests || exit 1; \
	done

# Run benchmarks for all modules
benchmarks:
	@echo "Running benchmarks for all modules..."
	@for dir in $(COMPONENTS); do \
		$(MAKE) -C $$dir benchmarks || exit 1; \
	done

# Build all modules
build:
	@echo "Building all modules..."
	@for dir in $(COMPONENTS); do \
		$(MAKE) -C $$dir build || exit 1; \
	done

# Format all code
format:
	@echo "Formatting code..."
	@gofmt -s -w .
	@if command -v golines >/dev/null 2>&1; then \
		golines -w .; \
	else \
		echo "golines not installed, skipping line wrapping"; \
	fi
	@for dir in $(COMPONENTS); do \
		$(MAKE) -C $$dir format || exit 1; \
	done

# Lint all code
lint:
	@echo "Linting code..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
		for dir in $(COMPONENTS); do \
			$(MAKE) -C $$dir lint || exit 1; \
		done; \
	else \
		echo "golangci-lint not installed, skipping lint"; \
	fi

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@for dir in $(COMPONENTS); do \
		$(MAKE) -C $$dir clean || exit 1; \
	done
	@find . -type f -name "*.test" -delete
	@find . -type f -name "coverage.out" -delete
	@find . -type f -name "coverage.html" -delete
	@find . -type d -name "vendor" -exec rm -rf {} + 2>/dev/null || true
	@rm -rf bin/
