clean:
	@go clean
	@rm -Rf ./bin

build: clean
	@env GOARCH=arm64 GOOS=linux go build -ldflags="-s -w" -o bin/post-account src/functions/post-account/main.go
	@env GOARCH=arm64 GOOS=linux go build -ldflags="-s -w" -o bin/get-account src/functions/get-account/main.go
	@env GOARCH=arm64 GOOS=linux go build -ldflags="-s -w" -o bin/post-transaction src/functions/post-transaction/main.go

test:
	@go test -v ./...

format:
	@go fmt ./...

run-dev:
	@serverless offline --useDocker start --host 0.0.0.0