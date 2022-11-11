package tx

import (
	"context"
	"world_generator_processor_service/internal/genders/entities"

	"github.com/olehmushka/postgres-golang"
)

type Tx interface {
	BeginTx(ctx context.Context) (TxClient, error)
	Commit(ctx context.Context, tx TxClient) error
	Rollback(ctx context.Context, tx TxClient) error

	CreateGender(ctx context.Context, tx TxClient, g entities.Gender) error
	QueryGenderByName(ctx context.Context, tx TxClient, name string) (entities.Gender, error)
	CreateGenderAcceptance(ctx context.Context, tx TxClient, g entities.GenderAcceptance) error
	QueryGenderAcceptanceByName(ctx context.Context, tx TxClient, name string) (entities.GenderAcceptance, error)
}

type TxClient interface {
	postgres.Transaction
}
