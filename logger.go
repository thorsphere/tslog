// Copyright (c) 2023-2026 thorsphere
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tslog

// Import standard library packages, tserr and tsfio.
import (
	"fmt"      // fmt
	"io"       // io
	"log/slog" // slog
	"os"       // os

	"github.com/thorstenrie/tserr" // tserr
	"github.com/thorstenrie/tsfio" // tsfio
)

// Logger contains a log.logger for logging and the minimum level for logging.
// The minimum level for logging is set with SetLevel.
type Logger struct {
	minLvl   *slog.LevelVar // minimum level for logging
	logger   *slog.Logger   // for logging
	out      io.Writer      // file for logging
	outFn    tsfio.Filename // filename for logging
	outOwned bool           // whether the output file is owned by the logger and should be closed by it
}

// New creates a new logger with default minimum level Info for logging. To alter
// the minimum level for logging use SetLevel. Logging is set to Stdout. To
// change logging output use SetOutput.
func New() *Logger {
	// Set output to os.Stdout
	o := os.Stdout
	// Create a new slog.LevelVar for the minimum level for logging
	l := new(slog.LevelVar)
	// Set the minimum level for logging to defaultMinLvl
	l.Set(slog.Level(defaultMinLvl))
	// Create a new slog.Handler for JSON logging with output o and minimum level l
	h := slog.NewJSONHandler(o, &slog.HandlerOptions{AddSource: false, Level: l, ReplaceAttr: level})
	// Return a new Logger with the created slog.Handler, output o and
	// output filename StdoutLogger
	return &Logger{
		minLvl:   l,
		logger:   slog.New(h),
		out:      o,
		outFn:    StdoutLogger,
		outOwned: false,
	}
}

// SetLevel sets the logging level. All levels equal or higher than the set level
// are logged. All log messages with levels below the set level are discarded.
// SetLevel returns an error for undefined levels, otherwise nil. If the provided
// level is lower than Trace level, the lowest level, the minimum level is set
// to Trace. If the provided level is higher than Fatal level, the highest level,
// the minimum level is set to Fatal.
func (l *Logger) SetLevel(level Level) error {
	// Initially set error e to nil
	var e error = nil
	// If level is lower than Trace, the lowest level, return an error and
	// set the minimum level to Trace.
	if level < traceLevel {
		// Set error to not existent
		e = tserr.NotExistent(fmt.Sprintf("log level %d", level))
		// Set minimum level to Trace level
		l.minLvl.Set(slog.Level(traceLevel))
		// If level is higher than Fatal, the highest level, return an error and
		// set the minimum level to Fatal.
	} else if level > fatalLevel {
		// Set error to not existent
		e = tserr.NotExistent(fmt.Sprintf("log level %d", level))
		// Set minimum level to Fatal level
		l.minLvl.Set(slog.Level(fatalLevel))
	} else {
		// Set minimum level to provided level
		l.minLvl.Set(slog.Level(level))
	}
	// Return e
	return e
}

// SetOutput sets the logging output to fn. Special loggers are
// 'stdout' for logging to Stdout (default)
// 'discard' for no logging
// 'tmp' for logging to tslog_* in the temporary directory
// If SetOuput returns an error, logging is set to Stdout
func (l *Logger) SetOutput(fn tsfio.Filename) error {

	if err := l.closeOut(); err != nil {
		// If closeOut returns an error, set logging output to Stdout and return an error
		l.setStdout()
		// Return error
		return tserr.Op(&tserr.OpArgs{Op: "close log output", Fn: string(l.outFn), Err: err})
	}

	// Handle special loggers
	switch fn {
	case DiscardLogger:
		// discard, no logging
		l.noLogger()
		// Return nil
		return nil
	case StdoutLogger:
		// Logging to Stdout
		l.setStdout()
		// Return nil
		return nil
	case TmpLogger:
		// Define pattern for the temporary file
		p := fmt.Sprintf("%v_*", defaultPattern)
		// Create temporary file for logging
		f, err := os.CreateTemp(os.TempDir(), p)
		// If it fails, return an error
		if err != nil {
			// Set logging output to Stdout
			l.setStdout()
			// Return error
			return tserr.Op(&tserr.OpArgs{Op: "create temp file", Fn: p, Err: err})
		}
		l.setFile(f, TmpLogger, true)
		// Return nil
		return nil
	}

	// Check filename using tsfio.CheckFile
	if err := tsfio.CheckFile(fn); err != nil {
		// If the check fails, set logging output to Stdout and return an error
		l.setStdout()
		// Return error
		return tserr.Check(&tserr.CheckArgs{F: string(fn), Err: err})
	}

	// Open file with filename fn
	f, e := tsfio.OpenFile(fn)
	// If OpenFile fails, set logging output to Stdout and return an error
	if e != nil {
		// Set logging output to Stdout
		l.setStdout()
		// Return error
		return tserr.Op(&tserr.OpArgs{Op: "open file", Fn: string(fn), Err: e})
	}
	// Set logging to file f with filename fn.
	l.setFile(f, fn, true)
	// Return nil
	return nil
}

// Trace logs a message at Trace level.
func (l *Logger) Trace(msg string) {
	l.tryLog(traceLevel, msg)
}

// Debug logs a message at Debug level.
func (l *Logger) Debug(msg string) {
	l.tryLog(debugLevel, msg)
}

// Info logs a message at Info level.
func (l *Logger) Info(msg string) {
	l.tryLog(infoLevel, msg)
}

// Warn logs a message at Warn level.
func (l *Logger) Warn(msg string) {
	l.tryLog(warnLevel, msg)
}

// Error logs error err at Error level.
func (l *Logger) Error(err error) {
	l.tryLog(errorLevel, err.Error())
}

// Fatal logs error err at Fatal level.
func (l *Logger) Fatal(err error) {
	l.tryLog(fatalLevel, err.Error())
}
