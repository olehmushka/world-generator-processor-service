package storage

import (
	"context"
	"world_generator_processor_service/core/storage"
	"world_generator_processor_service/internal/genders/entities"
	"world_generator_processor_service/internal/genders/repos/tx"
)

type Adapter interface {
	BeginTx(ctx context.Context) (tx.TxClient, error)
	Commit(ctx context.Context, tx tx.TxClient) error
	Rollback(ctx context.Context, tx tx.TxClient) error

	CreateGenderTx(ctx context.Context, tx tx.TxClient, g entities.Gender) error
	QueryGenderByNameTx(ctx context.Context, tx tx.TxClient, name string) (entities.Gender, error)
	QueryGenders(ctx context.Context, origin entities.Origin, paginate storage.Pagination) ([]entities.Gender, error)
	CountGenders(ctx context.Context, origin entities.Origin) (int, error)

	CreateGenderAcceptanceTx(ctx context.Context, tx tx.TxClient, g entities.GenderAcceptance) error
	QueryGenderAcceptanceByNameTx(ctx context.Context, tx tx.TxClient, name string) (entities.GenderAcceptance, error)
	QueryGenderAcceptances(ctx context.Context, origin entities.Origin, paginate storage.Pagination) ([]entities.GenderAcceptance, error)
	CountGenderAcceptances(ctx context.Context, origin entities.Origin) (int, error)
}
