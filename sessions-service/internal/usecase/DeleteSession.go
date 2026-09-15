package usecase

import (
	"context"
	"AuthService/sessions-service/internal/repository"
)

type DeleteSessionConfig struct {
	SessionRepo repository.SessionRepository
}

type DeleteSessionUC struct {
	config DeleteSessionConfig
}

func NewDeleteSessionUC(conf DeleteSessionConfig) *DeleteSessionUC {
	return &DeleteSessionUC{
		config: conf,
	}
}

func (u *DeleteSessionUC) DeleteSession(
	ctx context.Context, 
	sessionId string,
) error {
	err := u.config.SessionRepo.Delete(ctx, sessionId)

	if err != nil {
		return err
	}

	return nil
}