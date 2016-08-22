package golog_test

import (
	"fmt"
	"testing"

	"github.com/chapsuk/golog"
	test "github.com/smartystreets/goconvey/convey"
)

func TestTextFormatter(t *testing.T) {

	test.Convey("Check text formatter", t, func() {
		out := &FakeWriter{}
		log := golog.NewLogger(out, &golog.TextFormatter{DateFormat: " "}, golog.Context{})

		test.Convey("write with default date format", func() {
			log.SetFormatter(&golog.TextFormatter{})
			log.Info("info message")

			test.So(out.log, test.ShouldHaveLength, 1)
		})

		test.Convey("write without context", func() {
			log.Info("info message")
			test.So(out.log, test.ShouldContain, "   INFO  []  info message\n")
		})

		test.Convey("write with context (without sort) [foo: foo, err: err, int: 1]", func() {
			log.InfoCtx(golog.Context{
				"foo": foo,
				"err": fmt.Errorf("error"),
				"int": 1,
			}, "this log message")
			// only check length
			test.So(out.log, test.ShouldHaveLength, 1)
			for _, v := range out.log {
				fmt.Print(v)
			}
		})
		test.Convey("write with trace", func() {
			log.SetTraceLevel(golog.DebugLevel)
			test.So(out.log, test.ShouldBeEmpty)
			log.Debug("trace debug")
			test.So(out.log, test.ShouldNotBeEmpty)
		})
	})
}
