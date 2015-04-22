package main

import "fmt"

func main(){
    a := make(map[int]string)
    a[1] = "a"
    a[2] = "b"
    fmt.Println(a)
    fmt.Println(a[1])
    fmt.Println(len(a))
    delete(a, 1)
    fmt.Println(a)
    _, b := a[10]
    fmt.Println(b)
}
