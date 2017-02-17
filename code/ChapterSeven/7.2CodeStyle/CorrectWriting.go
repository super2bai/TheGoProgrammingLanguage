package main

func main() {
	Foo(1, 2)
}

func Foo(a, b int) (ret int, err error) {
	if a > b {
		ret = a
	} else {
		ret = b
	}
	return ret, nil
}
