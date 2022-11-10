package engine

import (
	"context"
	"world_generator_processor_service/internal/genders/entities"
)

type Adapter interface {
	LoadNativeGenders(context.Context) ([]entities.Gender, error)
	LoadNativeGenderAcceptances(context.Context) ([]entities.GenderAcceptance, error)
	LoadInfluences(context.Context) ([]entities.Influence, error)
}
