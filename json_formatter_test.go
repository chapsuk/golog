package golog_test

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"
	"testing"
	"time"

	"github.com/chapsuk/golog"

	test "github.com/smartystreets/goconvey/convey"
)

type JSONLogMessage struct {
	Time  string `json:"_t"`
	Level string `json:"_l"`
	Msg   string `json:"_m"`
	Trace string `json:"_trace"`
}

type JSONLogMessageWithContext struct {
	Time  string `json:"_t"`
	Level string `json:"_l"`
	Msg   string `json:"_m"`
	// Context
	Name   string      `json:"name"`
	Age    int         `json:"age"`
	Weight float64     `json:"weight"`
	Money  float64     `json:"money"`
	Childs bool        `json:"childs"`
	Brain  bool        `json:"brain"`
	Car    interface{} `json:"car"`
	Err    string      `json:"err"`
	Uint8  uint8       `json:"unit8"`
	Uptr   uintptr     `json:"uptr"`
	F32    float32     `json:"f32"`
}

var dateFormat = time.RFC3339

func TestJSONFormatter(t *testing.T) {
	test.Convey("Check JSONFormatter", t, func() {
		log := golog.New()
		b := log.Pool.Get()
		defer log.Pool.Put(b)

		test.Convey("unmarshal log message", func() {
			f := golog.JSONFormatter{DateFormat: dateFormat}
			buf := f.Format(b, golog.DebugLevel, golog.Context{}, foo, []byte{})
			res := JSONLogMessage{}
			err := json.Unmarshal(buf.Bytes(), &res)
			if err != nil {
				t.Error(err)
			}

			logTime, err := time.Parse(dateFormat, res.Time)
			if err != nil {
				t.Error(err)
			}

			test.So(res.Level, test.ShouldEqual, golog.LevelToString(golog.DebugLevel))
			test.So(logTime, test.ShouldHappenOnOrBefore, time.Now())
			test.So(res.Msg, test.ShouldEqual, foo)
		})

		test.Convey("check context variables", func() {
			f := golog.JSONFormatter{}
			var uiptr uintptr
			var f32 float32
			ctx := golog.Context{
				"name":   "mak",
				"age":    21,
				"weight": 73.5,
				"money":  -11.3,
				"childs": false,
				"brain":  true,
				"car":    nil,
				"err":    fmt.Errorf("test error"),
				"unit8":  1,
				"uptr":   uiptr,
				"f32":    f32,
			}
			buf := f.Format(b, golog.InfoLevel, ctx, "prfile info", []byte{})
			res := JSONLogMessageWithContext{}
			err := json.Unmarshal(buf.Bytes(), &res)
			if err != nil {
				t.Error(err)
			}

			test.So(res.Name, test.ShouldEqual, "mak")
			test.So(res.Age, test.ShouldEqual, 21)
			test.So(res.Weight, test.ShouldEqual, 73.5)
			test.So(res.Money, test.ShouldEqual, -11.3)
			test.So(res.Childs, test.ShouldBeFalse)
			test.So(res.Car, test.ShouldBeNil)
			test.So(res.Err, test.ShouldEqual, "test error")
			test.So(res.Uint8, test.ShouldEqual, 1)
			test.So(res.Uptr, test.ShouldEqual, uiptr)
			test.So(res.F32, test.ShouldEqual, f32)
		})

		test.Convey("check correct trace", func() {
			f := golog.JSONFormatter{DateFormat: dateFormat}
			trace := debug.Stack()
			buf := f.Format(b, golog.DebugLevel, golog.Context{}, foo, trace)
			golog.SetOutput(os.Stdout)
			res := JSONLogMessage{}
			err := json.Unmarshal(buf.Bytes(), &res)
			if err != nil {
				t.Error(err)
			}

			test.So(res.Trace, test.ShouldEqual, string(trace))
		})
	})
}
