FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download -x

COPY . .

RUN go build -o /tmp/auth-service ./auth-service/cmd

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /tmp/auth-service ./auth-service

EXPOSE 8080

CMD ["./auth-service"]