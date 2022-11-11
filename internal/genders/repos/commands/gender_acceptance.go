package commands

import (
	"context"
	"world_generator_processor_service/internal/genders/entities"

	"github.com/olehmushka/postgres-golang"
)

type GenderAcceptance interface {
	Create(ctx context.Context, g entities.GenderAcceptance) error
	CreateTx(ctx context.Context, client postgres.Writer, g entities.GenderAcceptance) error
}
