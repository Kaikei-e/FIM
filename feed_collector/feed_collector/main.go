package main

import (
	"context"
	"feed_collector/driver/db_driver"
	"feed_collector/router"
	"feed_collector/slogger"
)

func main() {
	ctx := context.Background()

	slogger.Init()
	slogger.Logger.Info("Application started")

	fimDB, err := db_driver.InitConn(ctx)
	if err != nil {
		slogger.Logger.Error("Failed to initialize the database connection.", "Caused by", err)
		return
	}
	defer fimDB.Close()

	router.Router(fimDB, ctx)
}
