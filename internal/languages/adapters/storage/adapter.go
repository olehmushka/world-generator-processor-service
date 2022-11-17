package storage

import (
	"context"
	"world_generator_processor_service/core/storage"
	"world_generator_processor_service/internal/languages/entities"
	"world_generator_processor_service/internal/languages/repos/tx"
)

type Adapter interface {
	BeginTx(ctx context.Context) (tx.TxClient, error)
	Commit(ctx context.Context, tx tx.TxClient) error
	Rollback(ctx context.Context, tx tx.TxClient) error

	CreateFamilyTx(ctx context.Context, tx tx.TxClient, f *entities.Family) error
	QueryFamilyBySlugTx(ctx context.Context, tx tx.TxClient, slug string) (*entities.Family, error)
	QueryFamilies(ctx context.Context, origin entities.Origin, paginate storage.Pagination) ([]*entities.Family, error)
	CountFamilies(ctx context.Context, origin entities.Origin) (int, error)

	CreateSubfamilyTx(ctx context.Context, tx tx.TxClient, sf *entities.Subfamily) error
	QuerySubfamilyBySlugTx(ctx context.Context, tx tx.TxClient, slug string) (*entities.Subfamily, error)
	QuerySubfamilies(ctx context.Context, origin entities.Origin, paginate storage.Pagination) ([]*entities.Subfamily, error)
	CountSubfamilies(ctx context.Context, origin entities.Origin) (int, error)

	CreateLanguageTx(ctx context.Context, tx tx.TxClient, lang *entities.Language) error
	QueryLanguageBySlugTx(ctx context.Context, tx tx.TxClient, slug string) (*entities.Language, error)
	QueryLanguages(ctx context.Context, origin entities.Origin, paginate storage.Pagination) ([]*entities.Language, error)
	CountLanguages(ctx context.Context, origin entities.Origin) (int, error)
}
