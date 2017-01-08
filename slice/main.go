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
	fmt.Println(mySlice)

	// multi dimensional slice

	records := make([][]string, 0)

	//var records [][]string
	// student 1
	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.00"
	student1[3] = "74.00"
	// store the record
	records = append(records, student1)
	// student 2
	student2 := make([]string, 4)
	student2[0] = "Gomez"
	student2[1] = "Lisa"
	student2[2] = "92.00"
	student2[3] = "96.00"
	// store the record
	records = append(records, student2)
	// print
	fmt.Println(records)

	// Three ways of making an slice
	// var student [] string
	// var students [][] string
	// student := [] string{}
	// students := [][] string{}
	// student := make([] string, 100)
	// students := make([][] string, 35)
}

//String can be slice because is made of runes
//runes are one byte to 4 byte size
