package logging

import (
	"log/slog"
	"os"

	"github.com/go-logr/logr"
)

var Log logr.Logger

func InitLogger(appName string, appModuleName string) {
	handler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.Level(-1),
	})
	logger := logr.FromSlogHandler(handler)
	Log = logger.WithName(appName)
	Log.WithValues(appName, appModuleName)
}
