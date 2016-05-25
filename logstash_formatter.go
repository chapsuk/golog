package golog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type LogstashFormatter struct{}

func (f *LogstashFormatter) Format(b *bytes.Buffer, lvl Level, ctx Context, msg string) *bytes.Buffer {
	b.WriteString(`{"@timestamp":"`)
	b.WriteString(time.Now().Format(time.RFC3339))
	b.WriteString(`","@version":1,"level":"`)
	b.WriteString(LevelToString(lvl))
	b.WriteString(`","`)

	for k, v := range ctx {
		b.WriteString(k)
		b.WriteString(`":`)
		f.appendValue(b, v)
		b.WriteString(`,"`)
	}

	b.WriteString(`message":"`)
	b.WriteString(msg)
	b.WriteString(`"}`)
	b.WriteByte('\n')
	return b
}

func (f *LogstashFormatter) appendValue(buf *bytes.Buffer, val interface{}) {
	if val == nil {
		buf.WriteString("null")
		return
	}

	if err, ok := val.(error); ok {
		buf.WriteString(`"`)
		buf.WriteString(err.Error())
		buf.WriteString(`"`)
		return
	}

	value := reflect.ValueOf(val)
	kind := value.Kind()
	if kind == reflect.Ptr {
		if value.IsNil() {
			buf.WriteString("null")
			return
		}
		value = value.Elem()
		kind = value.Kind()
	}
	switch kind {
	case reflect.Bool:
		if value.Bool() {
			buf.WriteString("true")
		} else {
			buf.WriteString("false")
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		buf.WriteString(strconv.FormatInt(value.Int(), 10))

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		buf.WriteString(strconv.FormatUint(value.Uint(), 10))

	case reflect.Float32:
		buf.WriteString(strconv.FormatFloat(value.Float(), 'g', -1, 32))

	case reflect.Float64:
		buf.WriteString(strconv.FormatFloat(value.Float(), 'g', -1, 64))

	default:
		var err error
		var b []byte
		if stringer, ok := val.(fmt.Stringer); ok {
			b, err = json.Marshal(stringer.String())
		} else {
			b, err = json.Marshal(val)
		}

		if err != nil {
			if s, ok := val.(string); ok {
				b, err = json.Marshal(s)
			} else if s, ok := val.(fmt.Stringer); ok {
				b, err = json.Marshal(s.String())
			} else {
				b, err = json.Marshal(fmt.Sprintf("%#v", val))
			}

			if err != nil {
				// should never get here, but JSONFormatter should never panic
				msg := "Could not Sprintf value"
				buf.WriteString(`"` + msg + `"`)
				return
			}
		}
		buf.Write(b)
	}
}
