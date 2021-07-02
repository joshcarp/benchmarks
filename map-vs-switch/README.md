# map-vs-switch
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/benchmarks/map-vs-switch
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkOptionSwitch-16    	254830975	         4.519 ns/op
BenchmarkOptionMap-16       	11139555	       116.6 ns/op
PASS
ok  	github.com/joshcarp/benchmarks/map-vs-switch	3.229s
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
