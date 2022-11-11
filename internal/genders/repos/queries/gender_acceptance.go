package queries

import (
	"context"
	"world_generator_processor_service/core/storage"
	"world_generator_processor_service/internal/genders/entities"

	"github.com/olehmushka/postgres-golang"
)

type GenderAcceptance interface {
	QueryGenderAcceptanceByName(ctx context.Context, name string) (entities.GenderAcceptance, error)
	QueryGenderAcceptanceByNameTx(ctx context.Context, client postgres.Reader, name string) (entities.GenderAcceptance, error)
	QueryGenderAcceptances(ctx context.Context, origin entities.Origin, paginate storage.Pagination) ([]entities.GenderAcceptance, error)
	CountGenderAcceptances(ctx context.Context, origin entities.Origin) (int, error)
}
