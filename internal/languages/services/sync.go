package services

type Sync interface {
	SyncLanguages()
	SyncSubfamilies()
	SyncFamilies()
}
