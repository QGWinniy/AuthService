package transport

import (
	"context"
	"ssh/internal/domain"
)

type ServerServiceGateway interface {
	GetUserServer(ctx context.Context, user domain.User, serverId int) (*domain.UserServer, error)
	CheckUserHasServer(ctx context.Context, userServer *domain.UserServer) (bool, error) // true - если у user есть сервера с там же ip
}