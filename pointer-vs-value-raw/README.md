# pointer-vs-value-raw
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/go-benchmarks/pointer-vs-value-raw
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkNewValue
BenchmarkNewValue-16                        	100000000	        11.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewPointer
BenchmarkNewPointer-16                      	100000000	        11.61 ns/op	       0 B/op	       0 allocs/op
BenchmarkFuncValue
BenchmarkFuncValue-16                       	99289747	        12.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkFuncPointer
BenchmarkFuncPointer-16                     	99713017	        12.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodValue
BenchmarkMethodValue-16                     	99990400	        12.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodPointer
BenchmarkMethodPointer-16                   	99777056	        12.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkFuncValueNoRef
BenchmarkFuncValueNoRef-16                  	99603031	        12.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkFuncPointerNoRef
BenchmarkFuncPointerNoRef-16                	99103092	        12.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodValueNoRef
BenchmarkMethodValueNoRef-16                	99628418	        12.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodPointerNoRef
BenchmarkMethodPointerNoRef-16              	98464777	        12.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedNewValue
BenchmarkUnexportedNewValue-16              	100000000	        11.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedNewPointer
BenchmarkUnexportedNewPointer-16            	100000000	        11.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedFuncValue
BenchmarkUnexportedFuncValue-16             	99017760	        12.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedFuncPointer
BenchmarkUnexportedFuncPointer-16           	99773804	        12.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedMethodValue
BenchmarkUnexportedMethodValue-16           	98733511	        12.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedMethodPointer
BenchmarkUnexportedMethodPointer-16         	99382329	        12.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedFuncValueNoRef
BenchmarkUnexportedFuncValueNoRef-16        	99248686	        12.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedFuncPointerNoRef
BenchmarkUnexportedFuncPointerNoRef-16      	99772816	        12.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedMethodValueNoRef
BenchmarkUnexportedMethodValueNoRef-16      	100000000	        12.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkUnexportedMethodPointerNoRef
BenchmarkUnexportedMethodPointerNoRef-16    	99583714	        12.07 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/joshcarp/go-benchmarks/pointer-vs-value-raw	24.601s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
