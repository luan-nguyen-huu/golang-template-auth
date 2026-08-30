.PHONY: help run dev build test migrate tidy docker-up docker-down clean

help:
	@echo "Available commands:"
	@echo "  make run         - Run application directly"
	@echo "  make dev         - Run application with Air hot reload"
	@echo "  make migrate     - Run database migrations"
	@echo "  make build       - Build binaries into bin/ directory"
	@echo "  make test        - Run all unit & integration tests"
	@echo "  make tidy        - Download and tidy Go dependencies"
	@echo "  make docker-up   - Start PostgreSQL and Redis via Docker Compose"
	@echo "  make docker-down - Stop Docker Compose containers"
	@echo "  make clean       - Remove built binaries and temporary files"

run:
	go run cmd/main/main.go

dev:
	./dev.sh

migrate:
	go run cmd/migrate/main.go

build:
	@mkdir -p bin
	go build -o bin/api cmd/main/main.go
	go build -o bin/migrate cmd/migrate/main.go

test:
	go test -v -race ./...

tidy:
	go mod tidy

docker-up:
	docker compose up -d

docker-down:
	docker compose down

clean:
	rm -rf bin/ tmp/
