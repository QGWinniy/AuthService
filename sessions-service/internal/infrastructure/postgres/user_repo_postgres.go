package postgres

import (
	"context"
	"database/sql"
	"AuthService/sessions-service/internal/domain"
	"AuthService/sessions-service/pkg/postgres"
)

type PostgresSessionRepository struct {
	DB *sql.DB
}

func NewPostgresSessionRepository() (*PostgresSessionRepository, error) {
	cnf := postgres.InitPostgresConfig()

	db, err := postgres.InitPostgresConnection(cnf)
	if err != nil {
		return nil, err
	}

	return &PostgresSessionRepository{
		DB: db,
	}, nil
}

func (r *PostgresSessionRepository) Close() error {
	return r.DB.Close()
}

const CreateSessionQuery = `
INSERT INTO sessions (
	session_id,
	user_id,
	data,
	expires_at
)
VALUES ($1, $2, $3, $4)
`

func (r *PostgresSessionRepository) Create(
	ctx context.Context,
	session *domain.Session,
) error {

	_, err := r.DB.ExecContext(
		ctx,
		CreateSessionQuery,
		session.SessionId,
		session.UserId,
		session.Data,
		session.ExpiresAt,
	)

	return err
}

const GetSessionQuery = `
SELECT
	session_id,
	user_id,
	data,
	expires_at
FROM sessions
WHERE session_id = $1
  AND expires_at > NOW()
`

func (r *PostgresSessionRepository) Get(
	ctx context.Context,
	sessionId string,
) (*domain.Session, error) {

	var session domain.Session

	err := r.DB.QueryRowContext(
		ctx,
		GetSessionQuery,
		sessionId,
	).Scan(
		&session.SessionId,
		&session.UserId,
		&session.Data,
		&session.ExpiresAt,
	)

	if err != nil {
		return nil, err
	}

	return &session, nil
}

const DeleteSessionQuery = `
DELETE FROM sessions
WHERE session_id = $1
`

func (r *PostgresSessionRepository) Delete(
	ctx context.Context,
	sessionId string,
) error {

	_, err := r.DB.ExecContext(
		ctx,
		DeleteSessionQuery,
		sessionId,
	)

	return err
}
