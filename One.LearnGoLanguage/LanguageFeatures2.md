### 1.2 语言特性
* 自动垃圾回收
* 更丰富的内置类型
* 函数多返回值
* 错误处理
* 匿名函数和闭包
* 类型和接口
* 并发编程
* 反射
* 语言交互性
	
#### 1.2.1 自动垃圾回收
   >手动管理内存缺点:
* 内存泄漏(程序所占用的内存会一直疯长,直至占用所有系统内存并导致程序崩溃，而如果泄漏的是系统资源的话,最终有可能导致系统崩溃)
* 由于指针的到处传递而无法确定何时可以释放该指针所指向的内存块.(如果释放了内存,其它地方还在使用这块内存的指针,那么这些指针就变成"野指针					          wild pointer"或者"悬空指针danling pointer",对这些指针进行的任何读写操作都会导致不可预料的后果)

>由于效率杰出,Apache、Nginx和MySQL用C和C++开发。
>内存检查工具:Rational Purify、Compuware BoundsChecker和因特尔的Parallel Inspector等
		
>从设计角度也衍生了类似于内存引用计数之类的方法("智能指针"),后续在Windows平台上标准化的COM出现的一个重要原因就是为了解决内存管理的难题
		
>到目前位置,内存泄漏的最佳解决方案是在语言级别引入自动垃圾回收算法(Garbage Collection,GC)
>
>所谓垃圾回收,即所有的内存分配动作都会被运行时记录,同时任何对该内存的使用也都会被记录,然后垃圾回收器也会对所有已经分配的内存惊醒跟踪监测,
		一旦发现有些内存已经不再被任何人使用,就阶段性的回收这些没人用的内存.
	
