# map-concurrent
```bash
=== RUN   TestConcurrentRead
--- PASS: TestConcurrentRead (0.00s)
=== RUN   TestConcurrentWrite
fatal error: concurrent map writes
fatal error: concurrent map writes
fatal error: concurrent map writes

goroutine 1123 [running]:
runtime.throw(0x1147154, 0x15)
	/usr/local/go/src/runtime/panic.go:1117 +0x72 fp=0xc00017b730 sp=0xc00017b700 pc=0x10363b2
runtime.mapassign_faststr(0x11211c0, 0xc0000aa000, 0x1143b38, 0x3, 0xc0000ac088)
	/usr/local/go/src/runtime/map_faststr.go:291 +0x3d8 fp=0xc00017b798 sp=0xc00017b730 pc=0x1013678
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:27 +0x50 fp=0xc00017b7d0 sp=0xc00017b798 pc=0x110a730
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1371 +0x1 fp=0xc00017b7d8 sp=0xc00017b7d0 pc=0x106c161
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1 [chan receive]:
testing.(*T).Run(0xc000082900, 0x11467db, 0x13, 0x114d698, 0x107f201)
	/usr/local/go/src/testing/testing.go:1240 +0x2da
testing.runTests.func1(0xc000128300)
	/usr/local/go/src/testing/testing.go:1512 +0x78
testing.tRunner(0xc000128300, 0xc000133de0)
	/usr/local/go/src/testing/testing.go:1194 +0xef
testing.runTests(0xc000120030, 0x12174a0, 0x2, 0x2, 0xc037076766b24fa0, 0x8bb2dae33e, 0x121ec20, 0x1145587)
	/usr/local/go/src/testing/testing.go:1510 +0x2fe
testing.(*M).Run(0xc000148000, 0x0)
	/usr/local/go/src/testing/testing.go:1418 +0x1eb
main.main()
	_testmain.go:45 +0x138

goroutine 34 [sleep]:
time.Sleep(0x5f5e100)
	/usr/local/go/src/runtime/time.go:193 +0xd2
runtime/pprof.profileWriter(0x116ee68, 0xc000124020)
	/usr/local/go/src/runtime/pprof/pprof.go:799 +0x69
created by runtime/pprof.StartCPUProfile
	/usr/local/go/src/runtime/pprof/pprof.go:784 +0x11f

goroutine 19 [runnable]:
sync.(*WaitGroup).Add(0xc00009c030, 0x1)
	/usr/local/go/src/sync/waitgroup.go:53 +0x14d
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite(0xc000082900)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:25 +0xbe
testing.tRunner(0xc000082900, 0x114d698)
	/usr/local/go/src/testing/testing.go:1194 +0xef
created by testing.(*T).Run
	/usr/local/go/src/testing/testing.go:1239 +0x2b3

goroutine 1218 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1239 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1224 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1223 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1245 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1256 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1222 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1251 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1221 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1220 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1219 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1194 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1216 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1235 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1215 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1217 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1233 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1190 [runnable]:
runtime.goexit1()
	/usr/local/go/src/runtime/proc.go:3365 +0x6c
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1372 +0x6
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1246 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1234 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1250 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1255 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1238 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1247 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1240 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1254 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1243 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1248 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1244 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1261 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1258 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1184 [running]:
	goroutine running on other thread; stack unavailable
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1252 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1191 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1193 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1242 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1253 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1236 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1260 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1259 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1241 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1249 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1257 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1192 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1225 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1139 [running]:
	goroutine running on other thread; stack unavailable
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1213 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1212 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1211 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1210 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1209 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1208 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1207 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1206 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1214 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1204 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1203 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1202 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1201 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1237 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1199 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1200 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1195 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1226 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1227 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1228 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1229 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1230 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1231 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1232 [runnable]:
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1139 [running]:
runtime.throw(0x1147154, 0x15)
	/usr/local/go/src/runtime/panic.go:1117 +0x72 fp=0xc0000b4f30 sp=0xc0000b4f00 pc=0x10363b2
runtime.mapassign_faststr(0x11211c0, 0xc0000aa000, 0x1143b38, 0x3, 0x0)
	/usr/local/go/src/runtime/map_faststr.go:211 +0x3f1 fp=0xc0000b4f98 sp=0xc0000b4f30 pc=0x1013691
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:27 +0x50 fp=0xc0000b4fd0 sp=0xc0000b4f98 pc=0x110a730
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1371 +0x1 fp=0xc0000b4fd8 sp=0xc0000b4fd0 pc=0x106c161
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea

goroutine 1184 [running]:
runtime.throw(0x1147154, 0x15)
	/usr/local/go/src/runtime/panic.go:1117 +0x72 fp=0xc000375730 sp=0xc000375700 pc=0x10363b2
runtime.mapassign_faststr(0x11211c0, 0xc0000aa000, 0x1143b38, 0x3, 0xc00015c088)
	/usr/local/go/src/runtime/map_faststr.go:211 +0x3f1 fp=0xc000375798 sp=0xc000375730 pc=0x1013691
github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite.func1(0xc0000aa000, 0xc00009c030)
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:27 +0x50 fp=0xc0003757d0 sp=0xc000375798 pc=0x110a730
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1371 +0x1 fp=0xc0003757d8 sp=0xc0003757d0 pc=0x106c161
created by github.com/joshcarp/go-benchmarks/map-concurrent.TestConcurrentWrite
	/Users/carpeggj/Documents/work/benchmarks/map-concurrent/test_test.go:26 +0xea
exit status 2
FAIL	github.com/joshcarp/go-benchmarks/map-concurrent	0.168s
FAIL
```
## Memory profile
![](mem.svg)
## CPU profile
![](cpu.svg)
