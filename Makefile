# Simple Makefile for a Go project
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

build-client:
	@echo "-- Building..."
	@cd ui; yarn run build;
	@echo "-- Build done. You can now serve the build using 'yarn run preview'"

build-api:
	@echo "-- Building..."
	@cd api; go build -o bin/main cmd/main.go
	@echo "-- Build done. You can find the binary in api/bin/"

build-api-win:
	@echo "Building app for Windows platform..."
	@cd api; go build -o bin/main.exe cmd/main.go
	@echo "-- Build done. You can find the executable in api/bin/"

dev-server: docker-up dev-api

# Run the Go application
dev-api:
	@cd api; go run cmd/main.go

dev-client:
	@cd ui; npm run dev

gen-proto: gen-proto-client gen-proto-api

# Generation through protobuf-ts library which is more up to date than grpc-web
gen-proto-client:
	@rm -rf ui/src/lib/proto
	@mkdir -p ui/src/lib/proto
	@cd ui; npx protoc --ts_out src/lib/proto --proto_path $(ROOT_DIR)/proto $(ROOT_DIR)/proto/providers_app_service.proto
	@echo "-- Protocol buffer messages compiled for client"

# Generate protocol buffer files for Go (API)
gen-proto-api:
	@protoc \
	--go_out=api \
	--go_opt=paths=source_relative \
    --go-grpc_out=api \
	--go-grpc_opt=paths=source_relative \
    proto/providers_app_service.proto
	@echo "-- Protocol buffer messages compiled for API"

# Create DB container
docker-up:
	@if docker compose -f docker/docker-compose.yml --env-file api/.env up -d --build 2>/dev/null; then \
		: ; \
	else \
		echo "-- Falling back to Docker Compose V1"; \
		docker-compose -f docker/docker-compose.yml --env-file api/.env up -d --build; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose -f docker/docker-compose.yml down 2>/dev/null; then \
		: ; \
	else \
		echo "-- Falling back to Docker Compose V1"; \
		docker-compose -f docker/docker-compose.yml down; \
	fi

# Test the application
test-api:
	@echo "-- Running all tests..."
	@cd api; go test ./internal/... -v

# Integrations Tests for the application
itest-api:
	@echo "-- Running integration tests..."
	@cd api; go test ./internal/database -v

# Clean the binary
clean-bin:
	@echo "-- Cleaning binaries..."
	@rm -f api/bin/main


# Live reload (Windows)
watch-api-win:
	cd api; ./watch.sh win;

# Live reload (Linux / OS X)
watch-api:
	cd api; ./watch.sh unix;

.PHONY: all-api dev-up-api build-api dev-api test-api clean-api watch-api itest-api docker-up docker-down watch-api-win
