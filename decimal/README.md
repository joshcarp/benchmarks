# decimal
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/benchmarks/decimal
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkDecimal
BenchmarkDecimal/ddAdd.decTest_bigFloat
BenchmarkDecimal/ddAdd.decTest_bigFloat-16         	     134	   8364651 ns/op	       2 B/op	       0 allocs/op
BenchmarkDecimal/ddAdd.decTest_anzDecimal
BenchmarkDecimal/ddAdd.decTest_anzDecimal-16       	       8	 148672238 ns/op	    3077 B/op	       0 allocs/op
BenchmarkDecimal/ddAdd.decTest_float64
BenchmarkDecimal/ddAdd.decTest_float64-16          	    1328	    912010 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecimal/ddMultiply.decTest_float64
BenchmarkDecimal/ddMultiply.decTest_float64-16     	    1237	   1301861 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecimal/ddMultiply.decTest_bigFloat
BenchmarkDecimal/ddMultiply.decTest_bigFloat-16    	     100	  13719989 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecimal/ddMultiply.decTest_anzDecimal
BenchmarkDecimal/ddMultiply.decTest_anzDecimal-16  	       6	 210692686 ns/op	    3960 B/op	      37 allocs/op
BenchmarkDecimal/ddAbs.decTest_anzDecimal
BenchmarkDecimal/ddAbs.decTest_anzDecimal-16       	       7	 168830923 ns/op	     888 B/op	      37 allocs/op
BenchmarkDecimal/ddAbs.decTest_float64
BenchmarkDecimal/ddAbs.decTest_float64-16          	    1011	   1261731 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecimal/ddAbs.decTest_bigFloat
BenchmarkDecimal/ddAbs.decTest_bigFloat-16         	     100	  14177999 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecimal/ddDivide.decTest_bigFloat
BenchmarkDecimal/ddDivide.decTest_bigFloat-16      	     100	  13011196 ns/op	   11472 B/op	     717 allocs/op
BenchmarkDecimal/ddDivide.decTest_anzDecimal
BenchmarkDecimal/ddDivide.decTest_anzDecimal-16    	       2	 723714040 ns/op	   15440 B/op	      39 allocs/op
BenchmarkDecimal/ddDivide.decTest_float64
BenchmarkDecimal/ddDivide.decTest_float64-16       	     769	   1898243 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/joshcarp/benchmarks/decimal	18.996s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
