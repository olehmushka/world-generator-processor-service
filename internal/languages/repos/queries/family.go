package queries

import (
	"context"
	"world_generator_processor_service/core/storage"
	"world_generator_processor_service/internal/languages/entities"

	"github.com/olehmushka/postgres-golang"
)

type Family interface {
	QueryFamilyBySlug(ctx context.Context, slug string) (*entities.Family, error)
	QueryFamilyBySlugTx(ctx context.Context, client postgres.Reader, slug string) (*entities.Family, error)
	QueryFamilies(ctx context.Context, origin entities.Origin, paginate storage.Pagination) ([]*entities.Family, error)
	CountFamilies(ctx context.Context, origin entities.Origin) (int, error)
}
