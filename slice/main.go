package main

import (
	"fmt"
)

func main() {
	mySlice := []string{"a", "b", "c", "d", "e"}
	mySlice2 := []string{"a", "b", "c", "d", "e"}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])
	fmt.Println("mySlice"[2])

	fmt.Printf("length %v\n", len(mySlice))
	fmt.Printf("capacity %v\n", cap(mySlice))

	mySlice = append(mySlice, "f")
	fmt.Printf("length %v\n", len(mySlice))
	fmt.Printf("capacity %v\n", cap(mySlice))

	mySlice = append(mySlice[3:], mySlice2[:3]...) // delete from slice and copy from another slice

}

//String can be slice because is made of runes
//runes are one byte to 4 byte size
