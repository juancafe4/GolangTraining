package main

import "fmt"

func main() {
	for i := 0; i < 500; i++ {
		fmt.Printf("%d \t %b \t %x  %q \n", i, i, i, i)
	}
}
