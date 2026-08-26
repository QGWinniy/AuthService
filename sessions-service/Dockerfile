FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /tmp/sessions-service ./sessions-service/cmd

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /tmp/sessions-service ./sessions-service

EXPOSE 50052

CMD ["./sessions-service"]