package services

import "context"

type Sync interface {
	SyncLanguages(ctx context.Context) error
	SyncSubfamilies(ctx context.Context) error
	SyncFamilies(ctx context.Context) error
}
