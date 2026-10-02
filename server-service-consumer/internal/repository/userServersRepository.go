package repository

import (
	"AuthService/server-service-consumer/internal/domain"
	"context"
)

type UserServersRepository interface {
	Create(ctx context.Context, userServer *domain.UserServer) error
	CheckUserHasServer(ctx context.Context, userServer *domain.UserServer) (bool, error) // true - если у user есть сервера с там же ip
}