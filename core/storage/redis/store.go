package redis

import (
	"context"
	"time"
)

type Store interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, value string, exp time.Duration) error
}
