# map-vs-switch
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/benchmarks/map-vs-switch
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkOptionSwitch
BenchmarkOptionSwitch-16    	130292881	        13.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkOptionMap
BenchmarkOptionMap-16       	 4490700	       282.1 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/joshcarp/benchmarks/map-vs-switch	4.508s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
