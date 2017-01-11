package main

import (
	"fmt"
)

func main() {
	// Unbuffer channel
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		fmt.Printf("Closing channel %T\n", c)
		close(c)
	}()

	// go func() {
	// 	for {
	// 		fmt.Println(<-c)
	// 	}
	// }()

	for n := range c {
		fmt.Println(n)
	}
	//time.Sleep(time.Second)
}
