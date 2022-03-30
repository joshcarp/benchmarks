# pointer-vs-value-raw
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/go-benchmarks/pointer-vs-value-raw
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkNewValue
BenchmarkNewValue-16                        	100000000	        11.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewPointer
BenchmarkNewPointer-16                      	100000000	        11.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkFuncValue
BenchmarkFuncValue-16                       	99598881	        12.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkFuncPointer
BenchmarkFuncPointer-16                     	99243926	        12.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodValue
BenchmarkMethodValue-16                     	100000000	        12.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodPointer
BenchmarkMethodPointer-16                   	99790224	        12.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkFuncValueNoRef
BenchmarkFuncValueNoRef-16                  	88488148	        13.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkFuncPointerNoRef
BenchmarkFuncPointerNoRef-16                	99897520	        12.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodValueNoRef
BenchmarkMethodValueNoRef-16                	99646180	        12.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodPointerNoRef
BenchmarkMethodPointerNoRef-16              	99989250	        12.04 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedNewValue
BenchmarkUnexportedNewValue-16              	100000000	        11.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedNewPointer
BenchmarkUnexportedNewPointer-16            	100000000	        11.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedFuncValue
BenchmarkUnexportedFuncValue-16             	99312376	        12.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedFuncPointer
BenchmarkUnexportedFuncPointer-16           	99326787	        12.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedMethodValue
BenchmarkUnexportedMethodValue-16           	99314604	        12.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedMethodPointer
BenchmarkUnexportedMethodPointer-16         	99785667	        12.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedFuncValueNoRef
BenchmarkUnexportedFuncValueNoRef-16        	99320440	        12.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedFuncPointerNoRef
BenchmarkUnexportedFuncPointerNoRef-16      	99632967	        12.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedMethodValueNoRef
BenchmarkUnexportedMethodValueNoRef-16      	99929565	        12.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedMethodPointerNoRef
BenchmarkUnexportedMethodPointerNoRef-16    	100000000	        12.05 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/joshcarp/go-benchmarks/pointer-vs-value-raw	24.471s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
