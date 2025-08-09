# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go logging package (`github.com/luciano-personal-org/logging`) that provides structured logging capabilities with debug options. It's part of a larger algo-platform project structure.

## Dependencies

- `github.com/go-logr/logr v1.4.2` - Main logging interface
- `github.com/luciano-personal-org/exception v0.1.18` - Custom exception handling

## Core Architecture

The logging package is built around a single `Logger` struct that wraps `logr.Logger` and provides:

- **Structured JSON logging** via `slog.NewJSONHandler`
- **Error handling** specifically for `exception.TradingError` types
- **Debug capabilities** with multiple levels (INFO, STACK, MEM, GC, BUILD, ALL)
- **Global logger instance** available as `Logging`

Key components:
- `InitLogger()` - Initializes the global logger with app name and module name
- `Info()` - Standard info logging
- `Error()` - Specialized error logging for TradingError types  
- `Debug()` - Multi-level debug logging with runtime stats

## Common Commands

### Build and Test
```bash
go build          # Build the package
go test ./...     # Run all tests
go vet ./...      # Run Go vet for static analysis
go fmt ./...      # Format code
```

### Module Management
```bash
go mod tidy       # Clean up dependencies
go mod download   # Download dependencies
```

## Development Notes

- The README.md references a different project (Go-Lang API Scaffold) and should be updated to reflect this logging package
- Debug functionality provides detailed runtime information including memory stats, GC stats, and build info
- Error logging expects `exception.TradingError` types with specific methods (`ErrorCode()`, `Details()`, `OriginalError()`)