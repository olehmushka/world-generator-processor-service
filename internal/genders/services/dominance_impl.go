package services

import "go.uber.org/fx"

type dominanceSrv struct{}

func NewDominanceSrv() Dominance {
	return &dominanceSrv{}
}

var DominanceSrvModule = fx.Options(
	fx.Provide(NewDominanceSrv),
)
