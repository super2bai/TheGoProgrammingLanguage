### 3.4 可见性

成员变量、成员方法首字母大写，则包外可见；首字母小写，包外不可见。
```go
type Rect struct {
	X, Y          float64
	Width, Height float64
}

func (r *Rect) area() float64 {
	return r.Width * r.Height
}
```
>Go语言中符号的可访问性是包级而不是类型级的。

上面例子中，尽管area()是Rect的内部方法，但同一包中的其他类型也都可以访问到它。这样可访问性很粗旷，很特别，但是非常实用。因为不用加friend关键字。