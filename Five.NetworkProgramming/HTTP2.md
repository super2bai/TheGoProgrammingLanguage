### 5.2 HTTP编程
>HTTP(HyperText Transfer Protocol,超文本传输协议)是互联网上应用最为广泛的一种网络协议，定义了客户端和服务端之间请求与响应的传输标准

Go语言标准库内建提供了`net/http`包，涵盖了HTTP客户端和服务端的具体实现。使用`net/http`包，可以很方便的编写HTTP客户端或服务端的程序。

阅读本节内容，读者需要具备如下知识点：
* [了解HTTP基础知识](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol)
* [了解Go语言中接口的用法](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/Three.ObjectOrientedProgramming/Interface5.md)

####5.2.1 HTTP客户端

Go内置的`net/http`包提供了最简洁的HTTP客户端实现，无需借助第三方网络通信库(比如`libcurl`)就可以直接使用HTTP中用得最多的`GET`和`POST`方式请求数据。

**1.基本方法**
`net/http`包的Client类型提供了如下几个方法，可以用最简洁的方式实现HTTP请求：
```go
func (c *Client) Get(url string) (resp *Response, err error)
func (c *Client) Post(url string, bodyType string, body io.Reader) (resp *Response, err error) 
func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error) 
func (c *Client) Head(url string) (resp *Response, err error)
func (c *Client) Do(req *Request) (resp *Response, err error) 
```

下面该要介绍这几个方法。
* `http.Get()`
 要请求一个资格资源，只需调用`http.Get()`方法(等价于`http.DefaultClient.Get()`)即可。
 
 ``` go
 resp, err := http.Get("http://example.com/")
	if err != nil {
		//处理错误
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout,resp.Body)
 ```
 上面这段代码请求一个网站首页，并将其网页内容打印到标准输出流中。

* `http.Post()`
要以POST的方式发送数据，也很简单，只需调用http.Post()方法并依次传递下面3个参数即可
	* 请求的目标URL
	* 将要POST数据的资源类型(`MIMEType`)
	* 数据的比特刘(`[]byte`形式)
	
```go 
		//上传图片
	resp, err := http.Post("http://example.com/upload", "image/jpeg", &imageDataBuf)
	if err != nil {
		//处理错误
		return
	}
	if resp.StatusCode != http.StatusOK {
		//处理错误
		return
	}
```	

* `http.PostForm()`
此方法实现了标准编码格式为`application/x-www-form-urlencoded`的表单提交。

```go
	//模拟HTML表单提交一篇新文章
	resp, err := http.PostForm("http://example.com/posts", url.Values{"title":{"article title"},"content":{"articel body"}})
	if err != nil {
		//处理错误
		return
	}
```

* `http.Head()`
此表明只请求目标URL的头部信息，即`HTTPHeader`而不返回`HTTPBody`。
Go内置的`net/http`包同样也提供了`http.Head()`方法，该方法同`http.Get()`方法一样，只需传入目标URL一个参数即可。

```go
//请求一个网站首页的HTTPHeader信息
resp, err := http.Head("http://example.com/")
```

* (*http.Client).Do()
在多数情况下，`http.Get()`和`http.PostForm()`就可以满足需求，但是如果我们发起的HTTP请求需要更多的定制信息，设定一些自定义的`Http Header`字段，比如：
	* 设定自定义的"User-Agent"，而不是默认的"Go http package"
	* 传递Cookie

此时可以使用`net/http`包`http.Client`对象的`Do()`方法来实现：

```go
req, err := http.NewRequest("GET", "http://example.com", nil)
//...
req.Header.Add("User-Agent", "GoBook Custom User-Agent")
//...
client := &http.Client{}
resp, err := client.Do(req)
//...
```