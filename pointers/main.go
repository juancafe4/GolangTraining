package main

import (
	"fmt"
)

func zero(x int) {
	fmt.Printf("%p\n", &x)
	x = 0
}

func zero2(z *int) {
	fmt.Printf("%p\n", z)
	*z = 0
}
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
	//It is passed it value not the address nothing changes
	x := 5
	fmt.Printf("%p\n", &x)
	zero(x)
	fmt.Println(x)

	// wtih pointers
	y := 5
	fmt.Printf("%p\n", &y)
	zero2(&y)
	fmt.Println(y)
}
