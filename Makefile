# Configurations

BIN_NAME=app

build:
	@echo "Starting building the api..."
	@go build -o ./bin/$(BIN_NAME) ./cmd/api/

run: build
	@./bin/$(BIN_NAME)