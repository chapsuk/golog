package golog

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// Entery log message
type Entery struct {
	logger *Logger
	ctx    Context
	mu     sync.Mutex
}

// NewEntery yield new log message entery
func NewEntery(l *Logger, ctx Context) *Entery {
	return &Entery{
		logger: l,
		ctx:    ctx,
	}
}

// WithContext return new entery with append context
func (e *Entery) WithContext(ctx Context) *Entery {
	e.mu.Lock()
	defer e.mu.Unlock()
	return &Entery{
		logger: e.logger,
		ctx:    mergeCtx(e.ctx, ctx),
	}
}

// GetContext return current log message context
func (e *Entery) GetContext() Context {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.ctx
}

// SetLevel set max log message level
func (e *Entery) SetLevel(lvl Level) {
	e.logger.SetLevel(lvl)
}

// SetOutput for log messages
func (e *Entery) SetOutput(o io.Writer) {
	e.logger.SetOutput(o)
}

// SetFormatter log message
func (e *Entery) SetFormatter(f Formatter) {
	e.logger.SetFormatter(f)
}

// Print info log message
func (e *Entery) Print(args ...interface{}) {
	e.Info(args...)
}

// Printf info log message, fmt.Sprintf style
func (e *Entery) Printf(f string, args ...interface{}) {
	e.Infof(f, args...)
}

// Println not supported, only for standart logger interface compatibility
func (e *Entery) Println(args ...interface{}) {
	e.log(WarnLevel, e.GetContext(), "sorry *ln functions not supported")
}

// Error write error log message
func (e *Entery) Error(args ...interface{}) {
	e.log(ErrorLevel, e.GetContext(), fmt.Sprint(args...))
}

// ErrorCtx write error log message with context
func (e *Entery) ErrorCtx(ctx Context, m string) {
	e.log(ErrorLevel, mergeCtx(e.GetContext(), ctx), m)
}

// Errorf write error log message, fmt.Sprintf style
func (e *Entery) Errorf(f string, args ...interface{}) {
	e.log(ErrorLevel, e.GetContext(), fmt.Sprintf(f, args...))
}

// ErrorfCtx write error log message with context, fmt.Sprintf style
func (e *Entery) ErrorfCtx(ctx Context, f string, args ...interface{}) {
	e.log(ErrorLevel, mergeCtx(e.GetContext(), ctx), fmt.Sprintf(f, args...))
}

// Debug write debug log message
func (e *Entery) Debug(args ...interface{}) {
	e.log(DebugLevel, e.GetContext(), fmt.Sprint(args...))
}

// DebugCtx write debug log message with context
func (e *Entery) DebugCtx(ctx Context, m string) {
	e.log(DebugLevel, mergeCtx(e.GetContext(), ctx), m)
}

// Debugf write debug log message, fmt.Sprintf style
func (e *Entery) Debugf(f string, args ...interface{}) {
	e.log(DebugLevel, e.GetContext(), fmt.Sprintf(f, args...))
}

// DebugfCtx write debug log message with context, fmt.Sprintf style
func (e *Entery) DebugfCtx(ctx Context, f string, args ...interface{}) {
	e.log(DebugLevel, mergeCtx(e.GetContext(), ctx), fmt.Sprintf(f, args...))
}

// Info write info log message
func (e *Entery) Info(args ...interface{}) {
	e.log(InfoLevel, e.GetContext(), fmt.Sprint(args...))
}

// InfoCtx write info log message with context
func (e *Entery) InfoCtx(ctx Context, m string) {
	e.log(InfoLevel, mergeCtx(e.GetContext(), ctx), m)
}

// Infof write info log message, fmt.Sprintf style
func (e *Entery) Infof(f string, args ...interface{}) {
	e.log(InfoLevel, e.GetContext(), fmt.Sprintf(f, args...))
}

// InfofCtx write info log message with context, fmt.Sprintf style
func (e *Entery) InfofCtx(ctx Context, f string, args ...interface{}) {
	e.log(InfoLevel, mergeCtx(e.GetContext(), ctx), fmt.Sprintf(f, args...))
}

