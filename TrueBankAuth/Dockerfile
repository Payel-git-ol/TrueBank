FROM golang:1.25.4 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o auth-service ./cmd/app

FROM ubuntu:22.04
WORKDIR /app
COPY --from=builder /app/auth-service .

EXPOSE 7070
EXPOSE 50051

CMD ["./auth-service"]
