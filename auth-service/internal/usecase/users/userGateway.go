package usecase_users

import (
	"context"
	"AuthService/auth-service/internal/domain"
)

type UserGateway interface {
	CreateUser(
		ctx context.Context,
		login string,
		passwordHash string,
	) (uint64, error)

	GetUserByLogin(
		ctx context.Context,
		login string,
	) (*domain.User, error)
}