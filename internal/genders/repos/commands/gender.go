package commands

import (
	"context"
	"world_generator_processor_service/internal/genders/entities"

	"github.com/olehmushka/postgres-golang"
)

type Gender interface {
	Create(ctx context.Context, g entities.Gender) error
	CreateTx(ctx context.Context, client postgres.Writer, g entities.Gender) error
}
