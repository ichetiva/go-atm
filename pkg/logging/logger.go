package logging

import (
	"log/slog"
	"os"
)

func NewLogger(level string) *slog.Logger {
	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: getLevel(level),
	})
	return slog.New(h)
}

func getLevel(level string) slog.Leveler {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "info":
		return slog.LevelInfo
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
