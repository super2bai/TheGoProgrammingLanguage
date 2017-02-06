
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
/**
data:输入，JSON格式的文本(比特序列)
v:目标输出容器，用于存放解码后的值
*/
func Unmarshal(data []byte, v interface{}) error
```
解码代码：
```go
fmt.Println("Unmarshal")
	bookJSON := []byte(`{
			"Title" : "Go语言编程",
			"Authors" : ["XuShiWei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan", "XuDaoli"],
			"Publisher" : "ituring.com.cn",
			"IsPublished" : true,
			"Price" : 9.99
		}`)
	var book Book
	err = json.Unmarshal(bookJSON, &book)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(book)
	//output:{Go语言编程 [XuShiWei HughLv Pandaman GuaguaSong HanTuo BertYuan XuDaoli] ituring.com.cn true 9.99}
```
Go时如何将JSON数据解码后的值一一准确无误的关联到一个数据结构中的相应字段呢？
>`json.Unmarshal()`函数会根据一个约定的顺序查找目标结构中的字段，如果找到一个即发生匹配。假设一个JSON对象有个名为"Foo"的索引，要将`Foo`所对应的值填充到目标结构体的目标字段上，`json.Unmarshal()`将会遵循如下顺序进行查找匹配：
* 包含Foo标签的字段
* 名为Foo的字段
* 名为Foo或者除了首字母其他字母不区分大小写的名为Foo的字段
这些字段在类型生命中必须都是以大写字母开头、可被导出的字段。

但是当JSON数据里面的结构和Go里面的目标类型的结构对不上时，会发生什么呢？
```go
bookWRONG := []byte(`{
			"Title" : "Go语言编程",
			"Sales" : 1
		}`)
	var gobookWRONG Book
	err = json.Unmarshal(bookWRONG, &gobookWRONG)
	if err != nil {
		fmt.Println(1)
		fmt.Println(err)
	} else {
		fmt.Println("gobookWRONG=", gobookWRONG)
	}
//output:gobookWRONG= {Go语言编程 []  false 0}
```
如果JSON中的字段在Go目标类型中不存在，`json.unmarshal()`函数在解码过程中会丢弃改字段。在上面的示例代码中，由于`Sales`字段并没有在`Book`类型中定义，所以会被忽略，只有`Tile`这个字段的值才会被填充到`gobookWRONG`中。

这个特性可以从同一段JSON数据中筛选指定的值填充到多个Go语言类型中，当然，前提是已知JSON数据的字段结构。这也同样意味着，目标类型中不可被导出的私有字段（非首字母大写）将不会受到解码转化的影响。

但，如果JSON的数据结构是未知的，应该如何处理呢？

#### 5.4.3 解码未知结构的JSON数据
我们已经知道，Go语言支持接口。在Go语言里，接口是一组预定义方法的组合，任何一个类型均可通过实现接口预定义的方法来实现，且无需显示声明，所以没有任何方法的空接口可以代表任何类型。换句话说，每一个类型其实都至少实现了一个空接口。

Go内建这样灵活的类型系统，向我们传达了一个很有价值的信息：空接口是通用类型。如果要解码一段未知结构的JSON，只需将这段JSON数据解码输出到一个空接口即可。在解码JSON数据的过程中，JSON数据里面的元素类型将做如下替换：
* 布尔值 -> `bool`
* 数值 -> `float64`
* 字符串 -> `string`
* 数组 -> `[]interface{}`
* 对象 -> `map[string]interface{}`
* `null` -> `nil`

在Go的标准库`encoding/json`包中，允许使用`map[string]interface{}`和`[]interface{}`类型的值来分别存放未知结构的JSON对象或数组，示例代码如下：
```go
/*
output:map[Title:Go语言编程 Authors:[XuShiWei HughLv Pandaman GuaguaSong HanTuo BertYuan XuDaoli] Publisher:ituring.com.cn IsPublished:true Price:9.99 Sales:100000]
*/
	bookJSON := []byte(`{
			"Title" : "Go语言编程",
			"Authors" : ["XuShiWei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan", "XuDaoli"],
			"Publisher" : "ituring.com.cn",
			"IsPublished" : true,
			"Price" : 9.99,
			"Sales":100000
		}`)
	var bookInterface interface{}
	err = json.Unmarshal(bookJSON, &bookInterface)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(bookInterface)
```
在上述代码中，`bookInterface`被定义为一个空接口。`json.Unmarshal()`函数将一个JSON对象解码到空接口`bookInterface`中，最终`bookInterface`将会是一个键值对的`map[string]interface{}`结构：
```go
map[string]interface{
			"Title" : "Go语言编程",
			"Authors" : ["XuShiWei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan", "XuDaoli"],
			"Publisher" : "ituring.com.cn",
			"IsPublished" : true,
			"Price" : 9.99,
			"Sales":100000
		}
```
要访问解码后的数据结构，需要先判断目标结构是否为预期的数据类型,然后，可以通过for循环搭配range语句一一访问解码后的目标数据：
```go
bk, ok := bookInterface.(map[string]interface{})
	if ok {
		for k, v := range bk {
			switch v2 := v.(type) {
			case string:
				fmt.Println(k, " is string ", v2)
			case int:
				fmt.Println(k, " is int ", v2)
			case bool:
				fmt.Println(k, " is bool ", v2)
			case []interface{}:
				fmt.Println(k, " is an array : ")

				for i, iv := range v2 {
					fmt.Println(i, iv)
				}
			default:
				fmt.Println(k, " is another type not handle yet")
			}
		}
	}

```
虽然有些繁琐，但的确是一种解码未知结构的JSON数据的安全方式。


#### 5.4.4 JSON的流式读写
Go内建的`encoding/json`包还提供 `Decoder`和`Encoder`两个类型，用于支持JSON数据的流式读写，并提供`NewDecoder()`和`NewEncoder()`两个函数来便于具体实现
```go
package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Panicln(err)
			return
		}
		for k := range v {
			if k != "Title" {
				delete(v, k)
			}
		}
		if err := enc.Encode(&v); err != nil {
			log.Panicln(err)
		}

	}
}
```
使用`Decoder`和`Encoder`对数据流进行处理可以应用得更为广泛些，比如读写`HTTP`连接、`WebSocket`或文件等，Go的标准库`net/rpc/jsonrpc`就是一个应用了`Decoder`和`Encoder`的实际例子。
