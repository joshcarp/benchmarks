# map-vs-switch
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/go-benchmarks/map-vs-switch
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkOptionSwitch
BenchmarkOptionSwitch-16    	138591686	         8.110 ns/op	       0 B/op	       0 allocs/op
BenchmarkOptionMap
BenchmarkOptionMap-16       	 5195742	       248.5 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/joshcarp/go-benchmarks/map-vs-switch	3.789s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
