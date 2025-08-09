# Go Logging Package

A comprehensive Go logging package that provides structured logging capabilities with advanced debug options and custom error handling.

## Overview

This logging package wraps the `logr.Logger` interface to provide:

- **Structured JSON logging** with source information
- **Custom error handling** for `TradingError` types
- **Advanced debug capabilities** with runtime statistics
- **Global logger instance** for easy access across your application

## Installation

```bash
go get github.com/luciano-personal-org/logging
```

## Dependencies

- `github.com/go-logr/logr v1.4.2` - Structured logging interface
- `github.com/luciano-personal-org/exception v0.1.18` - Custom exception handling

## Quick Start

### Basic Setup

```go
package main

import (
    "github.com/luciano-personal-org/logging"
)

func main() {
    // Initialize the global logger
    logging.InitLogger("my-app", "main-module")
    
    // Use the global logger instance
    logging.Logging.Info("Application started", "version", "1.0.0")
}
```

### Logging Methods

#### Info Logging
```go
logging.Logging.Info("User logged in", "userID", 123, "email", "user@example.com")
```

#### Error Logging
```go
// Requires a TradingError type
err := exception.NewTradingError("USER_NOT_FOUND", "User does not exist", originalErr)
logging.Logging.Error(err, "userID", 123)
```

#### Debug Logging
```go
// Basic debug
logging.Logging.Debug("Processing request", logging.DebugOptions{
    Enabled: true,
    Level:   logging.INFO,
}, "requestID", "req-123")

// Debug with memory statistics
logging.Logging.Debug("Memory usage check", logging.DebugOptions{
    Enabled: true,
    Level:   logging.MEM,
}, "operation", "data-processing")

// Debug with all information
logging.Logging.Debug("Full system debug", logging.DebugOptions{
    Enabled: true,
    Level:   logging.ALL,
})
```

## Debug Levels

The package supports multiple debug levels:

- `INFO`: Basic debug information
- `STACK`: Include stack trace
- `MEM`: Memory statistics (heap, stack usage)
- `GC`: Garbage collection statistics
- `BUILD`: Build information
- `ALL`: All available debug information

## API Reference

### Types

#### Logger
```go
type Logger struct {
    Logs logr.Logger
}
```

#### DebugOptions
```go
type DebugOptions struct {
    Enabled bool   // Enable/disable debug output
    Level   string // Debug level (INFO, STACK, MEM, GC, BUILD, ALL)
}
```

### Functions

#### InitLogger
```go
func InitLogger(appName string, appModuleName string)
```
Initializes the global logger with JSON output to stderr.

#### Info
```go
func (l *Logger) Info(msg string, keysAndValues ...interface{})
```
Logs an info message with optional key-value pairs.

#### Error
```go
func (l *Logger) Error(err exception.TradingError, keysAndValues ...interface{})
```
Logs an error with detailed error information from TradingError.

#### Debug
```go
func (logger *Logger) Debug(msg string, options DebugOptions, keysAndValues ...interface{})
```
Logs debug information with configurable detail levels.

### Global Instance

The package provides a global logger instance:

```go
var Logging Logger
```

Access it after initialization:

```go
logging.InitLogger("my-app", "my-module")
logging.Logging.Info("Ready to log!")
```

## Example Usage

```go
package main

import (
    "errors"
    "github.com/luciano-personal-org/logging"
    "github.com/luciano-personal-org/exception"
)

func main() {
    // Initialize logger
    logging.InitLogger("trading-app", "order-service")
    
    // Info logging
    logging.Logging.Info("Service started", 
        "port", 8080, 
        "environment", "production")
    
    // Error logging with custom error type
    originalErr := errors.New("database connection failed")
    tradingErr := exception.NewTradingError(
        "DB_CONNECTION_ERROR", 
        "Failed to connect to trading database", 
        originalErr)
    
    logging.Logging.Error(tradingErr, 
        "host", "db.trading.com", 
        "port", 5432)
    
    // Debug with memory stats during heavy processing
    logging.Logging.Debug("Processing large order batch", 
        logging.DebugOptions{
            Enabled: true,
            Level:   logging.MEM,
        }, 
        "batchSize", 10000,
        "operation", "bulk-insert")
}
```

## Development

### Building
```bash
go build
```

### Testing
```bash
go test ./...
```

### Formatting
```bash
go fmt ./...
go vet ./...
```

### Dependencies
```bash
go mod tidy
```

## License

See LICENSE file for details.
