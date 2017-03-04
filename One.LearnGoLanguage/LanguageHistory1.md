### 1.1语言简史
	贝尔实验室.多位诺贝尔奖获得者.现代科技至关重要的研究成果(晶体管、通信技术、数码相机的感光元件CCD和光电池)
	计算科学研究中欣的部门对操作系统和编程语言的贡献:
>1969年,Ken Thompson&Dennis Ritchie开发出Unix,衍生出了C语言.
	20世纪80年代,Plan 9操作系统研究项目,目的为解决Unix中的问题，发展出一个后续替代系统.
	
		之后几十年中,该项目演变出Inferno项目分支和Limbo语言.
		
			Limbo语言是用于开发运行在小型计算机上的分布式应用的编程语言,支持模块化编程,编译期和运行时的强类型检查,进程内基于具有类型的通信通道,
			原子性垃圾收集和简单的抽象数据类型.被设计为:即便是在没有硬件内存保护的小型设备上,也能安全运行。
			
			Limbo语言被认为是Go语言的前身,同一批人设计,Go语言从Limbo语言中继承了众多优秀的特性。
			
		Plan 9原班人马加入Google,创造出Go语言.
		
>2007年9月,Go语言还是这帮大牛20%自由时间的实验项目
2008年5月,Google全力支持这个项目,全身心投入Go语言的设计和开发工作中.
2009年11月,正式对外发布,此后两年内快速迭代,发展迅猛.
2012年3月28日,发布第一个正式版本.
		
    开源方式发布,BSD授权协议.任何人可以查看Go语言的所有源代码,并可以为Go语言发展贡献自己的力量。
    Google组建了一个独立的小组全职开发Go语言,在服务中逐步增加对Go语言的支持(GAE Google AppEngine)


----------


**主要作者:**

[Ken Thompson](http://en.wikipedia.org/wiki/Ken_Thompson): 设计了B语言和C语言,创建了Unix和Plan 9操作系统,1983年图灵奖得主,Go语言共同作者.

[Rob Pike](http://en.wikipedia.org/wiki/Rob_Pike): Unix小组的成员,参与Plan 9和Inferno操作系统,参与Limbo和Go语言的研发,《Unix编程环境》作者之一

Robert Griesemer:曾协助制作Java和HotSpot编译器和Chrome浏览器的JavaScript引擎V8.

[Russ Cox](http://swtch.com/~rsc/): 参与Plan 9操作系统的开发,Google Code Search项目负责人

Ian Lance Taylor:GCC社区的活跃任务,gold连接器和GCC过程间优化LTO的主要设计者,Zembu公司创始人

[Brad Fitzpatrick](http://en.wikipedia.org/wiki/Brad_Fitzpatrick): LiveJournal的创始人,memcached作者

