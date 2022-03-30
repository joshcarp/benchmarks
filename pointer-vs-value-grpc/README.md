# pointer-vs-value-grpc
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/go-benchmarks/pointer-vs-value-grpc
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkValue
BenchmarkValue-16      	    2079	    804733 ns/op	    9131 B/op	     172 allocs/op
BenchmarkPointer
BenchmarkPointer-16    	    1280	    866747 ns/op	    9195 B/op	     172 allocs/op
PASS
ok  	github.com/joshcarp/go-benchmarks/pointer-vs-value-grpc	4.359s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
