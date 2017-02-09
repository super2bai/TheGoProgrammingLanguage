### 5.1 Socket
在Go语言中编写网络程序时，我们看不到传统的编码形式。以前我们使用Socket编程时，会按照如下步骤展开：
* 建立Socket：使用`socket()`函数
* 绑定Socket：使用`bind()`函数
* 监听：使用`listen()`函数。或者连接：使用`connect()`函数
* 接受连接：使用`accept()`函数
* 接收：使用`receive()`函数。或者发送：使用`send()`函数

Go语言标准库对此过程进行了抽象和封装。无论我们期望使用什么协议建立什么形式的连接，都只需要调用`net.Dial()`即可。


#### 5.1.1 Dial()函数
`Dial()`函数的原型如下：
```go

/**
net:网络协议的名字
addr:IP地址或域名，ex:192.168.0.196:8026(端口号可选)

return:如果连接成功，返回连接对象，否则返回error
*/
func Dial(net,addr string)(Conn,error)
```

几种常见协议的调用方式。
* TCP链接
```go
conn, err := net.Dial("tcp","192.168.0.10:2100")
```
* UDP链接
```go
conn, err := net.Dial("udp","192.168.0.12:975")
```
* ICMP链接(使用协议名称)
```go
conn, err := net.Dial("ip4:icmp","www.baidu.com")
```
* ICMP链接(使用协议编号)
```go
conn, err := net.Dial("ip4:1","10.0.0.3")
```

[协议编号的含义](http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xml)


目前，`Dial()`函数支持如下几种网络协议：
* tcp
* tcp4(仅限IPv4)
* tcp6(仅限IPv6)
* udp
* udp(仅限IPv4)
* udp(仅限IPv6)
* ip
* ip4(仅限IPv4)
* ip6(仅限IPv6)

在成功建立连接后，就可以进行数据的发送和接收。发送数据时，使用`conn`的`Write()`成员方法，接收数据时使用`Read()`方法。

#### 5.1.2 ICMP示例程序

例子：使用`ICMP`协议向在线的主机发送一个问候，并等待主机返回。

[ICMP示例程序](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/code/ChapterFive/5.1.2ICMP/icmptest.go)

用法：
```bash
$ sudo su
$ go build icmptest.go
$ ./icmptest www.baidu.com

output:
[69 0 9 0 32 96 0 0 54 1 187 104 61 135 169 125 192 168 1 107 0 0 156 205 0 13 0 37 99]
Got response
Identifier matches
Sequence matches
Custom data matches
```

#### 5.1.3 TCP示例程序

例子：建立`TCP`链接来实现初步的HTTP协议，通过向网络主机发送HTTP Head请求，读取网络主机返回的信息。
[simplehttp](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/code/ChapterFive/5.1.3TCP/simplehttp.go)

用法：
```bash
$ go build simplehttp.go
$ ./simplehttp baidu.com:80

output:
HTTP/1.1 200 OK
Date: Sun, 29 Jan 2017 04:53:06 GMT
Server: Apache
Last-Modified: Tue, 12 Jan 2010 13:48:00 GMT
ETag: "51-47cf7e6ee8400"
Accept-Ranges: bytes
Content-Length: 81
Cache-Control: max-age=86400
Expires: Mon, 30 Jan 2017 04:53:06 GMT
Connection: Close
Content-Type: text/html
```

#### 5.1.4 更丰富的网络通信
实际上，`Dial()`函数是对`DialTCP()`、`DialUDP()`、`DialIP()`和`DialUnix()`的封装。也可以直接调用这些函数，它们的功能是一致的。这些函数的原型如下：
```go
func DialTCP(net string, laddr, raddr *TCPAddr) (*TCPConn, error) 
func DialUDP(net string, laddr, raddr *UDPAddr) (*UDPConn, error) 
func DialIP(netProto string, laddr, raddr *IPAddr) (*IPConn, error) 
func DialUnix(net string, laddr, raddr *UnixAddr) (*UnixConn, error) 
```
之前基于TCP发送HTTP请求，读取服务器返回的HTTP head的整个流程也可以使用[simplehttp2](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/code/ChapterFive/5.1.3TCP/simplehttp2.go)

`net.ResolveTCPAddr()`用于解析地址和端口号
`net.DialTCP()`用于建立链接
这两个函数都在`Dial()`中得到了封装。
此外，`net`包中还包含了一系列的工具函数，合理地使用这些函数可以更好地保障程序地质量。
```go
//验证IP地址有效性地代码
func net.ParseIP()
//创建子网掩码地代码
func IPv4Mask(a,b,c,d byte) IPMask
//获取默认子网掩码
func (ip IP)DefaultMask() IPMask
//根据域名查找IP的代码
func ResolveIPAddr(net,addr string)(*IPAddr, error)
func LookupHost(name string)(cname string,addrs []string,err error)
```