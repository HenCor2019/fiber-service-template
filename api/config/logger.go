package config

import (
	"log/slog"
	"os"
)

func InitLogger() *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	return logger
}
