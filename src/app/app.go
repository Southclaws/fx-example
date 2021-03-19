package app

import (
	"context"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/Southclaws/fx-example/src/api"
	"github.com/Southclaws/fx-example/src/config"
	"github.com/Southclaws/fx-example/src/db"
	"github.com/Southclaws/fx-example/src/logger"
)

func Start(ctx context.Context) {
	app := fx.New(
		fx.NopLogger,

		fx.Provide(
			config.New,
			db.NewDatabase,
		),

		logger.Build(),
		api.Build(),
	)

	err := app.Start(ctx)
	if err != nil {
		panic(err)
	}

	// Wait for context cancellation from interrupt signals set up in main().
	<-ctx.Done()

	// Graceful shutdown time is 30 seconds.
	ctx, cf := context.WithTimeout(context.Background(), time.Second*30)
	defer cf()

	if err := app.Stop(ctx); err != nil {
		zap.L().Error("fatal error occurred", zap.Error(err))
	}
}
