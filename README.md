# Golog

[![Build Status](https://travis-ci.org/chapsuk/golog.svg)](https://travis-ci.org/chapsuk/golog)
[![Coverage Status](https://coveralls.io/repos/github/chapsuk/golog/badge.svg?branch=master)](https://coveralls.io/github/chapsuk/golog?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/chapsuk/golog)](https://goreportcard.com/report/github.com/chapsuk/golog)

## Install

`go get github.com/chapsuk/golog`

## Usage

```go
package main

import (
    "github.com/chapsuk/golog"
)

func main() {
    // golog.SetFormatter(&golog.JSONFormatter)
    // golog.SetFormatter(&golog.LogstashFormatter{})
    golog.SetFormatter(&golog.TextFormatter{}) // not needed, TextFormatter by default
    golog.Info("Hallo")

    logger := golog.WithContext(golog.Context{
        "host":  "localhost",
        "place": "example",
    })
    logger.Warn("why so serious?!")

    logger.ErrorCtx(golog.Context{
        "uid": 666,
    }, "Omen detected")

    logger.Fatalf("The %s!", "end")
}
```

Output:

```bash
// JSONFormatter
{"_t":"2016-06-10T12:26:20+03:00", "_l":"INFO", "_m":"Hallo"}
{"_t":"2016-06-10T12:26:20+03:00", "_l":"WARN", "host":"localhost", "place":"example", "_m":"why so serious?!"}
{"_t":"2016-06-10T12:26:20+03:00", "_l":"ERROR", "host":"localhost", "place":"example", "uid":666, "_m":"Omen detected"}
{"_t":"2016-06-10T12:26:20+03:00", "_l":"FATAL", "host":"localhost", "place":"example", "uid":666, "_m":"The end!"}

// LogstashFormatter
{"@timestamp":"2016-06-10T12:26:35+03:00","@version":1,"level":"INFO","message":"Hallo"}
{"@timestamp":"2016-06-10T12:26:35+03:00","@version":1,"level":"WARN","host":"localhost","place":"example","message":"why so serious?!"}
{"@timestamp":"2016-06-10T12:26:35+03:00","@version":1,"level":"ERROR","host":"localhost","place":"example","uid":666,"message":"Omen detected"}
{"@timestamp":"2016-06-10T12:26:35+03:00","@version":1,"level":"FATAL","host":"localhost","place":"example","uid":666,"message":"The end!"}

// TextFormatter
2016-06-10T12:22:01+03:00  INFO  []  Hallo
2016-06-10T12:22:01+03:00  WARN  [place: example, host: localhost]  why so serious?!
2016-06-10T12:22:01+03:00  ERROR  [uid: 666, place: example, host: localhost]  Omen detected
2016-06-10T12:22:01+03:00  FATAL  [place: example, host: localhost, uid: 666]  The end!
exit status 1
```

## Tests

1. Install [goconvey](https://github.com/smartystreets/goconvey)
1. Run `$GOPATH/bin/goconvey` for watch test result in browser, or `go test -v ./...`.   


## Benchmark

Benchmark [source](https://github.com/chapsuk/golog/tree/master/bench/bench_test.go).

```bash
≻ go test -v -bench=. -benchmem ./bench/bench_test.go 2>/dev/null
BenchmarkLog-4                    100000             31437 ns/op            5345 B/op        112 allocs/op
BenchmarkLogComplex-4              20000             50919 ns/op           10915 B/op        224 allocs/op
BenchmarkLogxi-4                  100000             17629 ns/op            2321 B/op         58 allocs/op
BenchmarkLogxiComplex-4            30000             40875 ns/op            7379 B/op        178 allocs/op
BenchmarkLogrus-4                  30000             49576 ns/op            8253 B/op        139 allocs/op
BenchmarkLogrusComplex-4           20000             50315 ns/op           11544 B/op        229 allocs/op
BenchmarkLog15-4                   30000             48401 ns/op            9634 B/op        204 allocs/op
BenchmarkLog15Complex-4            20000             96301 ns/op           12751 B/op        254 allocs/op
BenchmarkGolog-4                   50000             22442 ns/op            3060 B/op         52 allocs/op
BenchmarkGologComplex-4            30000             39815 ns/op            8438 B/op        184 allocs/op
BenchmarkZapSugar-4              1000000              2200 ns/op             350 B/op         24 allocs/op
BenchmarkZapSugarComplex-4        100000             17372 ns/op            2386 B/op         73 allocs/op
BenchmarkZap-4                   1000000              1884 ns/op             785 B/op          4 allocs/op
BenchmarkZapComplex-4             500000              2050 ns/op             576 B/op          5 allocs/op
PASS
ok      command-line-arguments  28.506s
```

## License

[MIT](https://github.com/chapsuk/golog/blob/master/LICENSE)
