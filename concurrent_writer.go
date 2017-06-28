package golog

import (
	"io"
	"sync"
)

// ConcurrentWriter write log message to output with output lock
type ConcurrentWriter struct {
	writer io.Writer
	mu     sync.Mutex
}

// NewConcurrentWriter yield new concurrent writer instance
func NewConcurrentWriter(w io.Writer) *ConcurrentWriter {
	return &ConcurrentWriter{writer: w}
}

// Write bytes to output with lock
func (w *ConcurrentWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.writer.Write(p)
}
