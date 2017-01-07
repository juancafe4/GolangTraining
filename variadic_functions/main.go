package main

import (
	"fmt"
)

func main() {
	n := average(43, 12, 36, 87, 90)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T\n", sf)

	total := 0.0

	// range returns the keys and the value
	// We do not need the key we put _
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
