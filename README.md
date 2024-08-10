# How to config logging in golang application with log/slog

## Step 1:
Import log/slog in your logging config file and create global variable for logger object:
```go
import (
    log/slog
)

var Logger *slog.Logger

```

## Step 2:
Create and open log-file:
```go
...
func LogInit() {
    file, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil { panic(err) }
...
```

## Step 3:
Create handler logs and logger with this handler:
```go
...
    handler := slog.NewTextHandler(file, nil)
    Logger := slog.New(handler)
}
...
```

## Step 4:
Import created logger in other package, initinitialize and set for default logger in slog:
Import:
```go
import (
	".../pkg/mylogger"
	"log/slog"
)
```

Initinitialize:
```go
mylogger.LogInit()
```

Set for default:
```go
slog.SetDefault(mylogger.Logger)
```
