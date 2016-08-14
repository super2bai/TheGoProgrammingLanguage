### 4.3 goroutine
>goroutine是Go语言中的轻量级线程实现，由Go运行时(runtime)管理。让函数并发执行，只需在调用函数前加`go`关键字，这次调用就会在一个新的goroutine中并发执行。当被调用的函数返回时，这个goroutine也就自动结束了。**如果这个函数有返回值，那么这个返回值回被丢弃**。

```go
package main

import "fmt"

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	for i := 0; i < 10; i++ {
		go Add(1, 1)
	}
}
```

之所以被丢弃，涉及到Go语言的程序执行机制：

>Go程序从初始化main package冰执行main()函数开始，当main ()函数返回时，程序退出，且程序并不等待其他goroutine(非主goroutine)结束。

对于上面的例子，主函数启动了10个goroutine，然后返回，这时程序就退出了，而被启动的执行Add(1, 1)的gorouine没有来得及执行，所以程序没有任何输出。要让主函数等待所有goroutine退出后再返回，涉及到goroutine之间通信的问题。见下一节。