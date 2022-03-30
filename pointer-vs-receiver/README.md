# pointer-vs-receiver
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/go-benchmarks/pointer-vs-receiver
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkMutateFoo
BenchmarkMutateFoo-16        	 6348207	       185.2 ns/op	      80 B/op	       1 allocs/op
BenchmarkDontMutateFoo
BenchmarkDontMutateFoo-16    	 5360542	       201.9 ns/op	      80 B/op	       1 allocs/op
PASS
ok  	github.com/joshcarp/go-benchmarks/pointer-vs-receiver	3.099s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
