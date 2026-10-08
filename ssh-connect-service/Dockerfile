FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY ssh-connect-service/go.mod ssh-connect-service/go.sum ./
RUN go mod download -x

COPY ssh-connect-service/ ./

RUN go build -o /tmp/ssh-connect-service ./cmd

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /tmp/ssh-connect-service ./ssh-connect-service

EXPOSE 8090

CMD ["./ssh-connect-service"]
