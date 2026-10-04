package usecase

import (
	"AuthService/server-service/internal/domain"
	"AuthService/server-service/internal/repository"
	"context"
)

type GetUserServersConfig struct {
	Repository repository.UserServersRepository
}

type GetUserServersUC struct {
	config GetUserServersConfig
}

func NewGetUserServersUC(
	conf GetUserServersConfig,
) *GetUserServersUC {
	return &GetUserServersUC{
		config: conf,
	}
}

func (uc *GetUserServersUC) GetUserServers(
	ctx context.Context,
	userID int64,
) ([]domain.UserServer, error) {
	return uc.config.Repository.GetUserServers(ctx, userID)
}
