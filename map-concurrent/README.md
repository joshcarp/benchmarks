# map-concurrent
```bash
=== RUN   TestConcurrentRead
--- PASS: TestConcurrentRead (0.05s)
=== RUN   TestConcurrentWrite
==================
WARNING: DATA RACE
Write at 0x00c0004ac000 by goroutine 63:
  runtime.mapassign_faststr()
      /usr/local/go/src/runtime/map_faststr.go:202 +0x0
  github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1()
      /Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:27 +0x5d

Previous write at 0x00c0004ac000 by goroutine 67:
  runtime.mapassign_faststr()
      /usr/local/go/src/runtime/map_faststr.go:202 +0x0
  github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1()
      /Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:27 +0x5d

Goroutine 63 (running) created at:
  github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite()
      /Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0x130
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1194 +0x202

Goroutine 67 (finished) created at:
  github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite()
      /Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0x130
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1194 +0x202
==================
==================
WARNING: DATA RACE
Write at 0x00c0004ae088 by goroutine 68:
  github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1()
      /Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:27 +0x72

Previous write at 0x00c0004ae088 by goroutine 67:
  github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1()
      /Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:27 +0x72

Goroutine 68 (running) created at:
  github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite()
      /Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0x130
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1194 +0x202

Goroutine 67 (finished) created at:
  github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite()
      /Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0x130
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1194 +0x202
==================
fatal error: concurrent map writes

goroutine 1304 [running]:
runtime.throw(0x11fb470, 0x15)
	/usr/local/go/src/runtime/panic.go:1117 +0x72 fp=0xc000369728 sp=0xc0003696f8 pc=0x1077b92
runtime.mapassign_faststr(0x11cee40, 0xc0004ac000, 0x11f7ddd, 0x3, 0x0)
	/usr/local/go/src/runtime/map_faststr.go:211 +0x431 fp=0xc000369790 sp=0xc000369728 pc=0x10547f1
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0004ac000, 0xc00049e030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:27 +0x5e fp=0xc0003697d0 sp=0xc000369790 pc=0x11b749e
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1371 +0x1 fp=0xc0003697d8 sp=0xc0003697d0 pc=0x10af961
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0x131

goroutine 1 [chan receive]:
testing.(*T).Run(0xc00008ac00, 0x11faaf7, 0x13, 0x1201b80, 0x1)
	/usr/local/go/src/testing/testing.go:1240 +0x610
testing.runTests.func1(0xc00008ac00)
	/usr/local/go/src/testing/testing.go:1512 +0xa7
testing.tRunner(0xc00008ac00, 0xc0000c5ce0)
	/usr/local/go/src/testing/testing.go:1194 +0x203
testing.runTests(0xc0000ca018, 0x1300f60, 0x2, 0x2, 0xc0374adf9ca02fa0, 0x8bb2e211a8, 0x13086c0, 0xc0000c5da0)
	/usr/local/go/src/testing/testing.go:1510 +0x613
testing.(*M).Run(0xc0000de000, 0x0)
	/usr/local/go/src/testing/testing.go:1418 +0x3b4
main.main()
	_testmain.go:45 +0x237

goroutine 21 [sleep]:
time.Sleep(0x5f5e100)
	/usr/local/go/src/runtime/time.go:193 +0xd2
runtime/pprof.profileWriter(0x122c430, 0xc0000ba020)
	/usr/local/go/src/runtime/pprof/pprof.go:799 +0x77
created by runtime/pprof.StartCPUProfile
	/usr/local/go/src/runtime/pprof/pprof.go:784 +0x18b

goroutine 1046 [runnable]:
sync.(*WaitGroup).Add(0xc00049e030, 0x1)
	/usr/local/go/src/sync/waitgroup.go:53 +0x2eb
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite(0xc00048b200)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:25 +0x105
testing.tRunner(0xc00048b200, 0x1201b80)
	/usr/local/go/src/testing/testing.go:1194 +0x203
created by testing.(*T).Run
	/usr/local/go/src/testing/testing.go:1239 +0x5d8

goroutine 1310 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0004ac000, 0xc00049e030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0x131
exit status 2
FAIL	github.com/joshcarp/go-benchmarks/map-concurrent	0.218s
FAIL
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
