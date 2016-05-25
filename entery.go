package golog

import (
	"fmt"
	"io"
	"sync"
)

type Entery struct {
	logger *Logger
	ctx    Context
	mu     sync.Mutex
}

func NewEntery(l *Logger, ctx Context) *Entery {
	return &Entery{
		logger: l,
		ctx:    ctx,
	}
}

func (e *Entery) AppendContext(ctx Context) *Entery {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.ctx = mergeCtx(e.ctx, ctx)
	return e
}

func (e *Entery) GetContext() Context {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.ctx
}

func (e *Entery) SetLevel(lvl Level) {
	e.logger.SetLevel(lvl)
}

func (e *Entery) SetOutput(o io.Writer) {
	e.logger.SetOutput(o)
}

func (e *Entery) SetFormatter(f Formatter) {
	e.logger.SetFormatter(f)
}

func (e *Entery) Error(m string) {
	e.log(ErrorLevel, e.GetContext(), m)
}

func (e *Entery) ErrorCtx(ctx Context, m string) {
	c := mergeCtx(e.GetContext(), ctx)
	e.log(ErrorLevel, c, m)
}

func (e *Entery) Errorf(f string, args ...interface{}) {
	e.log(ErrorLevel, e.GetContext(), fmt.Sprintf(f, args))
}

func (e *Entery) ErrorfCtx(ctx Context, f string, args ...interface{}) {
	c := mergeCtx(e.GetContext(), ctx)
	e.log(ErrorLevel, c, fmt.Sprintf(f, args))
}

func (e *Entery) Debug(m string) {
	e.log(DebugLevel, e.GetContext(), m)
}

func (e *Entery) DebugCtx(ctx Context, m string) {
	c := mergeCtx(e.GetContext(), ctx)
	e.log(DebugLevel, c, m)
}

func (e *Entery) Debugf(f string, args ...interface{}) {
	e.log(DebugLevel, e.GetContext(), fmt.Sprintf(f, args))
}

func (e *Entery) DebugfCtx(ctx Context, f string, args ...interface{}) {
	c := mergeCtx(e.GetContext(), ctx)
	e.log(DebugLevel, c, fmt.Sprintf(f, args))

}
func (e *Entery) Info(m string) {
	e.log(InfoLevel, e.GetContext(), m)
}

func (e *Entery) InfoCtx(ctx Context, m string) {
	c := mergeCtx(e.GetContext(), ctx)
	e.log(InfoLevel, c, m)
}

func (e *Entery) Infof(f string, args ...interface{}) {
	e.log(InfoLevel, e.GetContext(), fmt.Sprintf(f, args))
}

func (e *Entery) InfofCtx(ctx Context, f string, args ...interface{}) {
	c := mergeCtx(e.GetContext(), ctx)
	e.log(InfoLevel, c, fmt.Sprintf(f, args))
}

func (e *Entery) Warn(m string) {
	e.log(WarnLevel, e.GetContext(), m)
}

func (e *Entery) WarnCtx(ctx Context, m string) {
	c := mergeCtx(e.GetContext(), ctx)
	e.log(WarnLevel, c, m)
}

func (e *Entery) Warnf(f string, args ...interface{}) {
	e.log(WarnLevel, e.GetContext(), fmt.Sprintf(f, args))
}

func (e *Entery) WarnfCtx(ctx Context, f string, args ...interface{}) {
	c := mergeCtx(e.GetContext(), ctx)
	e.log(WarnLevel, c, fmt.Sprintf(f, args))
}

func (e *Entery) Fatal(m string) {
	e.log(FatalLevel, e.GetContext(), m)
	defer panic(fmt.Sprintf("fatal error: %s\n", m))
}

func (e *Entery) FatalCtx(ctx Context, m string) {
	c := mergeCtx(e.GetContext(), ctx)
	defer panic(fmt.Sprintf("fatal error: %s\n", m))
	e.log(FatalLevel, c, m)
}

func (e *Entery) Fatalf(f string, args ...interface{}) {
	m := fmt.Sprintf(f, args)
	defer panic(fmt.Sprintf("fatal error: %s\n", m))
	e.log(FatalLevel, e.GetContext(), m)
}

func (e *Entery) FatalfCtx(ctx Context, f string, args ...interface{}) {
	c := mergeCtx(e.GetContext(), ctx)
	m := fmt.Sprintf(f, args)
	defer panic(fmt.Sprintf("fatal error: %s\n", m))
	e.log(FatalLevel, c, m)
}

func (e *Entery) Panic(m string) {
	defer panic(fmt.Sprintf("panic error: %s\n", m))
	e.log(PanicLevel, e.GetContext(), m)
}

func (e *Entery) PanicCtx(ctx Context, m string) {
	c := mergeCtx(e.GetContext(), ctx)
	defer panic(fmt.Sprintf("panic error: %s\n", m))
	e.log(PanicLevel, c, m)
}

func (e *Entery) Panicf(f string, args ...interface{}) {
	m := fmt.Sprintf(f, args)
	defer panic(fmt.Sprintf("panic error: %s\n", m))
	e.log(PanicLevel, e.GetContext(), m)
}

func (e *Entery) PanicfCtx(ctx Context, f string, args ...interface{}) {
	c := mergeCtx(e.GetContext(), ctx)
	m := fmt.Sprintf(f, args)
	defer panic(fmt.Sprintf("panic error: %s\n", m))
	e.log(PanicLevel, c, m)
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
