package golog

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/chapsuk/golog/syslog"
)

// SyslogWriter write logs to syslog
type SyslogWriter struct {
	Writer  io.WriteCloser
	timeout time.Duration
	mu      sync.Mutex
	done    chan struct{}
}

// NewSyslogWriter return new SyslogWriter instance with concurrent writer to syslog
func NewSyslogWriter(network, addr, tag string, timeout time.Duration) *SyslogWriter {
	s := &SyslogWriter{
		timeout: timeout,
		done:    make(chan struct{}),
	}
	w, err := syslog.Dial(network, addr, syslog.LOG_USER, tag)
	if err != nil {
		s.Writer = os.Stdout
		std.Errorf("error connecting to syslog: %s", err.Error())

		go func() {
			t := time.NewTicker(s.timeout)
			defer t.Stop()

			for {
				select {
				case <-t.C:
					w, err := syslog.Dial(network, addr, syslog.LOG_USER, tag)
					if err != nil {
						std.Errorf("error connecting to syslog: %s", err.Error())
						continue
					}

					s.mu.Lock()
					s.Writer = w
					s.mu.Unlock()
					return
				case <-s.done:
					return
				}
			}
		}()
		return s
	}

	s.Writer = w
	return s
}

// Write writes data to syslog
func (w *SyslogWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if sl, ok := w.Writer.(*syslog.Writer); ok {
		err := sl.SetWriteDeadLine(time.Now().Add(w.timeout))
		if err != nil {
			return 0, err
		}
	}
	return w.Writer.Write(p)
}

// Close close connection
func (w *SyslogWriter) Close() {
	w.done <- struct{}{}
	w.Writer.Close()
}
