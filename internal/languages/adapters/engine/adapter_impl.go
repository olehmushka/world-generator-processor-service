package engine

import "go.uber.org/fx"

type adapter struct{}

func New() Adapter {
	return &adapter{}
}

var Module = fx.Options(
	fx.Provide(New),
)
