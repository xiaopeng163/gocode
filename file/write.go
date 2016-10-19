package main

import (
	"fmt"
	"os"
)

func main() {

	file, error := os.Create("new.txt")
	if error != nil {
		fmt.Println(error)
	}
	defer file.Close()
	line := []byte("line 1")
	file.Write(line)
    line 
}
