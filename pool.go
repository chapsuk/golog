package golog

import (
	"bytes"
	"sync"
)

// BufferPool is pool of buffers
type BufferPool struct {
	Pool sync.Pool
}

// NewBufferPool yield new buffer poll instance
func NewBufferPool() *BufferPool {
	return &BufferPool{
		Pool: sync.Pool{
			New: func() interface{} {
				b := bytes.NewBuffer(make([]byte, 128))
				b.Reset()
				return b
			},
		},
	}
}

// Get buffer from pool
func (p *BufferPool) Get() *bytes.Buffer {
	return p.Pool.Get().(*bytes.Buffer)
}

// Put clear buffer and put to pool
func (p *BufferPool) Put(b *bytes.Buffer) {
	b.Reset()
	p.Pool.Put(b)
}
