package unbuffer

import (
	"fmt"
	"time"
)

func unbuffer() {

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}

	}()

	time.Sleep(time.Second)

}
