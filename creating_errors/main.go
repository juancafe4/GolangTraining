package main

import (
	// "errors"
	"fmt"
	"log"
)

type SqurtError struct {
	lat, long string
	err       error
}

func (n *SqurtError) Error() string {
	return fmt.Sprintf("A negative value %v %v %v\n", n.lat, n.long, n.err)
}

// var SqurtError = errors.New("Error negative value")
func main() {
	_, err := Squrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func Squrt(f float64) (float64, error) {
	if f < 0 {
		sqrtErr := fmt.Errorf("Negative value for square root %v\n", f)
		return -1, &SqurtError{"50.2289 N", "99.4656 W", sqrtErr}
	}

	// Implementation
	return 42, nil
}
