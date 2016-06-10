# Golog

[![Build Status](https://travis-ci.org/chapsuk/golog.svg)](https://travis-ci.org/chapsuk/golog)
[![Coverage Status](https://coveralls.io/repos/github/chapsuk/golog/badge.svg?branch=master)](https://coveralls.io/github/chapsuk/golog?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/chapsuk/golog)](https://goreportcard.com/report/github.com/chapsuk/golog)

## Install

```
go get github.com/chapsuk/golog
```

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
```
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

```
BenchmarkLog-4          	  100000	     24277 ns/op	    4321 B/op	     100 allocs/op
BenchmarkLogComplex-4   	   30000	     42979 ns/op	    8098 B/op	     188 allocs/op

BenchmarkLogxi-4        	  100000	     18967 ns/op	    2321 B/op	      58 allocs/op
BenchmarkLogxiComplex-4 	   30000	     39623 ns/op	    5458 B/op	     154 allocs/op

BenchmarkLogrus-4       	   50000	     35746 ns/op	    6553 B/op	     137 allocs/op
BenchmarkLogrusComplex-4	   30000	     49963 ns/op	    9362 B/op	     209 allocs/op

BenchmarkLog15-4        	   30000	     54347 ns/op	    7970 B/op	     192 allocs/op
BenchmarkLog15Complex-4 	   20000	     80948 ns/op	   11210 B/op	     242 allocs/op

BenchmarkGolog-4        	  100000	     20129 ns/op	    1693 B/op	      32 allocs/op
BenchmarkGologComplex-4 	   30000	     42662 ns/op	    5150 B/op	     140 allocs/op

BenchmarkZap-4          	  500000	      2513 ns/op	       0 B/op	       0 allocs/op
BenchmarkZapComplex-4   	   50000	     28183 ns/op	    4129 B/op	     116 allocs/op
```

## License

[MIT](https://github.com/chapsuk/golog/blob/master/LICENSE)
