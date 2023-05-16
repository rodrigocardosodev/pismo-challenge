FROM golang:1.20.4-buster as dev

WORKDIR /usr/app

COPY . .

RUN go mod tidy

CMD ["go", "run", "./src/cmd/api/main.go"]