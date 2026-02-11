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
- **Tested**: Unit tests with high [code coverage](https://gocover.io/github.com/thorsphere/tslog)
- **Dependencies**: Only depends on [Go Standard Library](https://pkg.go.dev/std), [tsfio](https://gocover.io/github.com/thorsphere/tsfio) and [tserr](https://gocover.io/github.com/thorsphere/tserr)

## Usage

In the Go app, the package is imported with

```
import "github.com/thorsphere/tslog"
```

A tslog logger is based on type [Logger](https://pkg.go.dev/log#Logger) defined in Go Standard package [log](https://pkg.go.dev/log).

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
An error can be retrieved with package [tserr](https://gocover.io/github.com/thorsphere/tserr), func [New](https://pkg.go.dev/errors#New) or or with func [Errorf](https://pkg.go.dev/fmr#Errorf)

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
TraceLevel int = 1

// Debug: log detailed events for debugging of the app
DebugLevel int = 2

// Info: log an event under normal conditions of the app
InfoLevel int = 3

// Warn: log an unintended event, which is tried to be recovered and potentially
// impacting execution of the app
WarnLevel int = 4

// Error: log an unexpected event with at least one function of the app being not operable
ErrorLevel int = 5

// Fatal: log an unexpected critical event forcing a shutdown of the app
FatalLevel int = 6
```

The log level is set with

```
func (l *Logger) SetLevel(level int) error
```

## Output

The log messages are formatted in the JSON format. The root element is named `log`. Each log message has the field "level" which is a string respresentation of the log level, the field "message" and timestamp field "time". The timestamp has the format

```
// Layout for timestamp in the log message
timeLayout string = "2006-01-02 15:04:05 -0700 MST"
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

Output

```
{"log":{"level":"trace","message":"trace","time":"2023-01-21 12:51:56 +0100 CET"}}
{"log":{"level":"debug","message":"debug","time":"2023-01-21 12:51:56 +0100 CET"}}
{"log":{"level":"info","message":"info","time":"2023-01-21 12:51:56 +0100 CET"}}
{"log":{"level":"warn","message":"warn","time":"2023-01-21 12:51:56 +0100 CET"}}
{"log":{"level":"error","message":"error","time":"2023-01-21 12:51:56 +0100 CET"}}
{"log":{"level":"fatal","message":"fatal","time":"2023-01-21 12:51:56 +0100 CET"}}
```

## Links

[Godoc](https://pkg.go.dev/github.com/thorsphere/tslog)

[Gocover.io](https://gocover.io/github.com/thorsphere/tslog)

[Go Report Card](https://goreportcard.com/report/github.com/thorsphere/tslog)

[Open Source Insights](https://deps.dev/go/github.com%2Fthorsphere%2Ftslog)
