package main

import (
	"fmt"
)

func visit(numbers []int, cb func(int)) {
	for _, n := range numbers {
		cb(n)
	}
}
func main() {
	visit([]int{1, 2, 4, 6}, func(n int) {
		fmt.Println(n)
	})
}
