###2.6 错误处理

###2.6.1 error接口
```go
type error interface{
		Error() string
}
```
对于大多数函数，如果要反悔错误，大致上都可以定义为将error作为多种返回值中的最后一个，但这并非是强制要求。
```go
func main() {
	n, err := Foo(0)
	if err != nil {
		//
	} else {
		//
	}
}
func Foo(param int) (n int, err error) {
	//...
}
```

**自定义error类型**
* 首先，定义一个用于承载错误信息的类型。因为Go语言中接口的灵活性，根本不需要从error接口继承。
```go
package main

import (
	"os"
)
//类型转换
func main() {
	_, err := os.Stat("a.txt")
	if err != nil {
		if e, ok := err.(*os.PathError); ok && e.Err != nil {
			//获取PathError类型变量e中的其他信息并处理
		}
	}
}

type PathError struct {
	Op   string
	Path string
	Err  error
}

//实现了Error()方法
func (e *PathError) error() string {
	return e.Op + " " + e.Path + " " + e.Err.Error()
}
```

###2.6.2 defer
>当函数无论怎样返回，某资源必须释放时，可用这种与众不同、但有效的处理方式。传统的例子包括解锁互斥或关闭文件。

>这样延迟一个函数有双重优势：一是你永远不会忘记关闭文件，此错误在你事后编辑函数添加一个返回路径时常常发生。二是关闭和打开靠在一起，比放在函数尾要清晰很多。

>defer 在声明时不会立即执行，而是在函数 return 后，再按照 FILO （先进后出）的原则依次执行每一个 defer，一般用于异常处理、释放资源、清理数据、记录日志等。这有点像面向对象语言的析构函数，优雅又简洁，是 Golang 的亮点之一。

```go
defer xxx.Close()
defer func(){
	//destroy resource code
}()
```

###2.6.3 panic()和recover()
>Go语言引入了两个内置函数panic()和recover()以报告和处理运行时错误和程序中的错误场景

```go
func panic(interface{})
func recover() interface{}
```
>当在一个函数执行过程中调用`panic()`函数时，正常的函数执行流程将立即终止，但函数中之前使用`defer`关键字延迟执行的语句将正常展开执行，之后该函数将返回到调用函数，并导致逐层向上之行`panic`流程，直至所属的`goroutine`中所有正在执行的函数被终止。错误信息将被报告，包括在调用`panic()`函数时传入的参数，这个过程被称为**错误处理流程**

```go
panic(404)
panic("network broken")
panic(Error("file not exists"))
```

>`recover()`函数用于终止错误处理流程。一般情况下，`recover()` 应该在一个试用`defer`关键字的函数中执行以有效截取错误处理流程。如果在没有发生异常的`goroutine`中明确调用恢复过程(使用recover关键字)，会导致该`goroutine`所属的进程打印异常信息后直接退出。

```go
defer func(){
	if r:=recover();!=nil{			
	}
} 
```