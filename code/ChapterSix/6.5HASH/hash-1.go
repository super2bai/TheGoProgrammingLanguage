package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	TestString := "Hi,Pandaman!"

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(TestString))
	Result := Md5Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)

	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(TestString))
	Result = Sha1Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)

	/**
	output:

	d2372917b4fc4a5776124511cff870b6

	a2d861600e5b87946f8849f3d14d72426812856a
	*/
}
