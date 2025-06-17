# Makefile for WebAPI Go project

.PHONY: help build run test clean docker-up docker-down swagger deps

# Default target
help:
	@echo "Available commands:"
	@echo "  deps        - Install dependencies"
	@echo "  swagger     - Generate Swagger documentation"
	@echo "  build       - Build the application"
	@echo "  run         - Run the application"
	@echo "  test        - Run tests"
	@echo "  docker-up   - Start PostgreSQL with Docker Compose"
	@echo "  docker-down - Stop PostgreSQL containers"
	@echo "  clean       - Clean build artifacts"

# Install dependencies
deps:
	go mod tidy
	go mod download

# Generate Swagger documentation
swagger:
	swag init

# Build the application
build:
	go build -o bin/webapi main.go

# Run the application
run:
	go run main.go

# Run tests
test:
	go test ./...

# Start PostgreSQL with Docker Compose
docker-up:
	docker-compose up -d

# Stop PostgreSQL containers
docker-down:
	docker-compose down

# Clean build artifacts
clean:
	rm -rf bin/
	rm -rf docs/
