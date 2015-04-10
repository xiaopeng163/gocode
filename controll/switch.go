package main

import "fmt"

func main(){
    fmt.Print(Season(12))
}


func Season(month int) string {
    switch month {
        case 12, 1, 2: return "winter"
        case 3, 4, 5: return "spring"
        case 6, 7, 8: return "summer"
        case 9, 10, 11: return "autumn"
    }
    return "unknown"
}