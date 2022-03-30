# decimal
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/go-benchmarks/decimal
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkDecimal
BenchmarkDecimal/ddAdd.decTest_anzDecimal
BenchmarkDecimal/ddAdd.decTest_anzDecimal-16         	       1	1028243709 ns/op	    4296 B/op	       6 allocs/op
BenchmarkDecimal/ddAdd.decTest_float64
BenchmarkDecimal/ddAdd.decTest_float64-16            	      86	  13208797 ns/op	     157 B/op	       0 allocs/op
BenchmarkDecimal/ddAdd.decTest_bigFloat
BenchmarkDecimal/ddAdd.decTest_bigFloat-16           	       6	 231000830 ns/op	    1413 B/op	       0 allocs/op
BenchmarkDecimal/ddMultiply.decTest_anzDecimal
BenchmarkDecimal/ddMultiply.decTest_anzDecimal-16    	       1	1179225579 ns/op	    9080 B/op	      38 allocs/op
BenchmarkDecimal/ddMultiply.decTest_float64
BenchmarkDecimal/ddMultiply.decTest_float64-16       	      62	  18838072 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecimal/ddMultiply.decTest_bigFloat
BenchmarkDecimal/ddMultiply.decTest_bigFloat-16      	       4	 276929492 ns/op	      36 B/op	       0 allocs/op
BenchmarkDecimal/ddAbs.decTest_anzDecimal
BenchmarkDecimal/ddAbs.decTest_anzDecimal-16         	       1	1102807163 ns/op	   17704 B/op	      42 allocs/op
BenchmarkDecimal/ddAbs.decTest_float64
BenchmarkDecimal/ddAbs.decTest_float64-16            	      67	  17349717 ns/op	       2 B/op	       0 allocs/op
BenchmarkDecimal/ddAbs.decTest_bigFloat
BenchmarkDecimal/ddAbs.decTest_bigFloat-16           	       4	 277474951 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecimal/ddDivide.decTest_anzDecimal
BenchmarkDecimal/ddDivide.decTest_anzDecimal-16      	       1	2237572467 ns/op	   17272 B/op	      39 allocs/op
BenchmarkDecimal/ddDivide.decTest_float64
BenchmarkDecimal/ddDivide.decTest_float64-16         	      44	  25278812 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecimal/ddDivide.decTest_bigFloat
BenchmarkDecimal/ddDivide.decTest_bigFloat-16        	       4	 282055523 ns/op	   11472 B/op	     717 allocs/op
PASS
ok  	github.com/joshcarp/go-benchmarks/decimal	20.500s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
