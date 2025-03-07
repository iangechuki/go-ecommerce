include .envrc
MIGRATIONS_PATH = ./cmd/migrate/migrations

test:
	@go test ./... -v
build:
	@go build -o bin/go-commerce ./cmd/api/*.go
run: build
	@./bin/go-commerce
gen-docs:
	@swag init -g ./api/main.go -d cmd,internal
.PHONY: test

.PHONY: migration
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) up

.PHONY: migrate-down
migrate-down: 
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))
