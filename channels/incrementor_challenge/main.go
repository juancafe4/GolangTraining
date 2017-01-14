package main

import (
	"fmt"
	"strconv"
)

func main() {
	c := incrementor(2)
	var count int
	for v := range c {
		count++
		fmt.Println(v)
	}
	fmt.Println("Final Counter:", count)
}

func incrementor(n int) chan string {
	c := make(chan string)
	done := make(chan bool)
	for i := 0; i < n; i++ {
		go func(i int) {
			for j := 0; j < 20; j++ {
				c <- fmt.Sprint("Process: "+strconv.Itoa(i)+" printing:", i)
			}

			done <- true
		}(i)

	}

	go func() {
		for k := 0; k < n; k++ {
			<-done
		}
		close(c)
	}()
	return c
}
