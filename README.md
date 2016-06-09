# Golog

[![Build Status](https://travis-ci.org/chapsuk/golog.svg)](https://travis-ci.org/chapsuk/golog)
[![Coverage Status](https://coveralls.io/repos/github/chapsuk/golog/badge.svg?branch=master)](https://coveralls.io/github/chapsuk/golog?branch=master)

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

## Install

```
go get github.com/chapsuk/golog
```

## Usage

See [example](examples/main.go).

## Tests

1. Install [goconvey](https://github.com/smartystreets/goconvey)
1. Run `$GOPATH/bin/goconvey` for watch test result in browser, or `go test -v ./...`.   
