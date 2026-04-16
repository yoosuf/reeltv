.PHONY: build run test test-docker test-coverage docker-up docker-down docker-test-up docker-test-down migrate-up migrate-down seed clean help

# Variables
APP_NAME=reeltv-api
DOCKER_COMPOSE=docker-compose -f deployments/docker-compose.yml
DOCKER_TEST_COMPOSE=docker-compose -f docker-compose.test.yml
GO=go

# Build the application
build:
	cd backend && $(GO) build -o ../bin/$(APP_NAME) ./cmd/api

# Run the application locally
run:
	cd backend && $(GO) run ./cmd/api

# Run tests
test:
	cd backend && $(GO) test -v ./...

# Run tests with Docker test infrastructure
test-docker: docker-test-up
	cd backend && $(GO) test -v ./... -timeout 30s
	$(DOCKER_TEST_COMPOSE) down

# Run tests with coverage
test-coverage:
	cd backend && $(GO) test -v -coverprofile=../coverage.out ./...
	cd backend && $(GO) tool cover -html=../coverage.out -o ../coverage.html

# Start Docker services
docker-up:
	$(DOCKER_COMPOSE) up -d

# Stop Docker services
docker-down:
	$(DOCKER_COMPOSE) down

# Start Docker test services
docker-test-up:
	$(DOCKER_TEST_COMPOSE) up -d
	@echo "Waiting for test services to be ready..."
	@sleep 5

# Stop Docker test services
docker-test-down:
	$(DOCKER_TEST_COMPOSE) down

# View Docker logs
docker-logs:
	$(DOCKER_COMPOSE) logs -f

# Run database migrations (via Go application)
migrate-up:
	cd backend && $(GO) run ./cmd/api migrate up

# Rollback database migrations
migrate-down:
	cd backend && $(GO) run ./cmd/api migrate down

# Seed database with test data
seed:
	cd backend && $(GO) run ./cmd/api seed

# Download dependencies
deps:
	cd backend && $(GO) mod download
	cd backend && $(GO) mod tidy

# Format code
fmt:
	cd backend && $(GO) fmt ./...

# Run linter
lint:
	cd backend && golangci-lint run ./...

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

# Development workflow (start everything)
dev: docker-up deps migrate-up seed
	@echo "Development environment ready!"
	@echo "Run 'make run' to start the API server"

# Help
help:
	@echo "Available commands:"
	@echo "  build           - Build the application"
	@echo "  run             - Run the application locally"
	@echo "  test            - Run tests"
	@echo "  test-docker     - Run tests with Docker test infrastructure"
	@echo "  test-coverage   - Run tests with coverage"
	@echo "  docker-up       - Start Docker services"
	@echo "  docker-down     - Stop Docker services"
	@echo "  docker-test-up  - Start Docker test services"
	@echo "  docker-test-down - Stop Docker test services"
	@echo "  docker-logs     - View Docker logs"
	@echo "  migrate-up      - Run database migrations"
	@echo "  migrate-down    - Rollback database migrations"
	@echo "  seed            - Seed database with test data"
	@echo "  deps            - Download dependencies"
	@echo "  fmt             - Format code"
	@echo "  lint            - Run linter"
	@echo "  clean           - Clean build artifacts"
	@echo "  dev             - Start development environment"
	@echo "  help            - Show this help message"
