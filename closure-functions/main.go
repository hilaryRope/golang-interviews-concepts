package main

import "fmt"

/*
 A closure is a function value that references variables from outside its body.
The function may access and assign to the referenced variables; in this sense, the function is "bound" to the variables.
*/

func main() {
	f := incrementor()
	fmt.Println(f())
}

func incrementor() func() int {
	x := 0
	return func() int { // returns an anonymous function
		x++
		return x
	}
}
