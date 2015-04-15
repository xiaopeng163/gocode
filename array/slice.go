package main

import "fmt"

func main(){
    var a [10]int
    for i :=0; i < len(a); i++{
        a[i] = i
    }
    var b []int = a[0:3]
    for i :=1; i < len(b); i++ {
        fmt.Println(b[i])
}

}
