package golog

import (
	"io"
)

var std = New()

func AppendContext(ctx Context) ILogger {
	return std.AppendContext(ctx)
}

func GetContext() Context {
	return std.GetContext()
}

func SetLevel(lvl Level) {
	std.SetLevel(lvl)
}

func SetOutput(o io.Writer) {
	std.SetOutput(o)
}

func SetFormatter(f Formatter) {
	std.SetFormatter(f)
}

func Error(m string) {
	std.Error(m)
}

func ErrorCtx(ctx Context, m string) {
	std.ErrorCtx(ctx, m)
}

func Errorf(f string, args ...interface{}) {
	std.Errorf(f, args)
}

func ErrorfCtx(ctx Context, f string, args ...interface{}) {
	std.ErrorfCtx(ctx, f, args)
}

func Debug(m string) {
	std.Debug(m)
}

func DebugCtx(ctx Context, m string) {
	std.DebugCtx(ctx, m)
}

func Debugf(f string, args ...interface{}) {
	std.Debugf(f, args)
}

func DebugfCtx(ctx Context, f string, args ...interface{}) {
	std.DebugfCtx(ctx, f, args)
}

func Info(m string) {
	std.Info(m)
}

func InfoCtx(ctx Context, m string) {
	std.InfoCtx(ctx, m)
}

func Infof(f string, args ...interface{}) {
	std.Infof(f, args)
}

func InfofCtx(ctx Context, f string, args ...interface{}) {
	std.InfofCtx(ctx, f, args)
}

func Warn(m string) {
	std.Warn(m)
}

func WarnCtx(ctx Context, m string) {
	std.WarnCtx(ctx, m)
}

func Warnf(f string, args ...interface{}) {
	std.Warnf(f, args)
}

func WarnfCtx(ctx Context, f string, args ...interface{}) {
	std.WarnfCtx(ctx, f, args)
}

func Fatal(m string) {
	std.Fatal(m)
}

func FatalCtx(ctx Context, m string) {
	std.FatalCtx(ctx, m)
}

func Fatalf(f string, args ...interface{}) {
	std.Fatalf(f, args)
}

func FatalfCtx(ctx Context, f string, args ...interface{}) {
	std.FatalfCtx(ctx, f, args)
}

func Panic(m string) {
	std.Panic(m)
}

func PanicCtx(ctx Context, m string) {
	std.PanicCtx(ctx, m)
}

func Panicf(f string, args ...interface{}) {
	std.Panicf(f, args)
}

func PanicfCtx(ctx Context, f string, args ...interface{}) {
	std.PanicfCtx(ctx, f, args)
}
