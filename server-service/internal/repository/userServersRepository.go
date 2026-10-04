package repository

import (
	"AuthService/server-service/internal/domain"
	"context"
)

type UserServersRepository interface {
	GetUserServers(ctx context.Context, userId int64) ([]domain.UserServer, error)
	GetUserServer(ctx context.Context, id int) (*domain.UserServer, error)
}
