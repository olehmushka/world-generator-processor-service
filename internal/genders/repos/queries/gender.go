package queries

import (
	"context"
	"world_generator_processor_service/core/storage"
	"world_generator_processor_service/internal/genders/entities"

	"github.com/olehmushka/postgres-golang"
)

type Gender interface {
	QueryGenderByName(ctx context.Context, name string) (entities.Gender, error)
	QueryGenderByNameTx(ctx context.Context, client postgres.Reader, name string) (entities.Gender, error)
	QueryGenders(ctx context.Context, origin entities.Origin, paginate storage.Pagination) ([]entities.Gender, error)
	CountGenders(ctx context.Context, origin entities.Origin) (int, error)
}
