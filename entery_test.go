package golog_test

import (
	"fmt"
	"github.com/chapsuk/golog"
	test "github.com/smartystreets/goconvey/convey"
	"testing"
)

type ErrorWriter struct{}

func (w *ErrorWriter) Write(p []byte) (int, error) {
	return 0, fmt.Errorf("%s", "error")
}

func TestEntery(t *testing.T) {
	test.Convey("Check entery methods", t, func() {
		out := &FakeWriter{}
		log := golog.NewLogger(out, &FakeFormatter{}, golog.Context{})
		e := golog.NewEntery(log, golog.Context{})

		test.Convey("check append context", func() {
			ctx := golog.Context{
				"foo": foo,
				"bar": bar,
			}
			ne := e.AppendContext(ctx)

			test.So(e.GetContext(), test.ShouldResemble, ctx)
			test.So(ne.GetContext(), test.ShouldResemble, ctx)
		})

		test.Convey("check setOutput", func() {
			newOut := &FakeWriter{}
			e.SetOutput(newOut)

			e.Info("info message")

			test.So(newOut.log, test.ShouldContain, "info message")
			test.So(out.log, test.ShouldNotContain, "info message")
		})

		test.Convey("check setFormatter", func() {
			e.SetFormatter(&golog.JSONFormatter{DateFormat: " "}) // hack for empty time field

			e.Info("info message")
			test.So(out.log, test.ShouldContain, `{"_t":" ", "_l":"INFO", "_m":"info message"}`)

			log.Error("error message")
			test.So(out.log, test.ShouldContain, `{"_t":" ", "_l":"ERROR", "_m":"error message"}`)
		})

		test.Convey("check set level", func() {
			e.SetLevel(golog.ErrorLevel)

			test.So(log.Level, test.ShouldEqual, golog.ErrorLevel)

			e.Info("info message")
			test.So(out.log, test.ShouldBeEmpty)

			log.Debug("debug message")
			test.So(out.log, test.ShouldBeEmpty)
		})

		test.Convey("when writer return error, print message to stdout: ", func() {
			e.SetOutput(&ErrorWriter{})

			e.Info("info message")
			test.So(true, test.ShouldBeTrue)
		})
	})
}
