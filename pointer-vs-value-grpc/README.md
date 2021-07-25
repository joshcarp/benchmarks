# pointer-vs-value-grpc
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/go-benchmarks/pointer-vs-value-grpc
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkValue
BenchmarkValue-16      	    1602	    864118 ns/op	    8935 B/op	     171 allocs/op
BenchmarkPointer
BenchmarkPointer-16    	    1370	    966972 ns/op	    8956 B/op	     172 allocs/op
PASS
ok  	github.com/joshcarp/go-benchmarks/pointer-vs-value-grpc	4.615s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
