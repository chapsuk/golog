package golog_test

import (
	"github.com/chapsuk/golog"
	test "github.com/smartystreets/goconvey/convey"
	"testing"
)

const (
	foo = "foo"
	bar = "bar"
	baz = "baz"
	bee = "bee"
	bap = "bap"
)

type FakeWriter struct {
	log []string
}

func (fw *FakeWriter) Write(m []byte) (int, error) {
	fw.log = append(fw.log, string(m))
	return len(m), nil
}

func TestCuncurrentWriter(t *testing.T) {
	test.Convey("Write many messages", t, func() {
		fw := &FakeWriter{}
		w := golog.NewConcurrentWriter(fw)
		w.Write([]byte(foo))
		w.Write([]byte(bar))
		w.Write([]byte(baz))

		test.So(fw.log, test.ShouldContain, foo)
		test.So(fw.log, test.ShouldContain, bar)
		test.So(fw.log, test.ShouldContain, baz)
	})
}
