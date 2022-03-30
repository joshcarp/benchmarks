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
BenchmarkFuncValue-16                       	99755755	        12.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkFuncPointer
BenchmarkFuncPointer-16                     	99119487	        12.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodValue
BenchmarkMethodValue-16                     	99165169	        12.25 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodPointer
BenchmarkMethodPointer-16                   	94586066	        12.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkFuncValueNoRef
BenchmarkFuncValueNoRef-16                  	95687622	        12.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkFuncPointerNoRef
BenchmarkFuncPointerNoRef-16                	96082884	        12.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodValueNoRef
BenchmarkMethodValueNoRef-16                	88197932	        12.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodPointerNoRef
BenchmarkMethodPointerNoRef-16              	99025407	        12.21 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedNewValue
BenchmarkUnexportedNewValue-16              	100000000	        11.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedNewPointer
BenchmarkUnexportedNewPointer-16            	100000000	        11.70 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedFuncValue
BenchmarkUnexportedFuncValue-16             	98311483	        12.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedFuncPointer
BenchmarkUnexportedFuncPointer-16           	97381294	        12.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedMethodValue
BenchmarkUnexportedMethodValue-16           	98216671	        12.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedMethodPointer
BenchmarkUnexportedMethodPointer-16         	99461654	        12.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedFuncValueNoRef
BenchmarkUnexportedFuncValueNoRef-16        	99608091	        12.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedFuncPointerNoRef
BenchmarkUnexportedFuncPointerNoRef-16      	99658279	        12.17 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedMethodValueNoRef
BenchmarkUnexportedMethodValueNoRef-16      	99149527	        12.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedMethodPointerNoRef
BenchmarkUnexportedMethodPointerNoRef-16    	99347756	        12.10 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/joshcarp/go-benchmarks/pointer-vs-value-raw	24.540s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
