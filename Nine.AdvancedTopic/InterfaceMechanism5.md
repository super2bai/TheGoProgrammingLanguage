### 9.5 接口机理
曾经深入研究过C++语言中的虚函数以及函数重载原理的读者，可能对于C++中引入的虚标和虚标指针还有深刻的印象。因为C++中并没有真正的接口，而只有纯虚函数和纯虚类，因此虚函数的原理就可以认为是C++版本的接口原理。要深入理解这些细节，需要认真读的书还是那本[《深度探索C++对象模型》](https://www.amazon.cn/深度探索C-对象模型-斯坦利-B-李普曼/dp/B006QXQXTM/ref=sr_1_1?ie=UTF8&qid=1488654328&sr=8-1&keywords=深度探索C%2B%2B对象模型)。总而言之，C++的整个接口机制是基本原理，非常简单，实现细节非常复杂，实现的功能非常强大，要全部掌握也就非常有难度。

我们已经在第三章中详细介绍了Go语言接口的特性和使用方法，本节中，将以尽量简洁明了的方式来解释Go语言这种"非侵入式"接口的实现原理。

接口的主要用法包含：
* 从类型赋值到接口
* 接口之间赋值
* 接口查询
* 等等

而原理剖析也会主要覆盖这几个功能。

读者可以[下载源码](https://github.com/Lynn--/gobook/tree/master/chapter9/interface)，对照本节理解Go语言的接口机制。

#### 9.5.1 类型赋值给接口
对于Go语言的使用者而言，Go语言接口的非侵入式具有相当的神秘色彩，比如先看这个最简单的接口[使用示例](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/code/ChapterNine/9.5InterfaceMechanism/interface-1.go)。

对于学过其他面向对象编程语言(比如C++)的读者而言，已经习惯了由明确的继承关系来确定类型和接口之间的关联，现在看到上述示例中`ISpeaker`和`SimpleSpeaker`没有任何的关联约定，就会产生疑惑，为什么编译器不报编译错误呢？很显然，Go语言采取了一个与C++等语言不通的机制。

一个核心的问题就是：从机器的角度如何判断一个`SimpleSpeaker`类型实现了`ISpeaker`接口的所有方法？一个简单的逻辑就是需要获取这个类型的所有方法集合(集合A)，并获取该接口包含的所有方法集合(集合B)，然后判断集合B是否为集合A的子集，是则意味着`SimpleSpeaker`类型实现了`ISpeaker`接口。

可以用以下的[数据结构](https://github.com/qiniu/gobook/blob/master/chapter9/interface/interface-3.c)来描述Go语言类型管理方法的方式：
```C
typedef struct _MemberInfo {
	const char* tag;
	void* addr;
} MemberInfo;

typedef struct _TypeInfo {
	MemberInfo* members;
} TypeInfo;
```
在以上的两个数据结构中，`_MemberInfo`结构体对应于一个具体的方法，将方法名和方法地址对应起来。而`_TypeInfo`对应一个类型，每个类型包含一个`_MemberInfo`类型的数组。

现在再列出接口的方法描述方式：
```C
typedef struct _InterfaceInfo {
	const char** tags;
} InterfaceInfo;

typedef struct _ITbl {
	InterfaceInfo* inter;
	TypeInfo* type;
	//...
} ITbl;
```
每个接口的数据结构都包含两个基本信息：
* 本接口的接口方法表(`InterfaceInfo`)
* 所指向的具体实现类型的类型信息(`TypeInfo`)

有了类型和接口的数据结构后，就可以回头定义出 `SimpleSpeaker`和`ISpeaker`的具体数据。

`ISpeaker`接口的底层表现如下：
```C
typedef struct _ISpeakerTbl { 	
	InterfaceInfo* inter;													      
	TypeInfo* type; 
	int (*Speak)(void* this); 
} ISpeakerTbl; 

typedef struct _ISpeaker {
	ISpeakerTbl* tab; 
	void* data; 
} ISpeaker; 

const char* g_Tags_ISpeaker[] = {  
	"Speaker()",	
	NULL 
};

InterfaceInfo g_InterfaceInfo_ISpeaker = { 
	g_Tags_ISpeaker
};
```
每个接口都会包含一个指向接口表的指针，而接口表将方法名和方法的调用地址对应起来。
下面是`SimpleSpeaker`类型的底层表达方法：
```C
typedef struct _SimpleSpeaker {
	 char Message[256];
} A;

void SimpleSpeaker_Speak(A* this) {		 
	printf("I am speaking... %s\n", this->Message) 
}

MemberInfo g_Members_SimpleSpeaker[] = {   
	{ "Speak()", SimpleSpeaker_Speak },
	{ NULL, NULL } 
}; 

TypeInfo g_TypeInfo_SimpleSpeaker = {			
	g_Members_SimpleSpeaker
};
```
现在可以很容易判断`SimpleSpeaker`是否实现了`ISpeaker`接口：只需要将`g_Members_SimpleSpeaker`数组和`g_Tags_ISpeaker`数组的内容进行字符串比对即可。因为两者都包含了完整名称为`Speak()`的方法，因此`SimpleSpeaker`实现了`ISpeaker`。

Go语言可以在编译器获取足够多的信息并进行代码的优化。比如对于这个类型赋值到接口的场景，编译器可以先谈哦那个过以上的逻辑判断是否该类型和该接口之间可以赋值，之后专门为`SimpleSpeaker`类型生成一个全局的`ISpeaker`接口表，具体如下所示：
```C
ISpeakerTbl g_Itbl_ISpeaker_SimpleSpeaker = {  
	&g_InterfaceInfo_ISpeaker,
	&g_TypeInfo_SimpleSpeaker, 
	(int (*)(void* this))SimpleSpeaker_Speak 
};
```
对于例子中这行类型到接口的赋值和调用语句：
```C
speaker = &SimpleSpeaker{"Hello"}
speaker.Speak() 
```
对应的底层实现会接近如下的写法
```C
// 这时候的SimpleSpeaker只是一个纯数据接口
SimpleSpeaker* unnamed = NewSimpleSpeaker("Hello"); ISpeaker p = { 
	&g_Itbl_ISpeaker_SimpleSpeaker, 
	unnamed 
}; 
p.tbl-&gt;Speak(p.data) 
```
可以看到，这种明确的可以在编译器确定的工作，就没必要到运行期进行动态的类型查询和转换。

为了让读者能够比较完整地理解这个过程，在这里再提供了一份完整可执行的代码，供读者运行并观察运行效果:
* [Go语言版本](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/code/ChapterNine/9.5InterfaceMechanism/interface-2.go)
* [C语言版本](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/code/ChapterNine/9.5InterfaceMechanism/interface-2.c)

#### 9.5.2 接口查询
接口查询是一个在软件开发中非常常见的使用场景，比如一个拿着`IReader`接口的开发者，在某些时候会需要知道`IReader`所对应的类型是否也实现了`IReadWriter`接口，这样它可以切换到`IReadWriter`接口，然后调用该接口的Write()`方法写入数据。

在Go语言的使用中，这个过程非常简单，具体代码如下：
```go
var reader IReader = NewReader()
if writer, ok := reader.(IReadWriter); ok{
	writer.Write()
}
```
那么到底接口查询是如何被支持的呢？
在9.5.1节中，已经大致介绍了Go语言可以采取的接口匹配流程。在使用接口查询的时候，这个机制可以派上用场了。

按Go语言的定义，接口查询其实是在做接口方法查询，只要该类型实现了某个接口的所有方法，就可以认为该类型实现了此接口。相比类型赋值给接口时可以做的编译器优化，运行期接口查询就只能老老实实的做一次接口匹配了。下面来看一下基本的匹配过程：
```C
typedef struct _ITbl {
	InterfaceInfo* inter;
	TypeInfo* type; 
	//... 
} ITbl; 

ITbl* MakeItbl(InterfaceInfo* intf, TypeInfo* ti) {
	size_t i, n = MemberCount(intf);
	ITbl* dest = (ITbl*)malloc(n * sizeof(void*) + sizeof(ITbl)); 
	void** addrs = (void**)(dest + 1);
	for (i = 0; i < n; i++) { 
		addrs[i] = MemberFind(ti, intf->tags[i]);
		if (addrs[i] == NULL) {
			free(dest);
			return NULL;
		}
	}
	dest->inter = intf;
	dest->type = ti;
	return dest;
}
```
这是一个动态的接口匹配过程。这个流程就是按接口信息表中包含的方法名逐一查询匹配，如果发现传入的类型信息`ti`的方法列表时`inft`的方法列表的超集(即`intf`方法列表中的所有方法都存在于`ti`方法列表中)，则表示接口查询成功。

从这个过程可以看到，整个过程其实跟发起查询的那个源接口毫无关系，真正的查询时针对源接口所指向的具体类型以及目标接口。因为这个过程比较简洁、易懂，这里就不再列出完整的示例代码。

#### 9.5.3 接口赋值
与接口查询相比，其实还有另外一种简单一些的场景，叫接口赋值，那就是将一个接口直接赋值给另外一个接口，比如：
```go
var rw IReadWriter = ...
var r IReader = rw
```
这种赋值是否可以通过编译的判断依据时源接口和目标接口是否存在方法集合的包含关系。因为`IReadWriter`包含了`IReader`的所有方法，所以这种赋值过程时合法的。但是不能直接将·`IReader`接口赋值给`IReadWriter`，如果需要这种转换，就得用接口查询。

接口查询初看起来和我们描述的接口查询过程有些像，但因为接口赋值过程在编译器就可以确定，所以没必要动用消耗比较大的动态接口查询流程。可以认为接口赋值时接口查询的一种优化。在编译期，编译器就能判断是否可进行接口转换。如果可转换，编译器将为所有用到的接口赋值，生成各自的赋值函数：
```C
 IWriterTbl* Itbl_IWriter_From_IReadWriter(IReadWriterTbl* src) { 
	IWriterTbl* dest =(IWriterTbl*)malloc(sizeof(IWriterTbl)); 
	dest->inter = &g_InterfaceInfo_IWriter,
	dest->type = src->type; 
	dest->Write = src->Write; 
	return dest; 
}
```
这段代码没有做是否可以从`IReadWriter`接口转换到`IWriter`接口的判断，因为这是编译器在生成这个函数之前应该做的编译期动作。相关内容之前已经解释过，只需将这两个接口的方法集进行对比即可。

此时，关于接口机理的介绍就完成了。需要说明的是，介绍的只是其中一种可实现的途径，还存在大量其他的实现方法。如果读者有更好的想法，或者对本节有任何建议或问题，都欢迎与我们讨论。

