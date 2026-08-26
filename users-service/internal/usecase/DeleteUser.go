package usecase

import (
	"context"
	"reg/users-service/internal/repository"
)

type DeleteConfig struct {
	UsersRepo repository.UsersRepository
}

type DeleteUC struct {
	config DeleteConfig
}

func NewDeleteUC(conf DeleteConfig) *DeleteUC {
	return &DeleteUC{
		config: conf,
	}
}

func (u *DeleteUC) DeleteUser(
	ctx context.Context, 
	userId uint64,
) error {
	err := u.config.UsersRepo.Delete(ctx, userId)

	if err != nil {
		return err
	}

	return nil
}