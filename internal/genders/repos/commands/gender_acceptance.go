package commands

import (
	"context"
	"world_generator_processor_service/internal/genders/entities"
)

type GenderAcceptance interface {
	Create(ctx context.Context, g entities.GenderAcceptance) error
}
