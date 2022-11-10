package services

type Sync interface {
	SyncGenders() error
	SyncAcceptances() error
}
