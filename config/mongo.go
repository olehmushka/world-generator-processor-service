package config

type MongoDB struct {
	URL              string `env:"MONGODB_URL"`
	DBName           string `env:"MONGODB_DB_NAME"`
	Username         string `env:"MONGODB_USERNAME"`
	Password         string `env:"MONGODB_PASSWORD"`
	MaxBulkItemsSize int    `env:"MONGODB_MAX_BULK_ITEMS_SIZE"`
}
