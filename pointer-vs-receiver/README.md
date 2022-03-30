# pointer-vs-receiver
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/go-benchmarks/pointer-vs-receiver
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkMutateFoo
BenchmarkMutateFoo-16        	 7450215	       152.2 ns/op	      80 B/op	       1 allocs/op
BenchmarkDontMutateFoo
BenchmarkDontMutateFoo-16    	 6999441	       170.3 ns/op	      80 B/op	       1 allocs/op
PASS
ok  	github.com/joshcarp/go-benchmarks/pointer-vs-receiver	2.984s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
