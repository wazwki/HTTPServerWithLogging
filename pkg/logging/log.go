package logging

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func LogInit() {
	file, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	Logger = slog.New(slog.NewTextHandler(file, nil))
	Logger.Info("Logger start")
}
