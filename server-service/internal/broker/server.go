package broker

import (
	"AuthService/server-service/internal/domain"
	"context"
)

type ServerBrokerPublish interface {
	PublishServerCreationRequested(
		ctx context.Context,
		userServer *domain.UserServer,
	) error
}