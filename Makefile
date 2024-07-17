# ==================================================================================== #
#  MAKE FILE INSTRUCTION BY MAZELI VICTOR CHIDUBEM
#  01/05/2024
# ==================================================================================== #

DB_DRIVER := postgres
DB_URL := postgresql://root:secret@localhost:5467/beds_core?sslmode=disable
DB_MIGRATIONS_DIR :=./internal/db/migrations
DB_MIGRATE_EXEC := migrate
DB_MIGRATE_ARGS := -database $(DB_DRIVER) -url $(DB_URL) -path $(DB_MIGRATIONS_DIR) -verbose

.PHONY: help
help: ## Display this help message
	@echo "Usage: make <target>"
	@echo
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


# ==================================================================================== #
# SQL MIGRATIONS
# ==================================================================================== #

.PHONY: create-migration
create-migration: ## Create a new migration file for a table, e.g., make create-migration table_name=bug_report
ifdef table_name
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest create -ext sql -dir $(DB_MIGRATIONS_DIR) -seq $(table_name)
else
	@echo "Please provide a table_name argument, e.g., make create-migration table_name=bug_report"
endif

.PHONY: migrate-force
migrate-force:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -database=$(DB_URL) -path=$(DB_MIGRATIONS_DIR)  force ${version}


.PHONY: migrate-up
migrate-up: ## Migrate the database schema up to the latest version
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -database=$(DB_URL) -path=$(DB_MIGRATIONS_DIR)  up

.PHONY: migrate-down
migrate-down: ## Rollback the database schema by one migration
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -database=$(DB_URL) -path=$(DB_MIGRATIONS_DIR) down

.PHONY: migrate-drop
migrate-drop:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -database=$(DB_URL)  -path=$(DB_MIGRATIONS_DIR) drop -f

.PHONY: migrate-to
migrate-to: ## Migrate the database schema to a specific version, e.g., make migrate-to version=1
ifdef version
	 go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -database=$(DB_URL) -path=$(DB_MIGRATIONS_DIR)  goto $(version)
else
	@echo "Please provide a version argument, e.g., make migrate-to version=1"
endif

# ==================================================================================== #
# RUN TESTS
# ==================================================================================== #

.PHONY: test
test:
	@GOFLAGS="-count=1" go test -v -cover -race -vet=off ./...

# ==================================================================================== #
# CODE QUALITY CHECK
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

.PHONY: audit
audit: 
	go vet ./...
	golangci-lint run --fix
	go mod verify

.PHONY: lint
lint:
	golangci-lint run

.PHONY: lint-fix
lint-fix:
	golangci-lint run --fix


# ==================================================================================== #
# SQLC ORM CLI
# ==================================================================================== #
.PHONY: sqlc-gen
sqlc-gen:
	sqlc generate


# ==================================================================================== #
# BUILD
# ==================================================================================== #

.PHONY: docker-compose-up
docker-compose-up:
	docker compose down
	docker compose up -d --build

.PHONY: docker-compose-down
docker-compose-down:
	docker compose down


.PHONY: build-artifact
build-artifact:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api ./cmd/api/.


# ==================================================================================== #
# MOCK GEN
# ==================================================================================== #
mock:
	mockgen -package mockdb -destination internal/db/mock/store.go github.com/joey1123455/beds-api/internal/db/sqlc Store