### 9.4 `goroutine`机理
我们在第四章中已经详细介绍了如何使用`goroutine`编写各种并发程序，并介绍了该Go语言特性的强大之处。从根本上来说`goroutine`就是一种Go语言版本的协程(coroutine)。因此，要理解`goroutine`的运作机理，关键就是理解传统意义上协程的工作机理。此处，本节标题也可以改名为"协程机理"，因为它并不专门针对Go语言。

回头看看，协程这个术语应该是随着Lua语言的流行而流行起来的，但要刨根究底的话，协程第一次出现在1963年，用于汇编编程。最先实现了协程的语言应该是Simula和Modula-2(恐怕已经没多少读者知道这两门语言到底是怎么回事)。Lua和Go语言则可以算是近几年在语言层面支持协程的典型代表，但实际上支持协程的语言有三四十种之多，比如C#也在内部支持协程。因为本节不是谈协程轶事，所以关于协程的历史细节不再展开，有兴趣的读者可以自己去危机百科上查看。

#### 9.4.1 协程
协程，也有人称之为轻量级线程，具备以下几个特点：
* 能够在单一的系统线程中模拟多个任务的并发执行
* 在一个特定的时间，只有一个任务在运行，即并非真正的并行
* 被动的任务调度方式，即任务没有主动强占时间片的说法。当一个任务正在执行时，外部没有办法中止它。要进行任务切换，只能通过由该任务自身调用`yield()`来主动出让CPU使用权
* 每个协程都有自己的堆栈和局部变量

每个协程都包含3种运行状态：
* **运行**
* **挂起**则表示该协程尚未执行完成，但出让了时间片，以后有机会时会由调度器继续执行
* **停止**通常表示该协程已经执行完成(包括遇到问题后明确推出执行的情况)

