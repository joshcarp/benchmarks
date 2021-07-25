# pointer-vs-value-raw
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/go-benchmarks/pointer-vs-value-raw
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkNewValue
BenchmarkNewValue-16                        	75891573	        19.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewPointer
BenchmarkNewPointer-16                      	84328722	        16.72 ns/op	       0 B/op	       0 allocs/op
BenchmarkFuncValue
BenchmarkFuncValue-16                       	79138206	        17.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkFuncPointer
BenchmarkFuncPointer-16                     	78653467	        16.94 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodValue
BenchmarkMethodValue-16                     	74656876	        21.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodPointer
BenchmarkMethodPointer-16                   	74676859	        19.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkFuncValueNoRef
BenchmarkFuncValueNoRef-16                  	77773449	        20.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkFuncPointerNoRef
BenchmarkFuncPointerNoRef-16                	64547005	        18.04 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodValueNoRef
BenchmarkMethodValueNoRef-16                	82151716	        17.29 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodPointerNoRef
BenchmarkMethodPointerNoRef-16              	69237117	        18.04 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedNewValue
BenchmarkUnexportedNewValue-16              	94773913	        19.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedNewPointer
BenchmarkUnexportedNewPointer-16            	79575310	        21.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedFuncValue
BenchmarkUnexportedFuncValue-16             	67700712	        22.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedFuncPointer
BenchmarkUnexportedFuncPointer-16           	77470975	        20.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedMethodValue
BenchmarkUnexportedMethodValue-16           	78326869	        19.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedMethodPointer
BenchmarkUnexportedMethodPointer-16         	77463553	        19.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedFuncValueNoRef
BenchmarkUnexportedFuncValueNoRef-16        	72819846	        17.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedFuncPointerNoRef
BenchmarkUnexportedFuncPointerNoRef-16      	77781202	        18.02 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedMethodValueNoRef
BenchmarkUnexportedMethodValueNoRef-16      	72206122	        17.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedMethodPointerNoRef
BenchmarkUnexportedMethodPointerNoRef-16    	75587145	        19.35 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/joshcarp/go-benchmarks/pointer-vs-value-raw	30.893s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
