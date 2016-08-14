### 4.6 多核并行化
>在执行一些昂贵的计算任务时，希望能够尽量利用现代服务器普遍具备的多核特性来尽量将任务并行化，从而达到降低总计算时间的目的。此时需要了解CUP核心的数量，并针对性地分解计算任务到多个goroutine中去并行运行。

模拟一个完全可以并行的计算任务：计算N哥整数的综合。可以将所有整数分成M份，M即CPU的个数。让每个CPU开始计算分给它的那份计算任务，最后将每个CPU的计算结果再做一次累加，这样就可以得到所有N个整数的总和。

```go
package main

import (
	"fmt"
	"runtime"
)

type Vector []float64

func (v Vector) DoSome(i, n int, u Vector, c chan float64) {

	var sum float64
	for ; i < n; i++ {
		sum += u[i]
	}
	c <- sum
}

func (v *Vector) DoAll(u Vector) {
	NCPU := runtime.NumCPU()
	fmt.Println(NCPU)
	runtime.GOMAXPROCS(NCPU) //设置使用多个CPU核心
	c := make(chan float64, NCPU) //根据自己电脑的CPU产生对应个数的管道

	for i := 0; i < NCPU; i++ {
		go v.DoSome(i*len(u)/NCPU, (i+1)*len(u)/NCPU, u, c)
	}

	var sum float64 = 0.00
	for i := 0; i < NCPU; i++ {
		sum += <-c
	}
	fmt.Println(sum)
}

func main() {
	var v Vector
	u := []float64{1.00, 2.00, 3.00, 4.00, 5.00, 6.00}

	v.DoAll(u)
}
```
计算过程中其实只有一个CPU核心处于繁忙状态，官方的答案是：这是当前版本的Go便一起还不能很智能地发现和利用多核的优势。虽然确实创建了多个goroutine，并且从运行状态看这些goroutine也都在并行运行，但实际上所有这些goroutine都运行在同一个CPU核心上，在一个goroutine的到时间片执行的时候，其他goroutine都会处于等待状态。