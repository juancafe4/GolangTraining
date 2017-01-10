package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

//You can create your own type
type foo int

type Person struct {
	First string
	Last  string
	Age   int // non exported method when lowercase
}

type Student struct {
	Person
	Gpa float64 // exported method
}

// The reciever Person struct is getting thaat method
func (p Person) fullName() string {
	return p.First + "  " + p.Last
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
	fmt.Println(p1.First, p1.Last, p1.Age)
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

	fmt.Println(p3.First)
	fmt.Println(p3)
	// Marshal
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Printf("%v\n", string(bs))

	// fields must exported so they can exported to JSON

	// Unmarshal
	jsonPerson := []byte(`{"First": "John", "Last": "Cena", "Age": 34}`)

	var p4 Person

	json.Unmarshal(jsonPerson, &p4)

	fmt.Println("Printing person 4", p4)

	//Encoding
	json.NewEncoder(os.Stdout).Encode(p1)
	//Decoding

	var p5 Person

	rdr := strings.NewReader(`{"First": "John", "Last": "Cena", "Age": 34}`)

	fmt.Printf("%T \n", rdr)

	json.NewDecoder(rdr).Decode(&p5)

	fmt.Println("JSON decoded", p5)
}
