package main

import (
	"fmt"
)

func main() {
	mySlice := []string{"a", "b", "c", "d", "e"}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])
	fmt.Println("mySlice"[2])

	fmt.Printf("length %v\n", len(mySlice))
	fmt.Printf("capacity %v\n", cap(mySlice))

	mySlice = append(mySlice, "f")
	fmt.Printf("length %v\n", len(mySlice))
	fmt.Printf("capacity %v\n", cap(mySlice))
}

//String can be slice because is made of runes
//runes are one byte to 4 byte size
