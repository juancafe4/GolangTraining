package main

import (
	"fmt"
)

//You can create your own type
type foo int

type Person struct {
	first string
	last  string
	age   int
}

type Student struct {
	Person
	gpa float64
}

func main() {
	var myAge foo
	myAge = 44
	fmt.Printf("%T %v \n", myAge, myAge)

	p1 := Person{"James", "Bond", 20}

	fmt.Println(p1.first, p1.last, p1.age)

	student := Student{
		Person: Person{
			first: "John",
			last:  "Smith",
			age:   23,
		},
		gpa: 2.3,
	}

	fmt.Println(student)
}
