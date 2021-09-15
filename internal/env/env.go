package env

import (
	"github.com/joeshaw/envdecode"
)

// Conf represents all of the required environment variables
type Config struct {
	HTTPReadTimeout   int    `env:"HTTP_READ_TIMEOUT,default=5"`
	HTTPWriteTimeout  int    `env:"HTTP_WRITE_TIMEOUT,default=5"`
	HTTPAddress       string `env:"HTTP_ADDRESS,default=:80"`
	HTTPClientTimeout int    `env:"HTTP_CLIENT_TIMEOUT,default=5"`
	LogLevel          string `env:"LOG_LEVEL,default=debug"`
	ItunesEndpoint    string `env:"ITUNES_ENDPOINT"`
}

// MustLoad attempt to marshal the env vars. If and are missing or incorrect, method panics
func MustLoad() Config {
	var cfg Config
	if err := envdecode.Decode(&cfg); err != nil {
		panic(err)
	}
	return cfg
}
