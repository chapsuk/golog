package golog

import (
	"io"
)

var std = New()

// WithContext yield new logger instance with context
// new context append to current context
func WithContext(ctx Context) ILogger {
	return std.WithContext(ctx)
}

// GetContext return current global log context
func GetContext() Context {
	return std.GetContext()
}

// GetLevel return current max log level
func GetLevel() Level {
	return std.GetLevel()
}

// SetLevel set max log level
func SetLevel(lvl Level) {
	std.SetLevel(lvl)
}

// SetOutput for logger messages
func SetOutput(o io.Writer) {
	std.SetOutput(o)
}

// SetFormatter set log message formatter
func SetFormatter(f Formatter) {
	std.SetFormatter(f)
}

// Print info log message
func Print(args ...interface{}) {
	std.Print(args...)
}

// Printf info log message, fmt.Sprintf style
func Printf(f string, args ...interface{}) {
	std.Printf(f, args...)
}

// Println not supported, only for standart logger interface compatibility
func Println(args ...interface{}) {
	std.Println(args...)
}

// Error write error log message
func Error(args ...interface{}) {
	std.Error(args...)
}

// ErrorCtx write error log message with context
func ErrorCtx(ctx Context, m string) {
	std.ErrorCtx(ctx, m)
}

// Errorf write error log message fmt.Printf style
func Errorf(f string, args ...interface{}) {
	std.Errorf(f, args...)
}

// ErrorfCtx write error log message with context, fmt.Printf style
func ErrorfCtx(ctx Context, f string, args ...interface{}) {
	std.ErrorfCtx(ctx, f, args...)
}

// Debug write debug log message
func Debug(args ...interface{}) {
	std.Debug(args...)
}

// DebugCtx write debug log message with context
func DebugCtx(ctx Context, m string) {
	std.DebugCtx(ctx, m)
}

// Debugf write debug log message, fmt.Sprtinf style
func Debugf(f string, args ...interface{}) {
	std.Debugf(f, args...)
}

// DebugfCtx write debug log message with context, fmt.Sprtinf style
func DebugfCtx(ctx Context, f string, args ...interface{}) {
	std.DebugfCtx(ctx, f, args...)
}

// Info write info log message
func Info(args ...interface{}) {
	std.Info(args...)
}

// InfoCtx write info log message with context
func InfoCtx(ctx Context, m string) {
	std.InfoCtx(ctx, m)
}

// Infof write info log message, fmt.Sprtinf style
func Infof(f string, args ...interface{}) {
	std.Infof(f, args...)
}

// InfofCtx write info log message with context, fmt.Sprtinf style
func InfofCtx(ctx Context, f string, args ...interface{}) {
	std.InfofCtx(ctx, f, args...)
}

// Warn write warning log message
func Warn(args ...interface{}) {
	std.Warn(args...)
}

// WarnCtx write warning log message with context
func WarnCtx(ctx Context, m string) {
	std.WarnCtx(ctx, m)
}

// Warnf write warning log message, fmt.Sprintf style
func Warnf(f string, args ...interface{}) {
	std.Warnf(f, args...)
}

// WarnfCtx write warning log message with context, fmt.Sprintf style
func WarnfCtx(ctx Context, f string, args ...interface{}) {
	std.WarnfCtx(ctx, f, args...)
}

// Fatal write fatal log message and call os.Exit
func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

// FatalCtx write fatal log message with context and call os.Exit
func FatalCtx(ctx Context, m string) {
	std.FatalCtx(ctx, m)
}

// Fatalf write fatal log message and call os.Exit, fmt.Sprtinf style
func Fatalf(f string, args ...interface{}) {
	std.Fatalf(f, args...)
}

// FatalfCtx write fatal log message with context and call os.Exit, fmt.Sprtinf style
func FatalfCtx(ctx Context, f string, args ...interface{}) {
	std.FatalfCtx(ctx, f, args...)
}

// Fatalln not supported, only for standart logger interface compatibility
// call os.Exit anyway
func Fatalln(args ...interface{}) {
	std.Fatalln(args...)
}

// Panic write panice log message and throw panic
func Panic(args ...interface{}) {
	std.Panic(args...)
}

// PanicCtx write panice log message with context and throw panic
func PanicCtx(ctx Context, m string) {
	std.PanicCtx(ctx, m)
}

// Panicf write panice log message and throw panic, fmt.Sprintf style
func Panicf(f string, args ...interface{}) {
	std.Panicf(f, args...)
}

// PanicfCtx write panice log message with context and throw panic, fmt.Sprintf style
func PanicfCtx(ctx Context, f string, args ...interface{}) {
	std.PanicfCtx(ctx, f, args...)
}

// Panicln not supported, only for standart logger interface compatibility
// throw panic anyway
func Panicln(args ...interface{}) {
	std.Panicln(args...)
}
