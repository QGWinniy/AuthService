package postgres

import (
	"context"
	"database/sql"

	"reg/users-service/internal/domain"
	"reg/users-service/pkg/postgres"
)

type PostgresUserRepository struct {
	DB *sql.DB
}

func NewPostgresUserRepository() (*PostgresUserRepository, error) {
	cnf := postgres.InitPostgresConfig()

	db, err := postgres.InitPostgresConnection(cnf)
	if err != nil {
		return nil, err
	}

	return &PostgresUserRepository{
		DB: db,
	}, nil
}

func (r *PostgresUserRepository) Close() error {
	return r.DB.Close()
}

const CreateUserQuery = `
INSERT INTO users (login, password_hash)
VALUES ($1, $2)
RETURNING id
`

func (r *PostgresUserRepository) Create(
	ctx context.Context,
	user *domain.User,
) (uint64, error) {

	var userID uint64

	err := r.DB.QueryRowContext(
		ctx,
		CreateUserQuery,
		user.Login,
		user.PasswordHash,
	).Scan(&userID)

	if err != nil {
		return 0, err
	}

	return userID, nil
}

const GetUserByIDQuery = `
SELECT
	id,
	login,
	password_hash
FROM users
WHERE id = $1
`

func (r *PostgresUserRepository) GetByID(
	ctx context.Context,
	id uint64,
) (*domain.User, error) {

	var user domain.User

	err := r.DB.QueryRowContext(
		ctx,
		GetUserByIDQuery,
		id,
	).Scan(
		&user.Id,
		&user.Login,
		&user.PasswordHash,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

const GetUserByLoginQuery = `
SELECT
	id,
	login,
	password_hash
FROM users
WHERE login = $1
`

func (r *PostgresUserRepository) GetByLogin(
	ctx context.Context,
	login string,
) (*domain.User, error) {

	var user domain.User

	err := r.DB.QueryRowContext(
		ctx,
		GetUserByLoginQuery,
		login,
	).Scan(
		&user.Id,
		&user.Login,
		&user.PasswordHash,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

const DeleteUserQuery = `
DELETE FROM users
WHERE id = $1
`

func (r *PostgresUserRepository) Delete(
	ctx context.Context,
	id uint64,
) error {

	_, err := r.DB.ExecContext(
		ctx,
		DeleteUserQuery,
		id,
	)

	return err
}