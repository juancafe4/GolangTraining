package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int {
	return len(p)
}
func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}
func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
	s := []string{"Zeno", "John", "Al", "Jenny"}

	sort.StringSlice(s).Sort()
	// or sort.Strings(S)
	fmt.Println("String slice sorted", s)

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Sort(sort.Reverse(sort.IntSlice(n)))

	fmt.Println("Integers sorted: ", n)
	// Sort reverse impletes the interface to reverse sort

}
