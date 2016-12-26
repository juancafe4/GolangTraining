package main

import (
	"fmt"
)

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a // reference

	fmt.Println(b)
	fmt.Println(*b) // deference

	*b = 45

	fmt.Println(a)

	//Using no pointer example
}
