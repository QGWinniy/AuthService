package userserverpostgres

import (
	"database/sql"
	"ssh/internal/domain"
	"ssh/pkg/postgres"
)

type UserServerPostgres struct {
	DB *sql.DB
}

func NewUserServerPostgres() (*UserServerPostgres, error) {
	cnf := postgres.InitPostgresConfig()
	db, err := postgres.InitPostgresConnection(cnf)

	if err != nil {
		return nil, err
	}

	return &UserServerPostgres{
		DB: db,
	}, nil
}

func (r *UserServerPostgres) Close() error {
	return r.DB.Close()
}

const GetUserServerByID = `
	SELECT *
	FROM user_servers
	WHERE id = $1
`

func (r *UserServerPostgres) GetById(idServer int) (*domain.UserServer, error) {
	userServer := &domain.UserServer{}

	err := r.DB.QueryRow(
		GetUserServerByID,
		idServer,
	).Scan(
		&userServer.ID,
		&userServer.Addr,
		&userServer.Name,
		&userServer.UserName,
		&userServer.Password,
		&userServer.UserId,
	)

	if err != nil {
		return nil, err
	}

	return userServer, err
}