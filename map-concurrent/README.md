# map-concurrent
```bash
=== RUN   TestConcurrentRead
--- PASS: TestConcurrentRead (0.00s)
=== RUN   TestConcurrentWrite
fatal error: concurrent map writes

goroutine 1048 [running]:
runtime.throw(0x1147154, 0x15)
	/usr/local/go/src/runtime/panic.go:1117 +0x72 fp=0xc00034af30 sp=0xc00034af00 pc=0x10363b2
runtime.mapassign_faststr(0x11211c0, 0xc00007c300, 0x1143b38, 0x3, 0xc000134088)
	/usr/local/go/src/runtime/map_faststr.go:211 +0x3f1 fp=0xc00034af98 sp=0xc00034af30 pc=0x1013691
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc00007c300, 0xc000014220)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/map_vs_switch_test.go:27 +0x50 fp=0xc00034afd0 sp=0xc00034af98 pc=0x110a730
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1371 +0x1 fp=0xc00034afd8 sp=0xc00034afd0 pc=0x106c161
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/map_vs_switch_test.go:26 +0xea

goroutine 1 [chan receive]:
testing.(*T).Run(0xc000370180, 0x11467db, 0x13, 0x114d698, 0x107f201)
	/usr/local/go/src/testing/testing.go:1240 +0x2da
testing.runTests.func1(0xc000001380)
	/usr/local/go/src/testing/testing.go:1512 +0x78
testing.tRunner(0xc000001380, 0xc00011bde0)
	/usr/local/go/src/testing/testing.go:1194 +0xef
testing.runTests(0xc00000c048, 0x12174a0, 0x2, 0x2, 0xc036fd598fa44d90, 0x8bb2d62d2d, 0x121ec20, 0x1145587)
	/usr/local/go/src/testing/testing.go:1510 +0x2fe
testing.(*M).Run(0xc000010180, 0x0)
	/usr/local/go/src/testing/testing.go:1418 +0x1eb
main.main()
	_testmain.go:45 +0x138

goroutine 4 [sleep]:
time.Sleep(0x5f5e100)
	/usr/local/go/src/runtime/time.go:193 +0xd2
runtime/pprof.profileWriter(0x116ee68, 0xc00000e030)
	/usr/local/go/src/runtime/pprof/pprof.go:799 +0x69
created by runtime/pprof.StartCPUProfile
	/usr/local/go/src/runtime/pprof/pprof.go:784 +0x11f

goroutine 1022 [runnable]:
sync.(*WaitGroup).Add(0xc000014220, 0x1)
	/usr/local/go/src/sync/waitgroup.go:53 +0x14d
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite(0xc000370180)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/map_vs_switch_test.go:25 +0xbe
testing.tRunner(0xc000370180, 0x114d698)
	/usr/local/go/src/testing/testing.go:1194 +0xef
created by testing.(*T).Run
	/usr/local/go/src/testing/testing.go:1239 +0x2b3

goroutine 1122 [runnable]:
github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc00007c300, 0xc000014220)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/map_vs_switch_test.go:26
created by github.com/joshcarp/benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/map_vs_switch_test.go:26 +0xea
exit status 2
FAIL	github.com/joshcarp/benchmarks/map-concurrent	0.228s
FAIL
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
