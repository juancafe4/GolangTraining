package main

import (
	"fmt"
)

func main() {
	var x [58]string // array
	// var x[] string slice

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])

	// for _, i := range x {
	// 	fmt.Println(i)
	// }

}
