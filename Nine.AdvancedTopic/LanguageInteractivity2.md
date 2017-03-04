### 9.2 语言交互性
自C语言诞生以来，程序员们已经积累了无数代码库。即使后面还出现了众多时髦的新语言，有无数的代码库都还很偏执的只提供了C语言版本。因此，如何快捷方便的直接引用这些功能强大的C语言库，就成了所有现代语言都不得不重视的话题。比如，像Java这样非常重视面向对象身份的语言也都提供了JNI机制，以调用那些C代码库。

作为一门直接传承于C的语言，Go当然应该将与C语言的交互作为首要任务之一。Go的确也提供了这一功能，称为`Cgo`。

下面直接用一个来源与Go语言官方博客的例子来开始`Cgo`之旅。对于程序员来说，一段明了的[源代码](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/code/ChapterNine/9.2LanguageInteractivity/cgo1.go)可以比几页文字更好的说明问题。

这个例子的运行方法与第一章的`Hello world`示例没有区别，直接使用`go run`命令即可。
以上这个例子的整个逻辑看起来似乎很简单：导入了一个名为`C`的包，然后在函数中使用了`C`包包含的`random()`和`srandom()`函数，顺便还用了一个`C`包中提供的`uint`类型。

初看起来确实没有问题，但再仔细想一下，马上就会蹦出很多疑问来。
* Go语言标准库里的包名字都是小写的，这个名字大写的`C`包怎么看都不像是Go自带的，但也没有装过这个包，它到底是哪里来的呢?
* 为什么要在`import`前面写上那么奇怪的一段完全就是C语法的注释?这段注释是必需的吗？
* 不是说包内类型的可见性是由首个字母的大小写决定的吗？为什么这里能够使用`C`包里以小写字母开头的函数和类型呢？

如果能够提出以上这些问题，说明确定已经比较熟悉Go语言的语法了，如果还看不出来任何问题的话，建议抽空再复习一下本书前面的内容。不管如何，先继续`Cgo`之旅。

事实上，根本就不存在一个名为`C`的包。这个`import`语句其实就是一个信号，告诉`Cgo`它应该开始工作了。就是对这条`import`语句之前的块注释中的C源代码自动生成包装性质的Go代码。如果对以下这些概念有所了解，就相对比较容易理解`Cgo`这个声称Go代码的过程:
* Java的JNI
* .NET的P/Invoke
* RPC
* WebService的Proxy/Stub
* IDL语言
* WSDL语言等

不了解以上这些概念也没关系，因为函数调用从汇编的角度看，就是一个将参数按顺序压栈(push)，然后进行函数调用(call)的过程。`Cgo`生成的代码只不过是封装了这个压栈河调用的过程，从外面看起来就是一个普通的Go函数调用。
这时候该注意到`import`语句前紧跟的注释了。这个注释的内容是由意义的，而不是传统意义上的注释作用。这个例子例用的是一个块注释，实际上行注释也是没问题的，只要是紧贴在`import`语句之前即可。比如下面也是正确的`Cgo`写法：
```go
// #include <stdio.h>
// #include <stdlib.h>
import "C"
```

### 9.2.1 类型映射
在跨语言交互中，比较复杂的问题由两个：
* 类型映射
* 跨越调用边界传递指针所带来的对象声明周期和内存管理的问题

比如Go语言中的`string`类型需要跟C语言中的字符串组进行对应，并且要保证映射到C语言的对象的声明周期足够长，以避免在C语言执行过程中该对象就已经被垃圾回收。这一节先谈类型映射的问题。

对于C语言的原生类型，`Cgo`都会将其映射为Go语言中的类型
* `C.char`和`C.schar`->`signed char`
* `C.uchar`->`unsigned char`
* `C.short`和`C.ushort`->`unsigned short`
* `C.int`和`C.uint`->`unsigned int`
* `C.long`和`C.ulong`->`unsigned long`
* `C.longlong`->`long long`
* `C.ulonglong`->`unsigned long long`
* `C.float`和`C.double`
* `unsage.Pointer`->`void*指针类型`
* `struct_`->`struct`(Go: C.struct_person -> C:person struct )
* `union_`->`union`
* `enum_`->`enum`

如果C语言中的类型名称或变量名称与Go语言的关键字相同,`Cgo`会自动给这些名字加上下划线前缀。

#### 9.2.2 字符串映射
因为Go语言中有明确的`string`原生类型，而C语言中用字符串组表示，两者之间的转换时一个必须考虑的问题。`Cgo`提供了一系列函数来提供支持：
* `C.CString`
* `C.GoString`
* `C.GoStringN`
需要注意的是，每次转换都将导致一次内存复制，因此字符串内容其实是不可修改的(实际上，Go语言的`string`也不允许对其中的内容进行修改)。

由于`C.CString`的内存管理方式与Go语言自身的内存管理方式不兼容，设法期待Go语言可以做垃圾回收，因此在使用完后必须显式释放调用`C.CString`所生成的内存块，否则将导致严重的内存泄漏。结合之前已经学过的`defer`用法，所用到的`C.CString`的代码大致都可以写成如下的风格：
```go
var gostr string
//...初始化gostr
cstr := C.CString(gostr)
defer C.free(unsafe.Pointer(cstr))
//接下来大胆的使用cstr吧，因为保证可以被释放掉了
//C.sprintf(cstr, "content is: %d",123)
```

#### 9.2.3 C程序
在9.2节开头的示例中，块注释中只写了一条`include`语句，其实在这个块注释中，可以写任意合法的C源代码，而`Cgo`都会进行相应的处理并生成对应的Go代码。[稍微复杂一些的例子](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/code/ChapterNine/9.2LanguageInteractivity/cgo2.go)
这个块注释里就直接写了个C函数，它使用C标准库里的`printf()`打印了一句话。
还有另外一个问题，那就是如果这里的C代码需要依赖一个非C标准库的第三方库，怎么办呢？如果不解决的话必然会有连接时错误。`Cgo`提供了`#cgo`这样的伪C文法，让开发者有机会指定依赖的第三方库和编译选项。
下面的例子示范了`#cgo`的第一种用法：
```go
// #cgo CFLAGS: -DPNG_DEBUG=1
// #cgo linux CFLAGS: -DLINUX=1
// #cgo ldflags: -lpng
// #include <png.h>
import "C"
```
这个例子示范了如何使用`CFLAGS`来传入编译选项，使用`CFLAGS`来传入链接选项。`#cgo`还有另外一种更简单一些的用法，如下所示：
```go
// #cgo pkg-config: png cairo
// #include <png.h>
import "C"
```

#### 9.2.4 函数调用
对于常规的函数调用，开发者只要在运行`cgo`指令后查看以下生成的Go代码，就可以知道如何写对应的调用代码。有一点比较贴心的是，对于常规返回了一个值的函数，调用者可以用以下的方式顺便得到错误码：
```go
n, err := C.atoi("a234")
```
在传递数组类型的参数时需要注意，在Go语言中将第一个元素的地址作为整个数组的起始地址传入，这一点就不如C语言本身直接传入数组名字那么方便了。下面为一个传递数组的例子：
```go
n, err := C.f(&array[0])//需要显式指定第一个元素的地址
```

#### 9.2.5 编译Cgo
编译`Cgo`代码非常容易，不需要做任何特殊的处理。Go安装后，会自带一个`cgo`命令行工具，它用于处理所有带有`Cgo`代码的Go文件，生成Go语言版本的调用封装代码。而Go工具集对`cgo`命令行工具进行了良好的封装，使构建过程能够自动识别和处理带有`Cgo`代码的Go源代码文件，完全不给用户增加额外的工作负担。