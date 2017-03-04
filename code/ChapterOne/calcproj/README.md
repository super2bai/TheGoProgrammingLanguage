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
