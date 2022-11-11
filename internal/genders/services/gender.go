package services

import (
	"context"
	"world_generator_processor_service/core/storage"
	gendersEntities "world_generator_processor_service/internal/genders/entities"
)

type Gender interface {
	QueryGenders(ctx context.Context, origin gendersEntities.Origin, paginate storage.Pagination) ([]gendersEntities.Gender, int, error)
}
