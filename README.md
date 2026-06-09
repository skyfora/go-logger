# Go Logger for Skyfora

A reusable structured logger for Skyfora's Go projects. Built on [zap](https://github.com/uber-go/zap).

## Features

- JSON structured logging to stdout (default)
- Optional human-readable console output for local development or tests
- Optional rotated JSON file sinks
- Per-level file separation
- Service identity fields (`service`, `env`) injected at init

## Installation

Follow the instructions from [Go Official Website](https://go.dev/doc/faq#git_https)

Configure the ~/.sshconfig to use SSH on https url, by adding the following:

```bash
[url "ssh://git@github.com/"]
    insteadOf = https://github.com/
```

Or use password authentication by modifying the ~/.netrc file:

```
machine github.com
login <username>
password <password>
```

Then, you will also need to modify the GOPRIVATE environment variable:

```bash
(For Linux/MacOS)
export GOPRIVATE=github.com/skyfora/*
or
go env -w GOPRIVATE='github.com/skyfora/*'
```

Installation on the project:

```bash
go get github.com/skyfora/go-logger
```

## Usage

### Standard service init

```go
import (
    "github.com/skyfora/go-logger"
    "go.uber.org/zap"
)

func main() {
    logger.InitService("my-service", "prod", false)
    defer logger.Sync()

    logger.Info("started", zap.String("version", "1.0.0"))
}
```

### Manual configuration

```go
logger.Init(logger.Logger{
    WithStdout:  true,
    Debug:       false,
    ServiceName: "my-service",
    Environment: "prod",
})
defer logger.Sync()

logger.Info("Hello", zap.String("key", "value"))
logger.Debug("debug detail")
logger.Error("something failed", zap.Error(err))
```

Stdout emits one JSON object per line with keys: `timestamp`, `level`, `msg`, `caller`, plus any custom `zap.Field` values.

### Console output (tests / local)

```go
logger.Init(logger.Logger{
    WithStdout:    true,
    ConsoleStdout: true,
    Debug:         true,
})
```

### Separate level files

```go
logger.Init(logger.Logger{
    FilePath:      "logs/app.log",
    SeparateLevel: true,
    WithStdout:    true,
})
```

This writes JSON to rotated files per level (`logs/app.log.info`, `logs/app.log.error`, etc.) in addition to stdout.

### Silent logger (tests)

```go
logger.InitEmpty()
```
