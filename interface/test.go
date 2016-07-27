package main

import "fmt"

// People interface
type People interface {
	Speak(str string)
}

type Student struct {
	Name string
	Age  int
}

func (*Student) Speak() {
	fmt.Println("I am a student")
}

type Teacher struct {
	Name   string
	Age    int
	School string
}

func (*Teacher) Speak() {
	fmt.Println("I am a teacher")
}
