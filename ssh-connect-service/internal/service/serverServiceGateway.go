package service

import (
	"AuthService/ssh-connect-service/internal/domain"
	"context"
)

type ServerServiceGateway interface {
	GetUserServerByID(ctx context.Context, serverId int, JWT string) (*domain.UserServer, error)  // тут будет проверка что сервер принадлежит юзеру
}