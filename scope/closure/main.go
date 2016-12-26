package main

import (
	"fmt"
)

// It returns a function
func wrapper() func() int {
	y := 0
	return func() int {
		y++
		return y
	}
}
func main() {
	//Fucntions are first class
	x := 0
	increment := func() int {
		x++
		return x
	}
	// Here is calling the function returning the inner function
	// with y = 0
	increment2 := wrapper()

	fmt.Printf("%v\n", increment())
	fmt.Printf("%v\n", increment())

	fmt.Printf("%v\n", increment2())
	fmt.Printf("%v\n", increment2())

}
