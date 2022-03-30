# map-vs-switch
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/go-benchmarks/map-vs-switch
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkOptionSwitch
BenchmarkOptionSwitch-16       	27809196	        43.89 ns/op	       0 B/op	       0 allocs/op
BenchmarkOptionMap
BenchmarkOptionMap-16          	  322310	      4117 ns/op	    2324 B/op	       1 allocs/op
BenchmarkOptionGlobalMap
BenchmarkOptionGlobalMap-16    	15459291	        76.21 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/joshcarp/go-benchmarks/map-vs-switch	6.246s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
