FROM golang:1.25.4 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o user-service ./cmd/app

FROM ubuntu:22.04
WORKDIR /app
COPY --from=builder /app/user-service .

EXPOSE 5050
EXPOSE 50052

CMD ["./user-service"]
