# pointer-vs-value-grpc
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/go-benchmarks/pointer-vs-value-grpc
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkValue
BenchmarkValue-16      	    1891	    543385 ns/op	    9159 B/op	     172 allocs/op
BenchmarkPointer
BenchmarkPointer-16    	    2088	    554384 ns/op	    9109 B/op	     171 allocs/op
PASS
ok  	github.com/joshcarp/go-benchmarks/pointer-vs-value-grpc	2.830s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
