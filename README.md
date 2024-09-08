# **How to Set Up `log/slog` in a Golang HTTP Server**

# 1. Import `log/slog` in the logging configuration file and create a global variable for the logger object:

```go
import (
    "log/slog"
)

var Logger *slog.Logger
```

# 2. Create and open a log file:

```go
func LogInit() {
    file, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        panic(err)
    }
    ...
```

# 3. Set up the handler logs, define handler options, and create a logger with these handler options:

```go
    opts := &slog.HandlerOptions{
	AddSource: true,
	Level:     slog.LevelDebug,
    }
    handler := slog.NewTextHandler(file, opts)
    Logger = slog.New(handler)
}
```

# 4. Import the created logger in another package, initialize it, and set it as the default logger in `slog`:

## 4.1. Import:

```go
import (
	".../pkg/mylogger"
	"log/slog"
)
```

## 4.2. Initialization:

```go
mylogger.LogInit()
```

## 4.3. Set as the default logger:

```go
slog.SetDefault(mylogger.Logger)
```