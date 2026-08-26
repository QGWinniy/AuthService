package usecase

import (
	"context"
	"reg/users-service/internal/domain"
	"reg/users-service/internal/repository"
)

type CreteConfig struct {
	UsersRepo repository.UsersRepository
}

type CreateUC struct {
	config CreteConfig
}

func NewCreateUC(conf CreteConfig) *CreateUC {
	return &CreateUC{
		config: conf,
	}
}

func (u *CreateUC) CreateUser(
	ctx context.Context, 
	user *domain.User,
) (uint64, error) {
	UserId, err := u.config.UsersRepo.Create(ctx, user);

	if err != nil {
		return 0, err
	}

	return UserId, nil
}