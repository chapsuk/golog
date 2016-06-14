package golog_test

import (
	"encoding/json"
	"fmt"
	"github.com/chapsuk/golog"
	test "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

type LogstashLogMessage struct {
	Time    string `json:"@timestamp"`
	Level   string `json:"level"`
	Version int    `json:"@version"`
	Msg     string `json:"message"`
}

type LogstashLogMessageWithContext struct {
	Time    string `json:"@timestamp"`
	Level   string `json:"level"`
	Version int    `json:"@version"`
	Msg     string `json:"message"`
	// Context
	Name   string      `json:"name"`
	Age    int         `json:"age"`
	Weight float64     `json:"weight"`
	Money  float32     `json:"money"`
	Childs bool        `json:"childs"`
	Brain  bool        `json:"brain"`
	Car    interface{} `json:"car"`
	Err    string      `json:"err"`
	Uint8  uint8       `json:"unit8"`
	Uptr   uintptr     `json:"uptr"`
	F32    float32     `json:"f32"`
}

func TestLogstashFormatter(t *testing.T) {

	test.Convey("Check logstash formatter", t, func() {
		log := golog.New()
		b := log.Pool.Get()
		defer log.Pool.Put(b)

		test.Convey("unmarshal log message", func() {
			f := golog.LogstashFormatter{}
			buf := f.Format(b, golog.DebugLevel, golog.Context{}, foo)
			res := LogstashLogMessage{}
			err := json.Unmarshal(buf.Bytes(), &res)
			if err != nil {
				t.Error(err)
			}

			logTime, err := time.Parse(time.RFC3339Nano, res.Time)
			if err != nil {
				t.Error(err)
			}

			test.So(res.Level, test.ShouldEqual, golog.LevelToString(golog.DebugLevel))
			test.So(logTime, test.ShouldHappenOnOrBefore, time.Now())
			test.So(res.Msg, test.ShouldEqual, foo)
		})

		test.Convey("check context variables", func() {
			f := golog.LogstashFormatter{}
			var uiptr uintptr
			var f32 float32
			ctx := golog.Context{
				"name":   "mak",
				"age":    21,
				"weight": -100.25,
				"money":  73.5,
				"childs": false,
				"brain":  true,
				"car":    nil,
				"err":    fmt.Errorf("test error"),
				"unit8":  1,
				"uptr":   uiptr,
				"f32":    f32,
			}
			buf := f.Format(b, golog.InfoLevel, ctx, "prfile info")
			res := LogstashLogMessageWithContext{}
			err := json.Unmarshal(buf.Bytes(), &res)
			if err != nil {
				t.Error(err)
			}

			test.So(res.Name, test.ShouldEqual, "mak")
			test.So(res.Age, test.ShouldEqual, 21)
			test.So(res.Weight, test.ShouldEqual, -100.25)
			test.So(res.Money, test.ShouldEqual, 73.5)
			test.So(res.Childs, test.ShouldBeFalse)
			test.So(res.Car, test.ShouldBeNil)
			test.So(res.Err, test.ShouldEqual, "test error")
			test.So(res.Uint8, test.ShouldEqual, 1)
			test.So(res.Uptr, test.ShouldEqual, uiptr)
			test.So(res.F32, test.ShouldEqual, f32)
		})
	})
}
