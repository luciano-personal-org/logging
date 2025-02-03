package logging

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"runtime/debug"

	globaldebug "runtime/debug"

	"github.com/go-logr/logr"
	exception "github.com/luciano-personal-org/exception"
)

// DebugOptions is a type for debug options
const (
	INFO  = "INFO"  // Default
	STACK = "STACK" // Stack Trace
	MEM   = "MEM"   // Memory Stats
	GC    = "GC"    // GC Stats
	BUILD = "BUILD" // Build Info
	ALL   = "ALL"   // All Stats
)

type DebugOptions struct {
	Enabled bool
	Level   string
}

// isValidOption validates the debug option
func isValidOption(level string) bool {
	switch level {
	case INFO, STACK, MEM, GC, BUILD, ALL:
		return true
	default:
		return false
	}
}

type Level int

const (
	DEBUG Level = 1
)

// String returns the string representation of the SignalDirection.
func (d Level) String() string {
	return [...]string{"INFO", "DEBUG"}[d]
}

type Logger struct {
	Logs logr.Logger
}

var Logging Logger

func InitLogger(appName string, appModuleName string, logLevel string) {

	level := 0

	if logLevel == DEBUG.String() {
		level = int(DEBUG)
	}

	handler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.Level(level),
	})

	logger := logr.FromSlogHandler(handler)
	Logging.Logs = logger.WithName(appName)
	Logging.Logs.WithValues(appName, appModuleName)
}

func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
	l.Logs.Info(msg, keysAndValues...)
}

func (l *Logger) Error(err exception.TradingError, keysAndValues ...interface{}) {

	message := fmt.Sprintf("Error: %s, ErrorCode: %s, Details: %s", err.Error(), err.ErrorCode(), err.Details())

	l.Logs.Error(err.OriginalError(), message, keysAndValues...)
}

func (logger *Logger) Debug(msg string, options DebugOptions, keysAndValues ...interface{}) {

	logger.Logs.Info(msg, keysAndValues...)

	var level = options.Level
	// If Debug is disabled, return
	if !options.Enabled {
		return
	}
	// Default Level
	if !isValidOption(level) {
		msg := fmt.Sprintf("Invalid debug option: %s", level)
		logger.Logs.Error(errors.New(msg), "")
	}
	// Read GC Stats
	var gcStats globaldebug.GCStats
	globaldebug.ReadGCStats(&gcStats)
	// Read Mem Stats
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	// Read Build Info
	buildInfo, ok := globaldebug.ReadBuildInfo()

	// Business Message
	logger.Logs.WithCallDepth(2).Info(msg)
	// Stack Trace
	if level == STACK || level == ALL {
		logger.Logs.WithCallDepth(2).Info(fmt.Sprintf("Stack Trace:%s", debug.Stack()))
	}
	// Mem Stats
	if level == MEM || level == ALL {
		logger.Logs.WithCallDepth(2).Info(fmt.Sprintf("Alloc: %d bytes", memStats.Alloc))
		logger.Logs.WithCallDepth(2).Info(fmt.Sprintf("TotalAlloc: %d bytes", memStats.TotalAlloc))
		logger.Logs.WithCallDepth(2).Info(fmt.Sprintf("HeapAlloc: %d bytes", memStats.HeapAlloc))
		logger.Logs.WithCallDepth(2).Info(fmt.Sprintf("HeapSys: %d bytes", memStats.HeapSys))
		logger.Logs.WithCallDepth(2).Info(fmt.Sprintf("HeapIdle: %d bytes", memStats.HeapIdle))
		logger.Logs.WithCallDepth(2).Info(fmt.Sprintf("HeapInuse: %d bytes", memStats.HeapInuse))
		logger.Logs.WithCallDepth(2).Info(fmt.Sprintf("HeapReleased: %d bytes", memStats.HeapReleased))
		logger.Logs.WithCallDepth(2).Info(fmt.Sprintf("HeapObjects: %d", memStats.HeapObjects))
		logger.Logs.WithCallDepth(2).Info(fmt.Sprintf("StackInUse: %d bytes", memStats.StackInuse))
		logger.Logs.WithCallDepth(2).Info(fmt.Sprintf("StackSys: %d bytes", memStats.StackSys))
		logger.Logs.WithCallDepth(2).Info(fmt.Sprintf("NumGC: %d", memStats.NumGC))
	}
	// GC Stats
	if level == GC || level == ALL {
		logger.Logs.WithCallDepth(2).Info(fmt.Sprintf("LastGC: %v", gcStats.LastGC))
		logger.Logs.WithCallDepth(2).Info(fmt.Sprintf("NumGC: %v", gcStats.NumGC))
		logger.Logs.WithCallDepth(2).Info(fmt.Sprintf("PauseTotal: %v", gcStats.PauseTotal))
		logger.Logs.WithCallDepth(2).Info(fmt.Sprintf("Pause: %v", gcStats.Pause))
		logger.Logs.WithCallDepth(2).Info(fmt.Sprintf("PauseEnd: %v", gcStats.PauseEnd))
		logger.Logs.WithCallDepth(2).Info(fmt.Sprintf("PauseQuantiles: %v", gcStats.PauseQuantiles))
	}
	// Build Info
	if level == BUILD || level == ALL {
		if ok {
			logger.Logs.WithCallDepth(2).Info(fmt.Sprintf("Build Info: %s", buildInfo))
		}
	}

}
