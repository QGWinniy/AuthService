# REG

Микросервисное приложение на Go для регистрации и авторизации пользователей.

## Архитектура

```text
Client
  │
  ▼
auth-service :8080
  │
  ├── gRPC → users-service :50051 → PostgreSQL
  │
  └── gRPC → sessions-service :50052 → PostgreSQL