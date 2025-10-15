# Configurations

BIN_NAME=main

build:
	@echo "Starting building the api..."
	@go build -o ./bin/$(BIN_NAME) ./cmd/api/*.go

run: build
	@./bin/$(BIN_NAME)