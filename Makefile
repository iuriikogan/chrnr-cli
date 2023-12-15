# Use the Go compiler
GO := go

# Binary name
BINARY := chrnr-cli

# Build target
# This target compiles the main.go file and outputs the binary to the ./bin directory
build:
    $(GO) build -o ./bin/$(BINARY) ./main.go

# Run target
# This target runs the main.go file
run:
    $(GO) run ./main.go \

# Test target
# This target runs all the tests in the project with verbose output and coverage report
test:
    $(GO) test -v ./... -cover

# Benchmark test target
# This target runs all the benchmark tests in the project and generates CPU and memory profiles
benchmark:
    $(GO) test -bench=. ./... -benchmem -cpuprofile cpu.profile -memprofile mem

# This target builds the main.go file and outputs the binary to the bin directory
.PHONY: build run test benchmark
    $(GO) build -o bin/$(BINARY) src/main.go

# Clean target
# This target removes the generated binary, memory profile, and test files
clean: 
    rm -rf bin/$(BINARY)
    rm -rf mem.profile
    rm -rf /app/testfiles/