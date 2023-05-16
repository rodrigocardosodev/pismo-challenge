clean:
	@go clean
	@rm -Rf ./bin

build: clean ## Build binary
	@env GOARCH=arm64 GOOS=linux go build -ldflags="-s -w" -o bin/pismo-challenge src/cmd/api/main.go

integration-test: ## Run integration tests
	@go test -v ./src/adapters/...

unit-test: ## Run unit tests
	@go test -v ./src/application/services/...

format: ## Format code
	@go fmt ./...

run-dev: ## Run docker-compose
	@docker-compose up --build -d

stop-dev: ## Stop docker-compose
	@docker-compose down

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'