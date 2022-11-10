package services

import "go.uber.org/fx"

type syncSrv struct{}

func NewSyncSrv() Sync {
	return &syncSrv{}
}

var SyncSrvModule = fx.Options(
	fx.Provide(NewSyncSrv),
)

func (s *syncSrv) SyncTraditions() {}
