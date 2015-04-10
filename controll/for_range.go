package main

import "fmt"

func main(){
    var str string = "HelloWorld"
    for pos, char := range str {
        fmt.Printf("The position is %d, the char is %c\n", pos, char)
    }
}