// Warn write warning log message
func (e *Entery) Warn(args ...interface{}) {
	e.log(WarnLevel, e.GetContext(), fmt.Sprint(args...))
}

// WarnCtx write warning log message with context
func (e *Entery) WarnCtx(ctx Context, m string) {
	e.log(WarnLevel, mergeCtx(e.GetContext(), ctx), m)
}

// Warnf write warning log message, fmt.Sprintf style
func (e *Entery) Warnf(f string, args ...interface{}) {
	e.log(WarnLevel, e.GetContext(), fmt.Sprintf(f, args...))
}

// WarnfCtx write warning log message with context, fmt.Sprintf style
func (e *Entery) WarnfCtx(ctx Context, f string, args ...interface{}) {
	e.log(WarnLevel, mergeCtx(e.GetContext(), ctx), fmt.Sprintf(f, args...))
}

// Fatal write fatal log message and call os.Exit
func (e *Entery) Fatal(args ...interface{}) {
	e.log(FatalLevel, e.GetContext(), fmt.Sprint(args...))
	os.Exit(1)
}

// FatalCtx write fatal log message with context and call os.Exit
func (e *Entery) FatalCtx(ctx Context, m string) {
	e.log(FatalLevel, mergeCtx(e.GetContext(), ctx), m)
	os.Exit(1)
}

// Fatalf write fatal log message and call os.Exit, fmt.Sprintf style
func (e *Entery) Fatalf(f string, args ...interface{}) {
	e.log(FatalLevel, e.GetContext(), fmt.Sprintf(f, args...))
	os.Exit(1)
}

// FatalfCtx write fatal log message with context and call os.Exit, fmt.Sprintf style
func (e *Entery) FatalfCtx(ctx Context, f string, args ...interface{}) {
	e.log(FatalLevel, mergeCtx(e.GetContext(), ctx), fmt.Sprintf(f, args...))
	os.Exit(1)
}

// Fatalln not supported, only for standart logger interface compatibility
// call os.Exit anyway
func (e *Entery) Fatalln(args ...interface{}) {
	e.log(FatalLevel, e.GetContext(), "sorry *ln functions not supported")
	os.Exit(1)
}

// Panic write panic log message and throw panic
func (e *Entery) Panic(args ...interface{}) {
	m := fmt.Sprint(args...)
	e.log(PanicLevel, e.GetContext(), m)
	panic(m)
}

// PanicCtx write panic log message with context and throw panic
func (e *Entery) PanicCtx(ctx Context, m string) {
	e.log(PanicLevel, mergeCtx(e.GetContext(), ctx), m)
	panic(m)
}

// Panicf write panic log message and throw panic, fmt.Sprintf style
func (e *Entery) Panicf(f string, args ...interface{}) {
	m := fmt.Sprintf(f, args...)
	e.log(PanicLevel, e.GetContext(), m)
	panic(m)
}

// PanicfCtx write panic log message with context and throw panic, fmt.Sprintf style
func (e *Entery) PanicfCtx(ctx Context, f string, args ...interface{}) {
	m := fmt.Sprintf(f, args...)
	e.log(PanicLevel, mergeCtx(e.GetContext(), ctx), m)
	panic(m)
}

// Panicln not supported, only for standart logger interface compatibility
// throw panic anyway
func (e *Entery) Panicln(args ...interface{}) {
	e.log(ErrorLevel, e.GetContext(), "sorry *ln functions not supported")
	panic(fmt.Sprint(args))
}

func (e *Entery) log(lvl Level, ctx Context, msg string) {
	if e.isAvailableLvl(lvl) {
		b := e.logger.Pool.Get()
		defer e.logger.Pool.Put(b)
		m := e.logger.Formatter.Format(b, lvl, ctx, msg)
		_, err := e.logger.Writer.Write(m.Bytes())
		if err != nil {
			fmt.Printf("can`t write to log, message: " + m.String())
		}
	}
}

func (e *Entery) isAvailableLvl(lvl Level) bool {
	return e.logger.GetLevel() >= lvl
}

func mergeCtx(ctx1 Context, ctx2 Context) Context {
	for k, v := range ctx2 {
		ctx1[k] = v
	}
	return ctx1
}
