package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	srcFile, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer srcFile.Close()
	scanner := bufio.NewScanner(srcFile)
	scanner.Scan()
	fmt.Println(scanner.Text())
	scanner.Scan()
	fmt.Println(scanner.Text())
	scanner.Scan()
	fmt.Println(scanner.Bytes())

	srcFile.Seek(0, 0)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
