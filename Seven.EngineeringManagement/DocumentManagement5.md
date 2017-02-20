### 7.5 文档管理
程序包括：
* 代码
* 文档。

软件产品：
* 源代码（可选）
* 可执行程序
* 文档
* 服务

可以很容易看出，在一个软件交付的过程中，程序知识其中的一个基本环节，更多的工作是告诉用户如何部署、使用和维护软件，此时文档将起到关键性的作用。

对于程序员来说，所谓的文档，更多的是指代码中的注释、函数、接口的输入、输出、功能和参数说明，这些对于后续的维护和复用有着至关重要的作用。

在传统开发中，同步设计文档和代码是一件非常困难的事情。一旦开始有了一些细微的不一致，之后这个鸿沟将越来越大，并最终导致文档完全放弃。Go语言引入的规范做的比较彻底，让开发者完全甩掉注释语法的报复，专注于内容。

[示例代码-注释](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/code/ChapterSeven/src/foo/foo.go)

在这段代码里，添加了4条注释：
* 版权说明注释
* 包说明注释
* 函数说明注释
* 遗留问题说明

**注意：**
* 将项目路径添加到$GOPATH中,此处路径为：/Users/xxx/code/TheGoProgrammingLanguage/code/ChapterSeven
* foo.go在项目中的路径为：ChapterSeven/src/foo/foo.go

提取注释并展示文档：
```bash
$ go doc foo
$ go doc foo
package foo // import "foo"

Package foo implements a set of simple mathematical functions.These comments
are for demonstration purpose only.Nothing more.

If you have any question, please don't hesitate to add yourself to
golang-nuts@googlegroups.com

You can also visit golang.org for full Go documentation.

func Foo(a, b int) (ret int, err error)

BUG: #1 : I'm sorry but this code has an issue to be solved.

BUG: #2 : An issue assigned to another person.
```

已经演示了`go doc`命令提取包中的注释内容，并将其格式输出到终端窗口中。因为工程目录已经加入到`GOPATH`中，所以这个命令可以在任意位置运行。

虽然这个输出结果比较清晰，但考虑到有时候包里面的注释量非常大，所以更合适的查看方式是在浏览器窗口中，并且最好有交互功能。要大肠这样的效果也非常简单，只需修改命令行为：
```bash
godoc -http=:6060 -goroot="."
```
然后再访问http://localhost:6060/,点击顶部的`foo.go`，或者直接访问http://localhost:6060/pkg/foo/,就可以看到注释提取的效果。
端口可以切换为其他未被占用的端口号。

若要将注释提取为文档，要遵守如下的基本规则：
* 注释需要紧贴在对应的包声明和函数之前，能不有空行
* 注释如果要新起一个段落，应该用一个空白注释行隔开，因为直接换行书写会被认为是正常的段内折行
* 开发者可以直接在代码内用`//BUG(author):`的方式记录该代码片段中的遗留问题，这些遗留问题也会被抽取到文档中

Go语言是为开源项目而生的。从文档中可以看出，自动生成的文档包含源文件的跳转链接，通过它可以直接跳转到一个Go源文件甚至是一个特定的函数实现。如果开发者在看文档时觉得有障碍，可以直接跳转过去看代码，反正用Go语言写的代码通常都非常简洁、易懂，这样可以更有效的理解代码。
