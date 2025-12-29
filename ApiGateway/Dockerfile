FROM golang:1.25.4 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o api-gateway ./cmd/app

FROM ubuntu:22.04
WORKDIR /app
COPY --from=builder /app/api-gateway .

EXPOSE 8080

CMD ["./api-gateway"]
