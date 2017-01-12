package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	//c <- 1
	go func() {
		fmt.Println("here")
		c <- 1
	}()
	fmt.Println("hello")
	fmt.Println(<-c)

	c2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c2 <- i
		}
		close(c2)
	}()

	for v := range c2 {
		fmt.Println(v)
	}

}
