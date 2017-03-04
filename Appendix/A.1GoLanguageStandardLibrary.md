###A.1 Go语言标准库
一门语言是否能够比较快地受到开发者的欢迎，除了语法特性外，语言所附带的标准库的功能完整性和易用性也是一个非常重要的评判标准。假如Java只有一个编译器而没有JDK，C#没有对应的.NET Framework，那么很难想象这两门语言可以流行。

一个优秀的标准库应该能够解决大部分开发需求，只在极少情况下，比如解决比较专业的问 题或者特别复杂的问题时，才需要依赖第三方库。

Go语言的发布版本附带了一个非常强大的标准库。如果能够快速定位相应的功能，开发者的幸福感会大大提高。我们希望本章内容能够帮助学习Go语言的读者尽量快速定位到相应的包。

不过归根到底，学习新事物还是一回生二回熟，希望读者在学习Go语言的过程中遇到任何 问题，都能够保持足够的耐心，在解决一个又一个问题的过程当中，发现越来越多Go语言的可爱之处。Go语言标准库为我们提供了源代码，且所有的包都有单元测试案例。我们在查看Go语言标准库文档时，可以随时单击库里的函数名跳转到对应的源代码。这些源代码具备相当高的参考价值，平时多看看对提高自己的Go语言开发水平会大有裨益。

Go标准库可以大致按其中库的功能进行以下分类，这个分类比较简单，不求准确，但求能够帮助开发者根据自己模糊的需求更快找到自己需要的包。
* **输入输出**。这个分类包括二进制以及文本格式在屏幕、键盘、文件以及其他设备上的输入输出等，比如二进制文件的读写。对应于此分类的包有
	*  `bufio`
	*  `fmt`
	*  `io`
	*  `log`
	*  `flag`(用于处理命令行参数)
	*  ... 
* **文本处理**。这个分类包括字符串和文本内容的处理，比如字符编码转换等。对应于此分类的包有
	* `encoding`
	* `bytes`
	* `strings`
	* `strconv`
	* `text`
	* `mime`
	* `unicode`
	* `regexp`
	* `index`
	* `path`(用于处理路径字符串)
	* ...
* **网络**。这个分类包括开发网络程序所需要的包，比如`Socket`编程和网站开发等。对应于此分类的包有：
	*  `net`
	*  `http`
	*  `expvar`
	*  ...
* **系统**。这个分类包含对系统功能的封装，比如对操作系统的交互以及原子性操作等。对应于此分类的包有
	* `os`
	* `syscall`
	* `sync`
	* `time`
	* `unsafe`
	* ...
* **数据结构与算法**。对应于此分类的包有
	* `math`
	* `sort`
	* `container`
	* `crypto`
	* `hash`
	* `archive`
	* `compress`
	* `image`(因为提供的图像编解码都是算法，所以也归入此类)
	* ...
* **运行时**。对应于此分类的包有：
	* `runtime`
	* `reflect`
	* `go`
	* ...

####A.1.1 常用包介绍
本节我们介绍Go语言标准库里使用频率相对较高的一些包。熟悉了这些包后，使用Go语言开发一些常规的程序将会事半功倍。
 * **fmt**。它实现了格式化的输入输出操作，其中的`fmt.Printf()`和`fmt.Println()`是开发者使用最为频繁的函数。
 * **io**。它实现了一系列非平台相关的IO相关接口和实现，比如提供了对os中系统相关的IO功能的封装。我们在进行流式读写（比如读写文件）时，通常会用到该包。
 *  **bufio**。它在io的基础上提供了缓存功能。在具备了缓存功能后，`bufio`可以比较方便地提供`ReadLine`之类的操作。
 *  **strconv**。本包提供字符串与基本数据类型互转的能力。
 *  **os**。本包提供了对操作系统功能的非平台相关访问接口。接口为Unix风格。提供的功能包括文件操作、进程管理、信号和用户账号等。
 *  **sync**。它提供了基本的同步原语。在多个goroutine访问共享资源的时候，需要使用sync 中提供的锁机制。
 *  **flag**。它提供命令行参数的规则定义和传入参数解析的功能。绝大部分的命令行程序都 需要用到这个包。
 *  **encoding/json**。JSON目前广泛用做网络程序中的通信格式。本包提供了对JSON的基本支持，比如从一个对象序列化为JSON字符串，或者从JSON字符串反序列化出一个具体的对象等。
 *  **http**。它是一个强大而易用的包，也是Golang语言是一门“互联网语言”的最好佐证。通过`http`包，只需要数行代码，即可实现一个爬虫或者一个Web服务器，这在传统语言中 是无法想象的。 
 
#### A.1.2 完整包列表 

完整的包列表见表A-1。

 目 录 |包| 概 述|
