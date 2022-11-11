package services

import "context"

type Sync interface {
	SyncGenders(ctx context.Context) error
	SyncGenderAcceptances(ctx context.Context) error
}
