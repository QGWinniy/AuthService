package usecase

import (
	"context"
	"reg/users-service/internal/domain"
	"reg/users-service/internal/repository"
)

type GetUserByLoginConfig struct {
	UsersRepo repository.UsersRepository
}

type GetUserByLoginUC struct {
	config GetUserByLoginConfig
}

func NewGetUserByLoginUC(conf GetUserByLoginConfig) *GetUserByLoginUC {
	return &GetUserByLoginUC{
		config: conf,
	}
}

func (u *GetUserByLoginUC) GetUserByLogin(
	ctx context.Context,
	login string,
) (*domain.User, error) {

	user, err := u.config.UsersRepo.GetByLogin(
		ctx,
		login,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}