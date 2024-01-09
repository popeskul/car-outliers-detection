package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server struct {
		Port            int `mapstructure:"port" envconfig:"APP_PORT"`
		HealthCheckPort int `mapstructure:"health_check_port" envconfig:"APP_HEALTH_CHECK_PORT"`
	} `mapstructure:"server"`
}

func New() (*Config, error) {
	config := Config{}

	if err := envconfig.Process("", &config); err != nil {
		return nil, err
	}

	return &config, nil
}
