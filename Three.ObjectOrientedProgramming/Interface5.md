###3.5 接口
Go语言的主要设计者之一Rob Pike曾经说过，如果智能选择一个Go语言的特性移植到其他语言，他会选择接口.
>在Go语言中，一个类只需要实现了接口要求的所有函数，就认为这个类实现了该接口。

###3.5.1 侵入式接口
> “侵入式”的主要表现在于实现类需要明确自己实现了某个接口。这种强制性的接口继承是面向对象编程思想发展过程中一个遭受相当多质疑的特性。正是因为这种不合理的设计，实现每个类时都需要纠结以下两个问题：
*  提供哪些接口好呢
*  如果两个类实现了相同的接口，应该把接口放到哪个包好呢

###3.5.2非侵入式接口
```go

type File struct {
	//...
}

func (f *File) Read(buf []byte) (n int, err error) {
	return 0, nil
}
func (f *File) Write(buf []byte) (n int, err error) {
	return 0, nil
}
func (f *File) Seek(off int64, whence int) (pos int64, err error) {
	return 0, nil
}
func (f *File) Close() error {
	return nil
}

type IFile interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Seek(off int64, whence int) (pos int64, err error)
	Close() error
}

type IReader interface {
	Read(buf []byte) (n int, err error)
}
type IWrite interface {
	Write(buf []byte) (n int, err error)
}
type ISeek interface {
	Seek(off int64, whence int) (pos int64, err error)
}
type IClose interface {
	Close() error
}

func main() {
	//	var file1 IFile = new(File)
	//	var file2 IReader = new(File)
	//	var file3 IWrite = new(File)
	//	var file4 IClose = new(File)
}
```

>Go语言的非侵入式接口，优点：
*  Go语言的标准库，再也不需要绘制类库的继承树图。在Go语言中，类的继承树并无意义，只需要知道这个类实现了哪些方法，每个方法的含义就足够了。
*  实现类的时候，只需要关心自己应该提供哪些方法，不用再纠结接口需要拆的多细才合理。接口由使用方按需求定义，而不用事先规划。
*  不用为了实现一个接口而导入一个包，因为多引用一个外部的包，就意味着更多的耦合。接口由使用方按自身需求来定义，无需关心是否有其他模块定义过类似的接口。

###3.5.3 接口赋值
>接口赋值在Go语言中分为如下两种情况：
* 将对象实例赋值给接口，这要求该对象实例实现了接口要求的所有方法。
```go
package main

import "fmt"

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	c := *a + b
	fmt.Print(c)
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

func main() {
	var a Integer = 1
	var b1 LessAdder = &a //OK
	var b2 LessAdder = a  //not OK
	fmt.Print(b.Less(1))
}
```

Go语言根据下面的函数

```go
func (a Integer) Less(b Integer) bool {
	return a < b
}

```
自动生成一个新的Less()方法

```go
func (a *Integer) Less(b Integer) bool {
	return (*a).Less(b)
}
```
这样，类型`*Integer`就既存在Less()方法，也存在Add()方法，满足LessAdder接口。而另一方面来说，根据`func (a *Integer) Add(b Integer) `这个函数无法自动生成
```go
func (a Integer) Add(b Integer) bool {
	return (&a).Add(b)
}
```
因为`(&a).Add()`改变的只是函数参数a，对外部实际要操作的对象并无影响，这不符合用户的预期。所以，Go语言不回自动为其生成该函数。因此，类型`Integer`只存在`Less()`方法，却少`Add()`方法，不满足`LessAdder`接口。

>Go语言规范：**The method set of any other named type T consists of all methods with receiver type T. The method set of the corresponding pointer type T is the set of all methods with receiver T or T (that is, it also contains the method set of T).**

也就是说`*Integer`实现了接口`LessAdder`的所有方法，而`Integer`只实现了`Less()`，所以不能赋值。

* 将一个接口赋值给另一个接口，在Go语言中，只要两个接口拥有相同的方法列表（次序不同不要紧），那么它们就是等同的，可以相互赋值；如果接口A的方法列表是接口B的方法列表的子集，那么接口B可以赋值给接口A，反之不能。

###3.5.4 接口查询
>在Go语言中，可以询问接口它指向的对象是否是某个类型。

```go
var file1 Writer = ...
if file6, ok := file1.(*File); ok {
	//...
}
```
判断file1接口指向的对象实例是否是*File 类型，如果是则执行特定代码。
>查询接口所指向的对象是否为某个类型的这种用法可以认为只是接口查询的一个特例。接口是对一组类型的公共特性的抽象，所以查询接口查的是一个群体，查询具体类型，是查询具体的个体。

###3.5.5 类型查询
>在Go语言中，还可以更加直截了当的询问接口指向的对象实例的类型

```go
var v1 interface{} = ...
switch v:= v1.(type) {
	case int:
	case string: 
	...
}
```
语言中的类型多的数不清，所以类型查询并部经常使用，它更多是个补充，需要配合接口查询使用。

```go
type Stringer interface {
	String() string
}

func Println(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
		case string:
		default:
			if v, ok := arg.(Stringer); ok {
				val := v.String()
				fmt.Println(val)
			} else {

			}
		}
	}
}
```

Go语言标准库的Println()比这个例子要复杂很多，这里只摘取其中关键的部分进行分析。对于内置类型，Println()采用穷举法，将每个类型转换为字符串进行打印。对于更一版的情况，首先确定该类型是否实现了String()方法，如果实现了，则用String()方法将其转换为字符串进行打印。否则，Println()利用反射功能来遍历对象的所有成员变量进行打印

利用反射也可以进行类型查询，详情可参阅reflect.TypeOf()方法的相关文档，9.1节中，也会探讨反射。

###3.5.6 接口组合
>可以认为接口组合是类型匿名组合的一个特定场景，只不过接口只包含方法，而不包含任何成员变量。

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}
type ReadWriter1 interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}
```
ReadWriter组合了Reader和Writer两个接口，它完全等同于ReadWriter1。因为这两种写法的表意完全相同，ReadWriter和ReadWriter1既能做Reader接口的所有事情，又能做Writer接口的所有事情。在Go语言包中，还有众多类似的组合接口，比如ReadWriteCloser、ReadWriteSeeker、ReadSeeker和WriteClose等

###3.5.7 Any类型
>由于Go语言中任何对象实例都满足空接口interface{}，所以interface{}看起来像是可以指向人和对象的Any类型。

```go
var v1 interface{} = 1
var v2 interface{} = "abc"
var v3 interface{} = &v2
var v4 interface{} = struct{ X int }{1}
var v5 interface{} = &struct{ X int }{1}
```
当函数可以接受任意的对象实例时，会将其声明为interface{},最典型的例子是标准库fmt中PrintXXX系类的函数，详情见标准库源码。