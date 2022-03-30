# decimal
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/go-benchmarks/decimal
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkDecimal
BenchmarkDecimal/ddAdd.decTest_float64
BenchmarkDecimal/ddAdd.decTest_float64-16         	      97	  12119767 ns/op	      42 B/op	       0 allocs/op
BenchmarkDecimal/ddAdd.decTest_bigFloat
BenchmarkDecimal/ddAdd.decTest_bigFloat-16        	       6	 182049291 ns/op	     944 B/op	       0 allocs/op
BenchmarkDecimal/ddAdd.decTest_anzDecimal
BenchmarkDecimal/ddAdd.decTest_anzDecimal-16      	       2	1010268956 ns/op	    9216 B/op	       1 allocs/op
BenchmarkDecimal/ddMultiply.decTest_anzDecimal
BenchmarkDecimal/ddMultiply.decTest_anzDecimal-16 	       1	1140133497 ns/op	    9104 B/op	      40 allocs/op
BenchmarkDecimal/ddMultiply.decTest_float64
BenchmarkDecimal/ddMultiply.decTest_float64-16    	      61	  17435928 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecimal/ddMultiply.decTest_bigFloat
BenchmarkDecimal/ddMultiply.decTest_bigFloat-16   	       4	 271896474 ns/op	    2048 B/op	       0 allocs/op
BenchmarkDecimal/ddAbs.decTest_anzDecimal
BenchmarkDecimal/ddAbs.decTest_anzDecimal-16      	       1	1129041380 ns/op	    9080 B/op	      38 allocs/op
BenchmarkDecimal/ddAbs.decTest_float64
BenchmarkDecimal/ddAbs.decTest_float64-16         	      68	  17484972 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecimal/ddAbs.decTest_bigFloat
BenchmarkDecimal/ddAbs.decTest_bigFloat-16        	       4	 277385747 ns/op	    5120 B/op	       0 allocs/op
BenchmarkDecimal/ddDivide.decTest_anzDecimal
BenchmarkDecimal/ddDivide.decTest_anzDecimal-16   	       1	2216328768 ns/op	   17272 B/op	      39 allocs/op
BenchmarkDecimal/ddDivide.decTest_float64
BenchmarkDecimal/ddDivide.decTest_float64-16      	      45	  25261115 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecimal/ddDivide.decTest_bigFloat
BenchmarkDecimal/ddDivide.decTest_bigFloat-16     	       4	 276561992 ns/op	   11472 B/op	     717 allocs/op
PASS
ok  	github.com/joshcarp/go-benchmarks/decimal	21.148s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
