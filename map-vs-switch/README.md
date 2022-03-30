# map-vs-switch
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/go-benchmarks/map-vs-switch
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkOptionSwitch
BenchmarkOptionSwitch-16       	25133077	        47.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkOptionMap
BenchmarkOptionMap-16          	  328146	      3681 ns/op	    2325 B/op	       1 allocs/op
BenchmarkOptionGlobalMap
BenchmarkOptionGlobalMap-16    	17213882	        63.00 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/joshcarp/go-benchmarks/map-vs-switch	4.949s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
