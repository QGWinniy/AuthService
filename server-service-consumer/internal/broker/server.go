package broker

import (
	"AuthService/server-service-consumer/internal/domain"
	"context"
)

type ServerBrokerPublish interface {
	PublishServerCreationRequested(
		ctx context.Context,
		userServer *domain.UserServer,
	) error
}

type ServerBrokerConsumer interface {
	PublishServerGetRequested(
		ctx context.Context,
	) (*domain.UserServer, error)
}