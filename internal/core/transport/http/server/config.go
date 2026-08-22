package transport_http_server

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addr         string        `envconfig:"ADDR" required:"true"`
	ShutdownTime time.Duration `envconfig:"SHUTDOWN_TIME" default:"30s"`
}

func NewConfig() (*Config, error) {
	var config Config
	err := envconfig.Process("SERVER", &config)

	if err != nil {
		return nil, fmt.Errorf("server config: %w: ", err)
	}

	return &config, nil
}

func NewConfigMust() *Config {
	config, err := NewConfig()
	if err != nil {
		panic(err)
	}

	return config
}
