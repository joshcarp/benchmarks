# map-concurrent
```bash
=== RUN   TestConcurrentRead
--- PASS: TestConcurrentRead (0.00s)
=== RUN   TestConcurrentWrite
fatal error: concurrent map writes
fatal error: concurrent map writes
fatal error: concurrent map writes

goroutine 1326 [running]:
runtime.throw(0x1147154, 0x15)
	/usr/local/go/src/runtime/panic.go:1117 +0x72 fp=0xc000407f30 sp=0xc000407f00 pc=0x10363b2
runtime.mapassign_faststr(0x11211c0, 0xc00049a000, 0x1143b38, 0x3, 0xc00049c088)
	/usr/local/go/src/runtime/map_faststr.go:291 +0x3d8 fp=0xc000407f98 sp=0xc000407f30 pc=0x1013678
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc00049a000, 0xc00048c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:27 +0x50 fp=0xc000407fd0 sp=0xc000407f98 pc=0x110a730
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1371 +0x1 fp=0xc000407fd8 sp=0xc000407fd0 pc=0x106c161
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1 [chan receive]:
testing.(*T).Run(0xc000482300, 0x11467db, 0x13, 0x114d698, 0x107f201)
	/usr/local/go/src/testing/testing.go:1240 +0x2da
testing.runTests.func1(0xc000082600)
	/usr/local/go/src/testing/testing.go:1512 +0x78
testing.tRunner(0xc000082600, 0xc0000bdde0)
	/usr/local/go/src/testing/testing.go:1194 +0xef
testing.runTests(0xc0000ac030, 0x12174a0, 0x2, 0x2, 0xc03707170e3e3d80, 0x8bb2da5d10, 0x121ec20, 0x1145587)
	/usr/local/go/src/testing/testing.go:1510 +0x2fe
testing.(*M).Run(0xc0000d2000, 0x0)
	/usr/local/go/src/testing/testing.go:1418 +0x1eb
main.main()
	_testmain.go:45 +0x138

goroutine 18 [sleep]:
time.Sleep(0x5f5e100)
	/usr/local/go/src/runtime/time.go:193 +0xd2
runtime/pprof.profileWriter(0x116ee68, 0xc0000b0020)
	/usr/local/go/src/runtime/pprof/pprof.go:799 +0x69
created by runtime/pprof.StartCPUProfile
	/usr/local/go/src/runtime/pprof/pprof.go:784 +0x11f

goroutine 1025 [runnable]:
sync.(*WaitGroup).Add(0xc00048c030, 0x1)
	/usr/local/go/src/sync/waitgroup.go:53 +0x14d
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite(0xc000482300)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:25 +0xbe
testing.tRunner(0xc000482300, 0x114d698)
	/usr/local/go/src/testing/testing.go:1194 +0xef
created by testing.(*T).Run
	/usr/local/go/src/testing/testing.go:1239 +0x2b3

goroutine 1412 [running]:
	goroutine running on other thread; stack unavailable
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1520 [running]:
	goroutine running on other thread; stack unavailable
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1526 [runnable]:
sync.(*WaitGroup).Add(0xc00048c030, 0xffffffffffffffff)
	/usr/local/go/src/sync/waitgroup.go:53 +0x14d
sync.(*WaitGroup).Done(...)
	/usr/local/go/src/sync/waitgroup.go:99
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc00049a000, 0xc00048c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:28 +0x87
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1527 [runnable]:
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc00049a000, 0xc00048c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1528 [runnable]:
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc00049a000, 0xc00048c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1529 [runnable]:
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc00049a000, 0xc00048c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1530 [runnable]:
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc00049a000, 0xc00048c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1531 [runnable]:
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc00049a000, 0xc00048c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1532 [runnable]:
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc00049a000, 0xc00048c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1533 [runnable]:
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc00049a000, 0xc00048c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1535 [runnable]:
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc00049a000, 0xc00048c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1536 [runnable]:
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc00049a000, 0xc00048c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1537 [runnable]:
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc00049a000, 0xc00048c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1538 [runnable]:
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc00049a000, 0xc00048c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1539 [runnable]:
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc00049a000, 0xc00048c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1540 [runnable]:
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc00049a000, 0xc00048c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1541 [runnable]:
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc00049a000, 0xc00048c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1542 [runnable]:
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc00049a000, 0xc00048c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1543 [runnable]:
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc00049a000, 0xc00048c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1544 [runnable]:
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc00049a000, 0xc00048c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1412 [running]:
runtime.throw(0x1147154, 0x15)
	/usr/local/go/src/runtime/panic.go:1117 +0x72 fp=0xc000330f30 sp=0xc000330f00 pc=0x10363b2
runtime.mapassign_faststr(0x11211c0, 0xc00049a000, 0x1143b38, 0x3, 0xc00049c088)
	/usr/local/go/src/runtime/map_faststr.go:211 +0x3f1 fp=0xc000330f98 sp=0xc000330f30 pc=0x1013691
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc00049a000, 0xc00048c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:27 +0x50 fp=0xc000330fd0 sp=0xc000330f98 pc=0x110a730
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1371 +0x1 fp=0xc000330fd8 sp=0xc000330fd0 pc=0x106c161
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1520 [running]:
runtime.throw(0x1147154, 0x15)
	/usr/local/go/src/runtime/panic.go:1117 +0x72 fp=0xc0004a8730 sp=0xc0004a8700 pc=0x10363b2
runtime.mapassign_faststr(0x11211c0, 0xc00049a000, 0x1143b38, 0x3, 0x0)
	/usr/local/go/src/runtime/map_faststr.go:211 +0x3f1 fp=0xc0004a8798 sp=0xc0004a8730 pc=0x1013691
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc00049a000, 0xc00048c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:27 +0x50 fp=0xc0004a87d0 sp=0xc0004a8798 pc=0x110a730
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1371 +0x1 fp=0xc0004a87d8 sp=0xc0004a87d0 pc=0x106c161
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea
exit status 2
FAIL	github.com/joshcarp/benchmarks/map-concurrent	0.046s
FAIL
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
