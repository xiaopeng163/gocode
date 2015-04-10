package main

import "fmt"

func main(){

    i := 0
    result := &i
    fmt.Println(*result)
    test(10, 2, result)
    fmt.Println(*result)
}

func test(a, b int, c *int) {
    *c = a + b
}
