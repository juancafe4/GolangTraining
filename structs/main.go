package main

import (
	"fmt"
)

//You can create your own type
type foo int

type person struct {
	first string
	last  string
	age   int
}

func main() {
	var myAge foo
	myAge = 44
	fmt.Printf("%T %v \n", myAge, myAge)

	p1 := person{"James", "Bond", 20}

	fmt.Println(p1.first, p1.last, p1.age)
}
