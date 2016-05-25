package golog_test

import (
	"github.com/chapsuk/golog"
	test "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLevels(t *testing.T) {
	test.Convey("Check level to string method", t, func() {
		test.So(golog.LevelToString(golog.DebugLevel), test.ShouldEqual, "DEBUG")
		test.So(golog.LevelToString(golog.InfoLevel), test.ShouldEqual, "INFO")
		test.So(golog.LevelToString(golog.WarnLevel), test.ShouldEqual, "WARN")
		test.So(golog.LevelToString(golog.ErrorLevel), test.ShouldEqual, "ERROR")
		test.So(golog.LevelToString(golog.FatalLevel), test.ShouldEqual, "FATAL")
		test.So(golog.LevelToString(golog.PanicLevel), test.ShouldEqual, "PANIC")
		test.So(golog.LevelToString(73), test.ShouldEqual, "UNKNOWN")
	})
}
