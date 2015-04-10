package main

import "fmt"

func main(){

    i := 1
    for {
        if i < 10 {
            fmt.Println(i)
            i += 1
        } else {
            break
        }
    }

}
