package bench

import (
	"encoding/json"
	L "log"
	"os"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/chapsuk/golog"
	log "github.com/mgutz/logxi/v1"
	"github.com/uber-go/zap"
	"gopkg.in/inconshreveable/log15.v2"
)

type M map[string]interface{}

var testObject = M{
	"foo": "bar",
	"bah": M{
		"int":      1,
		"float":    -100.23,
		"date":     "06-01-01T15:04:05-0700",
		"bool":     true,
		"nullable": nil,
	},
}

var pid = os.Getpid()

func toJSON(m map[string]interface{}) string {
	b, _ := json.Marshal(m)
	return string(b)
}

func BenchmarkLog(b *testing.B) {
	l := L.New(os.Stderr, "bench ", L.LstdFlags)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		debug := map[string]interface{}{"l": "debug", "key1": 1, "key2": "string", "key3": false}
		l.Printf(toJSON(debug))

		info := map[string]interface{}{"l": "info", "key1": 1, "key2": "string", "key3": false}
		l.Printf(toJSON(info))

		warn := map[string]interface{}{"l": "warn", "key1": 1, "key2": "string", "key3": false}
		l.Printf(toJSON(warn))

		err := map[string]interface{}{"l": "error", "key1": 1, "key2": "string", "key3": false}
		l.Printf(toJSON(err))
	}
	b.StopTimer()
}

func BenchmarkLogComplex(b *testing.B) {
	l := L.New(os.Stderr, "bench ", L.LstdFlags)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		debug := map[string]interface{}{"l": "debug", "key1": 1, "obj": testObject}
		l.Printf(toJSON(debug))

		info := map[string]interface{}{"l": "info", "key1": 1, "obj": testObject}
		l.Printf(toJSON(info))

		warn := map[string]interface{}{"l": "warn", "key1": 1, "obj": testObject}
		l.Printf(toJSON(warn))

		err := map[string]interface{}{"l": "error", "key1": 1, "obj": testObject}
		l.Printf(toJSON(err))
	}
	b.StopTimer()
}

func BenchmarkLogxi(b *testing.B) {
	stdout := log.NewConcurrentWriter(os.Stderr)
	l := log.NewLogger3(stdout, "bench", log.NewJSONFormatter("bench"))
	l.SetLevel(log.LevelDebug)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Debug("debug", "key", 1, "key2", "string", "key3", false)
		l.Info("info", "key", 1, "key2", "string", "key3", false)
		l.Warn("warn", "key", 1, "key2", "string", "key3", false)
		l.Error("error", "key", 1, "key2", "string", "key3", false)
	}
	b.StopTimer()
}

func BenchmarkLogxiComplex(b *testing.B) {
	stdout := log.NewConcurrentWriter(os.Stderr)
	l := log.NewLogger3(stdout, "bench", log.NewJSONFormatter("bench"))
	l.SetLevel(log.LevelDebug)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Debug("debug", "key", 1, "obj", testObject)
		l.Info("info", "key", 1, "obj", testObject)
		l.Warn("warn", "key", 1, "obj", testObject)
		l.Error("error", "key", 1, "obj", testObject)
	}
	b.StopTimer()

}

func BenchmarkLogrus(b *testing.B) {
	l := logrus.New()
	l.Out = os.Stderr
	l.Formatter = &logrus.JSONFormatter{}

	context := logrus.Fields{"_n": "bench", "_p": pid, "key": 1, "key2": "string", "key3": false}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.WithFields(context).Debug("debug")
		l.WithFields(context).Info("info")
		l.WithFields(context).Warn("warn")
		l.WithFields(context).Error("error")
	}
	b.StopTimer()
}

func BenchmarkLogrusComplex(b *testing.B) {
	l := logrus.New()
	l.Out = os.Stderr
	l.Formatter = &logrus.JSONFormatter{}

	context := logrus.Fields{"_n": "bench", "_p": pid, "key": 1, "obj": testObject}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.WithFields(context).Debug("debug")
		l.WithFields(context).Info("info")
		l.WithFields(context).Warn("warn")
		l.WithFields(context).Error("error")
	}
	b.StopTimer()
}

func BenchmarkLog15(b *testing.B) {
	l := log15.New(log15.Ctx{"_n": "bench", "_p": pid})
	l.SetHandler(log15.SyncHandler(log15.StreamHandler(os.Stderr, log15.JsonFormat())))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Debug("debug", "key", 1, "key2", "string", "key3", false)
		l.Info("info", "key", 1, "key2", "string", "key3", false)
		l.Warn("warn", "key", 1, "key2", "string", "key3", false)
		l.Error("error", "key", 1, "key2", "string", "key3", false)
	}
	b.StopTimer()

}

func BenchmarkLog15Complex(b *testing.B) {
	l := log15.New(log15.Ctx{"_n": "bench", "_p": pid})
	l.SetHandler(log15.SyncHandler(log15.StreamHandler(os.Stderr, log15.JsonFormat())))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Debug("debug", "key", 1, "obj", testObject)
		l.Info("info", "key", 1, "obj", testObject)
		l.Warn("warn", "key", 1, "obj", testObject)
		l.Error("error", "key", 1, "obj", testObject)
	}
	b.StopTimer()
}

func BenchmarkGolog(b *testing.B) {
	out := golog.NewConcurrentWriter(os.Stderr)
	l := golog.NewLogger(out, &golog.JSONFormatter{DateFormat: "15:04:05.000000"}, golog.Context{
		"_n": "bench",
		"_p": pid,
	})

	ctx := golog.Context{"key": 1, "key2": "string", "key3": false}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.DebugCtx(ctx, "debug")
		l.InfoCtx(ctx, "info")
		l.WarnCtx(ctx, "warn")
		l.ErrorCtx(ctx, "error")
	}
	b.StopTimer()
}

func BenchmarkGologComplex(b *testing.B) {
	out := golog.NewConcurrentWriter(os.Stderr)
	l := golog.NewLogger(out, &golog.JSONFormatter{}, golog.Context{
		"_n": "bench",
		"_p": pid,
	})

	ctx := golog.Context{"key": 1, "obj": testObject}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.DebugCtx(ctx, "debug")
		l.InfoCtx(ctx, "info")
		l.WarnCtx(ctx, "warn")
		l.ErrorCtx(ctx, "error")
	}
	b.StopTimer()
}

func BenchmarkZap(b *testing.B) {
	context := []zap.Field{
		zap.String("_n", "bench"),
		zap.Int("_p", pid),
	}
	logger := zap.NewJSON(zap.All, zap.Output(zap.Discard), zap.Fields(context...))
	pContext := []zap.Field{
		zap.Int("key1", 1),
		zap.String("key2", "string"),
		zap.Bool("key", false),
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Debug("debug", pContext...)
		logger.Info("info", pContext...)
		logger.Warn("warn", pContext...)
		logger.Error("error", pContext...)
	}
	b.StopTimer()
}

func BenchmarkZapComplex(b *testing.B) {
	context := []zap.Field{
		zap.String("_n", "bench"),
		zap.Int("_p", pid),
	}

	logger := zap.NewJSON(zap.All, zap.Output(zap.Discard), zap.Fields(context...))
	pContext := []zap.Field{
		zap.Int("key", 1),
		zap.Object("obj", testObject),
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Debug("debug", pContext...)
		logger.Info("info", pContext...)
		logger.Warn("warn", pContext...)
		logger.Error("error", pContext...)
	}
	b.StopTimer()
}
