package main

import "fmt"

func main(){
    s1 := []int{1, 2, 3, 4}
    s2 := []int{5, 6}
    copy(s1, s2)
    fmt.Print(s1)
}

