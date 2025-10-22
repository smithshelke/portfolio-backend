PROJECT_ROOT := $(shell pwd)

.PHONY: dev migrations migrate down clean

dev:
	@echo "Starting development environment..."
	docker compose up -d

migrations:
	@echo "Generating new migration..."
	@read -p "Enter migration name: " MIGRATION_NAME; \
	docker run --rm --network portfolio-backend_default -v $(PROJECT_ROOT):/src -w /src arigaio/atlas --env dev migrate diff $(MIGRATION_NAME) --dir file://db/migrations --to file://db/schema

migrate:
	@echo "Applying migrations..."
	docker run --rm --network portfolio-backend_default -v $(PROJECT_ROOT):/src -w /src arigaio/atlas --env dev migrate apply --dir file://db/migrations

down:
	@echo "Stopping development environment..."
	docker compose down

clean:
	@echo "Cleaning up..."
	docker compose down -v
	rm -rf db/migrations/*
	rm -f db/schema/schema.sql

sqlc-generate:
	@echo "Generating sqlc code..."
	docker run --rm -v $(PROJECT_ROOT):/src -w /src sqlc/sqlc generate

swagger-docs:
	@echo "Generating Swagger documentation..."
	swag init --output docs --parseDependency --parseInternal