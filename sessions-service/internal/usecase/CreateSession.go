package usecase

import (
	"context"
	"time"

	"AuthService/sessions-service/internal/config"
	"AuthService/sessions-service/internal/domain"
	"AuthService/sessions-service/internal/repository"

	"AuthService/sessions-service/pkg/jwt"
)

type CreateSessionConfig struct {
	SessionRepo repository.SessionRepository
	JWTManager  *jwt.Manager
}

type CreateSessionUC struct {
	config CreateSessionConfig
}

func NewCreateSessionUC(conf CreateSessionConfig) *CreateSessionUC {
	return &CreateSessionUC{
		config: conf,
	}
}

func (u *CreateSessionUC) CreateSession(
	ctx context.Context,
	userID uint64,
) (*domain.Session, error) {
	now := time.Now()
	expiresAt := now.Add(config.SessionDuration)

	token, jti, err := u.config.JWTManager.Create(
		userID,
		expiresAt,
	)
	if err != nil {
		return nil, err
	}

	session := &domain.Session{
		SessionId: jti,
		UserId:    userID,
		ExpiresAt: expiresAt,
		Jwt:     token,
	}

	if err := u.config.SessionRepo.Create(ctx, session); err != nil {
		return nil, err
	}

	return session, nil
}