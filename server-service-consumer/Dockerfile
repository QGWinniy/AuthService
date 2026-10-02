FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /tmp/server-service-consumer ./server-service-consumer/cmd

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /tmp/server-service-consumer ./server-service-consumer

EXPOSE 8089

CMD ["./server-service-consumer"]
