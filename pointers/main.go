package main

import "fmt"

func main() {
	x := 42
	println(x)
	println(&x)

	fmt.Printf("%v\t%T\n", &x, &x)

	y := &x
	fmt.Printf("%v\t%T\n", y, y)
	println(*y)  // the value of y as we are de-referencing
	println(*&y) // address
	println(&*y) // address

	*y = 43
	fmt.Println(x) // 43

	s := "sample"
	fmt.Printf("%v\t%T\n", &s, &s)
}
