package storage

import (
	"context"
	"world_generator_processor_service/core/storage"
	"world_generator_processor_service/internal/genders/entities"
	"world_generator_processor_service/internal/genders/repos/queries"
	"world_generator_processor_service/internal/genders/repos/tx"

	"go.uber.org/fx"
)

type adapter struct {
	genderQueries           queries.Gender
	genderAcceptanceQueries queries.GenderAcceptance
	transactions            tx.Tx
}

func New(
	genderQueries queries.Gender,
	genderAcceptanceQueries queries.GenderAcceptance,
	transactions tx.Tx,
) Adapter {
	return &adapter{
		genderQueries:           genderQueries,
		genderAcceptanceQueries: genderAcceptanceQueries,
		transactions:            transactions,
	}
}

var Module = fx.Options(
	fx.Provide(New),
)

func (a *adapter) BeginTx(ctx context.Context) (tx.TxClient, error) {
	return a.transactions.BeginTx(ctx)
}

func (a *adapter) Commit(ctx context.Context, tx tx.TxClient) error {
	return a.transactions.Commit(ctx, tx)
}

func (a *adapter) Rollback(ctx context.Context, tx tx.TxClient) error {
	return a.transactions.Rollback(ctx, tx)
}

func (a *adapter) CreateGenderTx(ctx context.Context, tx tx.TxClient, g entities.Gender) error {
	return a.transactions.CreateGender(ctx, tx, g)
}

func (a *adapter) QueryGenderByNameTx(ctx context.Context, tx tx.TxClient, name string) (entities.Gender, error) {
	return a.transactions.QueryGenderByName(ctx, tx, name)
}

func (a *adapter) QueryGenders(ctx context.Context, origin entities.Origin, paginate storage.Pagination) ([]entities.Gender, error) {
	return a.genderQueries.QueryGenders(ctx, origin, paginate)
}

func (a *adapter) CountGenders(ctx context.Context, origin entities.Origin) (int, error) {
	return a.genderQueries.CountGenders(ctx, origin)
}

func (a *adapter) CreateGenderAcceptanceTx(ctx context.Context, tx tx.TxClient, g entities.GenderAcceptance) error {
	return a.transactions.CreateGenderAcceptance(ctx, tx, g)
}

func (a *adapter) QueryGenderAcceptanceByNameTx(ctx context.Context, tx tx.TxClient, name string) (entities.GenderAcceptance, error) {
	return a.transactions.QueryGenderAcceptanceByName(ctx, tx, name)
}

func (a *adapter) QueryGenderAcceptances(ctx context.Context, origin entities.Origin, paginate storage.Pagination) ([]entities.GenderAcceptance, error) {
	return a.genderAcceptanceQueries.QueryGenderAcceptances(ctx, origin, paginate)
}

func (a *adapter) CountGenderAcceptances(ctx context.Context, origin entities.Origin) (int, error) {
	return a.genderAcceptanceQueries.CountGenderAcceptances(ctx, origin)
}
