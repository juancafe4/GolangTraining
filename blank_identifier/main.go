package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	a := "I am A"
	// b := "I am B" b must be used

	fmt.Print(a)

	res, _ := http.Get("http://www.google.com")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	fmt.Printf("%s\n", page)
}
