package config

type HTTPServer struct {
	Address  string `env:"HTTP_SERVER_ADDRESS"`
	BasePath string `env:"HTTP_SERVER_BASE_PATH"`

	ReadTimeoutInSec       int `env:"HTTP_SERVER_READ_TIMEOUT_IN_SECONDS"`
	ReadHeaderTimeoutInSec int `env:"HTTP_SERVER_READ_HEADER_TIMEOUT_IN_SECONDS"`
	WriteTimeoutInSec      int `env:"HTTP_SERVER_WRITE_TIMEOUT_IN_SECONDS"`
	IdleTimeoutInSec       int `env:"HTTP_SERVER_IDLE_TIMEOUT_IN_SECONDS"`
}
