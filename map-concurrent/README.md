# map-concurrent
```bash
=== RUN   TestConcurrentRead
--- PASS: TestConcurrentRead (0.04s)
=== RUN   TestConcurrentWrite
==================
WARNING: DATA RACE
Write at 0x00c0002a60c0 by goroutine 21:
  runtime.mapassign_faststr()
      /usr/local/go/src/runtime/map_faststr.go:202 +0x0
  github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1()
      /Users/carpeggj/Documents/work/go-benchmarks/map-concurrent/test_test.go:27 +0x50

Previous write at 0x00c0002a60c0 by goroutine 20:
  runtime.mapassign_faststr()
      /usr/local/go/src/runtime/map_faststr.go:202 +0x0
  github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1()
      /Users/carpeggj/Documents/work/go-benchmarks/map-concurrent/test_test.go:27 +0x50

Goroutine 21 (running) created at:
  github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite()
      /Users/carpeggj/Documents/work/go-benchmarks/map-concurrent/test_test.go:26 +0xc4
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1259 +0x22f
  testing.(*T).Run·dwrap·21()
      /usr/local/go/src/testing/testing.go:1306 +0x47

Goroutine 20 (finished) created at:
  github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite()
      /Users/carpeggj/Documents/work/go-benchmarks/map-concurrent/test_test.go:26 +0xc4
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1259 +0x22f
  testing.(*T).Run·dwrap·21()
      /usr/local/go/src/testing/testing.go:1306 +0x47
==================
==================
WARNING: DATA RACE
Write at 0x00c0002c0088 by goroutine 22:
  github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1()
      /Users/carpeggj/Documents/work/go-benchmarks/map-concurrent/test_test.go:27 +0x5c

Previous write at 0x00c0002c0088 by goroutine 20:
  github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1()
      /Users/carpeggj/Documents/work/go-benchmarks/map-concurrent/test_test.go:27 +0x5c

Goroutine 22 (running) created at:
  github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite()
      /Users/carpeggj/Documents/work/go-benchmarks/map-concurrent/test_test.go:26 +0xc4
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1259 +0x22f
  testing.(*T).Run·dwrap·21()
      /usr/local/go/src/testing/testing.go:1306 +0x47

Goroutine 20 (finished) created at:
  github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite()
      /Users/carpeggj/Documents/work/go-benchmarks/map-concurrent/test_test.go:26 +0xc4
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1259 +0x22f
  testing.(*T).Run·dwrap·21()
      /usr/local/go/src/testing/testing.go:1306 +0x47
==================
    testing.go:1152: race detected during execution of test
--- FAIL: TestConcurrentWrite (0.05s)
=== CONT  
    testing.go:1152: race detected during execution of test
FAIL
exit status 1
FAIL	github.com/joshcarp/go-benchmarks/map-concurrent	0.360s
FAIL
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
