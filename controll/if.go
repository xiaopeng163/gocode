package main

import "fmt"

var flag bool = true


func main(){
    if flag {
        fmt.Print("flag is true")
    } else {
        fmt.Print("flag is false")
    }
}


