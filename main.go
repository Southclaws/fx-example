package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"

	"github.com/Southclaws/fx-example/src/app"
)

func main() {
	// Load environment variables from a .env file.
	//nolint:errcheck
	godotenv.Load()

	// Cancel root context on interrupt signal.
	ctx, cf := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cf()

	app.Start(ctx)
}
