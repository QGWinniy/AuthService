package usecase

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"reg/sessions-service/internal/config"
	"reg/sessions-service/internal/domain"
	"reg/sessions-service/internal/repository"
	"time"
)

type CreateSessionConfig struct {
	SessionRepo repository.SessionRepository
}

type CreateSessionUC struct {
	config CreateSessionConfig
}

func NewCreateSessionUC(conf CreateSessionConfig) *CreateSessionUC {
	return &CreateSessionUC{
		config: conf,
	}
}

func RandomSessionID(n int) (string, error) {
    b := make([]byte, n)

    if _, err := rand.Read(b); err != nil {
        return "", err
    }

    return base64.RawURLEncoding.EncodeToString(b), nil
}

func (u *CreateSessionUC) CreateSession(
	ctx context.Context,
	userId uint64,
) (*domain.Session, error) {
	sessionId, err := RandomSessionID(config.SessionIdLength)
	if err != nil {
		return nil, err
	}

	session := &domain.Session{
		SessionId: sessionId,
		UserId: userId,
		ExpiresAt: time.Now().Add(config.SessionDuration),
	}

	err = u.config.SessionRepo.Create(ctx, session)

	if err != nil {
		return nil, err
	}

	return  session, nil
}



