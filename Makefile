test:
	@go test ./... -v
build:
	@go build -o bin/go-commerce ./cmd/api/*.go
run: build
	@./bin/go-commerce
gen-docs:
	@swag init -g ./api/main.go -d cmd,internal
.PHONY: test