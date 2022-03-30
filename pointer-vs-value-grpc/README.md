# pointer-vs-value-grpc
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/go-benchmarks/pointer-vs-value-grpc
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkValue
BenchmarkValue-16      	    1909	    546017 ns/op	    9161 B/op	     172 allocs/op
BenchmarkPointer
BenchmarkPointer-16    	    2119	    555806 ns/op	    9113 B/op	     172 allocs/op
PASS
ok  	github.com/joshcarp/go-benchmarks/pointer-vs-value-grpc	2.828s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
