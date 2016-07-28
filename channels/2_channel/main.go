package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	finish := make(chan bool)
	n := 10
	go func() {
		for i := 0; i < n; i++ {
			c <- i
		}
		finish <- true
	}()

	go func() {
		for i := 9; i > 0; i-- {
			c <- i

		}
		finish <- true

	}()

	go func() {
		<-finish
		<-finish
		close(c)

	}()

	for i := range c {
		fmt.Println(i)

	}

}
