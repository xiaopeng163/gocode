package main

import "fmt"

func main(){
    fmt.Print(Test(1, 2, 3, 4))
}

func Test(a ...int) int{
    result := 0
    for _, value := range a{
        result += value
    }
    return result
}
