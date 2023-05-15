clean:
	@go clean
	@rm -Rf ./bin

build: clean
	@env GOARCH=arm64 GOOS=linux go build -ldflags="-s -w" -o bin/pismo-challenge src/cmd/api/main.go

integration-test:
	@go test -v ./src/adapters/...

unit-test:
	@go test -v ./src/application/services/...

format:
	@go fmt ./...

run-dev:
	@docker-compose up --build -d

stop-dev:
	@docker-compose down