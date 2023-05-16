FROM gomicro/goose as migrations

WORKDIR /usr/app

COPY ./migrations ./migrations

CMD goose -dir=migrations postgres "host=$POSTGRES_HOST user=$POSTGRES_USER password=$POSTGRES_PASSWORD dbname=$POSTGRES_DB sslmode=disable" up

FROM golang:1.20.4-buster as app

WORKDIR /usr/app

COPY . .

RUN go mod tidy

RUN go build -o ./bin/api ./src/cmd/api/main.go

CMD ["./bin/api"]