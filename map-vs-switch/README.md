# map-vs-switch
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/go-benchmarks/map-vs-switch
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkOptionSwitch
BenchmarkOptionSwitch-16    	 6685606	       224.8 ns/op	       1 B/op	       0 allocs/op
BenchmarkOptionMap
BenchmarkOptionMap-16       	  479304	      2423 ns/op	    2328 B/op	       1 allocs/op
PASS
ok  	github.com/joshcarp/go-benchmarks/map-vs-switch	3.928s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
