### 7.8 单元测试
Go本身提供了一套轻量级的测试框架。符合规则的测试代码会在运行测试时被自动识别并执行。单元测试源文件的命名规则如下：**在需要测试的包下面创建以"_test"结尾的go文件，形如[^.]*_test.go**。

Go的单元测试函数分为两类：
* 功能测试函数`Test*(t *testing.T)`
* 性能测试函数`Benchmark*(t *testing.T)`
```go
func TestAdd(t *testing.T) 
func BenchmarkAdd(t *testing.T) {
```
测试工具会根据函数中的实际执行动作得到不同的测试结果。功能测试函数会根据测试代码执行过程中是否发生错误来返回不同的结果，而性能测试函数仅仅打印整个测试过程的花费时间。

我们在第一章中已经示范过功能测试的写法，现在关键是了解一下`testing.T`中包含的一系列函数。比如本例中使用`t.Errorf()`函数打印了一句错误信息后终止测试。虽然`testing.T`包含很多其他函数，但其实用`t.Errorf()`也能覆盖大部分的测试代码编写场景了：
```go
func TestAdd(t *testing.T){
	r := Add(1,2)
	//这里本该是3，故意改成2测试错误场景
	if r != 2 {
		t.Errorf("Add(1,2) faild. Got %d ,expected 3.", r)
	}
}
```
执行功能单元测试非常简单，直接执行`go test`命令即可。下面的代码用于对整个`simplemath`包进行单元测试：
```bash
$ go test simplemath
PASS 
oksimplemath0.013s
```
接下来介绍性能测试。先看一个例子
```go
func BenchmarkAdd(b *testing.B){
	for i := 0;i < b.N; i++ {
		Add(1,2)
	}
}
```
可以看出，性能测试与功能测试代码相比，最大的区别在于代码里的这个`for`循环，循环`b.N`次。写这个`for`循环的原因是为了能够让测试时间运行足够长的时间便于进行平均运行时间的计算。如果测试代码中一些准备工作的时间太长，也可以这样处理以明确排除这些准备工作所话费时间对于性能测试时间的影响：
```go
func BenchmarkAdd(b *testing.B){
	b.StopTimer()   //暂停计时器
	DoPreparation() //一个耗时比较长的准备工作，比如读文件
	b.StartTimer()  //开启计时器，之前的准备时间未计入总花费时间内
	for i := 0;i < b.N; i++ {
		Add(1,2)
	}
}
```
性能单元测试的执行与功能测试一样简单，只不过调用时需要增加`-test.bench`参数而已，具体代码如下所示：
```bash
$ go test-test.bench add.go
PASS
oksimplemath0.013s
```