# pointer-vs-value-grpc
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/go-benchmarks/pointer-vs-value-grpc
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkValue
BenchmarkValue-16      	    5400	    250723 ns/op	    8779 B/op	     171 allocs/op
BenchmarkPointer
BenchmarkPointer-16    	    5527	    228442 ns/op	    8770 B/op	     171 allocs/op
PASS
ok  	github.com/joshcarp/go-benchmarks/pointer-vs-value-grpc	4.270s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
