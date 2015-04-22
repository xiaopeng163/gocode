package main

import "fmt"

func main(){

    a := make([]int, 5, 10)
    fmt.Printf("%p\n", a)
    fmt.Println(a)
    fmt.Println(len(a))
    fmt.Println(cap(a))
    a = append(a, 6, 7, 8, 9, 10)
    fmt.Printf("%p\n", a)
    fmt.Println(a)
    fmt.Println(len(a))
    fmt.Println(cap(a))
    a = append(a, 11, 12)
    fmt.Printf("%p\n", a)
    fmt.Println(a)
    fmt.Println(len(a))
    fmt.Println(cap(a))
}
