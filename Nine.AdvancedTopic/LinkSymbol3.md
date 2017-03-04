### 9.3 链接符号
链接符号关心的是如何将语言文法使用的符号转化为链接期使用的符号，在常规情况下，链接期使用的符号不可见，但是在一些特殊情况下，需要关心这一点，比如：在用`gdb`调试的时候，要设置断点：`b<函数名>`，这里的`<函数名>`是指"链接符号"，而非平常看到的语言文法层面使用的符号。

例如，在C语言中，一般的函数原型如下：
```C
RetType Method(ArgType1 arg1, ArgType2 arg2,...)
```
这里`Method`是C语言文法层面使用的符号，但其"链接符号"为`_Method`，而不是`Method`。

又如在C++中，一般化的函数原型如下：
```C++
RetType Method(ArgType1 arg1, ArgType2 arg2,...)
RetType Namespace::Method(ArgType1 arg1, ArgType2 arg2,...)
// 名字空间下的方法，名字空间可以有多层，如 A::B::C
RetType Namespace::ClassType::Method(ArgType1 arg1, ArgType2 arg2,...)
//类成员方法
```
由于C++支持函数虫灾，故此语言的"链接符号"构成及其复杂，需要包括：
* `Namespace`
* `ClassType`
* `Method`
* `ArgType1 arg1, ArgType2 arg2,...`

因此一般情况下，C++的"链接符号"都非常长。另外，不同编译器厂商生成"链接符号"的规则并不统一，这是C++语言很大的问题。缺乏二进制级别的交互标准，意味着不同厂商生成的二进制模块(`.o`或`.so`)是不兼容的。因此多数情况下，C++语言的模块间交互使用C的机制，而不是自身的机制。

在Go语言中，一般化的函数原型如下：
```go
package Package
func Method(arg1 argType1, arg2 argType2,...)(ret1 RetType1, ret2 RetType2,...)
func (v ClassType)Method(arg1 argType1, arg2 argType2,...)(ret1 RetType1, ret2 RetType2,...)
func (this *ClassType)Method(arg1 argType1, arg2 argType2,...)(ret1 RetType1, ret2 RetType2,...)//这种可以认为是上一种情况的特例
```

由于Go语言并无重载，故此语言的"链接符号"由如下信息构成：
* `Package`。`Package`名可以是多层，例如A/B/C
* `ClassType`。很特别的是，Go语言中`ClassType`可以是指针，也可以不是
* `Method`

其"链接符号"的组成规则如下：
* `Package.Method`
* `Package.ClassType.Method`

这样说可能比较抽象，下面再举个实际的例子。
假设在`qbox.us/mockfs`模块中，有如下几个函数：
* `func New(cfg Config) *MockFS`
* `func (fs *MockFS) Mkdir(dir string) (code int,err error)`
* `func (fs MockFS) Foo(bar Bar)`

它们的链接符号分别为：
* `qbox.us/mockfs.New`
* `qbox.us/mockfs.*MockFS.Mkdir`
* `qbox.us/mockfs.MockFS.Foo`