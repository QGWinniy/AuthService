package usecase

import (
	"AuthService/ssh-connect-service/internal/service"
	consoletoken "AuthService/ssh-connect-service/pkg/consoleToken"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type MakeTokenConfig struct {
	RepoUserServer service.ServerServiceGateway
	ConsoleTokenManager *consoletoken.Manager
}

type MakeTokenUC struct {
	config MakeTokenConfig
}

func NewMakeTokenUC(conf MakeTokenConfig) *MakeTokenUC {
	return &MakeTokenUC{
		config: conf,
	}
}

func (u *MakeTokenUC) MakeToken(
	ctx context.Context,
	serverID int,
	jwt string,
) (string, string, error) {
	userServer, err := u.config.RepoUserServer.GetUserServerByID(
		ctx,
		serverID,
		jwt,
	)

	if err != nil {
		return "", "", fmt.Errorf("unauthorized %v", err)
	}

	sessionID := uuid.NewString()

	token, err := u.config.ConsoleTokenManager.Create(
		consoletoken.UserServer{
			ID:       userServer.ID,
			Addr:     userServer.Addr,
			UserName: userServer.UserName,
			Password: userServer.Password,
		},
		sessionID,
		time.Now().Add(time.Minute),
	)
	if err != nil {
		return "", "", fmt.Errorf("failed to create console token %v", err)
	}

	return sessionID, token, nil
}
