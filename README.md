# Go Logger for Skyfora

A reusable logger for Skyfora's Go project.

## Features

- Configurable logging for debug / production
- Configuration for separate level files

## Installation

```bash
go get github.com/skyfora/go-logger
```

## Usage

```go
import (
    "github.com/skyfora/go-logger"
)

func main() {
    logger.Init(logger.Logger{
        FilePath: "logs.log",
        WithStdout: true,
    })
    defer logger.Sync()

    logger.Info("Hello, World!")
    logger.Debug("Hello, World!")
    logger.Error("Hello, World!")
    logger.Fatal("Hello, World!")
    logger.Warn("Hello, World!")
}
```

## Usage with separate level files

```go
logger.InitEmpty()
```
