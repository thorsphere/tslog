// Package tslog implements logging that tries to keep it simple.
//
// The tslog package is a logging interface in Go that tries to keep it simple.
// It provides log levels Trace, Debug, Info, Warn, Error and Fatal.
// The log messages are formatted in JSON format to enable parsing.
// The predefined default logger is set to log to Stdout on Info level. A new
// logger instance can be created with New(). The output of a logger can be set
// to a specific file, a temporary file, to Stdout and to discard.
// All function calls return an error, if any.
//
// Copyright (c) 2023-2026 thorsphere
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tslog

// Import tsfio.
import (
	"log/slog"
	"time"

	"github.com/thorsphere/tsfio" // tsfio
)

// Level represents the severity of a logging event.
// The levels are Trace, Debug, Info, Warn, Error and Fatal.
type (
	Level slog.Level
)

// Strings for special loggers
const (
	StdoutLogger  tsfio.Filename = tsfio.Filename("stdout")  // Stdout
	DiscardLogger tsfio.Filename = tsfio.Filename("discard") // discard, no logging
	TmpLogger     tsfio.Filename = tsfio.Filename("tmp")     // temporary file
)

// Levels for log levels.
const (
	// Trace: log the execution of code of the app
	traceLevel Level = Level(slog.LevelDebug - 4)
	// Debug: log detailed events for debugging of the app
	debugLevel Level = Level(slog.LevelDebug)
	// Info: log an event under normal conditions of the app
	infoLevel Level = Level(slog.LevelInfo)
	// Warn: log an unintended event, which is tried to be recovered and potentially
	// impacting execution of the app
	warnLevel Level = Level(slog.LevelWarn)
	// Error: log an unexpected event with at least one function of the app being not operable
	errorLevel Level = Level(slog.LevelError)
	// Fatal: log an unexpected critical event forcing a shutdown of the app
	fatalLevel Level = Level(slog.LevelError + 4)
)

// Strings for log levels as string.
const (
	traceString string = "trace" // Trace level as string
	debugString string = "debug" // Debug level as string
	infoString  string = "info"  // Info level as string
	warnString  string = "warn"  // Warn level as string
	errorString string = "error" // Error level as string
	fatalString string = "fatal" // Fatal level as string
)

// Defaults for logging
const (
	// Layout for timestamp in the log message
	timeLayout string = time.RFC3339Nano
	// Root element for temporary file
	defaultPattern string = "tslog"
	// Default log level is InfoLevel
	defaultMinLvl Level = infoLevel
)

// Global logger to provide a predefined standard logger
var (
	globalLogger *Logger = New()
)

// Default returns the global predefined standard logger
func Default() *Logger {
	return globalLogger
}

// SetLevel sets the logging level. All levels equal or higher than the set level
// are logged. All log messages with levels below the set level are discarded.
// SetLevel returns an error for undefined levels, otherwise nil.
func SetLevel(level Level) error {
	return globalLogger.SetLevel(level)
}

// SetOutput sets the logging output to fn. Special loggers are
// 'stdout' for logging to Stdout (default)
// 'discard' for no logging
// 'tmp' for logging to tslog_* in the temporary directory
// If SetOuput returns an error, logging is set to Stdout
func SetOutput(fn tsfio.Filename) error {
	return globalLogger.SetOutput(fn)
}

// Trace logs a message at Trace level on the global predefined standard logger.
func Trace(msg string) {
	globalLogger.Trace(msg)
}

// Debug logs a message at Debug level on the global predefined standard logger.
func Debug(msg string) {
	globalLogger.Debug(msg)
}

// Info logs a message at Info level on the global predefined standard logger.
func Info(msg string) {
	globalLogger.Info(msg)
}

// Warn logs a message at Warn level on the global predefined standard logger.
func Warn(msg string) {
	globalLogger.Warn(msg)
}

// Error logs error err at Error level on the global predefined standard logger.
func Error(err error) {
	globalLogger.Error(err)
}

// Fatal logs error err at Fatal level on the global predefined standard logger.
func Fatal(err error) {
	globalLogger.Fatal(err)
}
