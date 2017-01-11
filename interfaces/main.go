package main

import (
	"fmt"
	"math"
	"reflect"
)

type Square struct {
	side float64
}
type Circle struct {
	radius float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

//value receiver
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Shape interface {
	area() float64
}

// Empty interface
func shape(a interface{}) {
	fmt.Println(a)
}
func info(s Shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(reflect.TypeOf(s).String() == "main.Circle")
}
func main() {
	s := Square{10}
	c := Circle{5}
	info(s)
	info(c)

	shape(s)

	ex := make([]interface{}, 3)

	ex[0] = "Hello"
	ex[1] = true
	ex[2] = 34

	fmt.Println(ex)
}

// Value receiver

// func (c Circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// can pass like a value or a pointer info(c) or info(&c)

// Pointer receiver

// func (c *Circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// pass be must as an address info(&c) or use pointer
// because the pointer receiver might be pointing to many circles
// and the compoler does not know where to convert is pointing when you use a value
// In value receiver the compiler knows where exactly where the circle is pointing
// because is just a value with a fixed address.
