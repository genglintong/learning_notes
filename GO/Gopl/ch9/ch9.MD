- sync.Once 
```
var loadIconsOnce sync.Once
var icons map[string]image.Image
// Concurrency-safe.
func Icon(name string) image.Image {
    loadIconsOnce.Do(loadIcons)
    return icons[name]
}
```
- 竞争检查器
 - 使用竞争检查器
 只要在go build，go run或者go test命令后面加上-race的flag，就会使编译器创建一个你的应用的“修改”版或者一个附带了能够记录所有运行期对共享变量访问工具的test，并且会记录下每一个读或者写共享变量的goroutine的身份信息

```text
=== RUN   TestConcurrent
https://godoc.org, 839.724308ms, 7165 bytes
https://golang.org, 843.868488ms, 11075 bytes
==================
WARNING: DATA RACE
Write at 0x00c000096db0 by goroutine 11:
  runtime.mapassign_faststr()
      /usr/local/go/src/runtime/map_faststr.go:202 +0x0
  gopl.io/ch9/memo1.(*Memo).Get()
      /Users/genglintong/go/goPath/src/gopl.io/ch9/memo1/memo.go:35 +0x1ce
  gopl.io/ch9/memotest.Concurrent.func1()
      /Users/genglintong/go/goPath/src/gopl.io/ch9/memotest/memotest.go:93 +0xde

Previous write at 0x00c000096db0 by goroutine 10:
  runtime.mapassign_faststr()
      /usr/local/go/src/runtime/map_faststr.go:202 +0x0
  gopl.io/ch9/memo1.(*Memo).Get()
      /Users/genglintong/go/goPath/src/gopl.io/ch9/memo1/memo.go:35 +0x1ce
  gopl.io/ch9/memotest.Concurrent.func1()
      /Users/genglintong/go/goPath/src/gopl.io/ch9/memotest/memotest.go:93 +0xde

Goroutine 11 (running) created at:
  gopl.io/ch9/memotest.Concurrent()
      /Users/genglintong/go/goPath/src/gopl.io/ch9/memotest/memotest.go:90 +0x10c
  command-line-arguments_test.TestConcurrent()
      /Users/genglintong/lintong/OpenSource/learning_notes/GO/Gopl/ch9/memo_test.go:18 +0xd9
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:909 +0x199

Goroutine 10 (finished) created at:
  gopl.io/ch9/memotest.Concurrent()
      /Users/genglintong/go/goPath/src/gopl.io/ch9/memotest/memotest.go:90 +0x10c
  command-line-arguments_test.TestConcurrent()
      /Users/genglintong/lintong/OpenSource/learning_notes/GO/Gopl/ch9/memo_test.go:18 +0xd9
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:909 +0x199
==================
https://godoc.org, 846.398423ms, 7165 bytes
https://golang.org, 886.238346ms, 11075 bytes
https://play.golang.org, 1.32524031s, 6013 bytes
https://play.golang.org, 1.337596857s, 6013 bytes
http://gopl.io, 1.718271603s, 4154 bytes
http://gopl.io, 1.724362812s, 4154 bytes
--- FAIL: TestConcurrent (1.73s)
    testing.go:853: race detected during execution of test
FAIL
FAIL    command-line-arguments  2.575s
FAIL

```

- Goroutines和线程
    - 9.8.1. 动态栈 
        1. OS 线程 栈大小为 2M 固定大小
        2. goroutine 栈大小最小为 2K 可动态扩展