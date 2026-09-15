package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"AuthService/sessions-service/internal/domain"
	redisConnect "AuthService/sessions-service/pkg/redis"
	"time"

	"github.com/go-redis/redis/v8"
)	

type RedisSessionRepository struct {
	DB *redis.Client
}

func NewRedisSessionRepository() (*RedisSessionRepository, error) {
	cfg := redisConnect.InitRedisConfig()

	db, err := redisConnect.InitRedisConnection(cfg)

		if err != nil {
		return nil, err
	}

	return &RedisSessionRepository{
		DB: db,
	}, nil
}

func (r *RedisSessionRepository) Close() error {
	return r.DB.Close()
}

func sessionKey(sessionID string) string {
	return fmt.Sprintf("session:%s", sessionID)
}

func (r *RedisSessionRepository) Create(
	ctx context.Context,
	session *domain.Session,
) error {
	data, err := json.Marshal(session)
	if err != nil {
		return err
	}

	ttl := time.Until(session.ExpiresAt)
	if ttl <= 0 {
		return errors.New("session already expired")
	}

	return r.DB.Set(
		ctx,
		sessionKey(session.SessionId),
		data,
		ttl,
	).Err()
}

func (r *RedisSessionRepository) Get(
	ctx context.Context,
	sessionID string,
) (*domain.Session, error) {

	data, err := r.DB.Get(
		ctx,
		sessionKey(sessionID),
	).Bytes()

	if err != nil {
		return nil, err
	}

	var session domain.Session

	if err := json.Unmarshal(data, &session); err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *RedisSessionRepository) Delete(
	ctx context.Context,
	sessionID string,
) error {
	return r.DB.Del(
		ctx,
		sessionKey(sessionID),
	).Err()
}