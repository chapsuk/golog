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

// ILogger interface
type ILogger interface {
	AppendContext(Context) ILogger
	GetContext() Context

	SetFormatter(Formatter)
	SetOutput(io.Writer)
	SetLevel(Level)

	Error(string)
	ErrorCtx(Context, string)
	Errorf(string, ...interface{})
	ErrorfCtx(Context, string, ...interface{})

	Debug(string)
	DebugCtx(Context, string)
	Debugf(string, ...interface{})
	DebugfCtx(Context, string, ...interface{})

	Info(string)
	InfoCtx(Context, string)
	Infof(string, ...interface{})
	InfofCtx(Context, string, ...interface{})

	Warn(string)
	WarnCtx(Context, string)
	Warnf(string, ...interface{})
	WarnfCtx(Context, string, ...interface{})

	Fatal(string)
	FatalCtx(Context, string)
	Fatalf(string, ...interface{})
	FatalfCtx(Context, string, ...interface{})

	Panic(string)
	PanicCtx(Context, string)
	Panicf(string, ...interface{})
	PanicfCtx(Context, string, ...interface{})
}

// Pool of buffers for help gc and reduce memory allocate
type Pool interface {
	Get() *bytes.Buffer
	Put(*bytes.Buffer)
}
