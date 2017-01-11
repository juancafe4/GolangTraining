package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo() // go makes run async
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 100; i++ {
		fmt.Println(i+1, "Foo!")
		time.Sleep(time.Duration(3 * time.Millisecond))
	}

	wg.Done()
}

func bar() {
	for i := 0; i < 100; i++ {
		fmt.Println(i+1, "Bar!")
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	wg.Done()
}
