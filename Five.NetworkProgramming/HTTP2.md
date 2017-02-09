### 5.2 HTTP编程
>HTTP(HyperText Transfer Protocol,超文本传输协议)是互联网上应用最为广泛的一种网络协议，定义了客户端和服务端之间请求与响应的传输标准

Go语言标准库内建提供了`net/http`包，涵盖了HTTP客户端和服务端的具体实现。使用`net/http`包，可以很方便的编写HTTP客户端或服务端的程序。

阅读本节内容，读者需要具备如下知识点：
* [了解HTTP基础知识](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol)
* [了解Go语言中接口的用法](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/Three.ObjectOrientedProgramming/Interface5.md)

#### 5.2.1 HTTP客户端

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

* `(*http.Client).Do()`
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

**2.高级封装**

除了之前介绍的基本HTTP操作，Go语言标准库也暴露了比较底层的HTTP相关库，让开发者可以基于这些库灵活定制HTTP服务器和使用HTTP服务。

* 自定义`http.Client`

前面使用的`http.Get()`、`http.Post()`、`http.PostForm()`和`http.Head()`方法其实都是在`http.DefaultClient`的基础上进行调用的，比如`http.Get()`等价于`http.DefaultClient.Get()`，以此类推。

`http.DefaultClient`在字面上就传达了一个信息，既然存在默认的Client，那么HTTP Client大概是可以自定义的。实际上确实如何，在`net/http`包中，的确提供了Client类型。`http.Client`支持的类型：

```go
type Client struct {
	// Transport specifies the mechanism by which individual
	// HTTP requests are made.
	// If nil, DefaultTransport is used.
	//Transport用于确定HTTP请求的创建机制。
	//如果为空，将会使用DefaultTransport
	Transport RoundTripper

	// CheckRedirect specifies the policy for handling redirects.
	// If CheckRedirect is not nil, the client calls it before
	// following an HTTP redirect. The arguments req and via are
	// the upcoming request and the requests made already, oldest
	// first. If CheckRedirect returns an error, the Client's Get
	// method returns both the previous Response and
	// CheckRedirect's error (wrapped in a url.Error) instead of
	// issuing the Request req.
	//
	// If CheckRedirect is nil, the Client uses its default policy,
	// which is to stop after 10 consecutive requests.
	//CheckRedirect定义重定向策略
	//如果CheckRedirect不为空，客户端将在跟踪HTTP重定向前调用该函数
	//两个参数req和via分别为即将发起的请求和已经发起的所有请求，最早的已发起请求在最前面
	//如果CheckRedirect返回错误，客户端将直接返回错误，不会再发起该请求
	//如果CheckRedirect为空，Client将采用一种默认策略，将在10个连续请求后终止
	CheckRedirect func(req *Request, via []*Request) error

	// Jar specifies the cookie jar.
	// If Jar is nil, cookies are not sent in requests and ignored
	// in responses.
	//如果Jar为空，Cookie将不会在请求中发送，并会在响应中被忽略
	Jar CookieJar

	// Timeout specifies a time limit for requests made by this
	// Client. The timeout includes connection time, any
	// redirects, and reading the response body. The timer remains
	// running after Get, Head, Post, or Do return and will
	// interrupt reading of the Response.Body.
	//
	// A Timeout of zero means no timeout.
	//
	// The Client cancels requests to the underlying Transport
	// using the Request.Cancel mechanism. Requests passed
	// to Client.Do may still set Request.Cancel; both will
	// cancel the request.
	//
	// For compatibility, the Client will also use the deprecated
	// CancelRequest method on Transport if found. New
	// RoundTripper implementations should use Request.Cancel
	// instead of implementing CancelRequest.
	Timeout time.Duration
}
```
在Go语言标准库中，`http.Client`类型包含了3个公开数据成员：
* `Transport RoundTripper`
* `CheckRedirect func(req *Request, via []*Request) err`
* `Jar CookieJar`

