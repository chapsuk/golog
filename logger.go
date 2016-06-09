package golog

import (
	"io"
	"os"
	"sync"
)

// Logger struct
type Logger struct {
	Level     Level
	Writer    io.Writer
	Formatter Formatter
	mu        sync.Mutex
	Pool      *BufferPool
	Context   Context
}

// New return default logger
func New() *Logger {
	return &Logger{
		Level:     DebugLevel,
		Formatter: &TextFormatter{},
		Pool:      NewBufferPool(),
		Writer:    NewCuncurrentWriter(os.Stdout),
		Context:   Context{},
	}
}

// NewLogger yield new logger instance
func NewLogger(out io.Writer, f Formatter, ctx Context) *Logger {
	return &Logger{
		Level:     DebugLevel,
		Formatter: f,
		Pool:      NewBufferPool(),
		Writer:    out,
		Context:   ctx,
	}
}

// WithContext yield new logger instance with context
// new context append to current context
func (l *Logger) WithContext(ctx Context) ILogger {
	l.mu.Lock()
	defer l.mu.Unlock()
	return &Logger{
		Level:     l.Level,
		Formatter: l.Formatter,
		Pool:      l.Pool,
		Writer:    l.Writer,
		Context:   mergeCtx(l.Context, ctx),
	}
}

// GetContext return current global log context
func (l *Logger) GetContext() Context {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.Context
}

// GetLevel return current max log level
func (l *Logger) GetLevel() Level {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.Level
}

// SetLevel set max log level
func (l *Logger) SetLevel(lvl Level) {
	l.mu.Lock()
	l.Level = lvl
	l.mu.Unlock()
}

// SetOutput for logger messages
func (l *Logger) SetOutput(o io.Writer) {
	l.mu.Lock()
	l.Writer = o
	l.mu.Unlock()
}

// SetFormatter set log message formatter
func (l *Logger) SetFormatter(f Formatter) {
	l.mu.Lock()
	l.Formatter = f
	l.mu.Unlock()
}

// Print info log message
func (l *Logger) Print(args ...interface{}) {
	NewEntery(l, l.GetContext()).Print(args...)
}

// Printf info log message, fmt.Sprintf style
func (l *Logger) Printf(f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).Printf(f, args...)
}

// Println not supported, only for standart logger interface compatibility
func (l *Logger) Println(args ...interface{}) {
	NewEntery(l, l.GetContext()).Print(args...)
}

// Error write error log message
func (l *Logger) Error(args ...interface{}) {
	NewEntery(l, l.GetContext()).Error(args...)
}

// ErrorCtx write error log message with context
func (l *Logger) ErrorCtx(ctx Context, m string) {
	NewEntery(l, l.GetContext()).ErrorCtx(ctx, m)
}

// Errorf write error log message fmt.Printf style
func (l *Logger) Errorf(f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).Errorf(f, args...)
}

// ErrorfCtx write error log message with context, fmt.Printf style
func (l *Logger) ErrorfCtx(ctx Context, f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).ErrorfCtx(ctx, f, args...)
}

// Debug write debug log message
func (l *Logger) Debug(args ...interface{}) {
	NewEntery(l, l.GetContext()).Debug(args...)
}

// DebugCtx write debug log message with context
func (l *Logger) DebugCtx(ctx Context, m string) {
	NewEntery(l, l.GetContext()).DebugCtx(ctx, m)
}

// Debugf write debug log message, fmt.Sprtinf style
func (l *Logger) Debugf(f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).Debugf(f, args...)
}

// DebugfCtx write debug log message with context, fmt.Sprtinf style
func (l *Logger) DebugfCtx(ctx Context, f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).DebugfCtx(ctx, f, args...)
}

// Info write info log message
func (l *Logger) Info(args ...interface{}) {
	NewEntery(l, l.GetContext()).Info(args...)
}

// InfoCtx write info log message with context
func (l *Logger) InfoCtx(ctx Context, m string) {
	NewEntery(l, l.GetContext()).InfoCtx(ctx, m)
}

// Infof write info log message, fmt.Sprtinf style
func (l *Logger) Infof(f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).Infof(f, args...)
}

// InfofCtx write info log message with context, fmt.Sprtinf style
func (l *Logger) InfofCtx(ctx Context, f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).InfofCtx(ctx, f, args...)
}

// Warn write warning log message
func (l *Logger) Warn(args ...interface{}) {
	NewEntery(l, l.GetContext()).Warn(args...)
}

// WarnCtx write warning log message with context
func (l *Logger) WarnCtx(ctx Context, m string) {
	NewEntery(l, l.GetContext()).WarnCtx(ctx, m)
}

// Warnf write warning log message, fmt.Sprintf style
func (l *Logger) Warnf(f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).Warnf(f, args...)
}

// WarnfCtx write warning log message with context, fmt.Sprintf style
func (l *Logger) WarnfCtx(ctx Context, f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).WarnfCtx(ctx, f, args...)
}

// Fatal write fatal log message and call os.Exit
func (l *Logger) Fatal(args ...interface{}) {
	NewEntery(l, l.GetContext()).Fatal(args...)
}

// FatalCtx write fatal log message with context and call os.Exit
func (l *Logger) FatalCtx(ctx Context, m string) {
	NewEntery(l, l.GetContext()).FatalCtx(ctx, m)
}

// Fatalf write fatal log message and call os.Exit, fmt.Sprtinf style
func (l *Logger) Fatalf(f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).Fatalf(f, args...)
}

// FatalfCtx write fatal log message with context and call os.Exit, fmt.Sprtinf style
func (l *Logger) FatalfCtx(ctx Context, f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).FatalfCtx(ctx, f, args...)
}

// Fatalln not supported, only for standart logger interface compatibility
// call os.Exit anyway
func (l *Logger) Fatalln(args ...interface{}) {
	NewEntery(l, l.GetContext()).Fatalln(args...)
}

// Panic write panice log message and throw panic
func (l *Logger) Panic(args ...interface{}) {
	NewEntery(l, l.GetContext()).Panic(args...)
}

// PanicCtx write panice log message with context and throw panic
func (l *Logger) PanicCtx(ctx Context, m string) {
	NewEntery(l, l.GetContext()).PanicCtx(ctx, m)
}

// Panicf write panice log message and throw panic, fmt.Sprintf style
func (l *Logger) Panicf(f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).Panicf(f, args...)
}

// PanicfCtx write panice log message with context and throw panic, fmt.Sprintf style
func (l *Logger) PanicfCtx(ctx Context, f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).PanicfCtx(ctx, f, args...)
}

// Panicln not supported, only for standart logger interface compatibility
// throw panic anyway
func (l *Logger) Panicln(args ...interface{}) {
	NewEntery(l, l.GetContext()).Panicln(args...)
}
