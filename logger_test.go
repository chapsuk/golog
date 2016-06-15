package golog_test

import (
	"fmt"
	"github.com/chapsuk/golog"
	test "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

type ErrorWriter struct{}

func (w *ErrorWriter) Write(p []byte) (int, error) {
	return 0, fmt.Errorf("%s", "error")
}

func TestLogger(t *testing.T) {
	test.Convey("Save context after WithContext called", t, func() {
		log1 := golog.NewLogger(os.Stderr, &golog.JSONFormatter{}, golog.Context{
			"field": "main",
			"name":  "log1",
		})
		log2 := log1.WithContext(golog.Context{
			"name": "log2",
			"src":  "foo",
		})

		test.So(golog.Context{
			"field": "main",
			"name":  "log1",
		}, test.ShouldResemble, log1.GetContext())
		test.So(golog.Context{
			"field": "main",
			"name":  "log2",
			"src":   "foo",
		}, test.ShouldResemble, log2.GetContext())
	})

	test.Convey("ErrorWriter. Writer return error: \n", t, func() {
		log := golog.New()
		log.SetOutput(&ErrorWriter{})

		log.Print(foo)
	})
}