其中**`Transport`**类型必须实现`http.RoundTripper`接口。`Transport`指定了执行一个HTTP请求的运行机制，倘若不指定具体的`Transport`，默认会使用`http.DefaultTransport`,这意味着`http.Transport`也是可以自定义的。`net/http`包中的`http.Transport`类型实现了`http.RoundTripper`接口。

**`CheckRedirect`**函数指定处理重定向的策略。当使用HTTP Client的`Get()`或者是`Head()`方法发送HTTP请求时，若响应返回的状态码为30x(比如301/302/303/307)，HTTP Client会在遵循跳转规则前先调用这个`CheckRedirect`函数。

**`Jar`**可用于在HTTP Client中设定Cookie,Jar的类型必须实现了`http.CookieJar`接口，该接口预定义的`SetCookies()` 和`Cookies()`两个方法。如果HTTP Client中没有设定Jar，Cookie将被忽略而不会发送到客户端。实际上，一般都用`http.SetCookie()`方法来设定Cookie

使用自定义的`http.Client`及其`Do()`方法，可以非常灵活地控制HTTP请求，比如发送自定义HTTP Header或是改写重定向策略等。创建自定义的HTTP Client非常简单，具体代码如下

```go
client := &http.Client{
	CheckRedirect: redirectPolicyFunc,
}
resp, err := client.Get("http.example.com")
//...
req, err := http.NewRequest("GET", "http://example.com", nil)
//...
req.Header.Add("User-Agent", "Our Custom User-Agent")
req.Header.Add("If-None-Match", `W/"TheFileEtag"`)
resp, err = client.Do(req)
```

* 自定义`http.Transport`


在`http.Client`类型的结构定义中，看到的第一个数据成员就是一个`http.Transport`对象，该对象指定执行一个`HTTP`请求时的运行规则。

```go
//定义了`http.Transport`类型中的公开数据成员
type Transport struct {
	//...other ignore code

	// Proxy specifies a function to return a proxy for a given
	// Request. If the function returns a non-nil error, the
	// request is aborted with the provided error.
	// If Proxy is nil or returns a nil *URL, no proxy is used.
	//Proxy指定用于针对特定请求返回代理的函数。
	//接收一个*Request类型的请求实例作为参数并返回一个最终的HTTP代理
	//如果该函数返回一个非空的错误，请求将终止并返回该错误
	//如果Proxy为空或者返回一个空的URL指针，将不使用代理
	Proxy func(*Request) (*url.URL, error)

	// Dial specifies the dial function for creating unencrypted
	// TCP connections.
	// If Dial is nil, net.Dial is used.
	//Dial指定具体的dail()函数用于创建TCP连接
	//如果Dial为空，将默认使用`net.Dial()`函数
	Dial func(network, addr string) (net.Conn, error)
	//...other ignore code
	// TLSClientConfig specifies the TLS configuration to use with
	// tls.Client. If nil, the default configuration is used.
	//SSL连接专用.指定用于`tls.Client`的TLS配置信息
	//如果为空则使用默认配置
	TLSClientConfig *tls.Config

	// TLSHandshakeTimeout specifies the maximum amount of time waiting to
	// wait for a TLS handshake. Zero means no timeout.
	TLSHandshakeTimeout time.Duration

	// DisableKeepAlives, if true, prevents re-use of TCP connections
	// between different HTTP requests.
	//是否取消长连接，默认值为false，即启用长连接
	DisableKeepAlives bool

	// DisableCompression, if true, prevents the Transport from
	// requesting compression with an "Accept-Encoding: gzip"
	// request header when the Request contains no existing
	// Accept-Encoding value. If the Transport requests gzip on
	// its own and gets a gzipped response, it's transparently
	// decoded in the Response.Body. However, if the user
	// explicitly requested gzip it is not automatically
	// uncompressed.
	//是否取消压缩(GZip)，默认值为false，即启用压缩
	DisableCompression bool

	// MaxIdleConnsPerHost, if non-zero, controls the maximum idle
	// (keep-alive) to keep per-host.  If zero,
	// DefaultMaxIdleConnsPerHost is used.
	//如果非零值，指定与每个请求的目标主机之间的最大非活跃连接(keep-alive)数量
	//如果该值为空，则使用DefaultMaxIdleConnsPerHost常量值
	MaxIdleConnsPerHost int

	//...other ignore code
}

// CloseIdleConnections closes any connections which were previously
// connected from previous requests but are now sitting idle in
// a "keep-alive" state. It does not interrupt any connections currently
// in use.
//用于关闭所有非活跃的连接
func (t *Transport) CloseIdleConnections() 
// RegisterProtocol registers a new protocol with scheme.
// The Transport will pass requests using the given scheme to rt.
// It is rt's responsibility to simulate HTTP request semantics.
//
// RegisterProtocol can be used by other packages to provide
// implementations of protocol schemes like "ftp" or "file".
//
// If rt.RoundTrip returns ErrSkipAltProtocol, the Transport will
// handle the RoundTrip itself for that one request, as if the
// protocol were not registered.
//该方法可用于注册并启用一个新的传输协议，比如WebSocket的传输协议标准(ws)，或者FTP、File协议等
func (t *Transport) RegisterProtocol(scheme string, rt RoundTripper)
// RoundTrip implements the RoundTripper interface.
//
// For higher-level HTTP client support (such as handling of cookies
// and redirects), see Get, Post, and the Client type.
//用于实现http.RoundTripper接口
func (t *Transport) RoundTrip(req *Request) (*Response, error)
```

