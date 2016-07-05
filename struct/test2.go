package main

import "fmt"

func main() {

	type foo int

	var x foo
	x = 1
	fmt.Printf("%T,%#v\n", x, x)

	var y foo = 2
	fmt.Println("%T, %#v\n", y, y)
	fmt.Println(x + y)
}
