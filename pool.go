package golog

import (
	"bytes"
	"sync"
)

type BufferPool struct {
	Pool sync.Pool
}

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

func (p *BufferPool) Get() *bytes.Buffer {
	return p.Pool.Get().(*bytes.Buffer)
}

func (p *BufferPool) Put(b *bytes.Buffer) {
	b.Reset()
	p.Pool.Put(b)
}
