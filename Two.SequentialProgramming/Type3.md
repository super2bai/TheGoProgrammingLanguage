###2.3 类型
Go语言内置以下这些**基础类型**
* 布尔类型：`bool`
* 整型：`int8`、`byte`、`int16`、`int`、`uint`、`uintptr`等。
* 浮点类型：`float32`、`float64`
* 复数类型：`complex64`、`complex128`
* 字符串：`string`
* 字符类型：`rune`
* 错误类型：`error` 详情见2.6
Go语言也支持以下这些**复合类型**:
* 指针：`pointer`
* 数组：`array`
* 切片：`slice`
* 字典：`map`
* 通道：`chan` 详情见4.5
* 结构体：`struct`
* 接口：`interface`详情见3.5

**在这些基础类型之上Go还封装了`int`、`uint`和`uintptr`等，这些类型的特点在于使用方便，但使用者不能对这些类型的长度坐任何假设。对于常规的开发来说，用`int`和`uint`就可以了，没必要用`int8`之类明确指定长度的类型，以免导致移植困难。**

###2.3.1 布尔类型
>关键字为`bool`，可赋值为预定义的`true`和`false`。布尔类型不能接受其他类型的赋值，不支持自动或强制的类型转换。

```go
var v1 bool
v1 = true
v2 := (1 == 2) //v2也会被推导为bool类型
var b bool
b = 1       //编译错误
b = bool(1) //编译错误

var c bool
c = (1 != 0) //编译正确
```
###2.3.2 整型
| 类型           | 长度(字节) | 值范围 |
| ------------- |-----------| ---------|
| int8          |     1     | -128~127 |
| uint8(即byte) |     1     | 0~255|
| int16         |     2     | -32768~32767 |
| uint16        |     2     | 0~65535 |
| int32         |     4     | -2147483648~2147483647 |
| uint32        |     4     | 0~4294967295 |
| int64 		| 	  8     |    -9223372036854775808~9223372036854775807 |
| uint64 		| 	  8     | 0~18446744073709551615 |
| int 		    | 平台相关   |    平台相关 |
| uint		    | 平台相关   |    平台相关 |
| uintprt	    | 同指针     |  在32位平台下位4字节，64位平台下位8字节 |

######1.类型表示
**int和int32在Go语言里被认为是两种不同的类型，编译器也不会自动做类型转换。如必要，需做强制类型转换，注意数据长度被截短而发生的数据精度损失(比如将浮点数强制转为整数)和值溢出(值超过转换的目标类型的值范围)问题**

######2.数值运算
>Go语言支持：+、-、*、/和%

######3.比较运算
>Go语言支持：>、<、==、>=、<=、!=
**两个不同类型的整型数不能直接比较，但个中类型的整型变量都可以直接与字面常量（literal）进行比较**

######4.位运算
|运算|含义|样例|
|------|------|
|x<<y|左移|124<<2 //结果为496
|x>>y|右移|124>>2//结果为31
|x^y|异或 |124^2//结果为126
|x&y| 与  |124&2//结果为0
| x\|y| 或|124\|2//结果为126
|^x|取反|^2//结果为-3

###2.3.3 浮点型
>浮点型用语表示包含小数点的数据，Go中采用IEEE-754标准的表达方式
######1.浮点数表示
>Go语言定义了两个类型`float32`和`float64`

```go
var fvalue1 float32
fvalue1 = 12

/**
会被自动设为float64，而不管赋给它的数字是否是用32位长度表示的
如果不加小数点，fvalue2会被推导为整型而不是浮点型
*/
fvalue2 := 12.0

fvalue1 = fvalue2 //编译错误，类型不同
fvalue1 = float32(fvalue2)//编译正确，强转
```

######2.浮点数比较
>因为浮点数不是一种精确的表达方式,不能用`==`来比较,会导致结果不稳定.



```go
import "math"


func IsEqual(f1, f2 float64) bool {
	return math.Abs(f1, f2) < p
}
```
p为用户自定义的比较精度,比如0.0001

