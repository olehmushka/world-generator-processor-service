package mongo

import (
	"context"
)

type Cursor interface {
	ID() int64
	Next(context.Context) bool
	TryNext(context.Context) bool
	Decode(any) error
	Err() error
	Close(context.Context) error
	All(context.Context, any) error
	RemainingBatchLength() int
}
