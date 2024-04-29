package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrorNorgateMath = errors.New("norgate math: square root of negative number")

func main() {
	fmt.Printf("%T\n", ErrorNorgateMath)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}

}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrorNorgateMath
	}
	return 42, nil
}
