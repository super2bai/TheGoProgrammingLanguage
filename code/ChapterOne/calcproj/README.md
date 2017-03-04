假设有这样一个场景：需要开发一个基于命令行的计算器程序。下面为此程序的基本用法
```bash
$ calc help
USAGE: calc command [arguments] ...

The commands are:
	add		Addition of two values.
	sqrt	Square root of a non-negative value.

$ calc sqrt 4 # 开根号
2
$ calc add 1 2 # 加法
3	
```
我们假设这个工程被分割为两部分：
* 可执行程序，名为calc，内部只有一个calc.go文件
* 算法库，名为simplemath,每个command对应于一个同名的go文件，比如add.go

则一个正常的工程目录组织应该如下所示:

* calcproj
	* src
		* calc
			* calc.go
		* simplemath
			* add.go
			* add_test.go
			* sqrt.go
			* sqrt_test.go
	* bin
	* pkg(包将会安装到此处)		

在上面的结构里，无后缀名的为目录。
xxx_test.go表示的是一个对于xxx.go的单元测试，这也是Go工程里的命名规则。

为了让读者能够动手实践，这里我们会列出所有源代码并以注释的方式解释关键内容。
需要注意的是，本示例主要用于示范功能管理，并不保证代码达到产品级质量。

-------------------------------------------------------------
为了能够构建这个工程，需要先把这个工程的根目录加入到环境变量`GOPATH`中。假设calcproj目录位于`~/goyard`下，则应编辑`~/.bashrc`文件，并添加下面这行代码：
```bash
export GOPATH=~/goyard/calcproj
```
然后执行以下命令应用该设置:
```bash
$ source ~/.bashrc
```
`GOPATH`和`PATH`环境变量一样，也可以接受多个路径，并且路径和路径之间用冒号分割。
设置完`GOPATH`后，现在开始构建工程。假设希望把生成的可执行文件放到`calcproj/bin`目录中，需要执行的一系列执行如下：
```bash
$ cd ~/goyard/calcproj
$ mkdir bin
$ cd bin
$ go build calc
```
顺利的话，将在该目录下发现生成的一个叫做calc的可执行文件，执行该文件以查看帮助信息并进行算数运算：
```bash
$ ./calc
USAGE: calc command [arguments] ...

The commands are:
	add		Addition of two values.
	sqrt	Square root of a non-negative value.

$ ./calc sqrt 9
Result: 3
$ ./calc add 2 3
Result: 5	
```
从上面的构建过程可以看到，真正的构建命令就一句:
```go
go build calc
```
这就是为什么说Go命令行工具是非常强大的。不需要写makefile，因为这个工具会分析，知道目标代码的编译结果应该是一个包还是一个可执行文件，并分析import语句以了解包的依赖关系，从而在编译calc.go之前先把依赖的simplemath编译打包好。Go命令行程序指定的目录结构规则让代码管理变得非常简单。

另外，在写simplemath包时，为每一个关键的函数编写了对应的单元测试代码，分别位于add_test.go和sqrt_test.go中。那么我们到底怎么运行这些单元测试呢？这也非常简单。因为已经设置了GOPATH,所以可以在任意目录下执行以下命令：
```bash
$ go test simplemath
ok simplemath0.014s
```
可以看到，运行结果列出了测试内容、测试结果和测试实践。如果故意把add_test.go的代码改成这样的错误场景：
```go
func TestAdd1(t *testing.T) {
	r := Add(1, 2)
	if r != 2 {//这里本该是3，故意改成2测试错误场景
		t.Errorf("Add(1, 2) failed. Got %d, expected 3.", r)
	}
}
```
然后再次执行单元测试，将得到如下的结果：
```bash
$ go test simplemath
=== RUN   TestAdd1
--- FAIL: TestAdd1 (0.00s)
	add_test.go:8: Add(1, 2) failed. Got 3, expected 3.
```
打印的错误信息非常简洁，却已经足够让开发者快速定位到问题代码所在的文件和行数，从而在最短的时间内确认是单元测试的问题还是程序的问题。