package main

import (
	"feed_collector/router"
	"feed_collector/slogger"
)

func main() {
	slogger.Init()
	slogger.Logger.Info("Application started")

	router.Router()
}
