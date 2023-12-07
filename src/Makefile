
# Go compiler
GO := go

# Binary name
BINARY := chrnr-cli

# Build target
build:
	$(GO) build -o $(BINARY)

# Run target
run:
	$(GO) run .

# Test target
test:
	$(GO) test -v ./...

# Benchmark test target
benchmark:
	$(GO) test -bench=. ./...

.PHONY: build run test benchmark
