# map-vs-switch
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/benchmarks/map-vs-switch
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkOptionSwitch
BenchmarkOptionSwitch-16    	100000000	        10.32 ns/op
BenchmarkOptionMap
BenchmarkOptionMap-16       	 4500883	       251.7 ns/op
PASS
ok  	github.com/joshcarp/benchmarks/map-vs-switch	2.792s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
