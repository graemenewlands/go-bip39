# Define the output directory for the binaries
BIN_DIR=bin

# Build flags (if any)
BUILD_FLAGS=

# Build the first Go program
build-mnemonic:
	@echo "Building mnemonic..."
	@mkdir -p $(BIN_DIR)
	go build $(BUILD_FLAGS) -o $(BIN_DIR)/mnemonic cmd/mnemonic/main.go

# Build the second Go program
build-ed209:
	@echo "Building ed209..."
	@mkdir -p $(BIN_DIR)
	go build $(BUILD_FLAGS) -o $(BIN_DIR)/ed209 cmd/ed25519/main.go

# Build both programs
build: build-mnemonic build-ed209

# Clean up binaries
clean:
	@echo "Cleaning up..."
	rm -rf $(BIN_DIR)

# Format the Go source code
fmt:
	go fmt ./...

# Vet the Go source code
vet:
	go vet ./...

# Tidy up Go modules
tidy:
	go mod tidy

# Test the Go source code
test:
	go test ./...

# Phony targets are not actual files
.PHONY: build build-app1 build-app2 run-app1 run-app2 clean fmt vet tidy test