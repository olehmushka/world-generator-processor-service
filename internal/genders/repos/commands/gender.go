package commands

import (
	"context"
	"world_generator_processor_service/internal/genders/entities"
)

type Gender interface {
	Create(ctx context.Context, g entities.Gender) error
}
