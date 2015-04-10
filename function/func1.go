package main

import "fmt"

func main(){
    fmt.Print(Sum(1, 2, 3))
}


func Sum(a, b, c int) int {
    return a + b + c
}