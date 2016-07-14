package golog

import (
	"bytes"
	"io"
)

// Level logging
type Level uint8

// Context is log message context
type Context map[string]interface{}

// Formatter format log message to needed output
type Formatter interface {
	Format(*bytes.Buffer, Level, Context, string) *bytes.Buffer
}

// StandartLogger is log package interface
type StandardLogger interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})

	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})
}

// ILogger interface
type ILogger interface {
	WithContext(Context) ILogger
	GetContext() Context

	SetFormatter(Formatter)
	SetOutput(io.Writer)
	SetLevel(Level)

	Print(...interface{})
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Fatal(...interface{})
	Panic(...interface{})

	Printf(string, ...interface{})
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Fatalf(string, ...interface{})
	Panicf(string, ...interface{})

	Println(...interface{})
	Debugln(...interface{})
	Infoln(...interface{})
	Warnln(...interface{})
	Errorln(...interface{})
	Fatalln(...interface{})
	Panicln(...interface{})

	DebugCtx(Context, ...interface{})
	InfoCtx(Context, ...interface{})
	WarnCtx(Context, ...interface{})
	ErrorCtx(Context, ...interface{})
	FatalCtx(Context, ...interface{})
	PanicCtx(Context, ...interface{})

	DebugfCtx(Context, string, ...interface{})
	InfofCtx(Context, string, ...interface{})
	WarnfCtx(Context, string, ...interface{})
	ErrorfCtx(Context, string, ...interface{})
	FatalfCtx(Context, string, ...interface{})
	PanicfCtx(Context, string, ...interface{})
}

// Pool of buffers for help gc and reduce memory allocate
type Pool interface {
	Get() *bytes.Buffer
	Put(*bytes.Buffer)
}
