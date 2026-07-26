# tslog

[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/thorsphere/tslog)](https://pkg.go.dev/mod/github.com/thorsphere/tslog)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/thorsphere/tslog)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/thorsphere/tslog)
![GitHub Top Language](https://img.shields.io/github/languages/top/thorsphere/tslog)
[![CodeFactor](https://www.codefactor.io/repository/github/thorsphere/tslog/badge)](https://www.codefactor.io/repository/github/thorsphere/tslog)
![OSS Lifecycle](https://img.shields.io/osslifecycle/thorsphere/tslog)

Go package for logging.

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
func Trace(msg string)
func Debug(msg string) 
func Info(msg string)
func Warn(msg string)
func Error(err error)
func Fatal(err error)
```

Log levels `Error` and `Fatal` receive an error for logging.
An error can be retrieved with package [tserr](https://pkg.go.dev/github.com/thorsphere/tserr), func [New](https://pkg.go.dev/errors#New) or with func [Errorf](https://pkg.go.dev/fmt#Errorf)

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
TraceLevel Level = Level(slog.LevelDebug - 4)

// Debug: log detailed events for debugging of the app
DebugLevel Level = Level(slog.LevelDebug)

// Info: log an event under normal conditions of the app
InfoLevel Level = Level(slog.LevelInfo)

// Warn: log an unintended event, which is tried to be recovered and potentially
// impacting execution of the app
WarnLevel Level = Level(slog.LevelWarn)

// Error: log an unexpected event with at least one function of the app being not operable
ErrorLevel Level = Level(slog.LevelError)

// Fatal: log an unexpected critical event forcing a shutdown of the app
FatalLevel Level = Level(slog.LevelError + 4)
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

## Documentation & Resources

- [Go Package Documentation](https://pkg.go.dev/github.com/thorsphere/tslog) — Complete API reference
- [Open Source Insights](https://deps.dev/go/github.com%2Fthorsphere%2Ftslog) — Dependency analysis

## ⚖️ License & Commercial Usage

Copyright (c) 2023-2026 thorsphere. All rights reserved.

This project is licensed under the **Functional Source License v1.1 (FSL-1.1-ALv2)**. 

* The use, modification, and redistribution of this Go package is completely free for private, educational, non-commercial, and internal purposes. 
* If you are a company or institution looking to use this package in a commercial product, service, or business environment, you must secure a commercial license.
* Each version of this software automatically converts to the fully open-source Apache License, Version 2.0 on the second anniversary of its release.

For full details, please see the [LICENSE](LICENSE.md) file.

### 💼 Commercial Licensing & Inquiries

To purchase a commercial license or discuss support options, please reach out directly:

* 📩 **Contact:** business at thorsphere dot com
* 💬 **Response Time:** Usually within a couple of business days.

*Please include your company name and a brief overview of your use case so I can provide the right licensing details.*
