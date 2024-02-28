# logrus-hook-rotate

[![Go Report Card](https://goreportcard.com/badge/github.com/andoma-go/logrus-hook-rotate)](https://goreportcard.com/report/github.com/andoma-go/logrus-hook-rotate)
[![GoDoc](https://godoc.org/github.com/andoma-go/logrus-hook-rotate?status.svg)](https://godoc.org/github.com/andoma-go/logrus-hook-rotate)

This is a simple hook for logrus [andoma-go/logrus](https://github.com/andoma-go/logrus) to write log files using [natefinch/lumberjack](https://github.com/natefinch/lumberjack)

## Usage

```go
import (
	"github.com/andoma-go/logrus"
	nested "github.com/andoma-go/logrus-formatter-nested"
	rotatehook "github.com/andoma-go/logrus-hook-rotate"
	"github.com/mattn/go-colorable"
)

log := &logrus.Logger{
	Out: colorable.NewColorableStdout(),
	Formatter: &nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		HideKeys:        true,
	},
	Hooks: make(logrus.LevelHooks),
	Level: logrus.InfoLevel,
}

hook := rotatehook.NewRotateHook(&rotatehook.Config{
	Filename:   "debug.log",
	MaxSize:    5,
	MaxAge:     30,
	MaxBackups: 10,
	LocalTime:  true,
	Compress:   true,
	Formatter: &nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		HideKeys:        true,
		NoColors:        true,
	},
	Level:   logrus.DebugLevel,
	Enabled: true,
})

log.AddHook(hook)

log.Infoln("just info message")
log.Debugln("first debug message")

hook.SetEnabled(false)
log.Debugln("second debug message")
```
