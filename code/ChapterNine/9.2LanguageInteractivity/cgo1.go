package main

import "fmt"

/*
#include <stdlib.h>
*/
import "C"

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

/**
output:
Random: 677741240
*/
func main() {
	Seed(100)
	fmt.Println("Random:", Random())
}
