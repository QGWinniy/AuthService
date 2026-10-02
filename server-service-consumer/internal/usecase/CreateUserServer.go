package usecase

import (
	"AuthService/server-service-consumer/internal/domain"
	"AuthService/server-service-consumer/internal/repository"
	"context"
	"fmt"
)

type CreateUserServerConfig struct {
	UserServerRepository repository.UserServersRepository
}

type CreateUserServerUC struct {
	config CreateUserServerConfig
}

func NewCreateUserServerUC(
	conf CreateUserServerConfig,
) *CreateUserServerUC {
	return &CreateUserServerUC{
		config: conf,
	}
}

func (uc *CreateUserServerUC) CreateUserServer(
	ctx context.Context,
	userServer *domain.UserServer,
) error {

	hasUserServer, err := uc.config.UserServerRepository.CheckUserHasServer(ctx, userServer)

	if err != nil {
		return err
	}

	if hasUserServer == true {
		return fmt.Errorf("user has server with this id")
	}

	err = uc.config.UserServerRepository.Create(ctx, userServer)

	return err
}
