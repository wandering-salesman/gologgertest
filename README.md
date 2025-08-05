# Function Context based Logger

This workspace demonstrates a structured logging package in Go using [zap](https://github.com/uber-go/zap), with dynamic per-function log levels and context-based metadata.

## Structure

- `loggerpackage`: Custom logger implementation.
- `loggertest`: Example usage with a `UserService`.

## Features

- Attach logger to context.
- Set log levels for specific functions at runtime.
- Add metadata to logs via context.
- Use sugared logger for easy formatting.

## Usage

```sh
cd loggertest
go run main.go
```

## Requirements

- Go 1.23.5+
- [zap](https://github.com/uber-go/zap)

See [`main.go`](./loggertest/main.go) for usage details.
