# map-concurrent
```bash
=== RUN   TestConcurrentRead
--- PASS: TestConcurrentRead (0.05s)
=== RUN   TestConcurrentWrite
==================
WARNING: DATA RACE
Write at 0x00c000256088 by goroutine 76:
  github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1()
      /Users/carpeggj/Documents/work/go-benchmarks/map-concurrent/test_test.go:27 +0x5c

Previous write at 0x00c000256088 by goroutine 77:
  github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1()
      /Users/carpeggj/Documents/work/go-benchmarks/map-concurrent/test_test.go:27 +0x5c

Goroutine 76 (running) created at:
  github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite()
      /Users/carpeggj/Documents/work/go-benchmarks/map-concurrent/test_test.go:26 +0xc4
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1259 +0x22f
  testing.(*T).Run·dwrap·21()
      /usr/local/go/src/testing/testing.go:1306 +0x47

Goroutine 77 (finished) created at:
  github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite()
      /Users/carpeggj/Documents/work/go-benchmarks/map-concurrent/test_test.go:26 +0xc4
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1259 +0x22f
  testing.(*T).Run·dwrap·21()
      /usr/local/go/src/testing/testing.go:1306 +0x47
==================
==================
WARNING: DATA RACE
Write at 0x00c00022e0c0 by goroutine 65:
  runtime.mapassign_faststr()
      /usr/local/go/src/runtime/map_faststr.go:202 +0x0
  github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1()
      /Users/carpeggj/Documents/work/go-benchmarks/map-concurrent/test_test.go:27 +0x50

Previous write at 0x00c00022e0c0 by goroutine 77:
  runtime.mapassign_faststr()
      /usr/local/go/src/runtime/map_faststr.go:202 +0x0
  github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1()
      /Users/carpeggj/Documents/work/go-benchmarks/map-concurrent/test_test.go:27 +0x50

Goroutine 65 (running) created at:
  github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite()
      /Users/carpeggj/Documents/work/go-benchmarks/map-concurrent/test_test.go:26 +0xc4
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1259 +0x22f
  testing.(*T).Run·dwrap·21()
      /usr/local/go/src/testing/testing.go:1306 +0x47

Goroutine 77 (finished) created at:
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
FAIL	github.com/joshcarp/go-benchmarks/map-concurrent	0.377s
FAIL
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
