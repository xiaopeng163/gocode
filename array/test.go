package main

import "fmt"

func main() {
    a := [4]string{"a", "b", "c", "d"}
    for i := range a {
        fmt.Println("Array item", i, "is", a[i])
    }
}