package main

import "fmt"

// Some information about heap and stack memory allocation

//go build -gcflags "-m" ./main.go
//# command-line-arguments
//./main.go:11:6: can inline square
//./main.go:7:13: inlining call to square
//./main.go:8:13: inlining call to fmt.Println
//./main.go:8:13: ... argument does not escape
//./main.go:8:14: *y escapes to heap

func main() {
	x := 10
	y := square(&x)
	fmt.Println(*y)
}

func square(x *int) *int {
	z := (*x) * (*x)
	return &z
}
