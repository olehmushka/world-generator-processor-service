package redis

import (
	"context"
	"fmt"
	"time"
	"world_generator_processor_service/config"

	"github.com/go-redis/redis/v8"
	"github.com/olehmushka/golang-toolkit/wrapped_error"
	"go.uber.org/fx"
)

type store struct {
	client *redis.Client
}

func NewStore(cfg *config.Config) (Store, error) {
	return NewStoreByParams(cfg.Redis.URL, cfg.Redis.Username, cfg.Redis.Password)
}

var StoreModule = fx.Options(
	fx.Provide(NewStore),
)

func NewStoreByParams(url, username, password string) (Store, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "parsing redis url error for storage")
	}
	if username != "" {
		opts.Username = username
	}
	if password != "" {
		opts.Password = password
	}

	client := redis.NewClient(opts)

	return &store{
		client: client,
	}, nil
}

func (s *store) Get(ctx context.Context, key string) (string, error) {
	val, err := s.client.Get(ctx, key).Result()
	if err != nil {
		if redis.Nil.Error() == err.Error() {
			return "", nil
		}
		return "", wrapped_error.NewInternalServerError(err, "can not get value from redis by key")
	}

	return val, nil
}

func (s *store) Set(ctx context.Context, key, value string, exp time.Duration) error {
	if err := s.client.Set(ctx, key, value, exp).Err(); err != nil {
		return wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not set value in redis (key=%s, value=%s)", key, value))
	}

	return nil
}
