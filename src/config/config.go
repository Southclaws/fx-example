package config

import (
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap/zapcore"
)

// Config represents environment variable configuration parameters
type Config struct {
	Production bool          `envconfig:"PRODUCTION" default:"false"`
	LogLevel   zapcore.Level `envconfig:"LOG_LEVEL"  default:"info"`
	ListenAddr string        `envconfig:"LISTEN_ADDR" default:"0.0.0.0:8080"`
}

func New() (c Config, err error) {
	if err = envconfig.Process("", &c); err != nil {
		return c, err
	}

	return
}
