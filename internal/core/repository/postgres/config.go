package core_pgx_pool

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host     string        `envconfig:"HOST" required:"true"`
	Port     string        `envconfig:"PORT" default:"5432"`
	User     string        `envconfig:"USER" required:"true"`
	Database string        `envconfig:"DATABASE" required:"true"`
	Password string        `envconfig:"PASSWORD" required:"true"`
	Timeout  time.Duration `envconfig:"TIMEOUT" default:"30s"`
}

func NewConfig() (*Config, error) {
	var config Config
	err := envconfig.Process("POSTGRES", &config)
	if err != nil {
		return nil, err
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
