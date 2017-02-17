### 7.1 Go命令行工具
Go作为一门新语言，除了语言本身的特性在高速发展外，其相关的工具链也在逐步完善。任何一门程序设计语言要能推广开来病投入到生产环境中，高效、易用、有好的开发环境都是必不可少的。

Go这个名字在本书内已经被用过两次：
* 本语言的名字：Go语言
* 并发编程：go关键字(实现了最核心的goroutine功能，使其成为高并发服务的语言)

在这里又成了一个Go语言包自带的命令行工具的名字。从这个名字我们可以了解到，这个工具在Go语言设计者心目中的重要地位。
为了避免名字上的困扰，在本章中我们将用**Gotool**来程序这个工具。

**基本用法**
在安装了Go语言的安装包后，就直接自带**Gotool**。可以运行`go version`来查看**Gotool**的版本，也就是当前安装的Go语言的版本：
```bash
$ go version
go version go1.7.5 darwin/amd64
```
**Gotool**的功能非常强大，可以运行`go help`查看它的功能说明：
```bash
$ go help
Go is a tool for managing Go source code.

Usage:

	go command [arguments]

The commands are:

	build       compile packages and dependencies
	clean       remove object files
	doc         show documentation for package or symbol
	env         print Go environment information
	fix         run go tool fix on packages
	fmt         run gofmt on package sources
	generate    generate Go files by processing source
	get         download and install packages and dependencies
	install     compile and install packages and dependencies
	list        list packages
	run         compile and run Go program
	test        test packages
	tool        run specified go tool
	version     print Go version
	vet         run go tool vet on packages

Use "go help [command]" for more information about a command.

Additional help topics:

	c           calling between Go and C
	buildmode   description of build modes
	filetype    file types
	gopath      GOPATH environment variable
	environment environment variables
	importpath  import path syntax
	packages    description of package lists
	testflag    description of testing flags
	testfunc    description of testing functions

Use "go help [topic]" for more information about that topic.
```
简而言之，**Gotool**可以完成以下这几类工作：
* 代码格式化
* 代码质量分析和修复
* 单元测试与性能测试
* 工程构建
* 代码文档的提取和展示
* 依赖包管理
* 执行其他的包含指令，比如`6g`等
我们会在随后的章节中一一展开**Gotool**这些用法，目前只需要记住这个工具的用法就能够开始专业的开发了。

