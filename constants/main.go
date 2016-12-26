package main

import (
	"fmt"
)

// Constant unchanged value
const p = "I am a constant p"
const (
	PI       = 3.1416
	Language = "GO"
)

const (
	A = iota // 0
	B = iota // 1
	C = iota // 2
)

const (
	D = iota // 0
	E = iota // 1
	F = iota // 2
)

func main() {
	const q = 42
	fmt.Println(p)
	fmt.Println(q)
}
