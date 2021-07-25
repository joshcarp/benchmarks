# map-vs-switch
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/go-benchmarks/map-vs-switch
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkOptionSwitch
BenchmarkOptionSwitch-16       	43833154	        31.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkOptionMap
BenchmarkOptionMap-16          	  424086	      2595 ns/op	    2324 B/op	       1 allocs/op
BenchmarkOptionGlobalMap
BenchmarkOptionGlobalMap-16    	85426786	        20.55 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/joshcarp/go-benchmarks/map-vs-switch	6.924s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
