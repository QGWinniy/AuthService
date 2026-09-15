package usecase_sessions

import (
	"context"
	"AuthService/auth-service/internal/domain"
)

type CreateConfig struct {
	SessionGateway SessionGateway
}

type CreateSessionUC struct {
	config CreateConfig
}

func NewCreateSessionUC(
	conf CreateConfig,
) *CreateSessionUC {
	return &CreateSessionUC{
		config: conf,
	}
}

func (u *CreateSessionUC) Create(
	ctx context.Context,
	userID uint64,
) (*domain.Session, error) {

	session, err := u.config.SessionGateway.CreateSession(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}

	return session, nil
}