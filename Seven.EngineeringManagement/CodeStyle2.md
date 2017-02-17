

### 7.2 代码风格
>“代码必须是本着写给人阅读的原则来编写，只不过顺便给机器执行而已。”

这段话来自《计算机程序设计与解释》，很精练的说明了代码风格的作用。代码风格，是一个与人相关、与机器无关的问题。代码风格的好坏，不影响编译器的工作，但是影响团队协同，影响代码的复用、演进以及缺陷修复。

Go语言很可能是第一个将代码风格强制统一的语言。一些对于其他语言的编译器完全忽视的问题，在Go语言编译器前就会被认为是编译错误，比如如果花括号新起一行摆放，就会看到一个醒目的编译错误。Go语言的这种做法简化了问题。

接下来介绍Go语言的编码规范，主要分两类：
* 由Go编译器进行强制的编码规范
* 由**Gotool**推行的非强制性编码风格建议。

其他的一些编码规范里会列出的细节，比如应该用Tab还是4个空格，这些不在本书的讨论范围之内。

#### 7.2.1 强制性编码规范
可以认为，由Go编译器进行强制的编码规范也是Go语言设计者认为最需要统一的代码风格，下面一一诠释。
**1.命名**
命名规则涉及**变量**、**常量**、**全局函数**、**结构**、**接口**、**方法**等的命名。Go语言从语法层面进行了以下限定：**任何需要对外暴露的名字必须以大写字母开头，不需要对外暴露的则应该以小写字母开头。**
软件开发行业最流行的两种命名法分别为：
* [驼峰命名法DoSomething和doSomething](https://zh.wikipedia.org/wiki/駝峰式大小寫)
* [下划线法do_something](http://wenku.baidu.com/link?url=9CTGWjNNZ2BC01wRnbagf_bFV-gJnTf8iY9kW96N1JuCt7gyYy6hSR6Vnmm1OZ9rlTjIsDoAdAD5WU7bAR0W6C1wLXbaX0kJLoACu331Mke)

而Go语言明确宣告了拥护驼峰命名法而排斥下划线命名法。驼峰命名法在`Java`和`C#`中得到了官方的支持和推荐，而下划线命名法则主要用在`C`语言里，比如Linux内核和驱动开发上。在开始Go语言编程时，还是忘记下划线命名法吧，避免写出不伦不类的名字

**2.排列**

Go语言甚至对代码的排列方式也进行了语法级别的检查，约定了代码块中花括号的明确摆放位置。下面先列出一个错误的写法：

[示例代码-错误版](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/code/ChapterSeven/7.2CodeStyle/WrongWriting.go)

这个写法对于众多在微软怀抱里长大的程序员们时最熟悉不过的了，但是在Go语言中会有编译错误

```bash
$ go build test.go
# command-line-arguments
./test.go:8: syntax error: unexpected semicolon or newline before {
```
通过上面的错误信息就能猜到，是左花括号`{`的位置出问题了。下面我们将上面的代码调整一下

[示例代码-正确版](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/code/ChapterSeven/7.2CodeStyle/CorrectWriting.go)

可以看到，`else`甚至都必须紧跟在之间的右花括号`}`后面并且不能换行。Go语言的这条规则基本上就保证了所有Go代码的逻辑结构写法是完全一致的，也不会再出现有洁癖的程序员在维护别人代码之前非要把所有花括号的位置都调整一遍的问题。

#### 7.2.2 非强制性编码规范
**Gotool**中包含了一个代码格式化的功能，这也是一般语言都无法想象的事情。下面来看看格式化工具的用法：
```bash
$ go help fmt
usage: go fmt [-n] [-x] [packages]

Fmt runs the command 'gofmt -l -w' on the packages named
by the import paths.  It prints the names of the files that are modified.

For more about gofmt, see 'go doc cmd/gofmt'.
For more about specifying packages, see 'go help packages'.

The -n flag prints commands that would be executed.
The -x flag prints commands as they are executed.

To run gofmt with specific options, run gofmt itself.

See also: go fix, go vet.
```

可以看出，用法非常简单。接下来试验一下它的格式化效果。
先看看故意制造的比较丑陋的代码
```go
package main

import "fmt"
func Foo(a, b int) (ret int, err error) {
if a > b {
return a,nil
	} else {
return b,nil
	}
return 0, nil
}

func 
main(){i,_:=Foo(1,2)
fmt.Println("Hello, 世界",i)}
```
由于IDE保存文件就会自动格式化，未能提供文件。
这段代码能够正常编译，也能正常运行，只不过丑陋的代码让人看不下去。现在我们用**Gotool**中的格式化功能美化一下(假设上述代码被保存为hello.go)
```bash
$ go fmt hello.go
hello.go
```
执行这个命令后，将会更新到`hello1.go`文件，此时再打开`hello1.go`看一下旧貌换新的代码
```go
func Foo(a, b int) (ret int, err error) {
	if a > b {
		return a, nil
	} else {
		return b, nil
	}
	return 0, nil
}

func main() {
	i, _ := Foo(1, 2)
	fmt.Println("Hello, 世界", i)
}
```
可以看到，格式化工具做了很多事情
* 调整了每条语句的位置
* 重新摆放花括号的位置
* 用制表符缩紧代码
* 添加空格

当然，格式化工具不知道怎么帮你改进命名，这就不要苛求了。
这个工具并非只能一次格式化一个文件，比如不带任何参数直接运行`go fmt`的话，就可以直接格式化当前目录下的所有`*.go`文件，或者也可以指定一个`GOPATH`中可以找到的报名。
只不过我们并不是非常推荐使用这个工具。毕竟，保持良好的编码风格应该是一开始写代码时就注意到的。第一次就写成符合规范的样子，以后也就不用再考虑如何美化的问题了。