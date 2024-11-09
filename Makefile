# Simple Makefile for a Go project

build-api:
	@echo "Building..."
	@cd api; go build -o bin/main cmd/main.go
	@echo "Done! You can find the binary in api/bin/"
	
build-api-win:
	@echo "Building app for Windows platform..."
	@cd api; go build -o bin/main.exe cmd/main.go
	@echo "Done! You can find the executable in api/bin/"

# Run the Go application
dev-api:
	@cd api; go run cmd/main.go

# Generate protocol buffer files for Go (API)
gen-grpc:
	protoc \
	--go_out=api \
	--go_opt=paths=source_relative \
    --go-grpc_out=api \
	--go-grpc_opt=paths=source_relative \
    proto/traffic_lights_service.proto

# Create DB container
docker-up:
	@if docker compose -f api/docker/docker-compose.yml --env-file api/.env up --build 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose -f api/docker/docker-compose.yml --env-file api/.env up --build; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose -f api/docker/docker-compose.yml down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose -f api/docker/docker-compose.yml down; \
	fi

# Test the application
test-api:
	@echo "Running all tests..."
	@cd api; go test ./internal/... -v

# Integrations Tests for the application
itest-api:
	@echo "Running integration tests..."
	@cd api; go test ./internal/database -v

# Clean the binary
clean-bin:
	@echo "Cleaning binaries..."
	@rm -f api/bin/main


# Live reload (Windows)
watch-api-win:
	cd api; ./watch.sh win;

# Live reload (Linux / OS X)
watch-api:
	cd api; ./watch.sh unix;

.PHONY: all-api dev-up-api build-api dev-api test-api clean-api watch-api itest-api docker-up docker-down watch-api-win
