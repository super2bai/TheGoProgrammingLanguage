### 4.7 出让时间片
可以在每个goroutine中控制何时主动出让时间片给其他goroutine，这可以使用`runtime`包中的Gosched()函数实现。

如果要比较精细的控制goroutine的行为，就必须比较深入的了解Go语言开发包中runtime包所提供的具体功能。

```go
package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
```

输出结果：
hello
world
hello

注意结果：
1、先输出了hello,后输出了world.
2、hello输出了2个，world输出了1个（因为第2个hello输出完，主线程就退出了，第2个world没机会了）

把代码中的runtime.Gosched()注释掉，执行结果是：
hello
hello

**因为say("hello")这句占用了时间，等它执行完，线程也结束了，say("world")就没有机会了。**

这里同时可以看出，go中的goroutins并不是同时在运行。事实上，如果没有在代码中通过runtime.GOMAXPROCS(n) 其中n是整数，指定使用多核的话，goroutins都是在一个线程里的，它们之间通过不停的让出时间片轮流运行，达到类似同时运行的效果。 

还需要学习一句话：

**当一个goroutine发生阻塞，Go会自动地把与该goroutine处于同一系统线程的其他goroutines转移到另一个系统线程上去，以使这些goroutines不阻塞**