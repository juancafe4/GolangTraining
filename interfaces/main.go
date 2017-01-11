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
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Shape interface {
	area() float64
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

}
