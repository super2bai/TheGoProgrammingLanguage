### 4.8 同步

考虑到即使成功的使用channel来作为通信手段，还是避免不来多个goroutine之间共享数据的问题，Go语言的设计者虽然对channelyou极高的期望，但也提供了妥善的资源锁方案。

### 4.8.1 同步锁
>Go语言包中的`sync`包提供了两种锁类型： `sync.Mutex` （互斥锁）和 `sync.RWMutex`（读写锁）。

>`Mutex`是最简单的一种锁类型,，同时也比较暴力，当一个goroutine获得了 Mutex后，其他goroutine就只能乖乖等到这个goroutine释放该Mutex。

>互斥锁是传统的并发程序对共享资源进行访问控制的主要手段。它由标准库代码包sync中的Mutex结构体类型代表。sync.Mutex类型（确切地说，是*sync.Mutex类型）只有两个公开方法——Lock和Unlock。顾名思义，前者被用于锁定当前的互斥量，而后者则被用来对当前的互斥量进行解锁。

```go
var mutex sync.Mutex
func write() {
	mutex.Lock()
	//...
	defer mutex.Unlock()
}
```

**对于同一个互斥锁的锁定操作和解锁操作总是应该成对的出现。如果我们锁定了一个已被锁定的互斥锁，那么进行重复锁定操作的Goroutine将会被阻塞，直到该互斥锁回到解锁状态。避免这种情况发生的最简单、有效的方式依然是使用defer语句。**

>`RWMutex`相对友好些，是经单的单写多读模型。在读锁占用的情况下，会阻止写，但不阻止读，也就是多个goroutine可同时获取读锁（调用`RLock()`;而写锁（调用`Lock()`））会阻止任何其他goroutine（无论读和写）进来，整个锁相当于由该goroutine独占。从RWMutex的实现看，RWMutex是基于Mutex实现的，只读锁的实现使用类似引用计数器的功能．

```go
type RWMutex struct {
	w           Mutex  // held if there are pending writers
	writerSem   uint32 // semaphore for writers to wait for completing readers
	readerSem   uint32 // semaphore for readers to wait for completing writers
	readerCount int32  // number of pending readers
	readerWait  int32  // number of departing readers
}
```

>对于这两种锁类型，任何一个Lock()或RLock()均需要保证对应有Unlock()或RUnlock()调用与之对应，否则可能导致等待该锁的所有goroutine处于饥饿状态，甚至可能导致死锁。

>  func (rw *RWMutex) Lock()　　写锁，如果在添加写锁之前已经有其他的读锁和写锁，则lock就会阻塞直到该锁可用，为确保该锁最终可用，已阻塞的 Lock 调用会从获得的锁中排除新的读取器，即写锁权限高于读锁，有写锁时优先进行写锁定。

### 4.8.1 全局唯一性操作
>对于从全局的角度只需药运行一次的代码，比如全局初始化操作，Go语言提供了一个`Once`类型来保证全局的唯一性操作。


```go
package main

import (
	"sync"
	"time"
)

var a string = "hello, world"
var once sync.Once

func setup() {
	print(a)
}

func doprint() {
	once.Do(setup)

}

func twoprint() {
	go doprint()
	go doprint()
}
func main() {
	twoprint()
	time.Sleep(3e9)
}
```
输出：hello, world


```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func main() {

	for i, v := range make([]string, 10) {
		once.Do(onces)
		fmt.Println("count:", v, "---", i)
	}
	for i := 0; i < 10; i++ {

		go func() {
			once.Do(onced)
			fmt.Println("213")
		}()
	}
	time.Sleep(4000)
}
func onces() {
	fmt.Println("onces")
}
func onced() {
	fmt.Println("onced")
}
```
输出：
onces
count:  --- 0
count:  --- 1
count:  --- 2
count:  --- 3
count:  --- 4
count:  --- 5
count:  --- 6
count:  --- 7
count:  --- 8
count:  --- 9

整个程序，只会执行onces()方法一次,onced()方法是不会被执行的。

为了更好的控制并行中的原子性操作，sync包中还包含一个`atomic`子包，它提供了对于一些基础数据类型的原子操作函数，保证多CPU对同一块内存的操作是原子的，原子操作直接有底层CPU硬件支持，因而一般要比基于操作系统API的锁方式效率高些。

**下面内容引用网页[原文地址](http://blog.csdn.net/zhijiayang/article/details/51727197)**

**CAS**

原子操作中最经典的CAS(compare-and-swap)在atomic包中是Compare开头的函数。

```go 
func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)
func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)
func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)
func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)
func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)
```

CAS的意思是判断内存中的某个值是否等于old值，如果是的话，则赋new值给这块内存。CAS是一个方法，并不局限在CPU原子操作中。 
CAS比互斥锁乐观，但是也就代表CAS是有赋值不成功的时候，调用CAS的那一方就需要处理赋值不成功的后续行为了。

这一系列的函数需要比较后再进行交换，也有不需要进行比较就进行交换的原子操作。

```go 
func SwapInt32(addr *int32, new int32) (old int32)
func SwapInt64(addr *int64, new int64) (old int64)
func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)
func SwapUint32(addr *uint32, new uint32) (old uint32)
func SwapUint64(addr *uint64, new uint64) (old uint64)
func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)
```

**增加或减少**

对一个数值进行增加或者减少的行为也需要保证是原子的，它对应于atomic包的函数就是

```go 
func AddInt32(addr *int32, delta int32) (new int32)
func AddInt64(addr *int64, delta int64) (new int64)
func AddUint32(addr *uint32, delta uint32) (new uint32)
func AddUint64(addr *uint64, delta uint64) (new uint64)
func AddUintptr(addr *uintptr, delta uintptr) (new uintptr)
```

**读取或写入**

当我们要读取一个变量的时候，很有可能这个变量正在被写入，这个时候，我们就很有可能读取到写到一半的数据。 
所以读取操作是需要一个原子行为的。在atomic包中就是Load开头的函数群。

```go 
func LoadInt32(addr *int32) (val int32)
func LoadInt64(addr *int64) (val int64)
func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)
func LoadUint32(addr *uint32) (val uint32)
func LoadUint64(addr *uint64) (val uint64)
func LoadUintptr(addr *uintptr) (val uintptr)
```

好了，读取我们是完成了原子性，那写入呢？也是同样的，如果有多个CPU往内存中一个数据块写入数据的时候，可能导致这个写入的数据不完整。 
在atomic包对应的是Store开头的函数群。

```go 
func StoreInt32(addr *int32, val int32)
func StoreInt64(addr *int64, val int64)
func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)
func StoreUint32(addr *uint32, val uint32)
func StoreUint64(addr *uint64, val uint64)
func StoreUintptr(addr *uintptr, val uintptr)
```
