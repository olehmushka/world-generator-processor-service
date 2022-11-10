package engine

import (
	"context"
	"world_generator_processor_service/core/engine"
	"world_generator_processor_service/internal/genders/entities"

	"github.com/olehmushka/world-generator-engine/influence"
	"go.uber.org/fx"
)

type adapter struct {
	engine engine.Engine
}

func New(eng engine.Engine) Adapter {
	return &adapter{
		engine: eng,
	}
}

var Module = fx.Options(
	fx.Provide(New),
)

func (a *adapter) LoadNativeGenders(context.Context) ([]entities.Gender, error) {
	out := make([]entities.Gender, 0, 2)
	for _, g := range a.engine.LoadGenders() {
		out = append(out, entities.Gender{
			Name:   g.String(),
			Origin: entities.NativeOrigin,
		})
	}

	return out, nil
}

func (a *adapter) LoadNativeGenderAcceptances(context.Context) ([]entities.GenderAcceptance, error) {
	out := make([]entities.GenderAcceptance, 0, 2)
	for _, g := range a.engine.LoadGenderAcceptances() {
		out = append(out, entities.GenderAcceptance{
			Name:   g.String(),
			Origin: entities.NativeOrigin,
		})
	}

	return out, nil
}

func (a *adapter) LoadInfluences(context.Context) ([]entities.Influence, error) {
	out := make([]entities.Influence, 3)
	for _, i := range a.engine.LoadInfluences() {
		switch i {
		case influence.StrongInfluence:
			out = append(out, entities.StrongInfluence)
		case influence.ModerateInfluence:
			out = append(out, entities.ModerateInfluence)
		case influence.WeakInfluence:
			out = append(out, entities.WeakInfluence)
		}
	}

	return out, nil
}
