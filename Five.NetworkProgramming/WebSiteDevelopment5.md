###5.5 网站开发

在这一节中，我们将循序渐进地讲解怎样使用Go进行Web开发，本节将围绕一个简单的相册程序进行，尽管示例程序比较简单，但体现的都是使用Go开发网站的基础关键环节，旨在系统了解机基于原生的Go语言开发Web应用程序的基本思路及其相关细节的具体实现。

####5.5.1 最简单的网站程序
从最简单的网站程序入手。
第1章中编写的最简单的Hello World示例程序，稍微调整几行代码，将其改造成一个可用浏览器打开并能在网页中显示"Hello,world"的小程序。
```go
package main

import (
	"io"
	"log"
	"net/http"
)

/*
w:HTTP请求的目标路径"/hello"，该参数值可以是字符串，也可以是字符串形式的正则表达式

r:指定具体的毁掉方法，比如helloHandler
*/
func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello,world!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

```
我们引入了Go语言标准库中的`net/http`包，主要用于提供Web服务，响应并处理客户端(浏览器)的HTTP请求。同时，使用`io`包而不是`fmt`包来输出字符串，这样源文件编译成可执行文件后，体积要小很多，运行起来也更省资源。
接下来，让我们简单的了解Go语言的`http`包在上述示例中所做的工作。

#### 5.5.2 `net/http`包简介
可以看到，在`main()`方法中调用了`http.HandleFunc()`，该方法用于**分发请求**，即针对某一路径请求将其映射到指定的业务逻辑方法中。如果你有其他语言（比如Ruby、Python或者PHP等）的Web开发经验，可以将其形象的理解为提供类似URL路由或者URL映射之类的功能。

`helloHandler()`方法是`http.HandlerFunc`类型的实例，并传入`http.ResponseWriter`和`http.Request`作为其必要的两个参数。

`http.ResponseWriter`类型的对象用于**包装处理HTTP服务端的响应信息**。我们将字符串"Hello,world!"写入类型为`http.ResponseWriter`的`w`实例中，即可将该字符串数据发送到`HTTP`客户端。`http.Request`表示的是此次`HTTP`请求的一个数据结构体，即代表一个客户端，不过该示例中我们还未用到它。

`main()`方法中调用了`http.ListenAndServe()`，该方法用于在示例中监听`8080`端口，接收并调用内部程序来处理连接到此端口的请求。如果端口监听失败，会调用`log.Fatal()`方法输出异常出错信息。

正如你所见，`main()`方法中的短短两行即开启了一个`HTTP`服务，使用Go语言的`net/http`包搭建一个Web是如此简单！当然，`net/http`包的作用远不止这些，我们只用到其功能的一小部分。

####5.5.3 开发一个简单的相册网站
本节我们将综合之前介绍的网站开发相关知识，一步步介绍如何开发一个虽然简单但五脏俱全的相册网站。

**1.新建工程**
首先创建一个用于存放工程源代码的目录并切换到该目录中去，随后创建一个名为`photoweb.go`的文件，用于后面编辑我们的代码：
```bash
$ mkdir -p photoweb/uploads
$ cd photoweb
$ touch photoweb.go
```
示例程序实现功能：
* 支持图片上传
* 在网页中可以查看已上传的图片
* 能看到所有上传的图片列表
* 可以删除指定的图片
 
[示例代码](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/code/ChapterFive/5.5WebSiteDevelopment/photoweb)
