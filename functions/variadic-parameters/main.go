package main

import "fmt"

func main() {
	// demonstrating defer
	defer foo()
	bar()

	returnedSum := sum(1, 2, 3, 4)
	fmt.Println("The returned Sum as per the inputted numbers is ", returnedSum)

	// unfurling a slice
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi...)
	fmt.Println("unpacking the slice and the sum is ", x)
}

func sum(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\t", ii)

	n := 0
	for _, v := range ii {
		n += v
	}
	return n

}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
