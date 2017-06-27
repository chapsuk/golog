package golog_test

import (
	"github.com/chapsuk/golog"
	test "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSyslogWriter(t *testing.T) {

	test.Convey("failed write to syslog, connection errors", t, func() {
		log := golog.New()
		w := golog.NewSyslogWriter("tcp", "localhost", "test", 1)
		log.SetOutput(w)
		log.Print("is stderr")
	})
}