#### 1.2.2 更丰富的内置类型	
* 简单内置类型:数组、字符串
* 较新语言中内置的高级类型(Java、C#):数组、字符串
* 对于其它静态类型语言通常用库方式支持的:字典类型(map)   既然绝大多数开发者都需要用到这个类型,为什么还非要每一个人都写一行import语句来包含一个库
* 新增数据类型:数组切片(Slice),可以认为数组切片是一种可动态增长的数组.(类似于C++中的vector)
	
#### 1.2.3 函数多返回值
	目前主流语言中除Python外基本都不支持函数的多返回值功能
```go
	function getName(firstName,middleName,lastName,nickName string){
		return "May","M","Chen","Babe"
	}
```
	没有被明确赋值的返回值将保持默认的空值
```go
		fn,mn,ln,nn := getName()
```
	可以直接用下划线作为占位符来忽略其它不关心的返回值.
```go
		_,_,ln,nn := getName()
```
	
#### 1.2.4 错误处理
	* 引入defer关键字用于标准的错误处理流程,并提供了内置函数`panic`、`recover`完成异常的抛出与捕获(详情见第二章)
	* 大量减少代码量,无需仅仅为了程序安全性而添加大量嵌套`try-catch`语句(详情见2.6节)
		
#### 1.2.5 匿名函数和闭包
	在Go语言中,所有的函数也是值类型，可以作为参数传递.支持常规的匿名函数和闭包.
	


``` go
		f := func(x,y int) int {
			return x+y
		}
```
	
#### 1.2.6 类型和接口
	Go语言的类型定义非常接近于C语言中的机构(strut),直接沿用了`struct`关键字,不支持即成和重载,而只是支持了最基本的类型组合功能.(详情见第三章)
	"非侵入式"接口
```go
type Bired struct{
		...
	}
	func (b *Bird) Fly(){
		//以鸟的方式飞行
	}
	type IFly interface{
		Fly{}
	}
	func main(){
		var fly IFly =new(Bird)
		fly.Fly()
	}
```

	这种比较松散的对应关系可以大幅降低因为接口调整而导致的大量代码调整工作
#### 1.2.7 并发编程
通过使用`goroutine`而不是裸用操作系统的并发机制，以及使用消息传递来共享内存而不是使用共享内存来通信,更加轻盈和安全

通过在函数调用前使用关键字`go`,我们即可让该函数以`goroutine`方式执行.
	`goroutine`是一种比线程更佳轻盈、更省资源的`协程`.
>Go语言通过系统的线程来多路派遣这些函数的执行,使得每个用`go`关键字执行的函数可以运行成为一个单位协程.当一个协程阻塞的时候，调度起就会自动把其他协程安排到另外的线程中去执行,从而实现了程序无等待并行化运行.而且调度的开销非常小,一颗CPU调度的规模不下于每秒百万次，这使得能够创建大量的`goroutine`,从而可以很轻松的编写高并发程序.

>Go语言实现了CSP(通信顺序进程,Communicating Sequential Process)模型来作为`goroutine`间的推荐通信方式,在CSP模型中,一个并发系统由若干并行运行的顺序进程组成,每个进程不能对其他进行的变量赋值.进程间职能通过一对通信原语实现协作.Go语言用`channel(通道)`这个概念来轻巧地实现了CSP模型.

`channel(通道)`的使用方式比较接近UNIX系统中管道(pipe)的概念,可以方便地进行跨`goroutine`的通信

>由于一个进程内创建的所有goroutine运行在同一个内存地址空间中,因此如果不通的`goroutine`不得不去访问共享的内存变量，访问前应该先获得相应的读写锁.Go语言标准库中的`sync`包提供了完备的读写锁功能。

```go
package main

import "fmt"

func sum(values [] int,resultChan chan int){
	sum := 0
	for _, value := range values{
		sum += value
	}
	resultChan <- sum  //将计算结果发送到channel中
}

func main (){
	values := [] int{1,2,3,4,5,6,7,8,9,10}

	resultChan := make(chan int,2)
	go sum(values[:len(values)/2],resultChan)
	go sum(values[len(values)/2:],resultChan)
	sum1,sum2 := <-resultChan, <- resultChan //接收结果

	fmt.Println("Result:",sum1,sum2,sum1+sum2)
}
```

#### 1.2.8 反射
通过反射(reflection)可以获取对象类型的详细信息，并可动态操作对象。虽功能强大单代码可读性并不理想。若非必要，不推荐使用反射.**详情请见第九章.**

>Go语言的反射实现了反射的大部分功能，但没有内置类型工厂，故而无法做到像Java那样通过类型字符串创建对象实例。Go语言不推荐通过读取配置并根据类型名称创建对应的类型.

反射最常见的使用使用场景是做对象的序列化(serialization,有时候也叫做Marshal &Unmarshal).

>Go语言标准库的encoding/json、encoding/xml、encoding/gob、encoding/binary等包就大量依赖于反射功能来实现

```go
package main

import(
	"fmt"
	"reflect"
)

type Bird struct{
	Name string
	LifeExpectance int
}

func (b *Bird) Fly(){
	fmt.Println("I am flying...")
}

func main(){
	sparrow := &Bird("Sparrow",3)
	s:=reflect.ValueOf(sparrow).Elem()
	typeOfT := s.Type()
	for i:=0;i<s.NumField();i++{
		f:=s.Field(i)
		fmt.Printf("%d: %s %s = %v \n",i,typeOfT.Field(i).Name,f.Type(),f.Interface())
	}
}
```

#### 1.2.9 语言交互性
由于Go语言与C语言之间的天生联系,Go语言的设计者自然不会忽略如何重用现有C模块的问题，这个功能直接被命名为Cgo.Cgo即是语言特性,同时也是一个工具的名称.**详情请见第九章**

在Go代码中,可以按Cgo的特定语法混合编写C语言代码,然后Cgo工具可以将这些混合C代码提取并生成对于C功能的调用包装代码.开发者基本上可以完全忽略这个Go语言和C语言的边界是如何跨越的.

与Java中的JNI不通,Cgo的用法非常简单,下面代码演示如何在Go中调用C语言标准库的puts函数.

```go
package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "c"
import "unsage"


func main(){
	cstr:=C.CString("Hello,world")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))
}

```

	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
		