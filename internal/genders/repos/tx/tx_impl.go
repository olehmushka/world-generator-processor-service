package tx

import (
	"context"
	"world_generator_processor_service/internal/genders/entities"
	"world_generator_processor_service/internal/genders/repos/commands"
	"world_generator_processor_service/internal/genders/repos/queries"

	"github.com/olehmushka/postgres-golang"
	"go.uber.org/fx"
)

type tx struct {
	client                  postgres.Client
	genderComm              commands.Gender
	genderQueries           queries.Gender
	genderAcceptanceComm    commands.GenderAcceptance
	genderAcceptanceQueries queries.GenderAcceptance
}

func New(
	client postgres.Client,
	genderComm commands.Gender,
	genderQueries queries.Gender,
	genderAcceptanceComm commands.GenderAcceptance,
	genderAcceptanceQueries queries.GenderAcceptance,
) Tx {
	return &tx{
		client:                  client,
		genderComm:              genderComm,
		genderQueries:           genderQueries,
		genderAcceptanceComm:    genderAcceptanceComm,
		genderAcceptanceQueries: genderAcceptanceQueries,
	}
}

var Module = fx.Options(
	fx.Provide(New),
)

func (t *tx) BeginTx(ctx context.Context) (TxClient, error) {
	return t.client.BeginTx(ctx, postgres.TxOptions{})
}

func (t *tx) Commit(ctx context.Context, tx TxClient) error {
	return tx.Commit(ctx)
}

func (t *tx) Rollback(ctx context.Context, tx TxClient) error {
	return tx.Rollback(ctx)
}

func (t *tx) CreateGender(ctx context.Context, tx TxClient, g entities.Gender) error {
	return t.genderComm.CreateTx(ctx, tx, g)
}

func (t *tx) QueryGenderByName(ctx context.Context, tx TxClient, name string) (entities.Gender, error) {
	return t.genderQueries.QueryGenderByNameTx(ctx, tx, name)
}

func (t *tx) CreateGenderAcceptance(ctx context.Context, tx TxClient, ga entities.GenderAcceptance) error {
	return t.genderAcceptanceComm.CreateTx(ctx, tx, ga)
}

func (t *tx) QueryGenderAcceptanceByName(ctx context.Context, tx TxClient, name string) (entities.GenderAcceptance, error) {
	return t.genderAcceptanceQueries.QueryGenderAcceptanceByNameTx(ctx, tx, name)
}
