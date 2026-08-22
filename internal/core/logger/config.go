package core_logger

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	core_errors "github.com/qury53/todo-app2/internal/core/errors"
)

var avaibleLogLevels = map[int]bool{
	-4: true,
	0:  true,
	4:  true,
	8:  true,
}

type Config struct {
	LogFolder string `envconfig:"FOLDER" required:"true"`
	LogLevel  int    `envconfig:"LEVEL" default:"-4"` //(O)INFO (4) WARN  (8) ERROR
}

func NewConfig() (*Config, error) {
	var config Config
	err := envconfig.Process("LOG", &config)
	if err != nil {
		return nil, fmt.Errorf("logger config: %w", err)
	}

	if !avaibleLogLevels[config.LogLevel] {
		return nil, fmt.Errorf("invalid log level code: %w", core_errors.ErrInvalidArgument)
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
