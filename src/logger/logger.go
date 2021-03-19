package logger

import (
	"fmt"
	"os"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/Southclaws/fx-example/src/config"
	"github.com/Southclaws/fx-example/src/version"
)

func Build() fx.Option {
	return fx.Options(
		fx.Provide(func(cfg config.Config) *zap.Logger {
			var config zap.Config
			if cfg.Production {
				config = zap.NewProductionConfig()
				config.InitialFields = map[string]interface{}{"v": version.Version}
			} else {
				config = zap.NewDevelopmentConfig()
			}

			config.Level.SetLevel(cfg.LogLevel)
			config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

			logger, err := config.Build()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			return logger
		}),
		fx.Invoke(func(c config.Config, l *zap.Logger) {
			// Use our logger for globals too, even though it's passed to
			// dependents most of the time using DI, the global logger is used
			// in a couple of places during startup/shutdown.
			zap.ReplaceGlobals(l)
			if !c.Production {
				l.Info("logger configured in development mode")
			}
		}),
	)
}
