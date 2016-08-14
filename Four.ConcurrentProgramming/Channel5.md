### 4.5 channel
>channel是Go语言在**语言级别**提供的goroutine间的通信方式。channel是**进程内的通信方式**，因此通过channel传递对象的过程和调用函数时的参数传递行为比较一致，比如也可以传递指针等。channel是**类型相关的**，也就是说，一个channel只能传递一种类型的值，这个类型需要在声明channel时指定。

```go
package main

import "fmt"

func Count(ch chan int, index int) {
	fmt.Println("counting")
	ch <- index
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i], i)
	}

	for _, ch := range chs {
		fmt.Println(<-ch)

	}
}
```
在这个例子中，我们定义了一个包含10个channel的数组（名为chs），并把数组中的每个channel分配给10个不同的goroutine。在每个goroutine的打印函数完成后，通过`ch <- index`语句向对应的channel中写入一个数据。在这个channel被读取前，这个操作时阻塞的。

在所有的goroutine启动完成后，通过<-ch语句从10个channel中依次读取数据。在对应的channel写入数据前，这个操作也是阻塞的。

这样，就用channel实现了类似锁的功能，进而保证了所有goroutine完成后主函数才返回。

###4.5.1 基本语法
一般channel的声明形式为：
```go
var chanName chan ElementType
```
与一般的变量声明不同的地方仅仅是在类型之前加了chan关键字。ElementType指定这个channel所能传递的元素类型。
```go
//声明一个传递类型为int的channel
var ch chan int 	
//声明一个map，元素是bool型的channel
var m map[string] chan bool 
//声明并初始化一个int型的channel，通过内置函数make
ch1 := make(chan int) 

//写入---将一个数据写入（发送）至channel
ch <- value
//读取
value := <- ch
```
向channel写入数据通常会导致程序阻塞，直到有其他goroutine从这个channel中读取数据。

### 4.5.2 select
>Go语言直接在语言级别支持select关键字，用于处理异步IO问题。通过调用select()函数来监控一系列的文件句柄，一旦其中一个文件句柄发生了IO动作，该select()调用就会被返回。后来该机制也被用于实现高并发的Socket服务器程序。

>select的用法与switch相似，但每个case语句里必须是一个channel操作。

```go
select {
	case ch <- 2:
		fmt.Println("向ch写入数据")
	default:
		fmt.Println("channel is full !")
	}
```

###4.5.3  缓冲机制
对于需要传递大量数据的场景，需要为channel带上缓冲，而从达到消息队列的效果。
```go
c := make(chan int,1024)
```
上面的例子创建了一个大小为1024的int类型channel，即使没有读取方，写入方也可以一直往channel里写入，在缓冲区被填完之前都不会阻塞。

循环读取：
```go
for i := range c {
	fmt.Println("Received:", i)
}
```

###4.5.4 超时机制
在并发编程的通信过程中，最需要处理的就是超时问题，即向channel写数据时发现channel已满，活着从channel试图读取数据时发现channel为空。如果不正确处理这些情况，很可能会导致整个goroutine锁死。

```go
i := <-ch
```

永远都没有往ch中写数据，那么上述这个读取动作也将永远无法从ch中读取到数据，导致的结果就是整个goroutine永远阻塞并没有挽救的机会。

>Go语言没有提供直接的超时处理机制，但可以利用select机制。虽然select机制不是专为超时而设计的，却能很方便的解决超时问题。因为select的特点是只要其中一个case已经完成，程序就会继续往下执行，而不回考虑其他case的情况。

```go
package main

import (
	"fmt"
	"time"
	)

func main() {
	//首先，实现并执行一个匿名的超时等待函数
	timeout := make(chan bool, 1)
	go func(){
		time.Sleep(1e9) //等待一秒钟
		timeout <- true
	}()
	
	//然后把timeout这个channel利用起来
	select{
		case <- ch:
		fmt.Println("从ch中读取到数据")
		case <- timeout:
		fmt.Println("没有从ch中读取到数据，但从timeout中读取到了数据")
	}
}
```

这样使用select机制可以避免永久等待的问题，因为程序会在timeout中获取到一个数据后继续执行，无论对ch的读取是否还处于等待状态，从而达成1秒超时的效果。

###4.5.5 channel的传递

在Go语言中channel本身也是一个原生类型，因此channel本身在定义后也可以通过channel来传递。可以使用这个特性在*nix上非常常见的管道（pipe）特性。

管道也是使用非常广泛的一种设计模式，比如在处理数据时，我们可以采用管道设计，这样可以比较容易以插件的方式增加数据的处理流程。

假设在管道中传递的数据只是一个整型数，在实际的应用场景中这通常是一个数据块。

限定基本的数据结构，再写一个常规的处理函数。只要定义一系列PipeData的数据结构并一起传递给这个函数，就可以达到流式处理数据的目的。

```go
type PipeData struct {
	value   int
	handler func(int) int
	next    chan int
}

func handle(queue chan *PipeData) {
	for data := range queue {
		data.next <- data.handler(data.value)
	}
}
```

### 4.5.6 单向channel
单向channel只能用户发送或者接受数据。channel本身是支持读写的。单向channel只是对channel的一种使用限制。
```go
var ch1 chan int       //双向，支持读写
var ch2 chan<- float64 //单向，只支持写float64数据
var ch3 <-chan int     //单向，只支持读取int数据

ch4 := make(chan int)
ch5 := <-chan int(ch4) //类型转换，ch5是单向的读取channel
ch6 := chan<- int(ch4) //类型转换，ch6是单向的写入channel
```

从设计的角度考虑，所有的代码应该都遵循“最小权限原则”，从而避免没必要的使用滥用问题，进而导致程序失控。

单向channel的用法
```go
func Parse(ch <-chan int) {
	for value := range ch {
		fmt.Println("Parsing value", value)
	}
}
```
这个函数不回因为各种原因而对ch进行写，避免在ch中出现非期望的数据，从而很好的实践最小权限原则。

### 4.5.7 关闭channel
```go
close(ch)
```

判断是否已经被关闭
```go
x, ok := <-ch //ok为false则表示已经被关闭
```