package repository

import (
	"context"
	"reg/sessions-service/internal/domain"
)

type SessionRepository interface {
	Create(ctx context.Context, session *domain.Session) (error)
	Get(ctx context.Context, sessionId string) (*domain.Session, error)
	Delete(ctx context.Context, sessionId string) (error)
}