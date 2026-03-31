# tslog

[![Go Report Card](https://goreportcard.com/badge/github.com/thorsphere/tslog)](https://goreportcard.com/report/github.com/thorsphere/tslog)
[![CodeFactor](https://www.codefactor.io/repository/github/thorsphere/tslog/badge)](https://www.codefactor.io/repository/github/thorsphere/tslog)
![OSS Lifecycle](https://img.shields.io/osslifecycle/thorsphere/tslog)

[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/thorsphere/tslog)](https://pkg.go.dev/mod/github.com/thorsphere/tslog)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/thorsphere/tslog)
![Libraries.io dependency status for GitHub repo](https://img.shields.io/librariesio/github/thorsphere/tslog)

![GitHub release (latest by date)](https://img.shields.io/github/v/release/thorsphere/tslog)
![GitHub last commit](https://img.shields.io/github/last-commit/thorsphere/tslog)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/thorsphere/tslog)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/thorsphere/tslog)
![GitHub Top Language](https://img.shields.io/github/languages/top/thorsphere/tslog)
![GitHub](https://img.shields.io/github/license/thorsphere/tslog)

[Go](https://go.dev/) package for logging that tries to keep it simple ([KISS principle](https://en.wikipedia.org/wiki/KISS_principle)).

- **Simple**: Pre-defined global logger to Stdout without configuration and log levels Trace, Debug, Info, Warn, Error and Fatal.
- **Easy to parse**: The log messages are formatted in JSON format.
- **Flexible**: Logging can be configured to stdout (default), to a temp file, a specifically defined file or even discarded.
- **Tested**: Unit tests with high code coverage
- **Dependencies**: Only depends on [Go Standard Library](https://pkg.go.dev/std), [tsfio](https://pkg.go.dev/github.com/thorsphere/tsfio) and [tserr](https://pkg.go.dev/github.com/thorsphere/tserr)

## Usage

In the Go app, the package is imported with

```
import "github.com/thorsphere/tslog"
```

A tslog logger is based on type [Logger](https://pkg.go.dev/slog#Logger) defined in Go Standard package [slog](https://pkg.go.dev/slog).

## Default logger

The predefined default logger is set to log to Stdout on Info level. The default logger can be used with the external functions

```
func Trace(msg string) error
func Debug(msg string) error 
func Info(msg string) error
func Warn(msg string) error
func Error(err error) error
func Fatal(err error) error
```

Log levels `Error` and `Fatal` receive an error for logging.
An error can be retrieved with package [tserr](https://pkg.go.dev/github.com/thorsphere/tserr), func [New](https://pkg.go.dev/errors#New) or or with func [Errorf](https://pkg.go.dev/fmr#Errorf)

The default logger can be retrieved with

```
func Default() *Logger 
```

A new logger instance can be created with

```
func New() *Logger
```

## Configuration

A logger can be configured to log to stdout (default), a temporary file, a specific file or logging can be discarded (no logging).

The following configurations are available

- `stdout`: Log to Stdout (default)
- `tmp`: logging to `tslog_*` in the operating system temporary directory, where `*` stands for a random string (see [os.CreateTemp](https://pkg.go.dev/os#CreateTemp))
- `discard`: no logging
- `<filename>`: for logging to <filename>

Therefore, `stdout`, `tmp`, `discard` are reserved keywords. If none of the keywords apply, the provided argument will be
treated as filename. If and error occurs, then tslog will fall back to the default logging to Stdout.

The output is configured with

```
func (l *Logger) SetOutput(fn tsfio.Filename) error 
```

A logger can be configured to log from a specific level and any higher level. The levels are defined as

```
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
```

The log level is set with

```
func (l *Logger) SetLevel(level Level) error
```

## Output

The log messages are formatted in the JSON format. Each log message has the field "level" which is a string respresentation of the log level, the field "msg" and timestamp field "time". The timestamp has the format

```
// Layout for timestamp in the log message
timeLayout string = time.RFC3339Nano
```

## Example

```
package main

import (
    "errors"
    "github.com/thorsphere/tslog"
)

func main() {
    l := tslog.Default()
    l.SetLevel(tslog.TraceLevel)
    l.SetOutput("stdout")
    l.Trace("trace")
    l.Debug("debug")
    l.Info("info")
    l.Warn("warn")
    l.Error(errors.New("error"))
    l.Fatal(errors.New("fatal"))
}
```

[Go Playground](https://go.dev/play/p/lWrvK4UqDTD)

Output
```
{"time":"2009-11-10T23:00:00Z","level":"trace","msg":"trace"}
{"time":"2009-11-10T23:00:00Z","level":"debug","msg":"debug"}
{"time":"2009-11-10T23:00:00Z","level":"info","msg":"info"}
{"time":"2009-11-10T23:00:00Z","level":"warn","msg":"warn"}
{"time":"2009-11-10T23:00:00Z","level":"error","msg":"error"}
{"time":"2009-11-10T23:00:00Z","level":"fatal","msg":"fatal"}
```

## Links

[Godoc](https://pkg.go.dev/github.com/thorsphere/tslog)

[Go Report Card](https://goreportcard.com/report/github.com/thorsphere/tslog)

[Open Source Insights](https://deps.dev/go/github.com%2Fthorsphere%2Ftslog)
