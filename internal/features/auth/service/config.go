package auth_service

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	JWTSecretKey   string        `envconfig:"SECRET_KEY"`
	AccessTokenTTL time.Duration `envconfig:"ACCESS_TOKEN_TTL"`
	RefrshTokenTTL time.Duration `envconfig:"REFRESH_TOKEN_TTL"`
}

func NewConfig() (*Config, error) {
	var config Config
	if err := envconfig.Process(
		"JWT", &config,
	); err != nil {
		return nil, fmt.Errorf("failed to get make auth config: %w", err)
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
