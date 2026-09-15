package redis

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

func InitRedisConfig() Config {
	port, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	return Config{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     port,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	}
}

func InitRedisConnection(
	cfg Config,
) (*redis.Client, error) {

	addr := fmt.Sprintf(
		"%s:%d",
		cfg.Host,
		cfg.Port,
	)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		_ = client.Close()
		return nil, err
	}

	return client, nil
}