FROM golang:1.27-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /kv-server ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /kv-server /app/kv-server

EXPOSE 50051 50052 50053

ENTRYPOINT ["/app/kv-server"]
