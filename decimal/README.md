# decimal
```bash
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/benchmarks/decimal
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkDecimal
BenchmarkDecimal/ddAdd.decTest_anzDecimal
BenchmarkDecimal/ddAdd.decTest_anzDecimal-16         	       7	 163037731 ns/op
BenchmarkDecimal/ddAdd.decTest_float64
BenchmarkDecimal/ddAdd.decTest_float64-16            	    1334	   1038707 ns/op
BenchmarkDecimal/ddAdd.decTest_bigFloat
BenchmarkDecimal/ddAdd.decTest_bigFloat-16           	     129	   9572412 ns/op
BenchmarkDecimal/ddMultiply.decTest_anzDecimal
BenchmarkDecimal/ddMultiply.decTest_anzDecimal-16    	       8	 157448534 ns/op
BenchmarkDecimal/ddMultiply.decTest_float64
BenchmarkDecimal/ddMultiply.decTest_float64-16       	    1018	   1216165 ns/op
BenchmarkDecimal/ddMultiply.decTest_bigFloat
BenchmarkDecimal/ddMultiply.decTest_bigFloat-16      	     100	  11775678 ns/op
BenchmarkDecimal/ddAbs.decTest_anzDecimal
BenchmarkDecimal/ddAbs.decTest_anzDecimal-16         	       8	 163025933 ns/op
BenchmarkDecimal/ddAbs.decTest_float64
BenchmarkDecimal/ddAbs.decTest_float64-16            	     980	   1232110 ns/op
BenchmarkDecimal/ddAbs.decTest_bigFloat
BenchmarkDecimal/ddAbs.decTest_bigFloat-16           	     100	  10669494 ns/op
BenchmarkDecimal/ddDivide.decTest_anzDecimal
BenchmarkDecimal/ddDivide.decTest_anzDecimal-16      	       2	 675620078 ns/op
BenchmarkDecimal/ddDivide.decTest_float64
BenchmarkDecimal/ddDivide.decTest_float64-16         	     716	   1733190 ns/op
BenchmarkDecimal/ddDivide.decTest_bigFloat
BenchmarkDecimal/ddDivide.decTest_bigFloat-16        	     100	  13022178 ns/op
PASS
ok  	github.com/joshcarp/benchmarks/decimal	19.311s
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
