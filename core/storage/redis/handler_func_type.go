package redis

import "context"

type HandlerFunc func(context.Context, []byte) error
