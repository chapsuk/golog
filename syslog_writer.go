package golog

import (
	"io"
	"log/syslog"
	"time"
)

// SyslogWriter write logs to syslog
type SyslogWriter struct {
	Writer io.Writer
}

// NewSyslogWriter return new SyslogWriter instance with cuncurrent wirter to syslog
func NewSyslogWriter(addr, tag string, timeout int) *SyslogWriter {
	s := &SyslogWriter{}
	w, err := syslog.Dial("tcp", addr, syslog.LOG_USER, tag)

	if err != nil {
		std.Errorf("error connecting to syslog: %s", err.Error())

		go func() {
			for {
				time.Sleep(time.Second * time.Duration(timeout))

				w, err := syslog.Dial("tcp", addr, syslog.LOG_USER, tag)
				if err != nil {
					std.Errorf("error connecting to syslog: %s", err.Error())
					continue
				}

				s.Writer = NewCuncurrentWriter(w)
				break
			}
		}()

		return s
	}

	s.Writer = NewCuncurrentWriter(w)
	return s
}

func (w *SyslogWriter) Write(p []byte) (int, error) {
	return w.Writer.Write(p)
}
