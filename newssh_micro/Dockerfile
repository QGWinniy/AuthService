FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o newssh ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/newssh .
COPY --from=builder /app/internal/wsserver/static ./internal/wsserver/static

EXPOSE 8080

CMD ["./newssh"]