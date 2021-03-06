package golog

import (
	"bytes"
	"fmt"
	"time"
)

var _ Formatter = new(TextFormatter)

// TextFormatter formatter for console output
type TextFormatter struct {
	DateFormat string
}

// Format log message
func (f *TextFormatter) Format(
	b *bytes.Buffer,
	lvl Level,
	ctx Context,
	msg string,
	trace []byte,
) *bytes.Buffer {
	dateFormat := f.DateFormat
	if dateFormat == "" {
		dateFormat = time.RFC3339
	}

	b.WriteString(time.Now().Format(dateFormat))
	b.WriteString("  ")
	b.WriteString(LevelToString(lvl))
	b.WriteString("  ")

	b.WriteString("[")
	first := true
	for k, v := range ctx {
		if !first {
			b.WriteString(", ")
		} else {
			first = false
		}
		b.WriteString(k)
		b.WriteString(": ")

		switch value := v.(type) {
		case string:
			b.WriteString(value)
		case error:
			b.WriteString(value.Error())
		default:
			fmt.Fprint(b, value)
		}
	}
	b.WriteString("]  ")

	b.WriteString(msg)
	b.WriteByte('\n')
	if len(trace) > 0 {
		b.Write(trace)
	}
	return b
}
