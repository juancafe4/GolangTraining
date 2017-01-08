package main

import (
	"fmt"
)

func main() {
	// slice for lists
	// maps for dictionaries

	m := make(map[string]int)
	m["k1"] = 3
	m["k2"] = 5

	fmt.Printf("%T\n", m)
	fmt.Printf("%v\n", m)

	delete(m, "k2")

	fmt.Printf("After deleting... %v\n", m)

	val, exists := m["k1"]

	fmt.Println("Val: ", val)
	fmt.Println("exists?: ", exists)

	var myGreeting = map[string]string{
		"Jane": "Nice",
	}

	myGreeting["Tim"] = "Good Morning"
	myGreeting["Jenny"] = "Bonjour"

	fmt.Println(myGreeting)

	// Check if value exists
	if _, exists := myGreeting["Tim"]; exists {
		fmt.Println("It exists!")
	}
}
