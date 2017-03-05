# TheGoProgrammingLanguage
**《Go语言编程》---许式伟**读书纪录

[Go语言编程github](https://github.com/qiniu/gobook)

[Go语言编程书籍](https://www.amazon.cn/dp/B00932YRPA/ref=sr_1_1?ie=UTF8&qid=1486914755&sr=8-1&keywords=GO语言编程)

**环境**：

[MAC](http://www.apple.com/cn/mac)

[LiteIDE源码](https://github.com/visualfc/liteide)

[LiteIDE下载](http://www.golangtc.com/download/liteide)

[马克飞象](https://maxiang.io)

补充：
* 原书中有一些错误
* 必须操作但未在书中说明

### 目录
* [前言](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/Introduction/introduction.md)
* [第一章 初识Go语言](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/One.LearnGoLanguage)
	* [1.1 语言简史](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/One.LearnGoLanguage/LanguageHistory1.md)
	* [1.2 语言特性](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/One.LearnGoLanguage/LanguageFeatures2.md)
		* 1.2.1 自动垃圾回收
		* 1.2.2 更丰富的那只类型
		* 1.2.3 函数多返回值
		* 1.2.4 错误处理
		* 1.2.5 匿名函数和闭包
		* 1.2.6 类型和接口
		* 1.2.7 并发编程
		* 1.2.8 反射
		* 1.2.9 语言交互性
   * [1.3 第一个Go程序](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/One.LearnGoLanguage/FirstGoProgram3.md)
		* 1.3.1 代码解读
		* 1.3.2 编译环境准备
		* 1.3.3 编译程序
    * [1.4 开发工具的选择](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/One.LearnGoLanguage/ChooseTools4.md)
    * [1.5工程管理](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/One.LearnGoLanguage/EngineeringManagement5.md)
    * [1.6 问题追踪和调试](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/One.LearnGoLanguage/ProblemTrackingAndDebugging6.md)
		* 1.6.1 打印日志
		* 1.6.2 GDB调试
    * [1.7 如何寻求帮助](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/One.LearnGoLanguage/AskForHelp7.md)
	    * 1.7.1 邮件列表
	    * 1.7.2 网站资源
	* [1.8 小结](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/One.LearnGoLanguage/Summary8.md)
* [第二章 顺序编程](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Two.SequentialProgramming)
	* [2.1 变量](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Two.SequentialProgramming/Variable1.md)
		* 2.1.1 变量声明
		* 2.1.2 变量初始化
		* 2.1.3 变量赋值
		* 2.1.4 匿名变量
	* [2.2 常量](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Two.SequentialProgramming/Constant2.md)
		* 2.2.1 字面常量
		* 2.2.2 常量定义
		* 2.2.3 预定义常量
		* 2.2.4 枚举	
	* [2.3 类型](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Two.SequentialProgramming/Type3.md)
		* 2.3.1 布尔类型
		* 2.3.2 整型
		* 2.3.3 浮点型
		* 2.3.4 复数类型
		* 2.3.5 字符串
		* 2.3.6 自负类型
		* 2.3.7 数组
		* 2.3.8 数组切片
		* 2.3.9 map	
	* [2.4 流程控制](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Two.SequentialProgramming/ControlFlow4.md)
		* 2.4.1 条件语句
		* 2.4.2 选择语句
		* 2.4.3 循环语句
		* 2.4.4 跳转语句	
	* [2.5 函数](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Two.SequentialProgramming/Fuction5.md)
		* 2.5.1 函数定义
		* 2.5.2 函数调用
		* 2.5.3 不定参数
		* 2.5.4 多返回值
		* 2.5.5 匿名函数与闭包		
	* [2.6 错误处理](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Two.SequentialProgramming/ErrorHandling6.md)
		* 2.6.1 error接口
		* 2.6.2 defer
		* 2.6.3 panic()和recover()
	* [2.7 完整示例](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Two.SequentialProgramming/CompleteExample7.md)
		* 2.7.1 程序结构
		* 2.7.2 主程序
		* 2.7.3 算法实现
		* 2.7.4 主程序
		* 2.7.5 构建
	* [2.8 小结](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Two.SequentialProgramming/Summary8.md)
* [第三章 面向对象编程](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Three.ObjectOrientedProgramming)
	* [3.1  类型系统](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Three.ObjectOrientedProgramming/TypeSystem1.md)
		* 3.1.1 为类型添加方法
		* 3.1.2 值语义和引用语义
		* 3.1.3 结构体
	* [3.2 初始化](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Three.ObjectOrientedProgramming/Initialization2.md)
	* [3.3 匿名组合](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Three.ObjectOrientedProgramming/AnonymousCombination3.md)
	* [3.4 可见性](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Three.ObjectOrientedProgramming/Visibility4.md)
	* [3.5 接口](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Three.ObjectOrientedProgramming/Interface5.md)
		* 3.5.1 其他语言的接口
		* 3.5.2 非侵入式接口
		* 3.5.3 接口赋值
		* 3.5.4 接口查询
		* 3.5.5 类型查询
		* 3.5.6 接口组合
		* 3.5.7 Any类型
	* [3.6 完整示例](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Three.ObjectOrientedProgramming/CompleteExample6.md)
		* 3.6.1 音乐库
		* 3.6.2 音乐播放
		* 3.6.3 主程序
		* 3.6.4 构建运行
		* 3.6.5 遗留问题
	* [3.7 小结](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Three.ObjectOrientedProgramming)		
* [第四章 并发编程](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Four.ConcurrentProgramming)
	* [4.1 并发基础](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Four.ConcurrentProgramming/ConcurrentBasis1.md)
	* [4.2 协程](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Four.ConcurrentProgramming/Routine2.md)
	* [4.3 goroutine](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Four.ConcurrentProgramming/Goroutine3.md)
	* [4.4 并发通信](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Four.ConcurrentProgramming/ConcurrentCommunication4.md)
	* [4.5 channel](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Four.ConcurrentProgramming/Channel5.md)
		* 4.5.1 基本语法
		* 4.5.2 select
		* 4.5.3缓冲机制
		* 4.5.4 超时机制
		* 4.5.5 channel的传递
		* 4.5.6 单向channel
		* 4.5.7 关闭channel
	* [4.6 多核并行化](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Four.ConcurrentProgramming/Multi-coreParallelization6.md)
	* [4.7 出让时间片](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Four.ConcurrentProgramming/TransferTime7.md)
	* [4.8 同步](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Four.ConcurrentProgramming/)
		* 4.8.1 同步锁
		* 4.8.2 全局唯一性操作
	* [4.9 完整示例](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Four.ConcurrentProgramming/Synchronization8.md)
		* 4.9.1 简单IPC狂减
		* 4.9.2 中央服务器
		* 4.9.3 主程序
		* 4.9.4 运行程序
	* [4.10 小结](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Four.ConcurrentProgramming/Summary10.md)
* [第五章 网络编程](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Five.NetworkProgramming/)
	* [5.1 Socket编程](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Five.NetworkProgramming/Socket1.md)
		* 5.1.1 Dial函数
		* 5.1.2 ICMP示例程序
		* 5.1.3 TCP示例程序
		* 5.1.4 更丰富的网络通信
	* [5.2 HTTP 编程](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Five.NetworkProgramming/HTTP2.md)
		* 5.2.1 HTTP客户端
		* 5.2.2 HTTP服务端
	* [5.3 RPC 编程](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Five.NetworkProgramming/RPC3.md)
		* 5.3.1 Go语言中的RPC支持与处理
		* 5.3.2 Gob简介
		* 5.3.3 设计优雅的RPC接口
	* [5.4 JSON处理](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Five.NetworkProgramming/JSON4.md)
		* 5.4.1 编码为JSON格式
		* 5.4.2 编码JSON数据
		* 5.4.3 解码未知结构的JSON数据
		* 5.4.4 JSON的流式读写
	* [5.5 网站开发](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Five.NetworkProgramming/WebSiteDevelopment5.md)
		* 5.5.1 最简单的网站程序
		* 5.5.2 net/http包简介
		* 5.5.3 开发一个简单的相册网站
	* [5.6 小结](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Five.NetworkProgramming/Summary6.md)	
* [第六章 安全编程](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Six.SecurityProgramming/)
	* [6.1 数字加密](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Six.SecurityProgramming/DataEncryption1.md)
	* [6.2 数字签名](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Six.SecurityProgramming/DigitalSignature2.md)
	* [6.3 数字证书](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Six.SecurityProgramming/DigitalCertificate3.md)
	* [6.4 PKI体系](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Six.SecurityProgramming/PKISystem4.md)
	* [6.5 加密通信](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Six.SecurityProgramming/HashFunctionForGo5.md)
	* [6.6 加密通信](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Six.SecurityProgramming/EncryptedCommunication6.md)
		* 6.6.1 加密通信流程
		* 6.6.2 支持HTTPS的Web服务器
		* 6.6.3 支持HTTPS的文件服务器
		* 6.6.4 基于SSL/TLS的ECHO程序
	* [6.7 小结](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Six.SecurityProgramming/Summary7.md)	
* [第七章 工程管理](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Seven.EngineeringManagement)
	* [7.1 Go命令行工具](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Seven.EngineeringManagement/CommandLineTool1.md)
	* [7.2 代码风格](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Seven.EngineeringManagement/CodeStyle2.md)
		* 7.2.1 强制性编码规范
		* 7.2.2 非强制性编码规范
	* [7.3 远程import支持](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Seven.EngineeringManagement/RemoteImportSupport3.md)
	* [7.4 工程组织](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Seven.EngineeringManagement/EngineeringOrganization4.md)
		* 7.4.1 GOPATH
		* 7.4.2 目录结构
	* [7.5 文档管理](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Seven.EngineeringManagement/DocumentManagement5.md)
	* [7.6 工程构建](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Seven.EngineeringManagement/ConstructionOfTheProject6.md)
	* [7.7 跨平台开发](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Seven.EngineeringManagement/CrossPlatformDevelopment7.md)
		* 7.7.1 交叉编译
		* 7.7.2 Android支持
	* [7.8 单元测试](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Seven.EngineeringManagement/UnitTest8.md)
	* [7.9 打包分发](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Seven.EngineeringManagement/PackageDistribution9.md)
	* [7.10 小结](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Seven.EngineeringManagement/Summary10.md)
* [第八章 开发工具](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Eight.DevelopmentTools)
	* [8.1 选择开发工具](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Eight.DevelopmentTools/ChooseDevelopmentTool1.md)
	* [8.2 gedit](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Eight.DevelopmentTools/Gedit2.md)
		* 8.2.1 语法高亮
		* 8.2.2 编译环境
	* [8.3 Vim](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Eight.DevelopmentTools/Vim3.md)
	* [8.4 Eclipse](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Eight.DevelopmentTools/Eclipse4.md)
	* [8.5 Notepad++](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Eight.DevelopmentTools/NotePad++5.md)
		* 8.5.1 语法高亮
		* 8.5.2 编译环境
	* [8.6 LiteIDE](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Eight.DevelopmentTools/LiteIDE6.md)
	* [8.7 小节]	(https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Eight.DevelopmentTools/Summary7.md)
* [第九章 进阶话题](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Nine.AdvancedTopic)
	* [9.1 反射](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Nine.AdvancedTopic/Reflection1.md)
		* 9.1.1 基本概念
		* 9.1.2 基本用法
		* 9.1.3 对结构的反射操作
	* [9.2 语言交互性](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Nine.AdvancedTopic/LanguageInteractivity2.md)
		* 9.2.1 类型映射
		* 9.2.2 字符串映射
		* 9.2.3 C程序
		* 9.2.4 函数调用
		* 9.2.5 编译Cgo
	* [9.3 链接符号](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Nine.AdvancedTopic/LinkSymbol3.md)
	* [9.4 goroutine 机理](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Nine.AdvancedTopic/GoroutineMechanism4.md)
		* 9.4.1 协程
		* 9.4.2 协程的C语言实现
		* 9.4.3 协程库概述
		* 9.4.4 任务
		* 9.4.5 任务调度
		* 9.4.6 上下文切换
		* 9.4.7 通信机制
	* [9.5 接口机理](https://github.com/Lynn--/TheGoProgrammingLanguage/tree/master/Nine.AdvancedTopic/InterfaceMechanism5.md)
		* 9.5.1 类型赋值给接口
		* 9.5.2 接口查询
		* 9.5.3 接口赋值
* [附录 A](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/Appendix/A.1GoLanguageStandardLibrary.md)