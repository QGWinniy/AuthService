package postgres

import (
	"AuthService/server-service/internal/domain"
	"AuthService/server-service/pkg/postgres"
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

const GetUserServers = `
	SELECT
		id,
		addr,
		name,
		username,
		password,
		userid
	FROM user_servers
	WHERE userid = $1
`

func (r *PostgresUserServersRepository) GetUserServers(
	ctx context.Context,
	userId int64,
) ([]domain.UserServer, error) {
	rows, err := r.DB.QueryContext(
		ctx,
		GetUserServers,
		userId,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	userServers := make([]domain.UserServer, 0)

	for rows.Next() {
		var userServer domain.UserServer

		err = rows.Scan(
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

		userServers = append(userServers, userServer)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return userServers, nil
}

const GetUserServer = `
	SELECT
		id,
		addr,
		name,
		username,
		password,
		userid
	FROM user_servers
	WHERE userid = $1 AND id = $2
`

func (r *PostgresUserServersRepository) GetUserServer(
	ctx context.Context,
	userID int,
	serverID int,
) (*domain.UserServer, error) {
	userServer := &domain.UserServer{}

	err := r.DB.QueryRowContext(
		ctx,
		GetUserServer,
		userID,
		serverID,
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

	return userServer, nil
}
