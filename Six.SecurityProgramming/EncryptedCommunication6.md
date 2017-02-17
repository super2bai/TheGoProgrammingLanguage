### 6.6 加密通信

### 6.6 加密通信
一般的`HTTPS`是基于`SSL(Secure Sockets Layer)`协议。`SSL`是网景公司开发的位于`TCP`与`HTTP`之间的透明安全协议，通过`SSL`，可以把`HTTP`包数据以非对称加密的形式往返于浏览器和站点之间，从而避免被第三方非法获取。

目前，伴随着电子商务的兴起，`HTTPS`获得了广泛的应用。由IETF(Internet Engineering Task Force)实现的`TLS(Transport Layer Security)`是建立于SSL v3.0之上的兼容协议，它们主要的**区别**在于**所支持的加密算法**。

#### 6.6.1 加密通信流程
当用户在浏览器中输入一个以`https`开头的网址时，便开启了浏览器与被访问站点之间的加密通信。下面以一个用户访问[qbox.me](https://qbox.me)为例，给读者展现一下`SSL/TLS`的工作方式。
* 在浏览器输入HTTP协议的网址
* 服务器向浏览器返回证书，浏览器检查该证书的合法性
* 验证合法性
* 浏览器使用证书中的公钥加密一个随机对称密钥，并将加密后的密钥和使用密钥加密后的请求URL一起发送到服务器
* 服务器用私钥解密随机对称密钥，并用获取的密钥解密加密的请求URL
* 服务器把用户请求的网页用密钥加密，并返回给用户
* 用户浏览器用密钥解密服务器发来的网页数据，并将其展示出来

上述流程都是依赖`SSL/TLS`层实现的。在实际开发中,`SSL/TLS`的实现和工作原理比较复杂，但基本流程与上面的过程一致。

`SSL`协议由两层组成，
**上层协议**：
* 握手协议
* 更改密码规格协议
* 警报协议

**下层协议**:
* 记录协议

握手协议建立在记录协议之上，在实际的数据传输开始前，用于在客户与服务器之间进行“握手”。
“握手”是一个协商过程。这个协议使得客户端和服务器能够互相鉴别身份，协商加密算法。在任何数据传输之前，必须先进行“握手”。

在“握手”完成之后，才能进行**记录协议**，它的**主要功能**是为高层协议提供：
* 数据封装
* 压缩
* 添加MAC
* 加密
* ...

#### 6.6.2  支持HTTPS的Web服务器
Go语言目前实现了`TLS`协议的部分功能，已经可以提供最基础的安全层服务。以下代码为实现支持`TLS`的Web服务器。

[示例代码https.go](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/code/ChapterSix/6.6EncryptedCommunication/https.go)

[示例代码https2.go](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/code/ChapterSix/6.6EncryptedCommunication/https2.go)

#### 6.6.3 支持HTTPS的文件服务器
利用Go语言标准库中提供的完备封装，也可以很容易实现一个支持HTTPS的文件服务器，
[示例代码httpsfile.go](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/code/ChapterSix/6.6EncryptedCommunication/httpsfile.go)

#### 6.6.4 基于SSL/TLS的ECHO程序
在本章的最后，用一个完整的安全版ECHO程序来演示如何让Socket通信也支持HTTPS.
当然，ECHO程序支持HTTPS似乎没什么必要，但这个程序可以比较容易的改造成有实际价值的程序,
比如安全的聊天工具等.
下面首先实现这个超级ECHO程序的服务器端，[示例代码echoserver.go](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/code/ChapterSix/6.6EncryptedCommunication/echoserver.go)
再实现这个超级ECHO程序的客户端，[示例代码echoclient.go](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/code/ChapterSix/6.6EncryptedCommunication/echoclient.go)

需要注意的是，`SSL/TLS`协议只能运行于`TCP`之上，不能在`UDP`上工作，且`SSL/TLS`位于`TCP`与应用层协议之间,
因此所有基于`TCP`的应用层协议都可以透明的使用`SSL/TLS`为自己提供安全保障.
所谓透明的使用是指既不需要了解细节，也不需要专门处理该层的包，比如封装、解封等。

#### 6.7 小结
本章该要介绍了网络安全应用领域的相关知识点，以及GO语言对网络安全应用的全面支持,
同时还提供了具有一定使用价值的示例，让读者可以更加具体的理解相关的知识，并能基于
这些示例快速写出实用的程序。Go语言标准库的网络和加解密等相关的包在设计上都做了
一定程度的抽象，以大幅度提高易用性，提高开发效率。

因为本书的重点在于介绍Go语言的相关知识，所以对安全相关的知识就不做非常深入的诠释了，只是点到为止。
如果读者对安全编程有兴趣，可自行阅读网络安全的图书。