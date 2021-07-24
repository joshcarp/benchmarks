# pointer-vs-value-raw
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/go-benchmarks/pointer-vs-value-raw
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkNewValue
BenchmarkNewValue-16                        	1000000000	         0.8634 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewPointer
BenchmarkNewPointer-16                      	1000000000	         0.7018 ns/op	       0 B/op	       0 allocs/op
BenchmarkFuncValue
BenchmarkFuncValue-16                       	1000000000	         0.7020 ns/op	       0 B/op	       0 allocs/op
BenchmarkFuncPointer
BenchmarkFuncPointer-16                     	870921771	         1.329 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodValue
BenchmarkMethodValue-16                     	1000000000	         0.6360 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodPointer
BenchmarkMethodPointer-16                   	1000000000	         1.527 ns/op	       0 B/op	       0 allocs/op
BenchmarkFuncValueNoRef
BenchmarkFuncValueNoRef-16                  	1000000000	         1.076 ns/op	       0 B/op	       0 allocs/op
BenchmarkFuncPointerNoRef
BenchmarkFuncPointerNoRef-16                	1000000000	         1.965 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodValueNoRef
BenchmarkMethodValueNoRef-16                	1000000000	         0.8299 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodPointerNoRef
BenchmarkMethodPointerNoRef-16              	779971207	         1.317 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedNewValue
BenchmarkUnexportedNewValue-16              	1000000000	         1.182 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedNewPointer
BenchmarkUnexportedNewPointer-16            	1000000000	         1.378 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedFuncValue
BenchmarkUnexportedFuncValue-16             	1000000000	         0.7871 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedFuncPointer
BenchmarkUnexportedFuncPointer-16           	1000000000	         1.248 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedMethodValue
BenchmarkUnexportedMethodValue-16           	1000000000	         0.6898 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedMethodPointer
BenchmarkUnexportedMethodPointer-16         	999932420	         1.369 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedFuncValueNoRef
BenchmarkUnexportedFuncValueNoRef-16        	1000000000	         0.8663 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedFuncPointerNoRef
BenchmarkUnexportedFuncPointerNoRef-16      	1000000000	         1.267 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedMethodValueNoRef
BenchmarkUnexportedMethodValueNoRef-16      	1000000000	         0.7132 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedMethodPointerNoRef
BenchmarkUnexportedMethodPointerNoRef-16    	995335682	         1.187 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/joshcarp/go-benchmarks/pointer-vs-value-raw	24.232s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
