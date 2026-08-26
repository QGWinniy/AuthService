package usecase_users

import (
	"context"
	"reg/auth-service/internal/domain"
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