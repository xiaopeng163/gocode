package main

import "fmt"

var a int = 2

func main() {
	a := 1
	fmt.Printf("a is %d\n", a)
	test()
}

func test() {
	fmt.Printf("a is %d\n", a)
}
