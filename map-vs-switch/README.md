# map-vs-switch
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/benchmarks/map-vs-switch
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkOptionSwitch
BenchmarkOptionSwitch-16    	100000000	        12.88 ns/op
BenchmarkOptionMap
BenchmarkOptionMap-16       	 5617432	       320.9 ns/op
PASS
ok  	github.com/joshcarp/benchmarks/map-vs-switch	3.832s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
