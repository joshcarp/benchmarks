# pointer-vs-value
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/benchmarks/pointer-vs-value
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkValue
BenchmarkValue-16      	       1	23563411673 ns/op	874568648 B/op	17123729 allocs/op
BenchmarkPointer
BenchmarkPointer-16    	       1	22363352458 ns/op	873856744 B/op	17116937 allocs/op
PASS
ok  	github.com/joshcarp/benchmarks/pointer-vs-value	46.382s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
