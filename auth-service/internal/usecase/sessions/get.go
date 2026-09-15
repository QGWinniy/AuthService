package usecase_sessions

import (
	"context"
	"AuthService/auth-service/internal/domain"
)

type GetConfig struct {
	SessionGateway SessionGateway
}

type GetSessionUC struct {
	config GetConfig
}

func NewGetSessionUC(
	conf GetConfig,
) *GetSessionUC {
	return &GetSessionUC{
		config: conf,
	}
}

func (u *GetSessionUC) Get(
	ctx context.Context,
	sessionID string,
) (*domain.Session, error) {

	session, err := u.config.SessionGateway.GetSession(
		ctx,
		sessionID,
	)
	if err != nil {
		return nil, err
	}

	return session, nil
}