**自定义`http.Transport`**

```go
tr := &http.Transport{
		TLSClientConfig:    &tls.Config{RootCAs: pool},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://example.com")
```
Client和Transport在执行多个`goroutine`的并发过程中都是安全的，单处于性能考虑，应当创建一次后反复使用。


* 灵活的`http.RoundTripper`接口

在前面的两小节中，我们知道`HTTP Client`是可以自定义的，而`http.Client` 定义的第一个公开成员就是一个`http.Transport`类型的实例，且该成员所对应的类型必须实现`http.RoundTripper`接口。

```go 
type RoundTripper interface {
	// RoundTrip executes a single HTTP transaction, returning
	// a Response for the provided Request.
	//
	// RoundTrip should not attempt to interpret the response. In
	// particular, RoundTrip must return err == nil if it obtained
	// a response, regardless of the response's HTTP status code.
	// A non-nil err should be reserved for failure to obtain a
	// response. Similarly, RoundTrip should not attempt to
	// handle higher-level protocol details such as redirects,
	// authentication, or cookies.
	//
	// RoundTrip should not modify the request, except for
	// consuming and closing the Request's Body.
	//
	// RoundTrip must always close the body, including on errors,
	// but depending on the implementation may do so in a separate
	// goroutine even after RoundTrip returns. This means that
	// callers wanting to reuse the body for subsequent requests
	// must arrange to wait for the Close call before doing so.
	//
	// The Request's URL and Header fields must be initialized.
	//RoundTrip执行一个单一的HTTP事务，返回相应的相应信息
	//RoundTrip函数的实现不应该试图去理解响应的内容。
	//如果RoundTrip得到一个响应，无论该响应的HTTP状态码如何，都应将返回的err设置为nil。
	//非空的err只以为着没有成功获取到响应。
	//类似的，RoundTrip也不应识图处理更高级别的协议，比如重定向、认证和Cookie等。
	//RoundTrip不应修改响应内容，除非是为了理解Body内容。
	//每一个请求的URL和Header域都应被正确初始化
	RoundTrip(*Request) (*Response, error)
}
```

`http.ToundTripper`接口只定义了一个名为`RoundTrip`的方法。任何实现了`RoundTrip()`方法的类型即可实现`http.RoundTripper`接口。前面我们看到的`http.Transport`类型正是实现了`RoundTrip()`方法继而实现了该接口。

