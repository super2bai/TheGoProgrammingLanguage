###2.2 常量
>在Go语言中，常量是指编译期间就已知且不可改变的值。常量可以是树枝类型（包括整型、浮点型和复数类型）、布尔类型、字符串类型等。

###2.2.1  字面常量
>所谓字面常量 （literal），是指程序中硬编码的常量，如：
```go
-12
3.1415926		//浮点类型的常量
3.2+12i 		//复数类型的常量
true			//布尔类型的常量
"foo"			//字符串类型的常量
```
>Go语言的字面常量更接近自然语言中的常量概念，它是无类型的。

###2.2.2 常量定义
通过`const`关键字，可以给字面常量指定一个友好的名字
```go
const Pi float64 = 3.14159265358979323846
const zero = 0.0 //无类型浮点常量
const (
	size int64 = 1024
	eof        = -1 //无类型整型常量
)
const u, v float32 = 0, 3   //u =0.0,v=3.0,常量的多重赋值
const a, b, c = 3, 4, "foo" //a=3,b=4,c="foo",无类型整型和字符串常量

const mask = 1 << 3		//常量右值也可以是一个在编译期运算的常量表达式

//下面是错误示例！！！
const Home = os.GetEnv("HOME") //由于常量的赋值是一个编译期行为，所以右值不能出现任何需要在运行期才能得出结果的表达式
```
>Go的常量定义可以限定常量类型，但不是必需的。如果定义常量时没有指定类型，那么它与字面常量一样，是无类型常量

###2.3.3 与定义常量
>Go语言与定义了：`true`、`false`和`iota`
>`iota`比较特殊，可以被认为是一个可被编译器修改的常量，在每一个const关键字出现时被重置为0，然后在下一个const出现之前，每出现一个iota，其所代表的数字会自动增1
```go
const ( //iota被重设为0
	c0 = iota //c0=0
	c1 = iota //c1=1
	c2 = iota //c2=2
)

const (
	a = 1 << iota //a==1(iota在每个const开头被重设为0）
	b = 1 << iota //b=2
	c = 1 << iota //c=4
)

const (
	u         = iota * 42 //u==0
	v float64 = iota * 42 //v==42.0
	w         = iota * 42 //w==84
)
const x = iota //x==0(因为iota又被重设为0饿)
const y = iota //y==0(同上)
```
>如果两个const的赋值语句的表达式是一样的，那么刻意省略后一个赋值表达式。
```go
const (
	c0 = iota
	c1
	c2
)

const (
	a = 1 << iota
	b
	c
)
```
###2.2.4 枚举
>枚举指一系列相关的常量,Go语言并不支持众多其他语言明确支持的enum关键字。
`const`后跟一对圆括号的方式定义一组常量，这种定义法在Go语言中通畅用语定义枚举值
```go
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays //这个常量没有导出
)
```
**同Go语言的其他符号（symbol）一样，以大写字母开头的常量在包外可见。以上例子中numberOfDays为包内私有，其他符号则可被其他包访问。**