package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Consumer interface {
	Consume(context.Context)
	GetClient() *redis.Client
	GetHandler() HandlerFunc
}