----|------|----
&nbsp; |bufio|实现缓冲的I/O 
&nbsp; |bytes|提供了对字节切片操作的函数
&nbsp; |crypto|收集了常见的加密常数
&nbsp;|errors|实现了操作错误的函数
&nbsp;|Expvar|为公共变量提供了一个标准的接口，如服务器中的运算计数器
&nbsp;|flag|实现了命令行标记解析
&nbsp;|fmt|实现了格式化输入输出
&nbsp;|hash|提供了哈希函数接口
&nbsp;|html|实现了一个HTML5兼容的分词器和解析器
&nbsp;|image|实现了一个基本的二维图像库 
&nbsp;|io|提供了对I/O原语的基本接口
&nbsp;|log|它是一个简单的记录包，提供最基本的日志功能
&nbsp;|math|提供了一些基本的常量和数学函数
&nbsp;|mine|实现了部分的MIME规范
&nbsp;|net|提供了一个对UNIX网络套接字的可移植接口，包括TCP/IP、UDP域名解析和UNIX域套接字
&nbsp;|os|为操作系统功能实现了一个平台无关的接口
&nbsp;|path|实现了对斜线分割的文件名路径的操作
&nbsp;|reflect|实现了运行时反射，允许一个程序以任意类型操作对象
&nbsp;|regexp|实现了一个简单的正则表达式库
&nbsp;|runtime|包含与Go运行时系统交互的操作，如控制`goroutine`的函数
&nbsp;|sort|提供对集合排序的基础函数集
&nbsp;|strconv|实现了在基本数据类型和字符串之间的转换
&nbsp;|strings| 实现了操作字符串的简单函数
&nbsp;|sync| 提供了基本的同步机制，如互斥锁 
&nbsp;|syscall|包含一个低级的操作系统原语的接口
&nbsp;|testing|提供对自动测试Go包的支持
&nbsp;|time|提供测量和显示时间的功能
&nbsp;|unicode|Unicode编码相关的基础函数
archive| tar |实现对tar压缩文档的访问 
archive|zip|提供对ZIP压缩文档的读和写支持
compress|bzip2|实现了bzip2解压缩
compress|flate|实现了RFC 1951中所定义的DEFLATE压缩数据格式 compress|gzip|实现了RFC 1951中所定义的gzip格式压缩文件的读和写
compress| lzw|实 现了Lempel-Ziv-Welch编码格式的压缩的数据格 式 ，参见T. A. Welch, A Technique for High-Performance Data Compression, Computer, 17(6) (June 1984), pp 8-19
compress|zlib|实现了RFC 1950中所定义的zlib格式压缩数据的读和写
container|heap|提供了实现heap.Interface接口的任何类型的堆操作 container|list|实现了一个双链表
container|ring|实现了对循环链表的操作
crypto|aes|实现了AES加密（以前的Rijndael），详见美国联邦信息处理标准（197号文）
crypto|cipher|实现了标准的密码块模式，该模式可包装进低级的块加密实现中
crypto|des|实现了数据加密标准（Data Encryption Standard，DES）和三重数据加密算法（Triple Data Encryption Algorithm，TDEA），详见美国联邦信息处理标准（46-3号文） 
crypto|dsa|实现了FIPS 186-3所定义的数据签名算法（Digital Signature Algorithm）
crypto|ecdsa|实现了FIPS 186-3所定义的椭圆曲线数据签名算法（Elliptic Curve Digital Signature Algorithm）
crypto|elliptic|实现了素数域上几个标准的椭圆曲线
crypto|hmac|实现了键控哈希消息身份验证码（Keyed-Hash Message Authentication Code， HMAC），详见美国联邦信息处理标准（198号文） 
crypto|md5|实现了RFC 1321中所定义的MD5哈希算法
crypto|rand|实现了一个加密安全的伪随机数生成器
crypto|rc4|实现了RC4加密，其定义见Bruce Schneier的应用密码学（Applied Cryptography） 
crypto|rsa|实现了PKCS#1中所定义的RSA加密
crypto|sha1|实现了RFC 3174中所定义的SHA1哈希算法
crypto|sha256|实现了FIPS 180-2中所定义的SHA224和SHA256哈希算法
crypto|sha512| 实现了FIPS 180-2中所定义的SHA384和SHA512哈希算法
crypto|subtle| 实现了一些有用的加密函数，但需要仔细考虑以便正确应用它们
crypto|tls| 部分实现了RFC 4346所定义的TLS 1.1协议 
crypto|x509| 可解析X.509编码的键值和证书 
crypto|x509/pkix| 包含用于对X.509证书、CRL和OCSP的ASN.1解析和序列化的共享的、低级的结构
database|sql|围绕SQL提供了一个通用的接口 
database|sql/driver|定义了数据库驱动所需实现的接口，同sql包的使用方式 
debug| dwarf| 提供了对从可执行文件加载的DWARF调试信息的访问，这个包对于实现Go语言的调试器非常有价值 
debug|elf| 实现了对ELF对象文件的访问。ELF是一种常见的二进制可执行文件和共享库的 文件格式。Linux采用了ELF格式
debug|gosym| 访问Go语言二进制程序中的调试信息。对于可视化调试很有价值
debug|macho| 实现了对http://developer.apple.com/mac/library/documentation/DeveloperTools/Conceptual/MachORuntime/Reference/reference.html 所定义的Mach-O对象文件的访问
debug|pe|实现了对PE（Microsoft Windows Portable Executable）文件的访问
encoding| ascii85| 实现了ascii85数据编码，用于btoa工具和Adobe’s PostScript以及PDF文档格式 
encoding|asn1| 实现了解析DER编码的ASN.1数据结构，其定义见ITU-T Rec X.690 
encoding|base32| 实现了RFC 4648中所定义的base32编码
encoding|base64|实现了RFC 4648中所定义的base64编码 
encoding|binary| 实现了在无符号整数值和字节串之间的转化，以及对固定尺寸值的读和写
encoding|csv| 可读和写由逗号分割的数值（csv）文件 
encoding|gob| 管理gob流——在编码器（发送者）和解码器（接收者）之间进行二进制值交换
encoding|hex| 实现了十六进制的编码和解码
encoding| json| 实现了定义于RFC 4627中的JSON对象的编码和解码
encoding|pem| 实现了PEM（Privacy Enhanced Mail）数据编码
encoding|xml|实现了一个简单的可理解XML名字空间的XML 1.0解析器
go| ast| 声明了用于展示Go包中的语法树类型go|build| 提供了构建Go包的工具 
go|doc| 从一个Go AST（抽象语法树）中提取源代码文档 
go|parser| 实现了一个Go源文件解析器
go| printer| 实现了对AST（抽象语法树）的打印 go|scanner| 实现了一个Go源代码文本的扫描器 go|token| 定义了代表Go编程语言中词法标记以及基本操作标记（printing、predicates）的常量 hash |adler32| 实现了Adler-32校验和hash|crc32| 实现了32位的循环冗余校验或CRC-32校验和 
hash|crc64| 实现了64位的循环冗余校验或CRC-64校验和 
hash|fnv| 实现了Glenn Fowler、Landon Curt Noll和Phong Vo所创建的FNV-1和FNV-1a未加 密哈希函数 
html| template| 它自动构建HTML输出，并可防止代码注入
image| color| 实现了一个基本的颜色库
image| draw| 提供一些做图函数
image| gif| 实现了一个GIF图像解码器 
image|jpeg| 实现了一个JPEG图像解码器和编码器 
image|png| 实现了一个PNG图像解码器和编码器 
index|suffixarray|通过构建内存索引实现的高速字符串匹配查找算法
io| ioutil| 实现了一些实用的I/O函数
log| syslog| 提供了对系统日志服务的简单接口 
Math|big|实现了多精度的算术运算（大数）
Math|cmplx| 为复数提供了基本的常量和数学函数 Math|rand| 实现了伪随机数生成器
mime| multipart| 实现了在RFC 2046中定义的MIME多个部分的解析
net| http| 提供了HTTP客户端和服务器的实现 net|mail| 实现了对邮件消息的解析
net|rpc| 提供了对一个来自网络或其他I/O连接的对象可导出的方法的访问 
net|smtp| 实现了定义于RFC 5321中的简单邮件传输协议（Simple Mail Transfer Protocol) net|textproto| 实现了在HTTP、NNTP和SMTP中基于文本的通用的请求/响应协议
net| url| 解析URL并实现查询转义 
net|http/cgi| 实现了定义于RFC 3875中的CGI（通用网关接口） 
net|http/fcgi| 实现了FastCGI协议 net|http/httptest| 提供了一些HTTP测试应用 net|http/httputil| 提供了一些HTTP应用函数，这些是对net/http包中的东西的补充，只不过相对 不太常用 
net|http/pprof |通过其HTTP服务器运行时提供性能测试数据，该数据的格式正是pprof可视化工 具需要的 
net|rpc/jsonrpc |为rpc包实现了一个JSON-RPC ClientCodec和ServerCodec 
os| exec| 可运行外部命令
os| user| 通过名称和id进行用户账户检查
path| filepath| 实现了以与目标操作系统定义文件路径相兼容的方式处理文件名路径 
regexp| syntax| 将正则表达式解析为语法树
runtime| debug| 包含当程序在运行时调试其自身的功能 
runtime|pprof| 以pprof可视化工具需要的格式写运行时性能测试数据
sync| atomic| 提供了低级的用于实现同步算法的原子级的内存机制 
testing| iotest| 提供一系列测试目的的类型，实现了Reader和Writer标准接口 
testing|quick| 实现了用于黑箱测试的实用函数
testing|script| 帮助测试使用通道的代码
text| scanner |为UTF-8文本提供了一个扫描器和分词器 
text|tabwriter| 实现了一个写筛选器（tabwriter.Writer），它可将一个输入的tab分割的列 翻译为适当对齐的文本
text| template| 数据驱动的模板引擎，用于生成类似HTML的文本输出格式 
text|template/parse| 为template构建解析树
text| unicode/utf16| 实现了UTF-16序列的的编码和解码 
text|unicode/utf8 |实现了支持以UTF-8编码的文本的函数和常数