package main

import "fmt"

type Dog struct {
	name string
}

func (d Dog) walk() {
	fmt.Println("My name is ", d.name, " and I am walking ")
}

func (d *Dog) run() {
	d.name = "Jeff"
	fmt.Println("My name is ", d.name, " and I am running ")
}

type Young interface {
	walk()
	run()
}

func youngRun(y Young) {
	y.run() // as defined with a pointer receiver
}

func main() {
	dog1 := Dog{"Jack"}
	dog1.walk()
	dog1.run()

	dog2 := &Dog{"Claire"}
	dog2.walk()
	dog2.run()

	youngRun(dog2) // this can be done as dog2's run is a pointer's receiver

}
