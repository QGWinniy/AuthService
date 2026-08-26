FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /tmp/users-service ./users-service/cmd

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /tmp/users-service ./users-service

EXPOSE 50051

CMD ["./users-service"]