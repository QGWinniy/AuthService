package usecase

import (
	"context"
	"AuthService/sessions-service/internal/domain"
	"AuthService/sessions-service/internal/repository"
)

type GetSessionConfig struct {
	SessionRepo repository.SessionRepository
}

type GetSessionUC struct {
	config GetSessionConfig
}

func NewGetSessionUC(conf GetSessionConfig) *GetSessionUC {
	return &GetSessionUC{
		config: conf,
	}
}

func (u *GetSessionUC) GetSession(
	ctx context.Context,
	session_id string,
) (*domain.Session, error) {
	session, err := u.config.SessionRepo.Get(ctx, session_id)

	if err != nil {
		return nil, err
	}

	return session, err
}



