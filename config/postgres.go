package config

type Postgres struct {
	Client           pgClient
	ReaderClient     pgReaderClient
	WriterCLient     pgWriterClient
	BatchItemMaxSize int `env:"POSTGRES_CLIENT_BATCH_ITEM_MAX_SIZE" default:"100"`
}

type pgClient struct {
	Username string `env:"POSTGRES_CLIENT_USERNAME"`
	Password string `env:"POSTGRES_CLIENT_PASSWORD"`
	DBName   string `env:"POSTGRES_CLIENT_DBNAME"`
	Host     string `env:"POSTGRES_CLIENT_HOST"`
	Port     int    `env:"POSTGRES_CLIENT_PORT"`
}

type pgReaderClient struct {
	Username string `env:"POSTGRES_READER_CLIENT_USERNAME"`
	Password string `env:"POSTGRES_READER_CLIENT_PASSWORD"`
	DBName   string `env:"POSTGRES_READER_CLIENT_DBNAME"`
	Host     string `env:"POSTGRES_READER_CLIENT_HOST"`
	Port     int    `env:"POSTGRES_READER_CLIENT_PORT"`
}

type pgWriterClient struct {
	Username string `env:"POSTGRES_WRITER_CLIENT_USERNAME"`
	Password string `env:"POSTGRES_WRITER_CLIENT_PASSWORD"`
	DBName   string `env:"POSTGRES_WRITER_CLIENT_DBNAME"`
	Host     string `env:"POSTGRES_WRITER_CLIENT_HOST"`
	Port     int    `env:"POSTGRES_WRITER_CLIENT_PORT"`
}
