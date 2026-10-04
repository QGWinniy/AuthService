package postgres

import (
	"AuthService/server-service-consumer/internal/domain"
	"AuthService/server-service-consumer/pkg/postgres"
	"context"
	"database/sql"
)

type PostgresUserServersRepository struct {
	DB *sql.DB
}

func NewPostgresRepository() (*PostgresUserServersRepository, error) {
	cnf := postgres.InitPostgresConfig()
	db, err := postgres.InitPostgresConnection(cnf)

	if err != nil {
		return nil, err
	}

	return &PostgresUserServersRepository{
		DB: db,
	}, nil
}

func (r *PostgresUserServersRepository) Close() error {
	return r.DB.Close()
}

const CheckUserHasServer = `
	SELECT EXISTS (
		SELECT 1
		FROM user_servers
		WHERE userid = $1
      		AND addr = $2
	);
` 

func (r *PostgresUserServersRepository) CheckUserHasServer(
	ctx context.Context,
	userServer *domain.UserServer,
) (bool, error) {
	var exists bool

	err := r.DB.QueryRowContext(
		ctx,
		CheckUserHasServer,
		userServer.ID,
		userServer.Addr,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}


const CreateUserServer = `
	INSERT INTO user_servers (
		addr,
		name,
		username,
		password,
		userid
	)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
`


func (r *PostgresUserServersRepository) Create(
	ctx context.Context,
	userServer *domain.UserServer,
) error {
	
	err := r.DB.QueryRowContext(
		ctx,
		CreateUserServer,
		userServer.Addr,
		userServer.Name,
		userServer.UserName,
		userServer.Password,
		userServer.UserId,
	).Scan(&userServer.ID)

	return err
}


