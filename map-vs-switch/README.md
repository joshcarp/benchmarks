# map-vs-switch
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/benchmarks/map-vs-switch
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkOptionSwitch-16    	261121752	         4.675 ns/op
BenchmarkOptionMap-16       	11682766	       103.7 ns/op
PASS
ok  	github.com/joshcarp/benchmarks/map-vs-switch	3.247s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
