package redis

import "context"

type Publisher interface {
	Publish(context.Context, string, []byte) error
}
