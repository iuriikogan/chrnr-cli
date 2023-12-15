# Go compiler
GO := go

# Binary name
BINARY := chrnr-cli

# Build target
build:
	$(GO) build -o ./bin/$(BINARY) ./main.go

# Run target
run:
	$(GO) run ./main.go \

# Test target
test:
	$(GO) test -v ./... -cover

# Benchmark test target
benchmark:
	$(GO) test -bench=. ./... -benchmem -cpuprofile cpu.profile -memprofile mem

clean: 
	rm -rf bin/$(BINARY)
	rm -rf mem.profile
	rm -rf app/

.PHONY: build run test benchmark
	$(GO) build -o bin/$(BINARY) src/main.go

