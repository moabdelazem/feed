include .envrc

# configurations & constants
BIN_NAME=main
MIGRATIONS_PATH=./cmd/migrate/migrations

.PHONY: build
build:
	@echo "Starting building the api..."
	@go build -o ./bin/$(BIN_NAME) ./cmd/api/*.go

run: build
	@./bin/$(BIN_NAME)

# migrations Commands
.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) down

# docker Commands
.PHONY: dev-compose-up
database-compose-up:
	@docker compose -f "compose.dev.yml" up -d 

.PHONY: dev-compose-down
database-compose-down:
	@docker compose -f "compose.dev.yml" down 