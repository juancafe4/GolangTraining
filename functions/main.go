package main

import (
	"fmt"
)

func main() {
	s := greet("Jane", "Doe")
	fmt.Println(s)

	fmt.Println(greet2("Jane", "Doe"))

	//Anynomous function
	func() {
		fmt.Println("I am driving!")
	}()

	fmt.Println(half(2))
	fmt.Println(half(1))

	fmt.Println(greatest(2, 5, 34, 2))
}

// func greet(fname, lname string) string {
// 	return fmt.Sprint(fname, lname)
// }

// or

func greet(fname string, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}

//multiple returns
func greet2(fname string, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(fname, lname)
}

//Ex1 and ex2
func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func greatest(n ...int) int {
	max := 0
	for _, v := range n {
		if max <= v {
			max = v
		}
	}
	return max
}
