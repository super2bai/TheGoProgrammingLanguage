### 6.6 加密通信

### 6.6 加密通信
一般的`HTTPS`是基于`SSL(Secure Sockets Layer)`协议。`SSL`是网景公司开发的位于`TCP`与`HTTP`之间的透明安全协议，通过`SSL`，可以把`HTTP`包数据以非对称加密的形式往返于浏览器和站点之间，从而避免被第三方非法获取。

目前，伴随着电子商务的兴起，`HTTPS`获得了广泛的应用。由IETF(Internet Engineering Task Force)实现的`TLS(Transport Layer Security)`是建立于SSL v3.0之上的兼容协议，它们主要的**区别**在于**所支持的加密算法**。

#### 6.6.1 加密通信流程
当用户在浏览器中输入一个以`https`开头的网址时，便开启了浏览器与被访问站点之间的加密通信。下面以一个用户访问https://qbox.me为例，给读者展现一下`SSL/TLS`的工作方式。
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