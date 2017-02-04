### 5.4 JSON处理
>JSON(JavaSript Object Notation)是一种比XML更轻量级的数据交换格式，在易于人们阅读和编写的同时，也易于程序解析和生成。尽管JSON是JavaSript的一个子集，但JSON采用完全独立于编程语言的文本格式，且表现为键/值怼集合的文本描述形式（类似一些编程语言中的字典结构），这使它成为较为理想的、跨平台、跨语言的数据交换语言。

开发者可以用JSON传输简单的字符串、数字、布尔值，也可以传输i 个数组，或者一个更复杂的复合结构。在Web开发领域中，JSON被广泛应用于Web服务端程序和客户端之间的数据通信，但也不仅仅局限于此，其应用范围非常广阔，比如作为Web Services API输出的标准格式，又或是用作程序网络通信中的远程过程调用(RPC)等。

关于JSON的更多信息，请访问[JSON官方网站](http://json.org/)查阅

Go语言内建对JSON的支持。使用Go语言内置的`encoding/json`标准库，开发者可以轻松使用Go程序生成和解析JSON格式的数据。在Go语言实现JSON的编码和解码时，遵循RFC4627协议标准。

#### 5.4.1 编码为JSON格式
使用`json.Marshal()`函数可以对一组数据进行JSON格式的编码。`json.Marshal()`函数的声明如下：
```go
func Marshal(v interface{}) ([]byte, error) 
```
假设有如下一个`Book`类型的结构体:
```go
type Book struct {
	Title       string
	Authors     string
	Publisher   string
	IsPublished bool
	Price       float64
}
```
并且有如下一个`Book`类型的实例对象：
```go 
////下面Authors赋值有个坑，前面必须带有[]string,否则会报错：missing type in composite literal
gobook := Book{
		"Go语言编程",
		[]string{"XuShiWei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan", "XuDaoli"},
		"ituring.com.cn",
		true,
		9.99,
	}
```
然后，我们可以使用`json.Marshal()`函数将gobook实例生成一段JSON格式的文本。
```go
b, err := json.Marshal(gobook)
```
如果编码成功，`err`将赋于零值nil，变量`b`将会是一个进行JSON格式化后的`[]byte`类型：
```go
b == []byte(`{
		"Title" : "Go语言编程",
		"Authors" : ["XuShiWei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan", "XuDaoli"],
		"Publisher" : "ituring.com.cn",
		"IsPublished" : true,
		"Price" : 9.99
	}`)
```
当我们调用`json.Marshal(gobook)`语句时，会递归遍历`gobook`对象，如果发现`gobook`这个数据结构实现了`json.Marshal`接口且包含有效的值，`Marshal()`就会调用其`MarshalJSON()`方法将该数据结构生成JSON格式的文本。

Go语言的大多数数据类型都可以转化为有效的JSON格式，但`channel`、`complex`和`函数`这几种类型除外。

如果转化前的数据结构中出现指针，那么将会转化指针所指向的值，如果指针指向的是零值，那么`null`将作为转化后的结果输出。

在Go中，JSON转化前后的数据类型映射如下：
* 布尔值->布尔值
* 浮点数和整型->常规数字
* 字符串->以`UTF-8`编码转化输出为`Unicode`字符集的字符串，特殊字符比如`<`将会被转义为`\u003c`
* 数组和切片->数组，`[]byte`->`Base64`编码后的字符串，`slice`类型的零值->`null`
* 结构体->`JSON`对象，并且只有结构体里面以大写字母开头的可被导出的字段才会被转化输出，而这些可导出的字段会作为JSON对象的字符串索引
* `map`类型的数据结构，类型必须是`map[string]T`(T可以是`encoding/json`包支持的任意数据类型)

#### 5.4.2 解码JSON数据
可以使用`json.Unmarshal()`函数将JSON格式的文本解码为Go里预期的数据结构。
`json.Unmarshal()`函数的原型如下：
```go
func Unmarshal(data []byte, v interface{}) error
```