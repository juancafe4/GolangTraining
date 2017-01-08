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

	val, ok := m["k1"]

	fmt.Println("Val: ", val)
	fmt.Println("Ok?: ", ok)

	var myGreeting = map[string]string{}

	myGreeting["Tim"] = "Good Morning"
	myGreeting["Jenny"] = "Bonjour"

	fmt.Println(myGreeting)
}
