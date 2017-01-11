package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

// Special time to make some set up befero excecuting
func init() {
	fmt.Println("Number of cores ", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
}
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
