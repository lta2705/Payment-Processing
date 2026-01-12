# Payment Processing System Makefile

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Binary names
BINARY_NAME=payment-processor
BINARY_UNIX=$(BINARY_NAME)_unix
TEST_PRODUCER=test-producer

# Directories
CMD_DIR=./cmd
SERVER_DIR=$(CMD_DIR)/server
TEST_PRODUCER_DIR=$(CMD_DIR)/test-producer

.PHONY: all build clean test deps run-server run-test-producer docker-up docker-down generate-wire

# Default target
all: clean deps generate-wire build

# Build the server application
build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(SERVER_DIR)

# Build test producer
build-test:
	$(GOBUILD) -o $(TEST_PRODUCER) -v $(TEST_PRODUCER_DIR)

# Build all binaries
build-all: build build-test

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -f $(TEST_PRODUCER)

# Download dependencies
deps:
	$(GOMOD) tidy
	$(GOMOD) download

# Generate wire dependencies
generate-wire:
	$(GOCMD) generate ./internal/app

# Run tests
test:
	$(GOTEST) -v ./...

# Run the server
run-server: build
	./$(BINARY_NAME)

# Run test producer
run-test: build-test
	./$(TEST_PRODUCER)

# Docker operations
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

# Development workflow
dev-start: docker-up
	@echo "Waiting for services to be ready..."
	@sleep 10
	@echo "Services ready. You can now run: make run-server"

dev-stop: docker-down

# Build for production
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v $(SERVER_DIR)

# Help
help:
	@echo "Available targets:"
	@echo "  all              - Clean, download deps, generate wire, and build"
	@echo "  build            - Build server application"
	@echo "  build-test       - Build test producer"
	@echo "  build-all        - Build all binaries"
	@echo "  build-linux      - Build for Linux production"
	@echo "  clean            - Clean build files"
	@echo "  deps             - Download dependencies"
	@echo "  generate-wire    - Generate wire dependencies"
	@echo "  test             - Run tests"
	@echo "  run-server       - Build and run server"
	@echo "  run-test         - Build and run test producer"
	@echo "  docker-up        - Start Docker services"
	@echo "  docker-down      - Stop Docker services"
	@echo "  docker-logs      - View Docker logs"
	@echo "  dev-start        - Start development environment"
	@echo "  dev-stop         - Stop development environment"
	@echo "  help             - Show this help"