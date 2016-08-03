###3.2  初始化
```go
rect1 := new(Rect)
rect2 := &Rect{}
rect3 := &Rect{0, 0, 100, 200}
rect4 := &Rect{width: 100, height: 200}

fmt.Println(rect1.Area())
fmt.Println(rect2.Area())
fmt.Println(rect3.Area())
fmt.Println(rect4.Area())
```
>在Go语言中，未进行显式初始化的变量都会被初始化为该类型的零值。在Go语言中没有构造函数的概念，对象的创建通常交由一个全局的创建函数来完成，以NewXXX来命名，表示“构造函数”

```go
func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}
```