#### 9.4.2 协程的C语言实现
为了更好的剖析协程的运行原理，在本节中将引入Go语言的作者之一拉斯考克斯(Russ Cox)在Go语言出世之前所设计实现的一个轻量级协程库[libtask](http://swtch.com/libtask)，可以自行到该页面下载完整的源代码。这个库的作者使用的是非常开放的授权协议，因此可以随意修改和使用这些代码，但必须保持该份代码所附带的版权声明。

虽然没有具体地比对`goroutine`实现代码和`libtask`的直接关系，但有足够充分的理由相信`goroutine`和用于`goroutine`之间通信的`channel`都是参照`libtask`库实现的(甚至可能直接使用这个库)。至于`go`关键字等Go语言特性，都可以将其认为只是为了便于开发者使用而设计的语法糖。

本节将对这个代码库做一次结构化的阅读，并在必要的地方贴出一些关键的代码段。相信读者在阅读完本节后，对于协程的原理会有比较全面的理解。理解了协程的概念，对于中缺使用Go语言的`goroutine`以及分析使用`goroutine`时遇到的各种问题都会大有帮助。

#### 9.4.3 协程库概述
这个`libtask`库实现了以下几个关键模块：
* 任务即任务管理
* 任务调度器
* 异步IO
* `channel`

这个静态库直接提供了一个`main()`入口函数作为协程的驱动，因此库的使用者只需按该库约定的规则实现任务函数`taskmain()`，启动后这些任务自然会被以协程的方式创建和调度执行。`taskmain()`函数的声明如下：
```C
void taskmain(int argc, char *argv[]);
```
在分析代码之前，可以先看一下例子[primes.c](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/code/ChapterNine/9.4LinkSymbol/primes.c)，该程序从命令行得到一个整型数作为质数的查找范围，比如用户输入了100，则该程序会列出0到100之间的所有质数。

下面讲这个C程序翻译为对应的[primes.go](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/code/ChapterNine/9.4LinkSymbol/primes.go)，让读者可以比较容易的理解这个例子。

两个程序的执行结果完全一致，会打印出2到100之间的所有质数。读者可以对比阅读者两份代码，从而大致了解`libtask`中对应于Go语言各种概念的实现方法。

#### 9.4.4 任务
从上面的例子可以看出，在实现了一个任务函数后，真要让这个函数假如到调度队列中，我们还需要显式调用`taskcreate()`函数。下面大致介绍以下任务的概念，以及`taskcreate()`到底做了哪些事情。

任务用以下的结果表达：
```C
struct Task
{
	char name[256];
	char state[256];
	Task *next;
	Task *prev;
	Task *allnext;
	Task *allprev;
	Context context;
	uvlong alarmtime;
	uint id;
	uchar *sdk;
	unit stksize;
	int exiting;
	int alltaskslot;
	int system;
	int ready;
	void (*startfn)(void*);
	void *startarg;
	void *udata;
};
```
可以看到，每一个任务需要保存以下者几个关键数据：
* 任务上下文，用于在切换任务时保持当前任务的运行环境
* 栈
* 状态
* 该任务所对应的业务函数(9.4.3节中的·`rimetask()`函数)
* 任务的调用参数
* 之前和之后的任务

下面再来看以下[任务](https://swtch.com/libtask/task.c)的创建过程：
```C
static int taskidgen;

static Task*
taskalloc(void (*fn)(void*), void *arg, uint stack)
{
	Task *t;
	sigset_t zero;
	uint x, y;
	ulong z;

	/* allocate the task and stack together */
	t = malloc(sizeof *t+stack);
	if(t == nil){
		fprint(2, "taskalloc malloc: %r\n");
		abort();
	}
	memset(t, 0, sizeof *t);
	t->stk = (uchar*)(t+1);
	t->stksize = stack;
	t->id = ++taskidgen;
	t->startfn = fn;
	t->startarg = arg;

	/* do a reasonable initialization */
	memset(&t->context.uc, 0, sizeof t->context.uc);
	sigemptyset(&zero);
	sigprocmask(SIG_BLOCK, &zero, &t->context.uc.uc_sigmask);

	/* must initialize with current context */
	if(getcontext(&t->context.uc) < 0){
		fprint(2, "getcontext: %r\n");
		abort();
	}

	/* call makecontext to do the real work. */
	/* leave a few words open on both ends */
	t->context.uc.uc_stack.ss_sp = t->stk+8;
	t->context.uc.uc_stack.ss_size = t->stksize-64;
#if defined(__sun__) && !defined(__MAKECONTEXT_V2_SOURCE)		/* sigh */
#warning "doing sun thing"
	/* can avoid this with __MAKECONTEXT_V2_SOURCE but only on SunOS 5.9 */
	t->context.uc.uc_stack.ss_sp = 
		(char*)t->context.uc.uc_stack.ss_sp
		+t->context.uc.uc_stack.ss_size;
#endif
	/*
	 * All this magic is because you have to pass makecontext a
	 * function that takes some number of word-sized variables,
	 * and on 64-bit machines pointers are bigger than words.
	 */
//print("make %p\n", t);
	z = (ulong)t;
	y = z;
	z >>= 16;	/* hide undefined 32-bit shift from 32-bit compilers */
	x = z>>16;
	makecontext(&t->context.uc, (void(*)())taskstart, 2, y, x);

	return t;
}

int
taskcreate(void (*fn)(void*), void *arg, uint stack)
{
	int id;
	Task *t;

	t = taskalloc(fn, arg, stack);
	taskcount++;
	id = t->id;
	if(nalltask%64 == 0){
		alltask = realloc(alltask, (nalltask+64)*sizeof(alltask[0]));
		if(alltask == nil){
			fprint(2, "out of memory\n");
			abort();
		}
	}
	t->alltaskslot = nalltask;
	alltask[nalltask++] = t;
	taskready(t);
	return id;
}

```
可以看到，这个过程其实就是创建并设置了一个`Task`对象，然后将这个对象添加到`alltask`列表中，接着将该`Task`对象的状态设置为就绪，表示该任务可以接受调度器的调用。

#### 9.4.5 任务调度
上面提到了任务列表`alltask`，那么到底就绪的这些任务时如何被调度的呢？可以看一下[调度器的实现](https://swtch.com/libtask/task.c)，整个代码量也不是很多：
```C
static void
taskscheduler(void)
{
	int i;
	Task *t;

	taskdebug("scheduler enter");
	for(;;){
		if(taskcount == 0)
			exit(taskexitval);
		t = taskrunqueue.head;
		if(t == nil){
			fprint(2, "no runnable tasks! %d tasks stalled\n", taskcount);
			exit(1);
		}
		deltask(&taskrunqueue, t);
		t->ready = 0;
		taskrunning = t;
		tasknswitch++;
		taskdebug("run %d (%s)", t->id, t->name);
		contextswitch(&taskschedcontext, &t->context);
//print("back in scheduler\n");
		taskrunning = nil;
		if(t->exiting){
			if(!t->system)
				taskcount--;
			i = t->alltaskslot;
			alltask[i] = alltask[--nalltask];
			alltask[i]->alltaskslot = i;
			free(t);
		}
	}
}
```
逻辑其实很简单，就是循环执行正在等待中的任务，直到执行完所有的任务后退出。读者可能会觉得奇怪，这个函数里根本没有调用任务所对应的业务函数的代码，那么那些代码到底是怎么执行的呢？最关键的是下面这一句调用：
```C
contextswitch(&taskschedcontext, &t->context);
```
接下来解释这到底发生了什么。

#### 9.4.6 上下文切换
要理解函数执行过程中的上下文切换，需要理解几个比较底层的Linux系统函数：
* `makecontext()`
* `swapcontext()`

可以简单分析一下下面这个[小例子](https://github.com/Lynn--/TheGoProgrammingLanguage/blob/master/code/ChapterNine/9.4LinkSymbol/context.c)来理解这一系列函数的作用。

主函数里的`swapcontext()`调用将导致`f2()`函数被调用，因为`ctx[2]`中填充的内容为`f2()`函数的执行信息。而在执行`f2()`的过程中又遇到一次`swapcontext()`调用，这次切换到了`f1()`函数。这也是先打印两个`start`信息而没有任何一个函数先结束的原因。

现在还在`f1()`函数中，继续执行，结果又遇到了一个`swapcontext()`由于第二个参数为`ctx[2]`，因此再次切换回到了`f2()`。由于之前`f2()`函数在执行`swapcontext()`时将那个时刻的上下文全部记录到了`ctxp[2]()`中，因此从`f1()`再次切换回来后，`f2()`的执行将从之前的那一行代码继续执行，在本例中即执行打印"finish f2"信息。这也是`f2()`先于`f1()`结束的原因。

有了这些知识后，再回头去看`libtask`关于上下文切换的代码，就更容易理解了。因为在`taskalloc()`中的最后一行，可以看到每一个任务的上下文被设置为`taskstart()`函数相关，所以一旦调用`swapcontext()`切换到任务所记录的上下文，则将会导致`taskstart()`函数被调用，从而在`taskstart()`函数中进一步调用真正的业务函数，比如上例中的`primetask()`函数就是这么被调用到的(被设置为任务的`startfn`成员)。

下面是`taskstart()`函数的具体实现代码：
```C
static void
taskstart(uint y, uint x)
{
	Task *t;
	ulong z;

	z = x<<16;	/* hide undefined 32-bit shift from 32-bit compilers */
	z <<= 16;
	z |= y;
	t = (Task*)z;

//print("taskstart %p\n", t);
	t->startfn(t->startarg);
//print("taskexits %p\n", t);
	taskexit(0);
//print("not reacehd\n");
}
```
到这里，上下文切换的原理基本上已经解释完毕，那么到底什么时候应该发生上下文切换呢？我们知道，在任务的执行过程中发生任务切换只会因为以下原因之一：
* 该任务的业务代码主动要求切换，即主动让出执行权
* 发送了IO，导致执行阻塞

先看第一种情况，即主动出让执行权。这一动作通过主动调用`taskyield()`来完成。在下面的代码中，`taskswitch()`切换上下文以具体做到任务切换，`taskready()`函数将一个具体的任务设置为等待执行状态，`taskyield()`则借助其他的函数完成执行权出让。
```C

void
taskswitch(void)
{
	needstack(0);
	contextswitch(&taskrunning->context, &taskschedcontext);
}

void
taskready(Task *t)
{
	t->ready = 1;
	addtask(&taskrunqueue, t);
}

int
taskyield(void)
{
	int n;
	
	n = tasknswitch;
	taskready(taskrunning);
	taskstate("yield");
	taskswitch();
	return tasknswitch - n - 1;
}
```
当发生IO事件时，程序会先让其他处于`yield`状态的任务先执行，待清理掉这些可以执行的任务后，开始调用`poll`来坚挺所有处于IO阻塞状态的`pollfd`，一旦有某些`pollfd`成功读写，则将对应的任务切换为可调度状态。此时，IO阻塞导致自动切换的过程就完整展现在面前了。

#### 9.4.7 通信机制
这一节内容和协程机理没有直接联系，但是因为`channel`总是伴随着`goroutine`出现，所以顺便了解一下`channel`的原理也颇有好处。幸运的是，`libtask`中也提供了`channel`的参考实现。

我们已经知道，`channel`是推荐的`goroutine`之间的通信方式。而实际上，"通信"这个术语并不太太适用。从根本上来说，`channel`只是一个数据结构，可以被写入数据，也可以被读取数据。所谓的发送数据到`channel`，或者从`channel`读取数据，说白了就是对一个数据结构的操作，仅此而已。

下面就来看看`channel`的[数据结构](https://github.com/majek/libtask/blob/master/task.h)：
```go
struct Alt
{
	Channel		*c;
	void		*v;
	unsigned int	op;
	Task		*task;
	Alt		*xalt;
};

struct Altarray
{
	Alt		**a;
	unsigned int	n;
	unsigned int	m;
};

struct Channel
{
	unsigned int	bufsize;
	unsigned int	elemsize;
	unsigned char	*buf;
	unsigned int	nbuf;
	unsigned int	off;
	Altarray	asend;
	Altarray	arecv;
	char		*name;
};
```
可以看到`channel`的基本组成如下：
* 内存缓存，用于存放元素
* 发送队列
* 接受队列

从以下这个`channel`的[创建函数](http://lsub.org/sys/src/libthread/channel.c)可以看出，分配的内存缓存就紧跟在`channel`结构之后：
```C
Channel*
chancreate(int elemsize, int elemcnt)
{
	Channel *c;

	if(elemcnt < 0 || elemsize <= 0)
		return nil;
	c = _threadmalloc(sizeof(Channel)+elemsize*elemcnt, 1);
	c->e = elemsize;
	c->s = elemcnt;
	_threaddebug(DBGCHAN, "chancreate %p", c);
	return c;
}
```
因为协程原则上不会出现多线程编程中经常遇到的资源竞争问题，所以这个`channel`的数据结构甚至在访问的时候都不用加锁(因为Go语言支持多CPU核心并发执行多个`goroutine`，会造成资源竞争，所以在必要的位置还是需要加锁的)。

在理解了数据结构后，基本上可以知道这个数据结果会如何用于处理发送和接收数据，所以这里就不再针对此主题展开讨论。
