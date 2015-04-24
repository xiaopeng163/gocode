package main

import "fmt"

func main(){
    fmt.Print(test1(1, 2))
}

func test1(a, b int)(x, y int){
    return a + 1, b + 1
}
