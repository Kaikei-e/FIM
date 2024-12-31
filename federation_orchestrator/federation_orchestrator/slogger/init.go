package slogger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func Init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	Logger = logger
}
