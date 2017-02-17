### 7.4 工程组织
在第一章中已经大致介绍过Go语言中约定的工程管理方式，这里将进一步解释其中的各个细节。

#### 7.4.1 GOPATH
`GOPATH`这个环境变量是在讨论工程组织之前必须提到的内容。**Gotool**的大部分功能其实已经不再针对当前目录，而是针对**包名**，于是如何才能定位到对应的源代码就落到了`GOPATH`身上。
假设现在本地硬盘上有3个Go代码工程，分别为`~/work/go-proj1`、`~/work2/goproj2`和`~/work3/work4/go-proj3`,那么`GOPATH`可以设置为如下内容：
```bash
export GOPATH=~/work/go-proj1:~/work2/goproj2:~/work3/work4/go-proj3
```
经过这样的设置后，可以在任意位置对以上3个功能进行构建。

#### 7.4.2 目录结构
以第一章介绍的`calcproj`工程为例介绍工程管理规范：
caclproj
----README
----AUTHORS
----bin
--------calc
----pkg
--------linux_amd64
------------simplemath.a
----src
--------calc
------------calc.go
--------simplemath
------------add.go
------------add_test.go
------------sqrt.go
------------sqrt_test.go
Go语言工程不需要任何工程文件，一个比较完整的工程会在根目录处放置这样几个文本文件。
* README  简单介绍本项目目标和关键的注意事项，通常第一次使用时就应该先阅读本文档
* LICENSE 本工程采用的分发协议，所有开元项目通常都有这个文件

说明文档并不是工程必需的，但如果有的话可以让使用者更快上手。另外，虽然是个文本文件，但现在其实也是可以表达富格式的。比如，使用`github.com`管理代码的开发者就可以使用`Mackdown`语法来写纯文本的文档，这样就可以显示有格式的内容。这不是本书的重点，有兴趣的读者可以查看`github.com`网站。

一个标准的Go语言工程保函以下几个目录：`src`、`pkg`、`bin`。目录`src`用于包含所有的源代码，是**Gotool**一个强制的规则，而`pkg`和`bin`则无需手动创建，如果必要**Gotool**在构建过程中会自动创建这些目录。

构建过程中**Gotool**对包结构的理解完全依赖于`src`下面的目录结构，比如对于上面的例子，**Gotool**会认为`src`下包含了两个包：
* calc
* simplemath

而且对两个包的路径都是一级的，即`simplemath`下的`*.go`文件将会被构建为一个名为`simplemath.a`的包。假如希望这个包的路径带有一个命名空间，比如在使用时希望以这样的方式导入
```go
import "myns/simplemath"
```
那么就系要将目录结构调整为如下格式：
calcproj
----README
----...
----src
--------myns
------------simplemath
----------------add.go
----------------add_test.go
----------------sqrt.go
----------------sqrt_test.go

就是在`src`下多了一级`simplemath`的父目录`myns`。这样`Gotool`就能知道该怎么管理编译后的包了，工程构建后对应的`simplemath`包的位置将会是`pkg/linux_amd64/myns/simplemath.a`。规则非常简单易懂，重要的是彻底摆脱了`Makefile`等专门为构建而写的工程文件，避免了随时同步工程文件和代码的工作量。

还看到`pkg`目录下有一个自动创建的`linux_amd64`目录，相关规则在介绍[跨平台开发](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/Seven.EngineeringManagement/Cross-platformDevelopment7.md)时会详细介绍。