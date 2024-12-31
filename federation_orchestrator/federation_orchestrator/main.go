package main

import (
	"context"
	"federation_orchestrator/router"
	"federation_orchestrator/slogger"
)

func main() {
	ctx := context.Background()

	slogger.Init()
	slogger.Logger.Info("Application started")

	router.Router(ctx)
	slogger.Logger.Info("Application stopped")
}
