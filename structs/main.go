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

// The reciever Person struct is getting thaat method
func (p Person) fullName() string {
	return p.first + "  " + p.last
}

//Overriding
func (p Person) Greeting() {
	fmt.Println("It is a person")
}

func (s Student) Greeting() {
	fmt.Println("It is a student")
}
func main() {
	var myAge foo
	myAge = 44
	fmt.Printf("%T %v \n", myAge, myAge)

	p1 := Person{"James", "Bond", 20}
	var p2 Person
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2)
	student := Student{Person{"John", "Smith", 23}, 2.3}
	// student := Student{
	// 	Person: Person{
	// 		first: "John",
	// 		last:  "Smith",
	// 		age:   23,
	// 	},
	// 	gpa: 2.3,
	// }

	fmt.Println(student)

	fmt.Println(student.fullName())

	p1.Greeting()
	student.Greeting()

	//pointers
	p3 := &Person{"John", "Lemon", 34}

	fmt.Println(p3.first)
	fmt.Println(p3)
}
