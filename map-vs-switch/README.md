# map-vs-switch
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/go-benchmarks/map-vs-switch
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkOptionSwitch
BenchmarkOptionSwitch-16       	 5401515	       195.6 ns/op	       1 B/op	       0 allocs/op
BenchmarkOptionMap
BenchmarkOptionMap-16          	  458192	      2535 ns/op	    2328 B/op	       1 allocs/op
BenchmarkOptionGlobalMap
BenchmarkOptionGlobalMap-16    	 5806516	       184.6 ns/op	       1 B/op	       0 allocs/op
PASS
ok  	github.com/joshcarp/go-benchmarks/map-vs-switch	4.090s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
