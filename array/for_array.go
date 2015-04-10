package main

import "fmt"

func main(){
    var a [10]int
    for i :=0; i < len(a); i++{
        a[i] = i
    }
    for _, value := range a{
        fmt.Println(value)
    }
}