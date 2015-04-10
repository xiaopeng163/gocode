package main

import "fmt"

var (
    a = false
    b = true
)

var c bool

var d bool = true

func main(){
    e := true
    fmt.Printf("a is %t\n", a)
    fmt.Printf("b is %t\n", b)
    fmt.Printf("c is %t\n", c)
    fmt.Printf("d is %t\n", d)
    fmt.Printf("e is %t\n", e)
}