###2.3.4 复数类型
>复数实际上由两个实数（在计算机中用浮点数表示）构成，一个表示实部（real），一个表示虚部（imag）。
######1.复数表示
```go
var value1 complex64	//由两个float32构成的复数类型
value1=3.2+12i
value2:=3.2+12i			//value2是complex128类型
value3:=complex(3.2,12) //value3结果同value2
```
######2.实部与虚部
>对于一个复数`z = complex(x ,  y)`,就可以通过Go语言内置函数`real(z)`获得该复数的实部，也就是x；通过`imag(z)`获得该复数的虚部，也就是y。

详情请见`math/coplx`标准库的文档

###2.3.5 字符串
>在Go语言中，字符串也是一种基本类型。详情请见标准库strings包。
```go
var str string
str = "Hello, 世界" //字符串赋值
ch := str[0]        //取字符串的第一个字符
len(str)			//获取字符串长度
var str1 ="你好"
str + str1			//字符串连接
for i := 0; i < len(str); i++ {
	s := str[i] //依据下标取字符串中的字符，类型为byte
	fmt.Println(i, s)
}
fmt.Println("以Unicode字符遍历")
for i, ch := range str {
	fmt.Println(i, ch)
}
```

**UTF-8中，中文字符占3个字节**

**以Unicode字符方式遍历时，每个字符的类型是rune（早期的Go语言用int类型表示Unicode字符），而不是byte**
>字符串的内容可以用数组下标的方式获取**字符串的内容不能在初始化后被修改**

>Go编译器支持UTF-8的源代码文件格式，如果包含非ANSI字符，保存源文件时编码格式必须选择UFT-8。


