package main

import "fmt"

func main(){
    a := map[int]string{1: "a", 2: "b"}
    for i, j := range a {
        fmt.Printf("%d - > %s\n", i, j)
    }
}