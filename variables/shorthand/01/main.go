package main

import (
	"fmt"
)

var i string = "Cowboy"

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
	fmt.Printf("%v \n", i)

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
	fmt.Printf("%T \n", g)
	fmt.Printf("%T \n", i)

	/* Zero Values */
	var j int
	var k string
	var l float64
	var m bool

	fmt.Printf("%v \n", j)
	fmt.Printf("%v \n", k)
	fmt.Printf("%v \n", l)
	fmt.Printf("%v \n", m)

}

/*Use this method */
