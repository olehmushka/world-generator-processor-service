package tx

import (
	"context"
	"world_generator_processor_service/internal/languages/entities"

	"github.com/olehmushka/postgres-golang"
)

type Tx interface {
	BeginTx(ctx context.Context) (TxClient, error)
	Commit(ctx context.Context, tx TxClient) error
	Rollback(ctx context.Context, tx TxClient) error

	CreateFamily(ctx context.Context, tx TxClient, f *entities.Family) error
	QueryFamilyBySlug(ctx context.Context, tx TxClient, slug string) (*entities.Family, error)
	CreateSubfamily(ctx context.Context, tx TxClient, sf *entities.Subfamily) error
	QuerySubfamilyBySlug(ctx context.Context, tx TxClient, slug string) (*entities.Subfamily, error)
	CreateLanguage(ctx context.Context, tx TxClient, lang *entities.Language) error
	QueryLanguageBySlug(ctx context.Context, tx TxClient, slug string) (*entities.Language, error)
}

type TxClient interface {
	postgres.Transaction
}
