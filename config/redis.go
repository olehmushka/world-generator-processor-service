package config

type Redis struct {
	URL      string `env:"REDIS_URL"`
	Username string `env:"REDIS_USERNAME"`
	Password string `env:"REDIS_PASSWORD"`
}
