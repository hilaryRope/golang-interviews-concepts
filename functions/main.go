package main

import "fmt"

const param = "Ilaria"

func main() {
	foo()            // no params, no receiver, no return value
	bar(param)       // one param, no receiver, no return value
	s := fuzz(param) // one param, no receiver, returns a string
	fmt.Println(s)
	s1, n := dogYears(param, 47) // 2 params, no receiver, returns two values
	fmt.Printf(s1, n)

}

func foo() {
	fmt.Println("I am from foo")
}

func bar(s string) {
	fmt.Println("My name is", s)

}

func fuzz(s string) string {
	return fmt.Sprint("My name is ", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years "), age
}