通常，我们可以在默认的`http.Transport`之上包一层`Transport`并实现`RoundTrip()`方法，[例子customtrans.go](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/code/ChapterFive/5.2.1HTTPClient/customtrans.go)

因为实现了`http.RoundTripper`接口的代码通常需要在多个goroutine中并发执行，因此我们必须确保实现代码的线程安全性。

* `HTTP Client`

综上示例讲解可以看到，Go语言标准库提供的HTTP Client是相当优雅的。一方面提供了极其简单的使用方式，另一方面又具备极大的灵活性。
Go语言标准库提供的HTTP Client被设计成上下两层结构。一层是上述提到`http.Client`类及其封装的基础方法，我们不妨称其为"业务层"，是因为调用方通畅只需要关心请求的业务逻辑本身，而无须关心非业务相关的技术细节，这些细节包括

* HTTP底层传输细节
* HTTP代理
* gzip压缩
* 连接池及其管理
* 认证(SSL或其他认证方法)
	
之所以`HTTP Client`可以做到这么好的封装行，是因为`HTTP Client`在底层抽象了`http.RoundTripper`接口，而`http.Transport`实现了该接口，从而能够处理更多的细节，我们不妨将其称为"传输层"。`HTTP Client`在业务层初始化`HTTP Method`、目标URL、请求参数、请求内容等重要信息后，经过"传输层","传输层"在业务层处理的基础上补充其他细节，然后再发起HTTP请求，接收服务端返回的`HTTP`响应。

#### 5.2.2 HTTP服务端

本节将介绍HTTP服务端结束，包括如何处理HTTP请求和HTTPS请求。

**1.处理HTTP请求**

使用`net/http`包提供的`http.ListenAndServe()`方法，可以在指定的地址进行监听，开启一个`HTTP`，服务端该方法的原型如下：

```go
/**
该方法用于在指定的TCP网络地址addr进行监听，然后调用服务端处理程序来处理传入的连接请求。

addr:监听地址
handler:服务端处理程序(通常为空)
*/
func ListenAndServe(addr string,handler Handler) err
```
其中handler为空，这意味着服务端调用http.DefaultServeMux进行处理，而服务端编写的业务逻辑处理程序http.Handle()或http.HandleFunc()默认注入http.DefaultServeMux中，具体代码如下：

```go
http.Handle("/foo",fooHandler)
http.HandleFunc("/bar",func(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello,%q",html.EscapeString(r.URL.Path))
})
log.Fatal(http.ListenAndServe(":8080",nil))
```

如果想更多地控制服务端的行为，可以自定义`http.Server`，代码如下：

```go
s := &http.Server{
		Addr:           ":8080",
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
log.Fatal(s.ListenAndServe())
```

**2.处理`HTTPS`请求**

`net/http`包还提供`http.ListenAndServeTLS()`方法，用于处理https连接请求：
```go
func ListenAndServeTLS(addr string,certFile string,keyFile string,handler Handler) error
```
`ListenAndServeTLS`和`ListenAndServe`的行为一致，区别在于只处理`HTTPS`请求。
此外，服务器上必须包含证书和与之匹配的私钥的相关文件，比如`certFile`对应的SSL证书文件存放路径，`keyFile`对应证书私钥文件路径。如果证书是由证书颁发机构签署的，`certFile`参数指定的路径必须是存放在服务器上的经由CA认证过的SSL证书。

开启SSL监听服务也很简单，代码如下所示：
```go
http.Handle("/foo",fooHandler)
http.HandleFunc("/bar",func(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello,%q",html.EscapeString(r.URL.Path))
})
log.Fatal(http.ListenAndServeTLS(":10443","cert.pem,"key.pem",nil))
```
或者是
```go
ss := &http.Server{
		Addr:           ":10443",
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
log.Fatal(ss.ListenAndServeTLS("cert.pem","key.pem"))
```
