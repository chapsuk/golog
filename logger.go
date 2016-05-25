package golog

import (
	"io"
	"os"
	"sync"
)

type Logger struct {
	Level     Level
	Writer    io.Writer
	Formatter Formatter
	mu        sync.Mutex
	Pool      *BufferPool
	Context   Context
}

func New() *Logger {
	return &Logger{
		Level:     DebugLevel,
		Formatter: &TextFormatter{},
		Pool:      NewBufferPool(),
		Writer:    NewCuncurrentWriter(os.Stdout),
		Context:   Context{},
	}
}

func NewLogger(out io.Writer, f Formatter, ctx Context) *Logger {
	return &Logger{
		Level:     DebugLevel,
		Formatter: f,
		Pool:      NewBufferPool(),
		Writer:    out,
		Context:   ctx,
	}
}

func (l *Logger) AppendContext(ctx Context) ILogger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Context = mergeCtx(l.Context, ctx)
	return l
}

func (l *Logger) GetContext() Context {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.Context
}

func (l *Logger) GetLevel() Level {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.Level
}

func (l *Logger) SetLevel(lvl Level) {
	l.mu.Lock()
	l.Level = lvl
	l.mu.Unlock()
}

func (l *Logger) SetOutput(o io.Writer) {
	l.mu.Lock()
	l.Writer = o
	l.mu.Unlock()
}

func (l *Logger) SetFormatter(f Formatter) {
	l.mu.Lock()
	l.Formatter = f
	l.mu.Unlock()
}

func (l *Logger) Error(m string) {
	NewEntery(l, l.GetContext()).Error(m)
}

func (l *Logger) ErrorCtx(ctx Context, m string) {
	NewEntery(l, l.GetContext()).ErrorCtx(ctx, m)
}

func (l *Logger) Errorf(f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).Errorf(f, args)
}

func (l *Logger) ErrorfCtx(ctx Context, f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).ErrorfCtx(ctx, f, args)
}

func (l *Logger) Debug(m string) {
	NewEntery(l, l.GetContext()).Debug(m)
}

func (l *Logger) DebugCtx(ctx Context, m string) {
	NewEntery(l, l.GetContext()).DebugCtx(ctx, m)
}

func (l *Logger) Debugf(f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).Debugf(f, args)
}

func (l *Logger) DebugfCtx(ctx Context, f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).DebugfCtx(ctx, f, args)
}

func (l *Logger) Info(m string) {
	NewEntery(l, l.GetContext()).Info(m)
}

func (l *Logger) InfoCtx(ctx Context, m string) {
	NewEntery(l, l.GetContext()).InfoCtx(ctx, m)
}

func (l *Logger) Infof(f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).Infof(f, args)
}

func (l *Logger) InfofCtx(ctx Context, f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).InfofCtx(ctx, f, args)
}

func (l *Logger) Warn(m string) {
	NewEntery(l, l.GetContext()).Warn(m)
}

func (l *Logger) WarnCtx(ctx Context, m string) {
	NewEntery(l, l.GetContext()).WarnCtx(ctx, m)
}

func (l *Logger) Warnf(f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).Warnf(f, args)
}

func (l *Logger) WarnfCtx(ctx Context, f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).WarnfCtx(ctx, f, args)
}

func (l *Logger) Fatal(m string) {
	NewEntery(l, l.GetContext()).Fatal(m)
}

func (l *Logger) FatalCtx(ctx Context, m string) {
	NewEntery(l, l.GetContext()).FatalCtx(ctx, m)
}

func (l *Logger) Fatalf(f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).Fatalf(f, args)
}

func (l *Logger) FatalfCtx(ctx Context, f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).FatalfCtx(ctx, f, args)
}

func (l *Logger) Panic(m string) {
	NewEntery(l, l.GetContext()).Panic(m)
}

func (l *Logger) PanicCtx(ctx Context, m string) {
	NewEntery(l, l.GetContext()).PanicCtx(ctx, m)
}

func (l *Logger) Panicf(f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).Panicf(f, args)
}

func (l *Logger) PanicfCtx(ctx Context, f string, args ...interface{}) {
	NewEntery(l, l.GetContext()).PanicfCtx(ctx, f, args)
}
