package golog

import (
	"io"
	"sync"
)

// CuncurrentWriter write log message to output with output lock
type CuncurrentWriter struct {
	writer io.Writer
	mu     sync.Mutex
}

// NewCuncurrentWriter yield new cuncurrent writer instace
func NewCuncurrentWriter(w io.Writer) *CuncurrentWriter {
	return &CuncurrentWriter{writer: w}
}

// Write bytes to output with lock
func (w *CuncurrentWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.writer.Write(p)
}
