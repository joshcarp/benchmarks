# decimal
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/go-benchmarks/decimal
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkDecimal
BenchmarkDecimal/ddAdd.decTest_anzDecimal
BenchmarkDecimal/ddAdd.decTest_anzDecimal-16         	       1	1442911424 ns/op	    8264 B/op	       6 allocs/op
BenchmarkDecimal/ddAdd.decTest_float64
BenchmarkDecimal/ddAdd.decTest_float64-16            	      72	  17096872 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecimal/ddAdd.decTest_bigFloat
BenchmarkDecimal/ddAdd.decTest_bigFloat-16           	       4	 265018015 ns/op	    2048 B/op	       0 allocs/op
BenchmarkDecimal/ddMultiply.decTest_anzDecimal
BenchmarkDecimal/ddMultiply.decTest_anzDecimal-16    	       1	1563451530 ns/op	     888 B/op	      37 allocs/op
BenchmarkDecimal/ddMultiply.decTest_float64
BenchmarkDecimal/ddMultiply.decTest_float64-16       	      49	  25779207 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecimal/ddMultiply.decTest_bigFloat
BenchmarkDecimal/ddMultiply.decTest_bigFloat-16      	       3	 388495390 ns/op	    2970 B/op	       2 allocs/op
BenchmarkDecimal/ddAbs.decTest_float64
BenchmarkDecimal/ddAbs.decTest_float64-16            	      48	  24127490 ns/op	       6 B/op	       0 allocs/op
BenchmarkDecimal/ddAbs.decTest_bigFloat
BenchmarkDecimal/ddAbs.decTest_bigFloat-16           	       3	 380678454 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecimal/ddAbs.decTest_anzDecimal
BenchmarkDecimal/ddAbs.decTest_anzDecimal-16         	       1	1600477426 ns/op	   17272 B/op	      39 allocs/op
BenchmarkDecimal/ddDivide.decTest_anzDecimal
BenchmarkDecimal/ddDivide.decTest_anzDecimal-16      	       1	2918094149 ns/op	   17272 B/op	      39 allocs/op
BenchmarkDecimal/ddDivide.decTest_float64
BenchmarkDecimal/ddDivide.decTest_float64-16         	      37	  39845562 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecimal/ddDivide.decTest_bigFloat
BenchmarkDecimal/ddDivide.decTest_bigFloat-16        	       3	 353384387 ns/op	   11472 B/op	     717 allocs/op
PASS
ok  	github.com/joshcarp/go-benchmarks/decimal	23.092s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
