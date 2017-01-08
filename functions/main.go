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
