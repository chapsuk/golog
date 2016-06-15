package golog

import (
	"fmt"
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
	return &Logger{
		Level:     l.GetLevel(),
		Formatter: l.Formatter,
		Pool:      l.Pool,
		Writer:    l.Writer,
		Context:   l.buildContext(ctx),
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
	l.Debug(args...)
}

// Printf info log message, fmt.Sprintf style
func (l *Logger) Printf(f string, args ...interface{}) {
	l.Debugf(f, args...)
}

// Println not supported, only for standart logger interface compatibility
func (l *Logger) Println(args ...interface{}) {
	l.Debugln(args...)
}

// Debug write debug log message
func (l *Logger) Debug(args ...interface{}) {
	l.log(DebugLevel, l.buildContext(nil), fmt.Sprint(args...))
}

// DebugCtx write debug log message with context
func (l *Logger) DebugCtx(ctx Context, args ...interface{}) {
	l.log(DebugLevel, l.buildContext(ctx), fmt.Sprint(args...))
}

// Debugf write debug log message, fmt.Sprtinf style
func (l *Logger) Debugf(f string, args ...interface{}) {
	l.log(DebugLevel, l.buildContext(nil), fmt.Sprintf(f, args...))
}

// DebugfCtx write debug log message with context, fmt.Sprtinf style
func (l *Logger) DebugfCtx(ctx Context, f string, args ...interface{}) {
	l.log(DebugLevel, l.buildContext(ctx), fmt.Sprintf(f, args...))
}

// Debugln print debug log msg
func (l *Logger) Debugln(args ...interface{}) {
	l.log(DebugLevel, l.buildContext(nil), fmt.Sprint(args...))
}

// Info write info log message
func (l *Logger) Info(args ...interface{}) {
	l.log(InfoLevel, l.buildContext(nil), fmt.Sprint(args...))
}

// InfoCtx write info log message with context
func (l *Logger) InfoCtx(ctx Context, args ...interface{}) {
	l.log(InfoLevel, l.buildContext(ctx), fmt.Sprint(args...))
}

// Infof write info log message, fmt.Sprtinf style
func (l *Logger) Infof(f string, args ...interface{}) {
	l.log(InfoLevel, l.buildContext(nil), fmt.Sprintf(f, args...))
}

// InfofCtx write info log message with context, fmt.Sprtinf style
func (l *Logger) InfofCtx(ctx Context, f string, args ...interface{}) {
	l.log(InfoLevel, l.buildContext(ctx), fmt.Sprintf(f, args...))
}

// Infoln write info log message
func (l *Logger) Infoln(args ...interface{}) {
	l.log(InfoLevel, l.buildContext(nil), fmt.Sprint(args...))
}

// Warn write warning log message
func (l *Logger) Warn(args ...interface{}) {
	l.log(WarnLevel, l.buildContext(nil), fmt.Sprint(args...))
}

// WarnCtx write warning log message with context
func (l *Logger) WarnCtx(ctx Context, args ...interface{}) {
	l.log(WarnLevel, l.buildContext(ctx), fmt.Sprint(args...))
}

// Warnf write warning log message, fmt.Sprintf style
func (l *Logger) Warnf(f string, args ...interface{}) {
	l.log(WarnLevel, l.buildContext(nil), fmt.Sprintf(f, args...))
}

// WarnfCtx write warning log message with context, fmt.Sprintf style
func (l *Logger) WarnfCtx(ctx Context, f string, args ...interface{}) {
	l.log(WarnLevel, l.buildContext(ctx), fmt.Sprintf(f, args...))
}

// Warnln write warning log message
func (l *Logger) Warnln(args ...interface{}) {
	l.log(WarnLevel, l.buildContext(nil), fmt.Sprint(args...))
}

// Error write error log message
func (l *Logger) Error(args ...interface{}) {
	l.log(ErrorLevel, l.buildContext(nil), fmt.Sprint(args...))
}

// Errorln write error log message
func (l *Logger) Errorln(args ...interface{}) {
	l.log(ErrorLevel, l.buildContext(nil), fmt.Sprint(args...))
}

// ErrorCtx write error log message with context
func (l *Logger) ErrorCtx(ctx Context, args ...interface{}) {
	l.log(ErrorLevel, l.buildContext(ctx), fmt.Sprint(args...))
}

// Errorf write error log message fmt.Printf style
func (l *Logger) Errorf(f string, args ...interface{}) {
	l.log(ErrorLevel, l.buildContext(nil), fmt.Sprintf(f, args...))
}

// ErrorfCtx write error log message with context, fmt.Printf style
func (l *Logger) ErrorfCtx(ctx Context, f string, args ...interface{}) {
	l.log(ErrorLevel, l.buildContext(ctx), fmt.Sprintf(f, args...))
}

// Fatal write fatal log message
func (l *Logger) Fatal(args ...interface{}) {
	l.log(FatalLevel, l.buildContext(nil), fmt.Sprint(args...))
}

// FatalCtx write fatal log message with context
func (l *Logger) FatalCtx(ctx Context, args ...interface{}) {
	l.log(FatalLevel, l.buildContext(ctx), fmt.Sprint(args...))
}

// Fatalf write fatal log message
func (l *Logger) Fatalf(f string, args ...interface{}) {
	l.log(FatalLevel, l.buildContext(nil), fmt.Sprintf(f, args...))
}

// FatalfCtx write fatal log message
func (l *Logger) FatalfCtx(ctx Context, f string, args ...interface{}) {
	l.log(FatalLevel, l.buildContext(ctx), fmt.Sprintf(f, args...))
}

// Fatalln not supported, only for standart logger interface compatibility
func (l *Logger) Fatalln(args ...interface{}) {
	l.log(FatalLevel, l.buildContext(nil), fmt.Sprint(args...))
}

// Panic write panice log message and throw panic
func (l *Logger) Panic(args ...interface{}) {
	m := fmt.Sprint(args...)
	l.log(PanicLevel, l.buildContext(nil), m)
	panic(m)
}

// PanicCtx write panice log message with context and throw panic
func (l *Logger) PanicCtx(ctx Context, args ...interface{}) {
	m := fmt.Sprint(args...)
	l.log(PanicLevel, l.buildContext(ctx), m)
	panic(m)
}

// Panicf write panice log message and throw panic, fmt.Sprintf style
func (l *Logger) Panicf(f string, args ...interface{}) {
	m := fmt.Sprintf(f, args...)
	l.log(PanicLevel, l.buildContext(nil), m)
	panic(m)
}

// PanicfCtx write panice log message with context and throw panic, fmt.Sprintf style
func (l *Logger) PanicfCtx(ctx Context, f string, args ...interface{}) {
	m := fmt.Sprintf(f, args...)
	l.log(PanicLevel, l.buildContext(ctx), m)
	panic(m)
}

// Panicln not supported, only for standart logger interface compatibility
// throw panic anyway
func (l *Logger) Panicln(args ...interface{}) {
	m := fmt.Sprint(args...)
	l.log(PanicLevel, l.buildContext(nil), m)
	panic(m)
}

func (l *Logger) log(lvl Level, ctx Context, msg string) {
	if l.isAvailableLvl(lvl) {
		b := l.Pool.Get()
		defer l.Pool.Put(b)
		m := l.Formatter.Format(b, lvl, ctx, msg)
		_, err := l.Writer.Write(m.Bytes())
		if err != nil {
			fmt.Printf("can`t write to log, message: " + m.String())
		}
	}
}

func (l *Logger) isAvailableLvl(lvl Level) bool {
	return l.GetLevel() >= lvl
}

func (l *Logger) buildContext(ctx Context) Context {
	n := Context{}
	for k, v := range l.GetContext() {
		n[k] = v
	}
	if ctx != nil {
		for k, v := range ctx {
			n[k] = v
		}
	}
	return n
}
