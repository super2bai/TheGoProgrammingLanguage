###3.3 匿名组合
>Go语言也提供了继承，但是采用了组合的文法，所以将其称为匿名组合

```go
type Base struct {
	Name string
}
type Foo struct {
	Base
}

func (base *Base) Foo() {}
func (base *Base) Bar() {}
func (foo *Foo) bar() {
	foo.Base.Bar()
}
```

>以上代码定义了一个Base类（实现了Foo()和Bar()两个成员方法），然后定义了一个Foo类，该类从Base类“继承”并改写了Bar()方法（该方法实现时先调用了基类的Bar()方法）。
>在“派生类”Foo没有改写“基类”Base的成员方法时，相应的方法就被“继承”，例如在上面的例子中，调用foo.Foo()和调用foo.Base.Bar()效果一致。

在Go语言官方网站提供的Effective Go中曾提到匿名组合的一个小价值。它匿名组合了一个log.Logger指针
```go
package main

import (
	"log"
	"os"
)

type Job struct {
	Command string
	*log.Logger
}

func NewJob(command string, logger *log.Logger) *Job {
	return &Job{command, logger}
}

func (job *Job) Start() {
	job.Println("Starting now....")
	job.Println("end")
}

func main() {
	mylog := log.New(os.Stdout, "", log.LstdFlags)
	j := NewJob("test", mylog)
	j.Start()
}
```
>对于Job的实现者来说，甚至根本就不用意识到log.Logger类型的存在，这就是匿名组合的美丽所在。

>需要注意的是，被组合的类型所包含的方法虽然都升级成了外部这个组合类型的方法，但其实它们被组合方法调用时接受者并没有改变。上例中即使组合后调用的方式变成了`job.Println()`，但`Println()`的接受者仍然是`log.Logger`指针，因此在Println()中不可能访问到job的其他成员方法和变量。

>毕竟被组合的类型并不知道自己会被什么类型组合，当然就没法在实现方法时去使用那个未知的“组合者”的功能了。

---
**命名冲突**
```go
type X struct {
	name string
}
type Y struct {
	X
	name string
}
```
name并不会冲突，所有的Y类型的name成员的访问都只会访问到最外层的那个name变量，X.name变量相当于被隐藏起来了。

```go
type Logger struct {
	Level int
}
type Y struct {
	*Logger
	Name string
	*log.Logger
}
```
会冲突。编译会报错：`duplicate field Logger`
>匿名组合类型相当于以其类型名称（去掉包名部分）作为成员变量的名字。按此规则，Y类型中就相当于两个名为Logger的成员，虽然类型不同。