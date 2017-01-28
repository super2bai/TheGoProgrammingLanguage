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