FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o loadtester ./main.go

FROM alpine:latest

COPY --from=builder /app/loadtester /loadtester

ENTRYPOINT ["/loadtester"]
