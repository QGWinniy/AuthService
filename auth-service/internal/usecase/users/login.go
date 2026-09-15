package usecase_users

import (
	"context"
	"AuthService/auth-service/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

type LoginConfig struct {
	UserGateway UserGateway
}

type LoginUC struct {
	Config LoginConfig
}

func NewLoginUC(conf LoginConfig) *LoginUC {
	return &LoginUC{
		Config: conf,
	}
}

func Compare(password string, hash string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
}

func (u *LoginUC) Login(
	ctx context.Context,
	login string,
	password string,
) (*domain.User, error) {

	user, err := u.Config.UserGateway.GetUserByLogin(
		ctx,
		login,
	)
	if err != nil {
		return nil, err
	}

	if err := Compare(
		password,
		user.PasswordHash,
	); err != nil {
		return nil, err
	}

	return user, nil
}