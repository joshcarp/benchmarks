# decimal
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/go-benchmarks/decimal
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkDecimal
BenchmarkDecimal/ddAdd.decTest_bigFloat
BenchmarkDecimal/ddAdd.decTest_bigFloat-16         	     121	   9920913 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecimal/ddAdd.decTest_anzDecimal
BenchmarkDecimal/ddAdd.decTest_anzDecimal-16       	       7	 172572843 ns/op	    2366 B/op	       0 allocs/op
BenchmarkDecimal/ddAdd.decTest_float64
BenchmarkDecimal/ddAdd.decTest_float64-16          	    1431	   1016364 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecimal/ddMultiply.decTest_anzDecimal
BenchmarkDecimal/ddMultiply.decTest_anzDecimal-16  	       7	 195734910 ns/op	    2371 B/op	      37 allocs/op
BenchmarkDecimal/ddMultiply.decTest_float64
BenchmarkDecimal/ddMultiply.decTest_float64-16     	    1194	   1403251 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecimal/ddMultiply.decTest_bigFloat
BenchmarkDecimal/ddMultiply.decTest_bigFloat-16    	     100	  13342424 ns/op	      81 B/op	       0 allocs/op
BenchmarkDecimal/ddAbs.decTest_anzDecimal
BenchmarkDecimal/ddAbs.decTest_anzDecimal-16       	       7	 187770945 ns/op	     888 B/op	      37 allocs/op
BenchmarkDecimal/ddAbs.decTest_float64
BenchmarkDecimal/ddAbs.decTest_float64-16          	    1057	   1304641 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecimal/ddAbs.decTest_bigFloat
BenchmarkDecimal/ddAbs.decTest_bigFloat-16         	     100	  12924507 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecimal/ddDivide.decTest_anzDecimal
BenchmarkDecimal/ddDivide.decTest_anzDecimal-16    	       2	 766553410 ns/op	   15732 B/op	      42 allocs/op
BenchmarkDecimal/ddDivide.decTest_float64
BenchmarkDecimal/ddDivide.decTest_float64-16       	     728	   1818841 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecimal/ddDivide.decTest_bigFloat
BenchmarkDecimal/ddDivide.decTest_bigFloat-16      	      90	  13111409 ns/op	   11472 B/op	     717 allocs/op
PASS
ok  	github.com/joshcarp/go-benchmarks/decimal	20.137s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
