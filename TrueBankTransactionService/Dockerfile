FROM golang:1.25.4 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o transaction-service ./cmd/app

FROM ubuntu:22.04
WORKDIR /app
COPY --from=builder /app/transaction-service .

EXPOSE 6060
EXPOSE 50053

CMD ["./transaction-service"]