>Go语言支持UTF-8和Unicode编码，对于其他编码，可以基于iconv库用Cgo包装一个，[开源项目](https://github.com/xushiwei/go-iconv)

###2.3.6 字符类型
>Go中支持两个字符类型，`byte`(实际上是`uint8`的别名)，代表UTF-8字符串的单个字节的值；`rune`，代表单个Unicode字符

>`rune`操作查阅Go标准库的unicode包。`unitcode/uft8`包也提供了UFT-8和Unicode之间的转换.

>出于简化语言的考虑，Go语言的多数API都假设字符串为UFT-8编码。尽管Unicode字符在标准库有支持，但实际使用较少

###2.3.7 数组
>数组是Go中最常用的数据结构之一。是指一系列同一类型数据的集合。数组中包含的每个数据被称为数组元素（element），一个数组包含的元素个数被称为数组长度。
```go
var a [32]byte                    //长度为32的数组，每个元素为一个字节
//编译通不过，暂时注释
//var b [2 * N]struct{ x, y int32 } //复杂类型数组	
var c [1000]*float64              //指针数组
var d [3][5]int                   //二维数组,三行五列,15个元素
var e [2][2][2]float64            //等同于[2]([2]([2]float64))
len(a)							  //获得数组长度
```
**Go中，数组长度在定义后就不可更改，在声明时长度可以为一个常量或者一个常量表达式（在编译器即可计算结果的表达式）***

######1. 元素访问
```go
arr := [5]int{1, 2, 3, 4, 5}
for i := 0; i < len(arr); i++ {
	fmt.Println(i, arr[i])
}
for index, value := range arr {
	fmt.Println(index, value)
}
```
######2.值类型
>在Go语言中数组是一个值类型（value type）。所有的值类型变量在赋值和作为参数传递时都将产生一次复制动作。在函数体中无法修改传入的数组的内容，因为函数内操作的只是传入数组的一个副本。
```go
func main() {
	fmt.Println("Hello World!")
	arr := [5]int{1, 2, 3, 4, 5}
	modify(arr)
	fmt.Println("main arr :", arr)
	//main arr : [1 2 3 4 5]
}
func modify(array [5]int) {
	array[0] = 10
	fmt.Println("modify array", array)
	//modify array [10 2 3 4 5]
}
```

###2.3.8 数组切片
> 数组切片（`slice`）就像一个指向数组的指针，但拥有自己的数据结构，不仅仅是指针。**可以随时动态扩充存放空间，并且可以被随意传递而不悔到值所管理的元素被重复复制**

-----

>数组切片的数据结构抽象为以下三个变量：
>* 一个指向原声数组的指针
>* 数组切片中的元素个数
>* 数组切片已分配的存储空间

######1.创建数组切片
* 基于数组
```go
	//定义数组
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//基于数组创建一个数组切片
	var mySlice []int = arr[:5]
	fmt.Println("elements of arr: ")
	for _, v := range arr {
		fmt.Println(v, "")
	}
	fmt.Println("elements of mySlice: ")
	for _, v := range mySlice {
		fmt.Println(v, "")
	}
```
**Go语言支持Array[first:last]**这样的方式来基于数组生成一个数组切片。
```go
	myArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mySlice :=myArray[:]
	mySlice=myArray[:5
	mySlice=myArray[5:]
```
* 直接创建
**Go语言提供的内置函数make()可以用于灵活的创建数组切片。**
```go
mySlice1 := make([]int, 5)       //初始个数为5，元素初始值为0
mySlice2 := make([]int, 5, 10)   //初始格式为5，元素初始值为0，并预留10个元素的存储空间
mySlice3 := []int{1, 2, 3, 4, 5} //直接创建并初始化包含5个元素的数组切片
```
######2.元素遍历
同数组
######3.动态增减元素
**可动态增减元素是数组切片比数组更强大的功能。**
>数组切片多了一个存储能力（capacity）的概念，即元素个数和分配的空间可以是两个不同的值，合理的设置存储能力的值，可以大幅度降低数组切片内部重新分配内存和半送内存快的频率，从而大大提高程序性能。`cap()`函数返回的是数组切片分配的空间大小，而`len()`函数返回的是数组切片中当前所存储的元素个数。

```go
mySlice := make([]int, 5, 10) //初始格式为5，元素初始值为0，并预留10个元素的存储空间
fmt.Println("len(mySlice)", len(mySlice))
fmt.Println("cap(mySlice)", cap(mySlice))
mySlice2 := append(mySlice, 1, 2, 3)
mySlice3 := append(mySlice, mySlice2...)
fmt.Println(mySlice2)
fmt.Println(mySlice3)
```
如果没有省略号，会有编译错误.因为按append()的语义，从第二个参数起的所有参数都是待附加的元素
在mySlice后追加3个元素，从而生成一个新的数组切片,append()的第二个参数其实是一个不定参数，可以按需求添加若干个元素,甚至直接将一个数组切片追加到另一个数组切片的末尾

######4.基于数组切片创建数组切片
```go
oldSlice := []int{1, 2, 3, 4, 5}
newSlice := oldSlice[:3] //基于前三个创建数组切片
```
**选择的oldSlice元素范围可以i 超过所包含的元素个数，newSlice可以基于oldSlice的前六个元素创建，虽然oldSlice只有五个元素。只要选择的范围不超过oldSlice存储能力，那么个创建程序就是合法的。newslice中超过oldSlice元素的部分都会填上0。**

######5.内容复制
>数组切片支持Go语言的另一个内置函数`copy()`，用于将内容从一个数组切片复制到另一个数组切片。如果加入的两个数组切片不一样大，就会按其中较小的数组切片的元素个数进行复制。

```go
slice1 := []int{1, 2, 3, 4, 5}
slice2 := []int{5, 4, 3}
copy(slice2, slice1) //只会复制slice1的前3个元素到slice2中
copy(slice1, slice2) //只会复制slice2的3个元素到slice1的前3个位置
fmt.Println(slice1)
fmt.Println(slice2)
```
###2.3.9 map
**Go语言中，使用map不需要引入任何库，并且用起来也更加方便。**
>map是一堆键值对的未排序集合。

#####1.声明
```go
var myMap map[string] PersonInfo
//var 变量名 map[key] value
```
######2.创建并赋值
```go
type PersonInfo struct {
	ID      string
	Name    string
	Address string
}
func main() {
	myMap = make(map[string]PersonInfo)
	myMap = make(map[string]PersonInfo, 100) //指定存储能力
	myMap = map[string]PersonInfo{"123": PersonInfo{"1", "Jack", "Rom 101"}}//赋值
}
```
######3.元素删除
>Go语言提供了一个内置函数 `delete()`，用于删除容器内的元素。

```go
delete(myMap,"123")
//如果没有key为"123"的键值对，这个调用将什么都不会发生，也不会有什么副作用。
//如果传入的map变量的值是nil，该调用将导致抛出异常（panic）。
```
######4.元素查找
```go
value, ok :=myMap["123"]
if ok {
	//找到了
}
```
>判断是否成功找到特定的键，只需查看第二个返回值ok。