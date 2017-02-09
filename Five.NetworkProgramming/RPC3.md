### 5.3 RPC编程
>RPC(Remote Procedure Call,远程过程调用)是一种通过网络从远程计算机程序上请求服务，而不需要了解底层网络细节的应用程序通信协议。RPC协议构建与TCP或UDP，或者是HTTP之上，允许开发者直接调用另一台计算机上的程序，而开发者无需额外地为这个调用过程编写网络通信相关代码，使得开发包括网络分布式程序在内的应用程序更加容易。

RPC采用**客户端--服务器(Client/Server)**的工作模式。**请求过程**就是一个客户端(Clinet),而服务提供程序就是一个服务器(Server)。当执行一个远程过程调用时，**客户端程序**首先发送一个带有参数的调用信息到服务端，然后等待服务端响应。在**服务端**，服务进程保持睡眠状态直到客户端的调用信息到达为止。当一个调用信息到达时，服务端**获得进程参数**，**计算出结果**，并向客户端**发送应答信息**，然后**等待下一个调用**。最后，客户端**接收来自服务端的应答信息**，**获得进程结果**，然后调用执行并继续进行。

#### 5.3.1 Go语言中的RPC支持与处理
在Go中，标准库提供的`net/rpc`包实现了RPC协议需要的相关细节，开发者可以很方便地使用该包编写RPC的服务端和客户端程序，这使得用Go语言开发的多个进程之间的通信变得非常简单。

`net/rpc`包允许RPC客户端程序通过网络或是其他I/O连接调用一个远端对象的公开方法(必须是大写字母开头、可外部调用的)。在RPC服务端，可将一个对象注册为客访问的服务，之后该对象的公开方法就能够以远程的方式提供访问。一个RPC服务端可以注册多个不同类型的对象，但不允许注册同一类型的多个对象。

一个对象中只有满足如下这些条件的方法，才能被RPC服务端设置为可供远程访问：
* 必须是在对象外部可公开调用的方法(首字母大写)
* 必须有两个参数，且参数的类型都必须是包外部可以访问的类型或者是Go内建支持的类型
* 第二个参数必须是一个指针
* 方法必须返回一个error类型的值
以上4个条件，可以简单地用如下一行代码表示：
```go
func(t *T) MethodName(argType T1,replyType *T2) error
```
在上面这行代码中，类型T、T1和T2默认会使用Go内置的`encoding/gob`包(稍后介绍)进行编码解码。

该方法(MethodName)的第一个参数表示由RPC客户端传入的参数，第二个参数表示要返回给RPC客户端的结果，该方法最后返回一个error类型的值。

RPC服务端可以通过调用`rpc.ServeConn`处理单个连接请求。多数情况下，通过TCP或是HTTP在某个网络地址上进行监听来创建该服务是个不错的选择。

在RPC客户端，Go的`net/rpc`包提供了便利的`rpc.Dial()`和`rpc.DialHTTP()`方法来与指定的RPC服务端建立连接。

在建立连接之后，Go的`net/rpc`包允许我们使用同步或者异步的方式接收RPC服务端的处理结果。调用RPC客户端的**`Call()`**方法则进行**同步处理**,这时候客户端程序安顺序执行，只有接受完RPC服务端的处理结果之后才可以继续执行后面的程序。当调用RPC客户端的**`Go()`**方法时，则可以进行**异步处理**，RPC客户端程序无需等待服务端的结果即可执行后面的程序，而当接收到RPC服务端的处理结果时，再对其进行相应的处理。

无论是调用RPC客户端的`Call()`或者是`Go()`方法，都必须指定要调用的服务及其方法名称，以及一个客户端传入参数的引用，还有一个用于接收处理结果参数的指针。

如果没有明确指定RPC传输过程中使用何种编码解码器，默认将使用Go标准库提供的`encoding/gob`包进行数据传输。
[RPC服务端和客户端交互的示例程序](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/code/ChapterFive/5.3.1RPC/rpcserver.go)

#### 5.3.2 Gob简介
>Gob是Go的一个**序列化数据结构的编码解码工具**，在Go标准库中内置`encoding/gob`包以供使用。一个数据结构使用Gob进行序列化之后，能够用于网络传输。与JSON或XML这种基于文本描述的数据交换语言不通，Gob是二进制编码的数据流，并且Gob流是可以自解释的，它在保证高效率的同时，也具备完整的表达能力。

作为针对Go的数据结构进行编码和解码的专用序列化方法，这意味着Gob无法跨语言使用。在Go的`net/rpc`包中，传输数据所需要用到的编码解码器，默认就是Gob。**由于Gob仅局限于使用Go开发的程序，这意味着我们只能用Go的RPC实现进程间通信**。然而，大多数时候，我们用Go编写RPC服务器（或客户端），可能更希望它是通用的，与语言无关的，无论是Python、Java货其他语言实现的RPC客户端，均可与之通信。

#### 5.3.3 设计优雅的RPC接口
Go的`net/rpc`很灵活，它在数据传输前后显示了编码解码器的接口定义。这意味着，开发者可以自定义数据的传输方式以及RPC服务端和客户端之间的交互行为。

RPC提供的编码解码器接口如下：
```go
// A ClientCodec implements writing of RPC requests and
// reading of RPC responses for the client side of an RPC session.
// The client calls WriteRequest to write a request to the connection
// and calls ReadResponseHeader and ReadResponseBody in pairs
// to read responses.  The client calls Close when finished with the
// connection. ReadResponseBody may be called with a nil
// argument to force the body of the response to be read and then
// discarded.
type ClientCodec interface {
	// WriteRequest must be safe for concurrent use by multiple goroutines.
	WriteRequest(*Request, interface{}) error
	ReadResponseHeader(*Response) error
	ReadResponseBody(interface{}) error

	Close() error
}

// A ServerCodec implements reading of RPC requests and writing of
// RPC responses for the server side of an RPC session.
// The server calls ReadRequestHeader and ReadRequestBody in pairs
// to read requests from the connection, and it calls WriteResponse to
// write a response back.  The server calls Close when finished with the
// connection. ReadRequestBody may be called with a nil
// argument to force the body of the request to be read and discarded.
type ServerCodec interface {
	ReadRequestHeader(*Request) error
	ReadRequestBody(interface{}) error
	// WriteResponse must be safe for concurrent use by multiple goroutines.
	WriteResponse(*Response, interface{}) error

	Close() error
}
```

接口`ClientCodec`定义了RPC客户端如何在一个RPC会话中发送请求和读取响应。客户端程序通过`WriteRequest()`方法将一个请求写入到RPC连接中，并通过`ReadResponseHeader()`和`ReadResponseBody()`读取服务器端的响应信息。当整个过程执行完毕后，再通过Close()方法来关闭该连接。

接口`ServerCodec`定义了RPC服务器如何在一个RPC绘画中接收请求并发送响应。服务器端程序通过`ReadRequestHeader()`和`ReadRequestBody()`方法从一个RPC连接中读取请求信息，然后再通过`WriteResponse()`方法向该连接中的RPC客户端发送响应。当完成该过程欧，通过`Close()`方法来关闭连接。

通过实现上述接口，我们可以自定义数据传输前后的编码解码方式，而不仅仅局限于Gob。同样，可以自定义RPC服务端和客户端的交互行为。实际上， Go标准库提供的`net/rpc/json`包，就是一套实现了`rpc.ClientCodec`和`rpc.ServerCodec`接口的`JSON-RPC`模块。