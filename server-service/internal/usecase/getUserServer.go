package usecase

import (
	"AuthService/server-service/internal/domain"
	"AuthService/server-service/internal/repository"
	"context"
)

type GetUserServerConfig struct {
	Repository repository.UserServersRepository
}

type GetUserServerUC struct {
	config GetUserServerConfig
}

func NewGetUserServerUC(
	conf GetUserServerConfig,
) *GetUserServerUC {
	return &GetUserServerUC{
		config: conf,
	}
}

func (uc *GetUserServerUC) GetUserServer(
	ctx context.Context,
	userID int,
	serverID int,
) (*domain.UserServer, error) {
	return uc.config.Repository.GetUserServer(ctx, userID, serverID)
}
