# decimal
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/go-benchmarks/decimal
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkDecimal
BenchmarkDecimal/ddAdd.decTest_float64
BenchmarkDecimal/ddAdd.decTest_float64-16         	      96	  15905259 ns/op	      43 B/op	       0 allocs/op
BenchmarkDecimal/ddAdd.decTest_bigFloat
BenchmarkDecimal/ddAdd.decTest_bigFloat-16        	       6	 181488123 ns/op	    2285 B/op	       0 allocs/op
BenchmarkDecimal/ddAdd.decTest_anzDecimal
BenchmarkDecimal/ddAdd.decTest_anzDecimal-16      	       2	 959526394 ns/op	    9432 B/op	       2 allocs/op
BenchmarkDecimal/ddMultiply.decTest_anzDecimal
BenchmarkDecimal/ddMultiply.decTest_anzDecimal-16 	       1	1107062047 ns/op	    9096 B/op	      39 allocs/op
BenchmarkDecimal/ddMultiply.decTest_float64
BenchmarkDecimal/ddMultiply.decTest_float64-16    	      67	  21035162 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecimal/ddMultiply.decTest_bigFloat
BenchmarkDecimal/ddMultiply.decTest_bigFloat-16   	       4	 271200919 ns/op	    2048 B/op	       0 allocs/op
BenchmarkDecimal/ddAbs.decTest_anzDecimal
BenchmarkDecimal/ddAbs.decTest_anzDecimal-16      	       1	1098235285 ns/op	    9368 B/op	      40 allocs/op
BenchmarkDecimal/ddAbs.decTest_float64
BenchmarkDecimal/ddAbs.decTest_float64-16         	      68	  21054261 ns/op	       2 B/op	       0 allocs/op
BenchmarkDecimal/ddAbs.decTest_bigFloat
BenchmarkDecimal/ddAbs.decTest_bigFloat-16        	       4	 274590038 ns/op	    5120 B/op	       0 allocs/op
BenchmarkDecimal/ddDivide.decTest_anzDecimal
BenchmarkDecimal/ddDivide.decTest_anzDecimal-16   	       1	2141023322 ns/op	   17272 B/op	      39 allocs/op
BenchmarkDecimal/ddDivide.decTest_float64
BenchmarkDecimal/ddDivide.decTest_float64-16      	      45	  29016813 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecimal/ddDivide.decTest_bigFloat
BenchmarkDecimal/ddDivide.decTest_bigFloat-16     	       4	 275350758 ns/op	   11472 B/op	     717 allocs/op
PASS
ok  	github.com/joshcarp/go-benchmarks/decimal	21.979s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
