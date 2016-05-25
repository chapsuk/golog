# Golog

Hybrid from [logrus](https://github.com/Sirupsen/logrus) 
and [logxi](https://github.com/mgutz/logxi/), thanks authors. ^_^

## Motivation

* fast as [logxi](https://github.com/mgutz/logxi/)
* static and private context for log message as [logrus](https://github.com/Sirupsen/logrus)
* without dependecies ([goconvey](https://github.com/smartystreets/goconvey) for tests)
* tests  

## Benchmark

Benchmark [source](https://github.com/mgutz/logxi/tree/master/v1/bench).

```
BenchmarkLog-4          	   50000	     28870 ns/op	    4321 B/op	     100 allocs/op
BenchmarkLogComplex-4   	   30000	     43073 ns/op	    8098 B/op	     188 allocs/op

BenchmarkLogxi-4        	  100000	     19857 ns/op	    2321 B/op	      58 allocs/op
BenchmarkLogxiComplex-4 	   50000	     39897 ns/op	    5458 B/op	     154 allocs/op

BenchmarkLogrus-4       	   30000	     42155 ns/op	    8106 B/op	     165 allocs/op
BenchmarkLogrusComplex-4	   30000	     51275 ns/op	   10843 B/op	     229 allocs/op

BenchmarkLog15-4        	   30000	     57870 ns/op	    7970 B/op	     192 allocs/op
BenchmarkLog15Complex-4 	   20000	     86362 ns/op	   11099 B/op	     242 allocs/op

BenchmarkGolog-4        	  100000	     23248 ns/op	    3136 B/op	      52 allocs/op
BenchmarkGologComplex-4 	   30000	     44190 ns/op	    6529 B/op	     152 allocs/op
```

## Install

```
go get github.com/chapsuk/golog
```

## Usage

See [example](examples/main.go).

## Tests

1. Install [goconvey](https://github.com/smartystreets/goconvey)
1. Run `$GOPATH/bin/goconvey` for watch test result in browser, or `go test ./...`.   
