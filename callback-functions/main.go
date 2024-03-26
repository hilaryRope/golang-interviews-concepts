package main

import "fmt"

// Callback functions are functions passed in as parameters in other functions

func main() {
	x := doOperation(1, 2, add)
	fmt.Println(x)

	sub := doOperation(4, 2, subtract)
	fmt.Println(sub)
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a + b
}

func doOperation(a int, b int, f func(int, int) int) int {
	return f(a, b)

}
