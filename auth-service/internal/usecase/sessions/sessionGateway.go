package usecase_sessions

import (
	"context"
	"AuthService/auth-service/internal/domain"
)

type SessionGateway interface {
	CreateSession(
		ctx context.Context,
		userID uint64,
	) (*domain.Session, error)

	GetSession(
		ctx context.Context,
		sessionID string,
	) (*domain.Session, error)
}

