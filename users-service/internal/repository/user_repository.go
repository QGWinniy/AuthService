package repository

import (
	"context"
	"AuthService/users-service/internal/domain"
)

type UsersRepository interface {
	Create(ctx context.Context, user *domain.User) (uint64, error)
	GetByID(ctx context.Context, id uint64) (*domain.User, error)
	GetByLogin(ctx context.Context, login string) (*domain.User, error)
	Delete(ctx context.Context, id uint64) error
	// CheckAuth(ctx context.Context, user *domain.User) (bool, error)
}