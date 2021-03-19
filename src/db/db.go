package db

import (
	"context"
	"database/sql"

	"go.uber.org/fx"
)

func NewDatabase(lc fx.Lifecycle) (*sql.DB, error) {
	database := sql.OpenDB(nil)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			// database.Connect()
			return nil
		},
		OnStop: func(context.Context) error {
			// database.Disconnect()
			return nil
		},
	})

	return database, nil
}
