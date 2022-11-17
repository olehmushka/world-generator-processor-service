package tx

import (
	"context"
	"world_generator_processor_service/internal/languages/entities"
	"world_generator_processor_service/internal/languages/repos/commands"
	"world_generator_processor_service/internal/languages/repos/queries"

	"github.com/olehmushka/postgres-golang"
	"go.uber.org/fx"
)

type tx struct {
	client           postgres.Client
	familyComm       commands.Family
	familyQueries    queries.Family
	subfamilyComm    commands.Subfamily
	subfamilyQueries queries.Subfamily
	languageComm     commands.Language
	languageQueries  queries.Language
}

func New(
	client postgres.Client,
	familyComm commands.Family,
	familyQueries queries.Family,
	subfamilyComm commands.Subfamily,
	subfamilyQueries queries.Subfamily,
	languageComm commands.Language,
	languageQueries queries.Language,
) Tx {
	return &tx{
		client:           client,
		familyComm:       familyComm,
		familyQueries:    familyQueries,
		subfamilyComm:    subfamilyComm,
		subfamilyQueries: subfamilyQueries,
		languageComm:     languageComm,
		languageQueries:  languageQueries,
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

func (t *tx) CreateFamily(ctx context.Context, tx TxClient, f *entities.Family) error {
	return t.familyComm.CreateTx(ctx, tx, f)
}

func (t *tx) QueryFamilyBySlug(ctx context.Context, tx TxClient, slug string) (*entities.Family, error) {
	return t.familyQueries.QueryFamilyBySlugTx(ctx, tx, slug)
}

func (t *tx) CreateSubfamily(ctx context.Context, tx TxClient, sf *entities.Subfamily) error {
	return t.subfamilyComm.CreateTx(ctx, tx, sf)
}

func (t *tx) QuerySubfamilyBySlug(ctx context.Context, tx TxClient, slug string) (*entities.Subfamily, error) {
	return t.subfamilyQueries.QuerySubfamilyBySlugTx(ctx, tx, slug)
}

func (t *tx) CreateLanguage(ctx context.Context, tx TxClient, lang *entities.Language) error {
	return t.languageComm.CreateTx(ctx, tx, lang)
}

func (t *tx) QueryLanguageBySlug(ctx context.Context, tx TxClient, slug string) (*entities.Language, error) {
	return t.languageQueries.QueryLanguageBySlugTx(ctx, tx, slug)
}
