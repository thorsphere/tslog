// Copyright (c) 2023-2026 thorsphere
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tslog

// Import standard library packages and tserr.
import (
	"context"  // context
	"io"       // io
	"log/slog" // slog
	"os"       // os

	"github.com/thorsphere/tsfio"  // tsfio
	"github.com/thorstenrie/tserr" // tserr
)

// closeOut closes the output file for logging, if it is owned by the logger and
// returns an error, if any. If the output file is not owned by the logger,
// e.g. os.Stdout or os.Stderr, closeOut does not close the file and returns nil.
func (l *Logger) closeOut() error {
	// Do not close Stdout or Stderr
	if !l.outOwned {
		// Do not close if io.Writer is not owned by the logger, e.g. os.Stdout or os.Stderr,
		// and return nil
		l.out = nil
		l.outFn = ""
		return nil
	}

	// If the output file has a Close method, close the file and return an error, if any
	if closer, ok := l.out.(io.Closer); ok {
		// If the file is os.Stdout or os.Stderr, do not close it and skip
		if f, ok := closer.(*os.File); ok && (f == os.Stdout || f == os.Stderr) {
			// Do not close os.Stdout or os.Stderr and skip
		} else if err := closer.Close(); err != nil { // Close the file and return an error, if any
			// Return an error for closing the file
			return tserr.Op(&tserr.OpArgs{Op: "close log file", Fn: string(l.outFn), Err: err})
		}
	}
	// Set l.out to nil
	l.out = nil
	// Set logging output filename to an empty string
	l.outFn = ""
	// Set l.outOwned to false
	l.outOwned = false
	// Return nil
	return nil
}

// setFile sets logging to a io.Writer f with filename fn and retrieves whether
// the file is owned by the logger.
func (l *Logger) setFile(f io.Writer, fn tsfio.Filename, o bool) {
	// Set l.out to f
	l.out = f
	// Set logging output filename to fn
	l.outFn = fn
	// Set l.outOwned to o
	l.outOwned = o
	// Create a new JSON handler for logging to file f with minimum level l.minLvl
	h := slog.NewJSONHandler(f, &slog.HandlerOptions{AddSource: false, Level: l.minLvl, ReplaceAttr: level})
	// Set l.logger to a new slog.Logger with the created handler
	l.logger = slog.New(h)
}

// setStdout sets logging to Stdout.
func (l *Logger) setStdout() {
	// Set logging output to Stdout
	l.setFile(os.Stdout, StdoutLogger, false)
}

// noLogger sets logging to discard logging.
func (l *Logger) noLogger() {
	// Set logging output to discard
	l.setFile(io.Discard, DiscardLogger, false)
}

// trylog logs message msg, if lvl is equal to or higher than the
// minimum log level.
func (l *Logger) tryLog(lvl Level, msg string) {
	l.logger.Log(context.Background(), slog.Level(lvl), msg)
}

// level implements ReplaceAttr of type HandlerOptions. It changes the way levels are
// printed for both the standard log levels and the custom log levels.
func level(groups []string, a slog.Attr) slog.Attr {
	// If the attribute is the log level, change its value to the string representation
	// of the log level
	if a.Key == slog.LevelKey {
		// Retrieve the log level as slog.Level
		lvl := Level(a.Value.Any().(slog.Level))
		// Return the string representation for log level lvl
		switch {
		case lvl <= TraceLevel:
			a.Value = slog.StringValue(traceString)
		case lvl <= DebugLevel:
			a.Value = slog.StringValue(debugString)
		case lvl <= InfoLevel:
			a.Value = slog.StringValue(infoString)
		case lvl <= WarnLevel:
			a.Value = slog.StringValue(warnString)
		case lvl <= ErrorLevel:
			a.Value = slog.StringValue(errorString)
		default:
			a.Value = slog.StringValue(fatalString)
		}
	}
	// If the attribute is the time, format it according to timeLayout
	if a.Key == slog.TimeKey {
		// Format the time using the timeLayout constant
		a.Value = slog.StringValue(a.Value.Time().Format(timeLayout))
	}
	// Return the attribute
	return a

}
