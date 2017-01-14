package main

import (
	"errors"
	"log"
)

func main() {
	_, err := Squrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func Squrt(f float64) (float64, error) {
	if f < 0 {
		return -1, errors.New("Negative value for square root")
	}

	// Implementation
	return 42, nil
}
