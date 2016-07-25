###2.4流程控制
Go语言支持如下几种流程控制语句
* 条件语句：`if`、`else`、`else if`
* 选择语句：`switch`、`case`、`select`
* 循环语句：`for`、`range`
* 跳转语句：`goto`
还添加如下关键字：`break`、`continue`、`fallthrough`

###2.4.1条件语句
```go
if a < 5 {
	...
} else {
	...
}
```
* 条件语句不需要使用括号将条件包含起来()
* 无论语句内有几条语句，花括号{}都必须存在
* 左花括号{必须和if或者else处于同一行
* 在if之后，条件语句之前，可以添加变量初始化语句，使用;间隔
* 在有返回值的函数中，不允许将“最终的”return语句包含在if...else...结构中，否则会编译失败：`function end without a return statement`Go编译器无法找到终止该函数的return语句。

###2.4.2
```go
switch i {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 2:
		fallthrough
	case 3:
		fmt.Println("3")
	case 4, 5, 6:
		fmt.Println("4,5,6")
	default:
		fmt.Println("default")
	}
```
i=0,    0
i=1,    1
i=2,    3
i=3,    3
i=4,   4,5,6
i=5,   4,5,6
i=6,   4,5,6   
i=other,   default
* 左花括号{必须与switch处于同一行
* 条件表达式不限制为常量或者整数
* 单个case中，可以出现多个结果选项
* Go语言不需要用break来明确退出一个case
* 只有在case中明确添加`fallthrough`关键字，才会继续执行紧跟的下一个case
* 可以不设定switch之后的条件表达式，在此种情况下，整个switch结果与多个if...else...的逻辑作用等同

###2.4.3 循环语句
Go仅支持`for`
```go
sum := 0
for i := 0; i < 10; i++ {
	sum += i
}
fmt.Println(sum)
```
无限循环
```go
sum := 0
for {
	sum++
	if sum > 100 {
		break
	}
}
fmt.Println(sum)
```
多重赋值
```go
a := []int{1, 2, 3, 4, 5}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a)
```
* 左花括号{必须与for处于同一行
* 允许在循环条件中定义和初始化变量，Go不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多个变量
* Go语言的for循环同样支持continue和break来控制循环，但是它提供了一个更高级的break，可以选择中段哪一个循环

```go
J:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 6 {
				break J //现在终止的是j 循环，而不是i的那个
			}
			fmt.Println(i)
		}
	}
}
```
###2.4.4 跳转语句
`goto`跳转到本函数内的某个标签
```go
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}
```
输出：0、1、2、3、4、5、6、7、8、9

