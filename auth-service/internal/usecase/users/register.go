package usecase_users

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type RegisterConfig struct {
	UserGateway UserGateway
}

type RegisterUC struct {
	config RegisterConfig
}

func NewRegisterUC(conf RegisterConfig) *RegisterUC {
	return &RegisterUC{
		config: conf,
	}
}

func Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (u *RegisterUC) Register(
	ctx context.Context,
	login string,
	password string,
) (uint64, error) {
	passwordHash, err := Hash(password)

	if err != nil {
		return 0, err
	}

	userID, err := u.config.UserGateway.CreateUser(
        ctx,
        login,
        passwordHash,
    )

    if err != nil {
        return 0, err
    }

	return userID, err
}