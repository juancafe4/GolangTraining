package main

import (
	"fmt"
)

// Orders matters in Go
func main() {
	// fmt.Println(x)
	// x := 42 // This does not work out of order
	//fmt.Println(x)

	//You can create a block for scope
	{
		y := 43
		fmt.Println(y)
	}
}

var y = 56 /*This works out of the function */
