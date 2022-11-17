package storage

import (
	"context"
	"world_generator_processor_service/core/storage"
	"world_generator_processor_service/internal/languages/entities"
	"world_generator_processor_service/internal/languages/repos/queries"
	"world_generator_processor_service/internal/languages/repos/tx"

	"go.uber.org/fx"
)

type adapter struct {
	familyQueries    queries.Family
	subfamilyQueries queries.Subfamily
	languageQueries  queries.Language
	transactions     tx.Tx
}

func New(
	familyQueries queries.Family,
	subfamilyQueries queries.Subfamily,
	languageQueries queries.Language,
	transactions tx.Tx,
) Adapter {
	return &adapter{
		familyQueries:    familyQueries,
		subfamilyQueries: subfamilyQueries,
		languageQueries:  languageQueries,
		transactions:     transactions,
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

func (a *adapter) CreateFamilyTx(ctx context.Context, tx tx.TxClient, f *entities.Family) error {
	return a.transactions.CreateFamily(ctx, tx, f)
}

func (a *adapter) QueryFamilyBySlugTx(ctx context.Context, tx tx.TxClient, slug string) (*entities.Family, error) {
	return a.transactions.QueryFamilyBySlug(ctx, tx, slug)
}

func (a *adapter) QueryFamilies(ctx context.Context, origin entities.Origin, paginate storage.Pagination) ([]*entities.Family, error) {
	return a.familyQueries.QueryFamilies(ctx, origin, paginate)
}

func (a *adapter) CountFamilies(ctx context.Context, origin entities.Origin) (int, error) {
	return a.familyQueries.CountFamilies(ctx, origin)
}

func (a *adapter) CreateSubfamilyTx(ctx context.Context, tx tx.TxClient, sf *entities.Subfamily) error {
	return a.transactions.CreateSubfamily(ctx, tx, sf)
}

func (a *adapter) QuerySubfamilyBySlugTx(ctx context.Context, tx tx.TxClient, slug string) (*entities.Subfamily, error) {
	return a.transactions.QuerySubfamilyBySlug(ctx, tx, slug)
}

func (a *adapter) QuerySubfamilies(ctx context.Context, origin entities.Origin, paginate storage.Pagination) ([]*entities.Subfamily, error) {
	return a.subfamilyQueries.QuerySubfamilies(ctx, origin, paginate)
}

func (a *adapter) CountSubfamilies(ctx context.Context, origin entities.Origin) (int, error) {
	return a.subfamilyQueries.CountSubfamilies(ctx, origin)
}

func (a *adapter) CreateLanguageTx(ctx context.Context, tx tx.TxClient, lang *entities.Language) error {
	return a.transactions.CreateLanguage(ctx, tx, lang)
}

func (a *adapter) QueryLanguageBySlugTx(ctx context.Context, tx tx.TxClient, slug string) (*entities.Language, error) {
	return a.transactions.QueryLanguageBySlug(ctx, tx, slug)
}

func (a *adapter) QueryLanguages(ctx context.Context, origin entities.Origin, paginate storage.Pagination) ([]*entities.Language, error) {
	return a.languageQueries.QueryLanguages(ctx, origin, paginate)
}

func (a *adapter) CountLanguages(ctx context.Context, origin entities.Origin) (int, error) {
	return a.languageQueries.CountLanguages(ctx, origin)
}
