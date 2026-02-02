.PHONY: run build test migrate-up migrate-down

run:
	go run cmd/api/main.go

build:
	go build -o bin/api cmd/api/main.go

test:
	go test ./...

migrate-up:
	@echo "Running migrations..."
	@psql -U postgres -d ecolife -f migrations/001_init_schema.up.sql || echo "Migration already applied or database not available"

migrate-down:
	@echo "Rolling back migrations..."
	@psql -U postgres -d ecolife -f migrations/001_init_schema.down.sql || echo "Migration rollback failed or database not available"

