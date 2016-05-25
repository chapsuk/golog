package main

import (
	log "github.com/chapsuk/golog"
	"os"
)

type Example struct {
	logger log.ILogger
}

func (e Example) WriteInfo(msg string) {
	e.logger.Info(msg)
}

func main() {
	msg := "error: why so serious?!"
	log.AppendContext(log.Context{
		"host":    "localhost",
		"channel": "Root",
	})

	log.ErrorCtx(log.Context{
		"formatter": "text",
	}, msg)

	log.SetFormatter(&log.JSONFormatter{})
	log.ErrorCtx(log.Context{
		"formatter": "json",
	}, msg)

	log.SetFormatter(&log.LogstashFormatter{})
	log.ErrorCtx(log.Context{
		"formatter": "logstash",
	}, msg)

	newLogger := log.NewLogger(os.Stdout, &log.TextFormatter{}, log.Context{
		"channel": "NewLogger",
	})
	newLogger.Info("hallo")
	log.Info("I`m here")

	ex := &Example{
		logger: newLogger,
	}
	ex.WriteInfo("mmmm")
}
