FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /tmp/server-service ./server-service/cmd

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /tmp/server-service ./server-service

EXPOSE 8088

CMD ["./server-service"]
