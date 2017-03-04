package main

import "fmt"

type ISpeaker interface {
	Speak()
}

type SimpleSpeaker struct {
	Message string
}

func (speak *SimpleSpeaker) Speak() {
	fmt.Println("I am speaking? ", speak.Message)
}

func main() {
	var speak ISpeaker
	speak = &SimpleSpeaker{"hello"}
	speak.Speak()
}
