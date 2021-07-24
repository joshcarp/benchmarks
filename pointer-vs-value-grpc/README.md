# pointer-vs-value-grpc
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/go-benchmarks/pointer-vs-value-grpc
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkValue
BenchmarkValue-16      	       1	22143511877 ns/op	874585968 B/op	17123989 allocs/op
BenchmarkPointer
BenchmarkPointer-16    	       1	22428452027 ns/op	873895824 B/op	17117716 allocs/op
PASS
ok  	github.com/joshcarp/go-benchmarks/pointer-vs-value-grpc	44.820s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
