package queries

import (
	"context"
	"world_generator_processor_service/core/storage"
	"world_generator_processor_service/internal/languages/entities"

	"github.com/olehmushka/postgres-golang"
)

type Language interface {
	QueryLanguageBySlug(ctx context.Context, slug string) (*entities.Language, error)
	QueryLanguageBySlugTx(ctx context.Context, client postgres.Reader, slug string) (*entities.Language, error)
	QueryLanguages(ctx context.Context, origin entities.Origin, paginate storage.Pagination) ([]*entities.Language, error)
	CountLanguages(ctx context.Context, origin entities.Origin) (int, error)
}
