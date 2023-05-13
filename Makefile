clean:
	@go clean
	@rm -Rf ./bin

build: clean
	@env GOARCH=arm64 GOOS=linux go build -ldflags="-s -w" -o bin/pismo-challenge src/cmd/api/main.go

test:
	@go test -v ./...

format:
	@go fmt ./...

run-migrations:
	@/go/bin/goose -dir=migrations postgres "host=${PG_HOST} user=${PG_USER} password=${PG_PASSWORD} dbname=${PG_DATABASE} sslmode=disable" up

run-dev:
	@docker-compose up --build