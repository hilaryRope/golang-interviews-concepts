package main

import (
	"fmt"
	"log"
	"strconv"
)

type Book struct {
	title  string
	author string
}

func (b Book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("The count is ", strconv.Itoa(int(c)))
}

// Wrapper log.Println inside a new function that takes fmt.Stringer
func logInfo(s fmt.Stringer) {
	log.Println("LOG FROM THE EXERCISE ON STRINGER INTERFACE", s.String())
}
func main() {

	b := Book{
		title:  "Being and Time",
		author: "Martin Heidegger",
	}

	var n count = 42

	fmt.Println(b)
	fmt.Println(n)

	logInfo(b)
	logInfo(n)
}
