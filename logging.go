package logging

import (
	"log/slog"
	"os"

	"github.com/go-logr/logr"
)

type Level int

const (
	INFO Level = iota
	DEBUG
)

// String returns the string representation of the SignalDirection.
func (d Level) String() string {
	return [...]string{"INFO", "DEBUG"}[d]
}

var Log logr.Logger

func InitLogger(appName string, appModuleName string, logLevel string) {
	handler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.Level(slog.LevelInfo),
	})
	logger := logr.FromSlogHandler(handler)
	Log = logger.WithName(appName)
	Log.WithValues(appName, appModuleName)
}
