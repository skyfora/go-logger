# Go Logger for Skyfora

A reusable logger for Skyfora's Go project.

## Features

- Configurable logging for debug / production
- Configuration for separate level files

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

Instalation on the project:

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
