package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Conversion

	var num1 = 10
	var num2 = 12.343534

	fmt.Println(num1 + int(num2))

	var c1 rune = 'a'
	var c2 int32 = 'b'

	fmt.Println(string(c1))
	fmt.Println(string(c2))

	fmt.Println(string([]byte{'h', 'e', 'l', 'l', 'o'}))

	var x = "12"

	z, _ := strconv.Atoi(x)

	fmt.Println(z)

	// Assertion
	var name interface{} = "Sidney"

	str, ok := name.(string)

	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Println("Value is not a string.")
	}

	var num3 interface{} = 5

	fmt.Printf("%T\n", num3)
	fmt.Printf("%T\n", num3.(int))
	//fmt.Printf("%T\n", int(num3)) Cannot do that use assertion
	//fmt.Printf("%v\n", num3+5) // not the same type
	fmt.Printf("%v\n", num3.(int)+5)
}
