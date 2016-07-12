package main

import (
	"fmt"
	"os"
)

func main() {
	argWithProg := os.Args
	argWithoutProg := os.Args[1:]
	arg := os.Args[3]

	fmt.Println(argWithProg)
	fmt.Println(argWithoutProg)
	fmt.Println(arg)
	fmt.Printf("%T\n", arg)

}
