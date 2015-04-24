package main

import "fmt"

func main(){

    a := 1
    b := 2
    fmt.Println(a, b)
    call_by_value(a, b)
    fmt.Println(a, b)
    call_bye_reference(&a, &b)
    fmt.Println(a, b)

}

func call_by_value(x, y int) {
    x += 1
    y += 1
    fmt.Println(x, y)
}


func call_bye_reference(x, y *int){
    *x += 1
    *y += 1
    fmt.Println(*x, *y)
}