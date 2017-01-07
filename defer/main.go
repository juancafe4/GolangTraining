package main

import (
	"fmt"
)

func world() {
	fmt.Println("World")
}

func hello() {
	fmt.Println("Hello")
}
func main() {
	// defer makes this function to excecute last at the
	// end of the main function
	defer world()
	hello()
}
