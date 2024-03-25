FROM golang:1.21.6 as builder

WORKDIR /fizzbuzz-app

COPY . .

RUN go mod download

RUN go install ./api/cmd/fizzbuzz-api

FROM debian:bookworm-slim

WORKDIR /fizzbuzz-app

COPY --from=builder /go/bin/fizzbuzz-api .
COPY .env ./.env

CMD ["/fizzbuzz-app/fizzbuzz-api"]
