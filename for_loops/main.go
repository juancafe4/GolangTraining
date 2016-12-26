package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

	sum := 0

	for sum <= 100 {
		sum++
	}

	// for {
	//   //infinite loop
	// }

	// no while loops no for loops

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, "-", j)
		}
	}
}
