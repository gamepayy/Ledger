# syntax=docker/dockerfile:1

FROM golang:1.19

RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY .env ./.env

COPY database ./database
COPY docs ./docs
COPY middleware ./middleware
COPY routers ./routers
COPY sql ./sql

RUN go build -o gamepayy_ledger .

ENV HOST 0.0.0.0
ENV PORT 80
EXPOSE 80

CMD ["./gamepayy_ledger"]
