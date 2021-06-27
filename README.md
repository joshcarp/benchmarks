# map vs switch

## map
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/map-vs-switch
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkOptionMap
BenchmarkOptionMap-16    	10679650	       112.1 ns/op
PASS

## switch
goos: darwin
goarch: amd64
pkg: github.com/joshcarp/map-vs-switch
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkOptionSwitch
BenchmarkOptionSwitch-16    	286718242	         3.990 ns/op
PASS