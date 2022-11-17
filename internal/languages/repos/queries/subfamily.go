package queries

import (
	"context"
	"world_generator_processor_service/core/storage"
	"world_generator_processor_service/internal/languages/entities"

	"github.com/olehmushka/postgres-golang"
)

type Subfamily interface {
	QuerySubfamilyBySlug(ctx context.Context, slug string) (*entities.Subfamily, error)
	QuerySubfamilyBySlugTx(ctx context.Context, client postgres.Reader, slug string) (*entities.Subfamily, error)
	QuerySubfamilies(ctx context.Context, origin entities.Origin, paginate storage.Pagination) ([]*entities.Subfamily, error)
	CountSubfamilies(ctx context.Context, origin entities.Origin) (int, error)
}
