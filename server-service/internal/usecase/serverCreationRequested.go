package usecase

import (
	"AuthService/server-service/internal/broker"
	"AuthService/server-service/internal/domain"
	"context"
)

type ServerCreationRequestedConfig struct {
	Broker broker.ServerBrokerPublish
}

type ServerCreationRequestedUC struct {
	config ServerCreationRequestedConfig
}

func NewServerCreationRequestedUC (
	conf ServerCreationRequestedConfig,
) *ServerCreationRequestedUC {
	return &ServerCreationRequestedUC{
		config: conf,
	}
}

func (uc *ServerCreationRequestedUC) ServerCreationRequested (
	ctx context.Context,
	userServer *domain.UserServer,
) error {
	// тут проверяем всякую хуйню можно почитать что нибудь из бд
	err := uc.config.Broker.PublishServerCreationRequested(
		ctx,
		userServer,
	)
	return err
}

